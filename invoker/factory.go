package invoker

import (
	"reflect"
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/utils"
)

func NewInvokerFactory() types.InvokerFactory {
	return &CoreInvokerFactory{}
}

func (fi *CoreInvokerFactory) Instance(method interface{}, build types.Builder) (invoker types.Invoker, err error) {
	defer utils.Recover(&err)

	invoker = NewInvoker(method, build)
	return
}

func NewNoArgsInvoker(method interface{}) types.Invoker{
	if src := utils.ValueOf(method); reflect.Func == src.Kind() {
		return NoParamInvoker(src)
	}
	panic(types.NewWithError(types.ErrTypeNotFunction, method))
}

func NewSimpleInvoker(method interface{}) types.Invoker{
	if src := utils.ValueOf(method); reflect.Func == src.Kind() {
		return SimpleInvoker(src)
	}
	panic(types.NewWithError(types.ErrTypeNotFunction, method))
}
