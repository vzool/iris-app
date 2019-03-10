package controllers

import (
	"net/http"
	"time"

	"golang.org/x/crypto/blake2b"

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

		ctx.SetCookieKV("token", token, func(option *http.Cookie) {
			option.HttpOnly = true
		})
	}
}

// VerifyToken Handler
func VerifyToken(ctx iris.Context) {

	token := ctx.PostValue("token")

	symmetricKey := []byte(app.SecretKey()) // Must be 32 bytes

	// Decrypt data
	var newJSONToken paseto.JSONToken
	var newFooter string
	err := paseto.NewV2().Decrypt(token, symmetricKey, &newJSONToken, &newFooter)

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
			"newJsonToken": newJSONToken,
		})
	}
}

// Blake2BHash Handler
func Blake2BHash(ctx iris.Context) {

	hash, err := blake2b.New256([]byte(app.SecretKey()))

	if err != nil {

		ctx.JSON(iris.Map{
			"status":   "ERROR",
			"message1": err.Error,
		})

	} else {

		ctx.JSON(iris.Map{
			"status":    "OK",
			"file-hash": hash.Sum([]byte("Hello World!!!")),
			"blake-256": blake2b.Sum256([]byte("Hello World!!!")),
			"blake-384": blake2b.Sum384([]byte("Hello World!!!")),
			"blake-512": blake2b.Sum512([]byte("Hello World!!!")),
		})
	}
}
