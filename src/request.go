package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/bitly/go-simplejson"
	"github.com/parnurzeal/gorequest"
	UUID "github.com/satori/go.uuid"
	"strconv"
	"strings"
	"time"
)

func computeHmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

func Send(url string, method string, authorization string, payload string) simplejson.Json{

	url = "https://picaapi.picacomic.com" + url
	secretKey := "~n}$S9$lGts=U)8zfL/R.PM9;4[3|@/CEsl~Kk!7?BYZ:BAa5zkkRBL7r|1/*Cr"
	apiKey := "C69BAF41DA5ABD1FFEDC6D2FEA56B"
	appVersion := "2.1.0.7"
	appChannel := "1"
	buildVersion := "40"
	accept := "application/vnd.picacomic.com.v1+json"
	appPlatform := "android"
	appUUID := UUID.NewV4().String()
	userAgent := "okhttp/3.8.1"
	host := "picaapi.picacomic.com"
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := strings.Replace(UUID.NewV4().String(), "-", "", -1)

	signature := strings.Replace(url, "https://picaapi.picacomic.com/", "", -1)
	signature = signature + timestamp + nonce + method + apiKey
	signature = strings.ToLower(signature)
	signature = computeHmacSha256(signature, secretKey)

	request := gorequest.New()

	var body string

	if method == "GET" {
		_, body, _ = request.Get(url).
			Set("api-key", apiKey).
			Set("app-version", appVersion).
			Set("app-channel", appChannel).
			Set("app-build-version", buildVersion).
			Set("accept", accept).
			Set("app-platform", appPlatform).
			Set("app-uuid", appUUID).
			Set("User-Agent", userAgent).
			Set("Host", host).
			Set("nonce", nonce).
			Set("time", timestamp).
			Set("signature", signature).
			Set("authorization", authorization).
			Set("Content-Type", "application/json; charset=UTF-8").
			End()
	}else if method == "POST"{
		_, body, _ = request.Post(url).
			Set("api-key", apiKey).
			Set("app-version", appVersion).
			Set("app-channel", appChannel).
			Set("app-build-version", buildVersion).
			Set("accept", accept).
			Set("app-platform", appPlatform).
			Set("app-uuid", appUUID).
			Set("User-Agent", userAgent).
			Set("Host", host).
			Set("nonce", nonce).
			Set("time", timestamp).
			Set("signature", signature).
			Set("authorization", authorization).
			Set("Content-Type","application/json; charset=UTF-8").
			Send(payload).
			End()
	}

	json, _ := simplejson.NewJson([]byte(body))

	return *json
}

func GetImage(url string, authorization string) gorequest.Response{

	url = "https://s3.picacomic.com/static" + url
	secretKey := "~n}$S9$lGts=U)8zfL/R.PM9;4[3|@/CEsl~Kk!7?BYZ:BAa5zkkRBL7r|1/*Cr"
	apiKey := "C69BAF41DA5ABD1FFEDC6D2FEA56B"
	appVersion := "2.1.0.7"
	appChannel := "1"
	buildVersion := "40"
	accept := "application/vnd.picacomic.com.v1+json"
	appPlatform := "android"
	appUUID := UUID.NewV4().String()
	userAgent := "okhttp/3.8.1"
	host := "s3.picacomic.com"
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := strings.Replace(UUID.NewV4().String(), "-", "", -1)

	signature := strings.Replace(url, "https://s3.picacomic.com/", "", -1)
	signature = signature + timestamp + nonce + "GET" + apiKey
	signature = strings.ToLower(signature)
	signature = computeHmacSha256(signature, secretKey)

	request := gorequest.New()

	resp, _, _ := request.Get(url).
		Set("api-key", apiKey).
		Set("app-version", appVersion).
		Set("app-channel", appChannel).
		Set("app-build-version", buildVersion).
		Set("accept", accept).
		Set("app-platform", appPlatform).
		Set("app-uuid", appUUID).
		Set("User-Agent", userAgent).
		Set("Host", host).
		Set("nonce", nonce).
		Set("time", timestamp).
		Set("signature", signature).
		Set("authorization", authorization).
		Set("Content-Type", "application/json; charset=UTF-8").
		End()

	return resp
}