package http

import (
	"github.com/kataras/iris"
	controller "github.com/vzool/iris-app/app/http/controllers"
)

func Routes(app *iris.Application) {

	app.Get("/ping", controller.GetPing)

	app.Get("/city", controller.CityIndex)
	app.Post("/city", controller.CityStore)
	app.Get("/city/{id:uint64}", controller.CityShow)
	app.Put("/city/{id:uint64}", controller.CityUpdate)
	app.Delete("/city/{id:uint64}", controller.CityDelete)

	app.Get("/task", controller.TaskIndex)
	app.Post("/task", controller.TaskStore)
	app.Get("/task/{id:uint64}", controller.TaskShow)
	app.Put("/task/{id:uint64}", controller.TaskUpdate)
	app.Delete("/task/{id:uint64}", controller.TaskDelete)
}
