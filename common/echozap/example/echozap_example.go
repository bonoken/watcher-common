package example

import (
	"github.com/bonoken/watcher-common/common/echozap"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"log"
	"net"
)

var (
	Logger *zap.SugaredLogger
)

func mainTest() {

	// echo
	serveEcho()

}

func serveEcho() {
	// echo
	e := echo.New()

	// Logger
	// example
	//e.Use(middleware.Logger())
	e.Use(echozap.Logger(Logger.Desugar()))

	e.Use(middleware.Recover())

	// Access-Control-Allow-Origin
	e.Use(middleware.CORS())

	// Http server (echo)
	listener, err := net.Listen("tcp", ":"+"8080")
	if err != nil {
		log.Fatal(err)
	}
	e.Listener = listener

	e.Logger.Fatal(e.StartServer(e.Server))

}
