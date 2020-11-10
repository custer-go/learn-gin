package classes

import (
	"gin-up/src/goft"
	"github.com/gin-gonic/gin"
)

// UserClass 
type UserClass struct {}

// NewUserClass UserClass generate constructor
func NewUserClass() *UserClass {
	return &UserClass{}
}

// UserList 控制器方法
func (this *UserClass) UserList() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"result": "success",
		})
	}
}

func (this *UserClass) Build(goft *goft.Goft) {
	goft.Handle("GET", "/user", this.UserList())
}
