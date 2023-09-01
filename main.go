package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"io"
	"net/http"
	"os"
)

type IPAddress struct {
	Origin string `json:"origin"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		ip, err := getIPAddress()
		if err != nil {
			return
		}
		_, err = w.Write([]byte(ip.Origin))
		if err != nil {
			return
		}
	})
	err := http.ListenAndServe(":"+os.Getenv("PORT"), r)
	if err != nil {
		return
	}
}

// 'https://httpbin.org/ip'  からIPアドレスを取得する
func getIPAddress() (IPAddress, error) {
	url := "https://httpbin.org/ip"
	req, err := http.NewRequest("GET", url, nil)
	client := new(http.Client)
	resp, err := client.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)
	if err != nil {
		return IPAddress{}, err
	}
	body, _ := io.ReadAll(resp.Body)
	var ipAddress IPAddress
	err = json.Unmarshal(body, &ipAddress)
	if err != nil {
		return IPAddress{}, err
	}
	return ipAddress, nil
}
