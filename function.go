package funcutils

import (
	"reflect"
	"sync"
)

type Function struct {
	Func func(any ...any)
	Args Args
}

type Args struct {
	ArgA interface{}
	ArgB []interface{}
}

func (f *Function) Run() {
	f.Func(f.Args.ArgA, f.Args.ArgB)
}

func (f *Function) RunWith(a interface{}, b ...interface{}) {

}

func (f *Function) RunWhenChannel(c chan bool) {
	<-c
	f.Run()
}
func (f *Function) RunWhenChannelWith(c chan bool, a interface{}, b ...interface{}) {
	<-c
	f.RunWith(a, b...)
}

func (f *Function) GoRun() func() {
	return func() {
		go func() {
			f.Func(f.Args.ArgA, f.Args.ArgB)
		}()
	}
}

func (f *Function) GoRunWithWG(wg *sync.WaitGroup) func() {
	return func() {
		go func() {
			defer wg.Done()
			f.Func(f.Args.ArgA, f.Args.ArgB)
		}()
	}
}

// Simply a semantic alias for the CreateFunc function
func NewFunction(f func(any ...interface{}), a interface{}, b ...interface{}) *Function {
	return CreateFunc(f, a, b...)
}

// CreateFunc creates a Function struct with the given function and arguments
func CreateFunc(f func(any ...interface{}), a interface{}, b ...interface{}) *Function {
	// Ensure b is always a slice of interface{}
	var argB []interface{}
	var isSlice bool
	if b == nil {
		argB = []interface{}{}
	} else {
		v := reflect.ValueOf(b)
		if v.Kind() == reflect.Slice {
			isSlice = true
		}
		if !isSlice {
			argB = []interface{}{b}
		} else {
			argB = b
		}
	}

	fnStruct := Function{
		Func: f,
		Args: Args{
			ArgA: a,
			ArgB: argB,
		},
	}
	return &fnStruct
}
