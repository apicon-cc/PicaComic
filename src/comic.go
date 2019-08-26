package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
)

func (s *Service) Categories(c *gin.Context) (int, interface{}){
	token := c.GetHeader("token")
	data := Send("/categories", "GET", token, "")

	if data.Get("code").MustInt() != 200 {
		return s.makeErrJSON(data.Get("code").MustInt(), data.Get("code").MustInt(), data.Get("message").MustString())
	}

	categories := data.Get("data")
	return s.makeSuccessJSON(categories)
}

func (s *Service) ComicInfo(c *gin.Context) (int, interface{}){
	token := c.GetHeader("token")
	id := c.Param("id")
	data := Send("/comics/" + id, "GET", token,"")

	if data.Get("code").MustInt() != 200 {
		return s.makeErrJSON(data.Get("code").MustInt(), data.Get("code").MustInt(), data.Get("message").MustString())
	}

	return s.makeSuccessJSON(data.Get("data").Get("comic"))
}

func (s *Service) ComicEpisode(c *gin.Context) (int, interface{}){
	token := c.GetHeader("token")
	id := c.Param("id")
	page, err := c.GetQuery("page")

	if !err {
		page = "1"
	}

	data := Send(fmt.Sprintf("/comics/%s/eps?page=%s", id, page), "GET", token,"")

	if data.Get("code").MustInt() != 200 {
		return s.makeErrJSON(data.Get("code").MustInt(), data.Get("code").MustInt(), data.Get("message").MustString())
	}

	return s.makeSuccessJSON(data.Get("data").Get("eps").Get("docs"))
}

func (s *Service) ComicDetail(c *gin.Context) (int, interface{}){
	token := c.GetHeader("token")
	bookID := c.Param("bookID")
	episodeID := c.Param("episodeID")
	page, err := c.GetQuery("page")

	if !err{
		page = "1"
	}

	data := Send(fmt.Sprintf("/comics/%s/order/%s/pages?page=%s", bookID, episodeID, page), "GET", token,"")

	if data.Get("code").MustInt() != 200 {
		return s.makeErrJSON(data.Get("code").MustInt(), data.Get("code").MustInt(), data.Get("message").MustString())
	}
	fmt.Println(data)

	return s.makeSuccessJSON(data.Get("data"))
}

func (s *Service) ComicMedia(c *gin.Context) gorequest.Response{
	token := c.GetHeader("token")
	url := c.Param("url")

	picture := GetImage("/" + url, token)
	return picture
}