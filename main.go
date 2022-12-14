package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/sing3demons/devopsgo/rest/handler"
)

func initDB() *sql.DB {
	connStr := os.Getenv("DB_CONNECTION")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	createTb := `
	CREATE TABLE IF NOT EXISTS news_articles (
	id SERIAL PRIMARY KEY,
	"title" text,
	"content" text,
	"author" text
	);
	`
	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("can't create table", err)
	}

	return db
}

func main() {
	_, err := os.Create("/tmp/live")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove("/tmp/live")

	db := initDB()
	h := handler.NewApplication(db)

	e := echo.New()
	e.GET("/", h.Greeting)
	e.GET("/healthz", func(c echo.Context) error { return c.JSON(http.StatusOK, "OK") })
	// Intentionally, not setup database at this moment so we ignore feature to access database
	e.GET("/news", h.ListNews)
	e.POST("/news", h.CreateNews)
	serverPort := ":" + os.Getenv("PORT")
	e.Logger.Fatal(e.Start(serverPort))
}
