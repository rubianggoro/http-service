package main

import (
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/rubianggoro/http-service/model"
)

func app(e *echo.Echo, store model.ArticleStore) {
	// curl http://localhost:8080/articles
	e.GET("/articles", func(c echo.Context) error {
		article := store.All()

		// Response
		return c.JSON(http.StatusOK, article)
	})

	// curl http://localhost:8080/articles/1
	e.GET("articles/:id", func(c echo.Context) error {
		// Given
		id, _ := strconv.Atoi(c.Param("id"))

		// Process
		article := store.Find(id)

		// Response
		return c.JSON(http.StatusOK, article)
	})

	// curl -d "title=Hello title&body=ini body article" http://localhost:8080/articles
	e.POST("/articles", func(c echo.Context) error {
		title := c.FormValue("title")
		body := c.FormValue("body")
		article, _ := model.CreateArticle(title, body)
		store.Save(article)
		return c.JSON(http.StatusOK, article)
	})

	// curl -X PUT -d "title=Hello title&body=ini body article" http://localhost:8080/articles/1
	e.PUT("articles/:id", func(c echo.Context) error {
		// Given
		id, _ := strconv.Atoi(c.Param("id"))

		// Process
		article := store.Find(id)
		article.Title = c.FormValue("title")
		article.Body = c.FormValue("body")

		// Persists
		store.Update(article)

		// Response
		return c.JSON(http.StatusOK, article)
	})

	// curl -X DELETE http://localhost:8080/articles/1
	e.DELETE("articles/:id", func(c echo.Context) error {
		// Given
		id, _ := strconv.Atoi(c.Param("id"))

		// Process
		article := store.Find(id)

		// Remove
		store.Delete(article)

		// Response
		return c.JSON(http.StatusOK, article)
	})
}

func main() {
	var store model.ArticleStore
	driver := "mysql"

	if driver == "inmemory" {
		store = model.NewArticleStoreInmemory()
	} else {
		store = model.NewArticleStoreMySQL()
	}

	e := echo.New()
	app(e, store)

	e.Logger.Fatal(e.Start(":8080"))
}
