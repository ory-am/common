package httprouter

import (
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"testing"
)

func TestGetterAndSetter(t *testing.T) {
	params := httprouter.Params{{"foo", "bar"}}
	ctx := NewContextFromRouterParams(context.Background(), params)
	assert.NotNil(t, ctx)
	assert.Equal(t, params, GetRouterParamsFromContext(ctx))

	assert.Equal(t, httprouter.Params{}, GetRouterParamsFromContext(context.Background()))
}
