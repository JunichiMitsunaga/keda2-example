package main

import (
	"fmt"
	"net/http"
	"os"
	"unsafe"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

var sl *zap.SugaredLogger

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
	sl.Info("Called")
	task := os.Getenv("WORKER_TASKS")
	return c.Blob(http.StatusOK, echo.MIMEApplicationJSON, Str2Bytes(fmt.Sprintf(`{"components":{"worker":{"tasks":'%s'}}} `, task)))
}

// Str2Bytes transfer string to bytes
func Str2Bytes(str string) []byte {
	vec := *(*[]byte)(unsafe.Pointer(&str))
	return vec
}
