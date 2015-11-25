package context

import (
	"github.com/julienschmidt/httprouter"
	"github.com/ory-am/common/handler"
	"golang.org/x/net/context"
)

var RouterParamKey handler.Key = 0

func NewContextFromRouterParams(ctx context.Context, ps httprouter.Params) context.Context {
	return context.WithValue(ctx, RouterParamKey, ps)
}

func GetRouterParamsFromContext(ctx context.Context) (ps httprouter.Params) {
	params, ok := ctx.Value(RouterParamKey).(httprouter.Params)
	if !ok {
		return httprouter.Params{}
	}
	return params
}
