package funcutils

import (
	"fmt"

	"gopkg.hlmpn.dev/pkg/go-logger"
	errutil "gopkg.hlmpn.dev/pkg/go-logger/errors"
)

func Recover() {
	if r := recover(); r != nil {
		logger.LogErrorf(fmt.Sprintf("Panic: %s", r))
		errutil.New(fmt.Sprintf("Panic: %s", r)).Unwrap()
	}
}

func TryCatch(fn func(), catch func(interface{})) error {
	defer func() {
		if r := recover(); r != nil {
			catch(r)
			errutil.New(fmt.Sprintf("Panic: %s", r)).Unwrap()
		}
	}()
	fn()
	return nil
}
