package main

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {

	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	// create client
	client := &http.Client{
		Timeout:   time.Second * 30,
		Transport: transport,
	}

	// req
	resp, err := client.Get("http://127.0.0.1:8080/bye")
	if err != nil {
		log.Fatalln("req err")
		return
	}

	defer resp.Body.Close()
	// 读取内容
	bds, _ := ioutil.ReadAll(resp.Body)

	log.Printf("data is %+v", string(bds))

}
