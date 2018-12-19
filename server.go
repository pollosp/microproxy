package main

import (
	"net/url"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Static("/", "static")

	url1, err := url.Parse(os.Getenv("PROXY"))
	if err != nil {
		e.Logger.Fatal(err)
	}
	targets := []*middleware.ProxyTarget{
		{
			URL: url1,
		},
	}
	g := e.Group("/api")
	g.Use(middleware.Proxy(middleware.NewRoundRobinBalancer(targets)))

	e.Logger.Info(e.Start(os.Getenv("ADDRESS")))
}
