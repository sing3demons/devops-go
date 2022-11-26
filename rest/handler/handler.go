package handler

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct {
	DB *sql.DB
}

func NewApplication(db *sql.DB) *handler {
	return &handler{db}
}

func (h *handler) Greeting(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

type NewsArticle struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func (h *handler) CreateNews(c echo.Context) error {
	var m NewsArticle
	if err := c.Bind(&m); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	row := h.DB.QueryRow("INSERT INTO news_articles (title, content, author) VALUES ($1,$2,$3) RETURNING id", m.Title, m.Content, m.Author)
	err := row.Scan(&m.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, m.ID)
}

func (h *handler) ListNews(c echo.Context) error {
	rows, err := h.DB.Query("SELECT * FROM news_articles")
	if err != nil {
		return err
	}
	defer rows.Close()

	var nn = []NewsArticle{}
	var n = NewsArticle{}

	for rows.Next() {
		err := rows.Scan(&n.ID, &n.Title, &n.Content, &n.Author)
		if err != nil {
			log.Fatal(err)
		}
		nn = append(nn, n)
	}

	return c.JSON(http.StatusOK, nn)
}
