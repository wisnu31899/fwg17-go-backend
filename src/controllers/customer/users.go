package customer

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/KEINOS/go-argonize"
	"github.com/gin-gonic/gin"
	"github.com/wisnu31899/fwg17-go-backend/src/lib"
	"github.com/wisnu31899/fwg17-go-backend/src/models"
	"github.com/wisnu31899/fwg17-go-backend/src/services"
)

func GetDetailUser(c *gin.Context) {
	// Mengambil data user yang sedang login dari konteks
	data := c.MustGet("id").(*models.User)
	id := data.Id

	// Mengambil detail user dari database berdasarkan id user yang sedang login
	user, err := models.FindOneUser(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusNotFound, &services.ResponseOnly{
				Success: false,
				Message: "user not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, &services.ResponseDetail{
		Success: true,
		Message: "get detail user successfully",
		Results: user,
	})
}

func UpdateUser(c *gin.Context) {
	// Mengambil data user yang sedang login dari konteks
	data := c.MustGet("id").(*models.User)
	id := data.Id

	// Mengikat data yang dikirim oleh klien ke variabel data
	c.Bind(&data)
	data.Id = id

	// Mendapatkan nama file foto lama sebelum menggantinya dengan foto baru
	oldUser, _ := models.FindOneUser(id)
	var oldPicture *string
	if oldUser.Picture != nil {
		oldPicture = oldUser.Picture
	}

	// Menangkap kedua nilai yang dikembalikan oleh lib.UploadFile
	picture, err := lib.UploadFile(c, "profile")
	if err != nil {
		if err.Error() == "file size exceeds the limit of 5MB" {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "file size exceeds the limit of 5MB",
			})
		} else if err.Error() == "file format is not supported" {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "file format is not supported. Only JPG or PNG files are allowed",
			})
		} else {
			c.JSON(http.StatusBadRequest, &services.ResponseOnly{
				Success: false,
				Message: "failed to upload picture",
			})
		}
		return
	}

	// Menetapkan nilai Picture ke dalam data.User
	data.Picture = &picture

	plain := []byte(data.Password)
	hash, err := argonize.Hash(plain)
	if err != nil {
		c.JSON(http.StatusBadRequest, &services.ResponseOnly{
			Success: false,
			Message: "Failed hash password",
		})
		return
	}

	data.Password = hash.String()

	user, err := models.UpdateUser(*data)

	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	// Memeriksa apakah foto lama ada
	if oldPicture != nil {
		// Menghapus foto lama dari folder penyimpanan
		err := os.Remove(*oldPicture)
		if err != nil {
			log.Println("Failed to delete old picture:", err)
		}
	}

	c.JSON(http.StatusOK, &services.ResponseDetail{
		Success: true,
		Message: "update user successfully",
		Results: user,
	})
}
