package main

/*
import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
*/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	ip, err := getPublicIPv4()
	if err != nil {
		fmt.Println("Error getting public IP address:", err)
		return
	}

	payload := map[string]interface{}{
		"uuid": "unique request identifier",
		"player": map[string]interface{}{
			"id":       "a1a2a3a4",
			"update":   true,
			"nickname": "nickname",
			"language": "en-GB",
			"currency": "EUR",
			"session": map[string]interface{}{
				"id": "111ssss3333rrrrr45555",
				"ip": ip,
			},
		},
		"config": map[string]interface{}{
			"brand": map[string]interface{}{
				"id":   "1",
				"skin": "1",
			},
		},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	apiURL := "https://skylinev88871.uat1.evo-test.com/ua/v1/skylinev88871001/test123"
	proxyURL, err := url.Parse("http://goproxy.u1s1.biz:16600")
	if err != nil {
		fmt.Println("Error parsing proxy URL:", err)
		return
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println(string(body))

}
