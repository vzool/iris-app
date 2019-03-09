package controllers

import (
	"github.com/kataras/iris"
	"github.com/vzool/iris-app/app"
	model "github.com/vzool/iris-app/app/models"
)

// TaskIndex Handler
func TaskIndex(ctx iris.Context) {

	db := app.DB()
	defer db.Close()

	data := []model.Task{}
	db.Find(&data)

	ctx.JSON(iris.Map{
		"status": "OK",
		"data":   data,
	})
}

// TaskStore Handler
func TaskStore(ctx iris.Context) {

	name := ctx.PostValue("name")

	db := app.DB()
	defer db.Close()

	data := model.Task{
		Name: name,
	}

	db.Create(&data)

	if data.ID != 0 {

		ctx.JSON(iris.Map{
			"status": "OK",
			"data":   data,
		})

	} else {

		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{
			"status":  "ERROR",
			"message": "Failed to save create record",
		})
	}
}

// TaskShow Handler
func TaskShow(ctx iris.Context) {

	id := ctx.Params().GetUint64Default("id", 0)

	db := app.DB()
	defer db.Close()

	data := model.Task{}

	db.Find(&data, id)

	if len(data.Name) > 0 {

		ctx.JSON(iris.Map{
			"status": "OK",
			"data":   data,
		})

	} else {

		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{
			"message": "Record not found",
		})
	}
}

// TaskUpdate Handler
func TaskUpdate(ctx iris.Context) {

	id := ctx.Params().GetUint64Default("id", 0)

	name := ctx.PostValue("name")

	db := app.DB()
	defer db.Close()

	data := model.Task{}
	db.Find(&data, id)

	if name != "" {

		if data.Name != name {

			data.Name = name
			db.Save(&data)
		}

		ctx.JSON(iris.Map{
			"status": "OK",
			"data":   data,
		})

	} else {

		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{
			"status":  "ERROR",
			"message": "Record not found",
		})
	}
}

// TaskDelete Handler
func TaskDelete(ctx iris.Context) {

	id := ctx.Params().GetUint64Default("id", 0)

	db := app.DB()
	defer db.Close()

	data := model.Task{}
	db.Find(&data, id)

	if len(data.Name) > 0 {

		db.Delete(&data, id)

		ctx.JSON(iris.Map{
			"status": "OK",
			"data":   data,
		})

	} else {

		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{
			"status":  "ERROR",
			"message": "Record not found",
		})
	}
}
