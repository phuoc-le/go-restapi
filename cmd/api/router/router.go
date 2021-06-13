package router

import (
	"github.com/phuoc-le/go-restapi/cmd/api/handlers/createuser"
	"github.com/phuoc-le/go-restapi/cmd/api/handlers/getuser"
	"github.com/phuoc-le/go-restapi/pkg/application"
	"github.com/julienschmidt/httprouter"
)

func Get(app *application.Application) *httprouter.Router {
	mux := httprouter.New()
	mux.GET("/users/:id", getuser.Do(app))
	mux.POST("/users", createuser.Do(app))
	return mux
}
