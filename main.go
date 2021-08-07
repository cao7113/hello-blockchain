package main

import (
	"github.com/cao7113/hellobc/api"
	"github.com/cao7113/hellobc/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

const contractAddr = "0x12B033a179238656C62D4a278528CA8910Fd2fdF"

func main() {
	client, err := ethclient.Dial(config.NodeURL)
	if err != nil {
		panic(err)
	}
	conn, err := api.NewApi(common.HexToAddress(contractAddr), client)
	if err != nil {
		panic(err)
	}

	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/greet/:message", func(c echo.Context) error {
		message := c.Param("message")
		reply, err := conn.Greet(&bind.CallOpts{}, message)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, reply)
	})

	e.GET("/hello", func(c echo.Context) error {
		reply, err := conn.Hello(&bind.CallOpts{})
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, reply) // Hello World
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
