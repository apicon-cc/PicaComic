package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/url"
)

func (s *Service) SearchKeyword(c *gin.Context) (int, interface{}){
	token := c.GetHeader("token")
	keyword, _ := c.GetQuery("keyword")
	page, _ := c.GetQuery("page")

	data := Send(fmt.Sprintf("/comics/search?page=%s&q=%s", page, url.QueryEscape(keyword)), "GET", token, "")

	if data.Get("code").MustInt() != 200 {
		return s.makeErrJSON(data.Get("code").MustInt(), data.Get("code").MustInt(), data.Get("message").MustString())
	}

	return s.makeSuccessJSON(data.Get("data").Get("comics"))
}
