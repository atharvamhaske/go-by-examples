package main

import (
	"log"
	"os"

	"github.com/gedex/go-instagram/instagram"
	"github.com/labstack/echo/v4"
)

func main() {
	clientID := os.Getenv("InstagramID")
	if clientID == "" {
		log.Fatalln("Please set InstagramID env variable")
	}

	instaClient := instagram.NewClient(nil)
	instaClient.ClientID = clientID

	app := NewApp(instaClient)

	e := echo.New()

	e.POST("/download/:username", func(c echo.Context) error {
		username := c.Param("username")

		go app.StartDownload(username)

		return c.JSON(200, map[string]string{
			"status": "started",
			"user":   username,
		})
	})

	log.Println("Server running on :8080")
	e.Logger.Fatal(e.Start(":8080"))
}

//lol this api services were stopped by meta
