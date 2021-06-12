package request

import (
	"context"
	"errors"
)

// ContextKey is type for read/write context
type ContextKey int

// CtxUserID is context key for getting id
const CtxUserID ContextKey = 1

// GetUserID reads user id string from request context
func GetUserID(ctx context.Context) (string, error) {
	v := ctx.Value(CtxUserID)
	userID, ok := v.(string)
	if !ok {
		return "", errors.New("could not parse userID header")
	}
	return userID, nil
}
