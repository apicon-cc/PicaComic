package main

import (
	"github.com/gin-gonic/gin"
)

func (s *Service) initRouter() {

	r := gin.Default()

	{
		//GET /authorization
		r.POST("/authorization", func(c *gin.Context) {
			c.JSON(s.Login(c))
		})
	}

	{
		//GET /categories
		r.GET("/categories", func(c *gin.Context){
			c.JSON(s.Categories(c))
		})

		//GET /search
		r.GET("/search", func(c *gin.Context){
			c.JSON(s.SearchKeyword(c))
		})
	}

	{
		//GET /comic/:id
		r.GET("/comic/:id", func(c *gin.Context){
			c.JSON(s.ComicInfo(c))
		})

		//GET /episode/:id
		r.GET("/episode/:id", func(c *gin.Context){
			c.JSON(s.ComicEpisode(c))
		})

		//GET /content/:bookID/order/:episodeID
		r.GET("/content/:bookID/order/:episodeID", func(c *gin.Context){
			c.JSON(s.ComicDetail(c))
		})

		//GET /image/:url
		r.GET("/image/:url", func(c *gin.Context){
			data := s.ComicMedia(c)
			c.DataFromReader(200, data.ContentLength,"image/png", data.Body, map[string]string{})
		})
	}

	s.Router = r
	err := s.Router.Run(s.Config.Get("server.port").(string))
	panic(err)
}