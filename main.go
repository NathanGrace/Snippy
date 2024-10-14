package main

import (
	"context"
	"fmt"
	"github.com/pocketbase/pocketbase/apis"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

type Appointments struct {
	Id      string `db:"id" json:"id"`
	StaffId string `db:"staff_id" json:"staff_id"`
}

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	startPocketbase()

	return nil
}

func startPocketbase() {
	pb := pocketbase.New()

	addCustomRoutes(pb)

	//serves static files from the provided public dir (if exists)
	//I think pb_public is where you could put your front end code
	//pb.OnBeforeServe().Add(func(e *core.ServeEvent) error {
	//	e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
	//	return nil
	//})

	if err := pb.Start(); err != nil {
		log.Fatal(err)
	}
}

func addCustomRoutes(pb *pocketbase.PocketBase) {

	pb.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/hello/:name", func(c echo.Context) error {
			name := c.PathParam("name")

			return c.JSON(http.StatusOK, map[string]string{"message": "Hello " + name})
		}) /* optional middlewares */

		return nil
	})

	pb.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/snippy/appointments", func(c echo.Context) error {

			appointments, err := pb.Dao().FindCollectionByNameOrId("appointments")
			if err != nil {
				return err
			}

			return c.JSON(http.StatusOK, appointments)
		})
		return nil
	})

	pb.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// attach a middleware globally to all routes
		//e.Router.Use(someMiddlereFunc)

		// attach multiple middlewares to a single route
		// each route will execute their own middlewares + the global ones
		e.Router.GET("/hello", func(c echo.Context) error {
			return c.String(200, "Hello world!")
		}, apis.ActivityLogger(pb), apis.RequireAdminAuth())
		return nil
	})

}
