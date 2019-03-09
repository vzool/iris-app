package controllers

import (
	"time"

	"github.com/gobuffalo/uuid"

	"github.com/joho/godotenv"
	"github.com/kataras/iris"
	"github.com/o1egl/paseto"
	"github.com/vzool/iris-app/app"
)

// GetPing Handler
func GetPing(ctx iris.Context) {

	env, _ := godotenv.Read()

	ctx.JSON(iris.Map{
		"message": "pong",
		"env":     env,
	})
}

// GetToken Handler
func GetToken(ctx iris.Context) {

	symmetricKey := []byte(app.SecretKey()) // Must be 32 bytes
	now := time.Now()
	exp := now.Add(24 * time.Hour)
	nbt := now

	uuid, _ := uuid.NewV4()

	jsonToken := paseto.JSONToken{
		Audience:   ctx.Path(),
		Issuer:     "iris-app",
		Jti:        uuid.String(),
		Subject:    "person@example.com",
		IssuedAt:   now,
		Expiration: exp,
		NotBefore:  nbt,
	}
	// Add custom claim	to the token
	jsonToken.Set("data", "this is a signed message")
	footer := "some footer"

	v2 := paseto.NewV2()

	// Encrypt data
	token, err := v2.Encrypt(symmetricKey, jsonToken, footer)

	if err != nil {

		ctx.JSON(iris.Map{
			"status":  "ERROR",
			"message": err.Error,
		})

	} else {

		ctx.JSON(iris.Map{
			"status": "OK",
			"token":  token,
		})
	}
}

// VerifyToken Handler
func VerifyToken(ctx iris.Context) {

	token := ctx.PostValue("token")

	symmetricKey := []byte(app.SecretKey()) // Must be 32 bytes

	// Decrypt data
	var newJsonToken paseto.JSONToken
	var newFooter string
	err := paseto.NewV2().Decrypt(token, symmetricKey, &newJsonToken, &newFooter)

	if err != nil {

		ctx.JSON(iris.Map{
			"status":  "ERROR",
			"message": err.Error,
		})

	} else {

		ctx.JSON(iris.Map{
			"status":       "OK",
			"token":        token,
			"newFooter":    newFooter,
			"newJsonToken": newJsonToken,
		})
	}
}
