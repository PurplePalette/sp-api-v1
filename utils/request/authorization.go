package request

import (
	"context"
	"errors"
)

// ContextKey is type for read/write context
type ContextKey int

// CtxUserId is context key for getting id
const CtxUserId ContextKey = 1

// GetUserId reads user id string from request context
func GetUserId(ctx context.Context) (string, error) {
	v := ctx.Value(CtxUserId)
	userId, ok := v.(string)
	if !ok {
		return "", errors.New("could not parse userId header")
	}
	return userId, nil
}
