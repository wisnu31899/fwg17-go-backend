package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/KEINOS/go-argonize"
	"github.com/gin-gonic/gin"
	"github.com/wisnu31899/fwg17-go-backend/src/lib"
	"github.com/wisnu31899/fwg17-go-backend/src/models"
	"github.com/wisnu31899/fwg17-go-backend/src/services"
)

type FormReset struct {
	Email           string `json:"email" form:"email" binding:"email"`
	Otp             string `form:"otp"`
	Password        string `form:"password"`
	ConfirmPassword string `form:"confirmPassword"`
}

func Login(c *gin.Context) {
	form := models.User{}
	err := c.ShouldBind(&form)

	if err != nil {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Invalid",
		})
		return
	}

	found, err := models.FindUserByEmail(form.Email)

	if err != nil {
		c.JSON(http.StatusUnauthorized, &services.ResponseOnly{
			Success: false,
			Message: "wrong email or password",
		})
		return
	}
	decoded, err := argonize.DecodeHashStr(found.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, &services.ResponseOnly{
			Success: false,
			Message: "wrong email or password",
		})
		return
	}

	plain := []byte(form.Password)
	if decoded.IsValidPassword(plain) {
		c.JSON(http.StatusOK, &services.ResponseOnly{
			Success: true,
			Message: "login successfully",
		})
		return
	} else {
		c.JSON(http.StatusUnauthorized, &services.ResponseOnly{
			Success: false,
			Message: "wrong email or password",
		})
		return
	}
}

func Register(c *gin.Context) {
	form := models.User{}

	err := c.ShouldBind(&form)

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Invalid",
		})
		return
	}

	plain := []byte(form.Password)
	hash, err := argonize.Hash(plain)

	if err != nil {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "failed to hased password",
		})
		return
	}

	form.Password = hash.String()

	_, err = models.CreateUser(form)

	if err != nil {
		if strings.HasSuffix(err.Error(), `unique constraint "users_email_key"`) {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "email already to app",
			})
			return
		}
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "failed to register",
		})
		return
	}

	c.JSON(http.StatusOK, &services.ResponseOnly{
		Success: true,
		Message: "register successfully",
	})
}

func ForgotPassword(c *gin.Context) {
	form := FormReset{}
	c.ShouldBind(&form)
	if form.Email != "" {
		found, _ := models.FindUserByEmail(form.Email)
		if found.Id != 0 {
			formReset := models.FormReset{
				Otp:   lib.RandomNumberStr(6),
				Email: found.Email,
			}
			models.CreateResetPassword(formReset)
			//kirim email
			fmt.Println(formReset.Otp)
			//email berakhir
			c.JSON(http.StatusOK, &services.ResponseOnly{
				Success: true,
				Message: "otp send your email",
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "failed to resetPassword",
			})
			return
		}
	}
	if form.Otp != "" {
		found, _ := models.FindByOtp(form.Otp)
		if found.Id != 0 {
			if form.Password == form.ConfirmPassword {
				foundUser, _ := models.FindUserByEmail(found.Email)
				data := models.User{
					Id: foundUser.Id,
				}

				hash, _ := argonize.Hash([]byte(form.Password))
				data.Password = hash.String()

				updated, _ := models.UpdateUser(data)
				message := fmt.Sprintf("reset password for %s success", updated.Email)
				c.JSON(http.StatusOK, &services.ResponseOnly{
					Success: true,
					Message: message,
				})
				models.DeleteResetPassword(found.Id)
				return
			} else {
				c.JSON(http.StatusBadRequest, &services.ResponseOnly{
					Success: false,
					Message: "confirmPassword does not match",
				})
				return
			}
		}
	}
	c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
		Success: false,
		Message: "internal server error",
	})
	return
}
