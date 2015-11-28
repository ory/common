package context

import (
	"github.com/go-errors/errors"
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

func FetchRouterParamsFromContext(ctx context.Context, keys ...string) (map[string]string, error) {
	res := make(map[string]string)
	var r string
	ps := GetRouterParamsFromContext(ctx)
	for _, key := range keys {
		r = ps.ByName(key)
		if len(r) == 0 {
			return map[string]string{}, errors.New(`Router param "` + key + `" not given.`)
		}
		res[key] = r
	}
	return res, nil
}
