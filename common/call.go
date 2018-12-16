package common

import (
	"errors"
	"reflect"
	utils "github.com/hmpz2004/RebateBot/utils"
)

const TAG_CALL = "common/call.go"

func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	
	utils.LogDebug(TAG_CALL, "in call.go Call()")
	utils.LogDebug(TAG_CALL, "name " + name)

	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}

	result = f.Call(in)
	return
}
