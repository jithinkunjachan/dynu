package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {

	user := os.Getenv("DYNU_USERNAME")
	if user == "" {
		log.Fatal("please provide DYNU_USERNAME")
	}
	pwd := os.Getenv("DYNU_PASSWORD")
	if pwd == "" {
		log.Fatal("please provide DYNU_PASSWORD")
	}

	host := os.Getenv("DYNU_HOST")
	if host == "" {
		host = "api.dynu.com"
	}

	h := md5.New()
	io.WriteString(h, pwd)
	u := url.URL{
		Scheme: "http",
		Host:   host,
		Path:   "nic/update",
	}
	q := u.Query()
	q.Add("username", user)
	q.Add("password", fmt.Sprintf("%x", h.Sum(nil)))
	u.RawQuery = q.Encode()

	c := http.Client{
		Timeout: time.Second * 20,
	}

	for {
		req, err := http.NewRequest(http.MethodGet, u.String(), nil)
		if err != nil {
			log.Fatalf("error while creating request %v", err)
		}

		resp, err := c.Do(req)
		if err != nil {
			log.Fatalf("error while quering dynu %v", err)
		}
		
		if resp.StatusCode != 200 {
			log.Fatalf("error while updating ip address status code [%v]", resp.StatusCode)
		}
		resp.Body.Close()
		log.Println("successfully updated")
		time.Sleep(time.Second * 50)
	}
}
