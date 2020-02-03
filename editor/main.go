// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gobuffalo/packr/v2"
	"github.com/labstack/echo"
	"github.com/urfave/cli/v2"
)

var server string
var instance string
var user string
var password string
var db string
var port string

func main() {
	app := new(cli.App)
	app.Name = "webserver go"
	app.Usage = "run apps and access http://localhost:8080"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "httpport",
			Aliases:     []string{"hp"},
			Value:       ":8080",
			Usage:       "http access port no",
			Destination: &port,
		},
	}

	app.Action = action
	err := app.Run(os.Args)
	if err != nil {
		log.Println(err)
	}
}

func action(c *cli.Context) error {
	ec := echo.New()
	box := packr.New("webapps", "./public")

	ec.GET("/*", echo.WrapHandler(http.StripPrefix("/", http.FileServer(box))))
	err := ec.Start(port)
	if err != nil {
		return err
	}
	return nil
}
