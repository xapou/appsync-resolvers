package resolvers

import (
	"encoding/json"
	"reflect"
)

type resolver struct {
	function interface{}
}

func (r *resolver) countArguments() int {
	return reflect.TypeOf(r.function).NumIn()
}

func (r *resolver) call(p json.RawMessage, identity Identity) (interface{}, error) {
	var args []reflect.Value
	var err error
	var returnValues []reflect.Value

	if r.countArguments() > 0 {
		pld := payload{p}
		args, err = pld.parse(reflect.TypeOf(r.function).In(0))

		if err != nil {
			return nil, err
		}
	}

	if r.countArguments() > 1 {
		var identityValue = reflect.ValueOf(identity)
		args = append(args, identityValue)
	}

	returnValues = reflect.ValueOf(r.function).Call(args)

	var returnData interface{}
	var returnError error

	if len(returnValues) == 2 {
		returnData = returnValues[0].Interface()
	}

	if err := returnValues[len(returnValues)-1].Interface(); err != nil {
		returnError = returnValues[len(returnValues)-1].Interface().(error)
	}

	return returnData, returnError
}
