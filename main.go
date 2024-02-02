package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wisnu31899/fwg17-go-backend/src/routers"
)

// type pageInfo struct {
// 	Page string `json:"page"`
// }

// type response struct {
// 	Success  bool        `json:"success"`
// 	Message  string      `json:"message"`
// 	PageInfo pageInfo    `json:"pageInfo"`
// 	Results  interface{} `json:"results"`
// }

// type User struct {
// 	Id       int    `json:"id"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

func main() {
	r := gin.Default()
	routers.Combine(r)
	r.GET("/:users", func(ctx *gin.Context) {})
	r.Run()
}
