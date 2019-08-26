package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (s *Service) Login(c *gin.Context) (int, interface{}){
	l := new(login)
	err := c.ShouldBindJSON(&l)
	if err != nil{
		return s.makeErrJSON(400, 40000, "Data format error.")
	}

	r := Send("/auth/sign-in", "POST", "", fmt.Sprintf(`{"email":"%s", "password":"%s"}`, l.Email, l.Password))
	return s.makeSuccessJSON(r.Get("data").Get("token").MustString())
}