package controllers

import (
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/vzool/iris-app/app"
	model "github.com/vzool/iris-app/app/models"
)

// CityIndex Handler
func CityIndex(ctx iris.Context) {

	db := app.DB()
	defer db.Close()

	data := []model.City{}
	db.Find(&data)

	ctx.JSON(iris.Map{
		"status": "OK",
		"data":   data,
	})
}

// CityStore Handler
func CityStore(ctx iris.Context) {

	db := app.DB()
	defer db.Close()

	name := ctx.PostValue("name")

	data := model.City{
		Name: name,
	}

	valid, err := govalidator.ValidateStruct(data)
	if err != nil {

		println("error: " + err.Error())

		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{
			"status":  "ERROR",
			"message": err.Error(),
		})

		return
	}

	if valid {

		db.Create(&data)

		if data.ID != 0 {

			ctx.JSON(iris.Map{
				"status": "OK",
				"data":   data,
			})

			return
		}
	}

	ctx.StatusCode(iris.StatusBadRequest)
	ctx.JSON(iris.Map{
		"status":  "ERROR",
		"message": "Failed to save create record",
	})
}

// CityShow Handler
func CityShow(ctx iris.Context) {

	id := ctx.Params().GetUint64Default("id", 0)

	db := app.DB()
	defer db.Close()

	data := model.City{}

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

// CityUpdate Handler
func CityUpdate(ctx iris.Context) {

	id := ctx.Params().GetUint64Default("id", 0)

	name := ctx.PostValue("name")

	db := app.DB()
	defer db.Close()

	data := model.City{}

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

// CityDelete Handler
func CityDelete(ctx iris.Context) {

	id := ctx.Params().GetUint64Default("id", 0)

	db := app.DB()
	defer db.Close()

	data := model.City{}

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
