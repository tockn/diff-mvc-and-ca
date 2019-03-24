package api

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

type ContextKey string

const (
	ginKey = "gin.context.key"
)

var ginContextMap map[string]*gin.Context

func init() {
	ginContextMap = map[string]*gin.Context{}
}

func addGinContext(ctx context.Context, c *gin.Context) context.Context {
	key := generateNewKey()

	ctx = setResKey(ctx, key)

	ginContextMap[key] = c

	return ctx
}

func getGinContext(ctx context.Context) *gin.Context {
	key := getResKey(ctx)
	res, ok := ginContextMap[key]
	if !ok {
		_, cancel := context.WithCancel(ctx)
		cancel()
	}
	return res
}

func setResKey(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, ContextKey(ginKey), value)
}

func deleteGinContext(ctx context.Context) {
	key := getResKey(ctx)
	if _, ok := ginContextMap[key]; ok {
		delete(ginContextMap, key)
	}
}

func getResKey(ctx context.Context) string {
	return getKey(ctx, ContextKey(ginKey))
}

func getKey(ctx context.Context, ctxKey ContextKey) string {
	key, _ := ctx.Value(ctxKey).(string)
	return key
}

func generateNewKey() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%d", rand.Int())
}
