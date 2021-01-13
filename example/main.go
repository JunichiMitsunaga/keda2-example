package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

var sl *zap.SugaredLogger

type Response struct {
	Components struct {
		Worker struct {
			Tasks int `json:"tasks"`
		} `json:"worker"`
	} `json:"components"`
}

func init() {
	// set logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sl = logger.Sugar()

}

func main() {
	// echo initialization
	e := echo.New()
	e.GET("/components", getComponents)
	if err := e.Start(":3003"); err != nil {
		sl.Errorf("ERROR: %v", err.Error())
	}
}

func getComponents(c echo.Context) error {
	res := &Response{}
	res.Components.Worker.Tasks, _ = strconv.Atoi(os.Getenv("WORKER_TASKS"))
	sl.Info(fmt.Sprintf(`response tasks :%d`, res.Components.Worker.Tasks))
	return c.JSON(http.StatusOK, res)
}
