package main

import (
	"fmt"
	"sky-tech/config"
	"sky-tech/handler"
	"sky-tech/middleware"
	"sky-tech/repository"
	"sky-tech/script"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())

	db := config.ConnectDB()
	defer db.Close()

	fmt.Println("DB connected")

	// Script to load data
	script.Ingest_metrics(db)

	// Retrieve data
	repo := repository.NewMetricRepository(db)
	h := handler.New(repo)

	e.GET("/metrics", h.GetMetrics)

	e.Logger.Fatal(e.Start(":8080"))

}
