package middleware

import (
	"fmt"

	"github.com/emvi/hide"

	"github.com/kataras/iris/context"
)

// HideID middleware to decode IDs
func HideID(ctx context.Context) {

	param := ctx.Params().Store

	for i := 0; i < param.Len(); i++ {

		entry, valid := param.GetEntryAt(i)

		if valid {

			id, err := hide.FromString(fmt.Sprint(entry))

			if err == nil {

				fmt.Println(entry, entry.Key, id)

				param.Set(entry.Key, id)
			}
		}
	}

	ctx.Next()
}
