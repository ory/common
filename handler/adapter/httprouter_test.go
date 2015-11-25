package adapter

import (
	"github.com/julienschmidt/httprouter"
	cc "github.com/ory-am/common/context"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"net/http"
	"testing"
)

func TestNewHttpRouterAdapter(t *testing.T) {
	called := false
	params := httprouter.Params{{"foo", "bar"}}
	NewHttpRouterAdapter(
		context.Background(),
	).ThenFunc(func(ctx context.Context, r http.ResponseWriter, w *http.Request) {
		called = true
		assert.Equal(t, params, cc.GetRouterParamsFromContext(ctx))
	})(nil, nil, params)
	assert.True(t, called)
}
