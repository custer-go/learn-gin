package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	. "topic/src"
)

func main() {
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("topicurl", TopicUrl)
	}

	v1 := router.Group("/v1/topics")
	{
		v1.GET("", GetTopicList)

		v1.GET("/:topic_id", GetTopicDetail)

		v1.Use(MustLogin())
		{
			v1.POST("", NewTopic)
			v1.DELETE("/:topic_id", DeleteTopic)
		}
	}

	router.Run() // 8080
}
