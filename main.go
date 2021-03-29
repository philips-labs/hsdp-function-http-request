package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

func main() {
	viper.SetEnvPrefix("request")
	viper.SetDefault("method", "POST")
	viper.SetDefault("body", "")
	viper.SetDefault("url", "")
	viper.SetDefault("username", "")
	viper.SetDefault("password", "")
	viper.AutomaticEnv()

	client := resty.New()

	r := client.R()

	username := viper.GetString("username")
	password := viper.GetString("password")
	if username != "" {
		r = r.SetBasicAuth(username, password)
	}
	body := viper.GetString("body")
	if body != "" {
		r = r.SetBody(body)
	}
	method := viper.GetString("method")
	url := viper.GetString("url")
	resp, err := r.Execute(method, url)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("%v", resp, err)
}
