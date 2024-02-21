package auth

import (
	"context"
	"fmt"

	"github.com/gmtstephane/go-template/models"
)

type userCtx struct{}

func User(c context.Context) models.User {
	u, ok := c.Value(userCtx{}).(models.User)
	fmt.Println(u, ok)
	if !ok {
		return models.User{}
	}
	return u
}

func Context(ctx context.Context, u models.User) context.Context {
	return context.WithValue(ctx, userCtx{}, u)
}
