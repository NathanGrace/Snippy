package main

import (
	"context"
	"fmt"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"log"
	"os"
	"os/signal"
)

type Users struct {
	Id       string `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Email    string `db:"email" json:"email"`
	Name     string `db:"name" json:"name"`
	Avatar   string `db:"avatar" json:"avatar"`
	Created  string `db:"created" json:"created"`
	Updated  string `db:"updated" json:"updated"`
}

type StaffMembers struct {
	Id             string `db:"id" json:"id"`
	Name           string `db:"name" json:"name"`
	Phone          string `db:"phone" json:"phone"`
	StaffId        string `db:"staff_id" json:"staff_id"`
	AdditionalInfo string `db:"additional_info" json:"additional_info"`
	Created        string `db:"created" json:"created"`
	Updated        string `db:"updated" json:"updated"`
}

type Appointments struct {
	Id                  string `db:"id" json:"id"`
	StaffId             string `db:"staff_id" json:"staff_id"`
	ClientId            string `db:"client_id" json:"client_id"`
	IsAvailable         bool   `db:"is_available" json:"is_available"`
	AdditionalInfo      string `db:"additional_info" json:"additional_info"`
	Service             string `db:"service" json:"service"`
	AppointmentDateTime string `db:"appointment_date_time" json:"appointment_date_time"`
	ClientName          string `db:"client_name" json:"client_name"`
	Created             string `db:"created" json:"created"`
	Updated             string `db:"updated" json:"updated"`
}

type Clients struct {
	Id             string `db:"id" json:"id"`
	Name           string `db:"name" json:"name"`
	Phone          string `db:"phone" json:"phone"`
	StaffId        string `db:"staff_id" json:"staff_id"`
	AdditionalInfo string `db:"additional_info" json:"additional_info"`
	Created        string `db:"created" json:"created"`
	Updated        string `db:"updated" json:"updated"`
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

	if err := pb.Start(); err != nil {
		log.Fatal(err)
	}
}

func addCustomRoutes(pb *pocketbase.PocketBase) {
	pb.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// serves static files from the provided public dir (if exists)
		se.Router.GET("/{path...}", apis.Static(os.DirFS("./pb_public"), false))

		return se.Next()
	})

	//old way before v23

	//pb.OnBeforeServe().Add(func(e *core.ServeEvent) error {
	//
	//	//BASIC EXAMPLE GIVEN BY DOCS
	//	e.Router.GET("/hello1/:name1", func(c echo.Context) error {
	//		name := c.PathParam("name1")
	//
	//		return c.JSON(http.StatusOK, map[string]string{"message": "Hello " + name})
	//	}) /* optional middlewares */
	//
	//	//same example with different syntax.
	//	_, _ = e.Router.AddRoute(echo.Route{
	//		Method: http.MethodGet,
	//		Path:   "/hello2/:name2",
	//		Handler: func(c echo.Context) error {
	//			name := c.PathParam("name2")
	//			return c.JSON(200, map[string]string{"message": "Hello " + name + "!"})
	//		},
	//		Middlewares: []echo.MiddlewareFunc{
	//			/* optional middlewares */
	//		},
	//	})
	//
	//	//me testing actual stuff
	//	e.Router.GET("/snippy/appointments", func(c echo.Context) error {
	//
	//		appointments := []Appointments{}
	//		err := pb.Dao().DB().
	//			NewQuery("SELECT * FROM appointments").
	//			All(&appointments)
	//		if err != nil {
	//			return err
	//		}
	//
	//		return c.JSON(http.StatusOK, map[string][]Appointments{"appointments": appointments})
	//	})
	//
	//	e.Router.GET("/snippy/apts", func(c echo.Context) error {
	//
	//		appointments := []Appointments{}
	//		err := pb.Dao().DB().
	//			NewQuery("SELECT * FROM appointments").
	//			All(&appointments)
	//
	//		//staff_id = user_id
	//		//is_available = false
	//		//appointment_date_time >= selected Day start
	//		//appointment_date_time <= selected day end
	//		if err != nil {
	//			return err
	//		}
	//
	//		return c.JSON(http.StatusOK, map[string][]Appointments{"appointments": appointments})
	//	})
	//
	//	e.Router.GET("/snippy/auth", func(c echo.Context) error {
	//
	//		appointments := []Appointments{}
	//		err := pb.Dao().DB().
	//			NewQuery("SELECT * FROM appointments").
	//			All(&appointments) //returns this appointments slice
	//
	//		if err != nil {
	//			return err
	//		}
	//
	//		return c.JSON(http.StatusOK, map[string][]Appointments{"appointments": appointments})
	//	})
	//
	//	return nil
	//})

}
