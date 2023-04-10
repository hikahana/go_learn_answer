package main

import (
	"net/http"

	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/hikahana/go_learn_answer/api/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/hikahana/go_learn_answer/api/controller"
	"github.com/hikahana/go_learn_answer/api/model"
)

type (
	Response struct {
		Int64  int64  `json:"int64"`
		String string `json:"string"`
		World  *Item  `json:"world"`
	}

	Item struct {
		Text string `json:"text"`
	}
)

// @title go_learn API
// @version 1.0
// @description go_learn API
// @host localhost:1323
// @BasePath /
// @schemes http

func main() {
	e := echo.New()
	db, _ := model.DB.DB()
	defer db.Close()

	arrowOrigins := []string{"*"}
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: arrowOrigins,
	}))

	e.POST("/users", controller.CreateUser)
	e.GET("/users", controller.GetUsers)
	e.GET("/users/:id", controller.GetUser)
	e.PUT("/users/:id", controller.UpdateUser)
	e.DELETE("/users/:id", controller.DeleteUser)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/checkswagger/", checkswagger)

	e.Logger.Fatal(e.Start(":1323"))


}

// @Summary checkswagger
// @Description checkswagger
// @ID hello-world
// @Accept  json
// @Produce  json
// @Success 200 {string} string "OK"
// @Router /hello-world [get]
func checkswagger(c echo.Context) error {
	return c.JSON(http.StatusOK, &Response{
		Int64:  1,
		String: "string",
		World: &Item{
			Text: "checkswagger",
		},
	})
}
