package middlewares

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/KEINOS/go-argonize"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/wisnu31899/fwg17-go-backend/src/models"
	"github.com/wisnu31899/fwg17-go-backend/src/services"
)

func Auth() (*jwt.GinJWTMiddleware, error) {
	godotenv.Load()
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "fwg17-go-backend",
		Key:         []byte(os.Getenv("APP_SECRET")),
		IdentityKey: "id",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			user := data.(*models.User)
			return jwt.MapClaims{
				"id":   user.Id,
				"role": user.Role,
			}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)

			var id int
			var role string

			if idFloat, ok := claims["id"].(float64); ok {
				id = int(idFloat)
			} else {
				// Handle error
				return nil
			}

			if roleStr, ok := claims["role"].(string); ok {
				role = roleStr
			} else {
				// Handle error
				return nil
			}

			// Buat variabel rolePointer dengan tipe *string dan alokasikan memori untuknya
			rolePointer := &role

			return &models.User{
				Id:   id,
				Role: rolePointer, // Gunakan rolePointer di sini
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			form := models.User{}
			err := c.ShouldBind(&form)

			if err != nil {
				return nil, err
			}

			found, err := models.FindUserByEmail(form.Email)

			if err != nil {
				return nil, err
			}

			decoded, err := argonize.DecodeHashStr(found.Password)
			if err != nil {
				return nil, err
			}

			plain := []byte(form.Password)
			if decoded.IsValidPassword(plain) {
				return &models.User{
					Id:   found.Id,
					Role: found.Role,
				}, nil
			} else {
				return nil, errors.New("invalid_password")
			}
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			// Periksa apakah data adalah nil atau bukan *models.User
			user, ok := data.(*models.User)
			if !ok || user == nil || user.Role == nil {
				// Jika data adalah nil atau bukan *models.User, atau Role adalah nil, maka kembalikan false
				return false
			}

			// Check if user's role is allowed to access the requested endpoint
			if user.Role == nil {
				// Jika Role adalah nil, kembalikan false
				return false
			}

			// Check user's role and grant access accordingly
			if *user.Role == "customer" { // Customer role
				// Customer hanya bisa mengakses /routers/customer
				if strings.HasPrefix(c.Request.URL.Path, "/customer") {
					return true
				}
			} else if *user.Role == "admin" { // Admin role
				// Admin bisa mengakses semua router
				return true
			} else {
				// Handle other roles if necessary
			}

			// Jika role tidak sesuai, maka akses ditolak
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			//response salah email atau password
			if strings.HasPrefix(c.Request.URL.Path, "/auth/login") {
				c.JSON(http.StatusUnauthorized, &services.ResponseOnly{
					Success: false,
					Message: "wrong email or password",
				})
				return
			}
			// Response Unauthorized diletakkan di sini, tetapi jika sudah terjadi Unauthorized, Anda harus menghentikan eksekusi Authorizator
			c.JSON(http.StatusUnauthorized, &services.ResponseOnly{
				Success: false,
				Message: "Unauthorized",
			})
		},
		LoginResponse: func(c *gin.Context, code int, token string, time time.Time) {
			c.JSON(http.StatusOK, &services.ResponseDetail{
				Success: true,
				Message: "login successfully and wellcome to app",
				Results: struct {
					Token string `json:"token"`
				}{
					Token: token,
				},
			})
		},
	})
	if err != nil {
		return nil, err
	}
	return authMiddleware, nil
}
