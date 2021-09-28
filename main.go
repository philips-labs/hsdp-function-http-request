package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func main() {
	listenString := ":8080"

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

	runner := func(c echo.Context) error {
		fmt.Printf("%v %v\n", method, url)
		resp, err := r.Execute(method, url)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return err
		}
		c.String(resp.StatusCode(), string(resp.Body()))
		return nil
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Any("/", runner)

	e.Logger.Fatal(e.Start(listenString))
}
