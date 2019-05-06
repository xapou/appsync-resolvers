package resolvers

import (
	"fmt"
	"reflect"
)

type resolver struct {
	function interface{}
}

func (r *resolver) countArguments() int {
	return reflect.TypeOf(r.function).NumIn()
}

func (r *resolver) call(c context) (interface{}, error) {
	var parameters []reflect.Value
	var returnValues []reflect.Value
	var arguments, source, identity reflect.Value

	// We have 3 possible signatures for the resolvers
	// A single parameter:
	// func handler(args *ArgsType) (*ReturnType, error)
	// 2 parameters:
	// func handler(args *ArgsType, identity *Identity) (*ReturnType, error)
	// 3 parameters:
	// func handler(args *ArgsType, source *SourceType, identity *Identity) (*ReturnType, error)

	apld := payload{c.Arguments}
	spld := payload{c.Source}
	ipld := payload{c.Identity}

	arguments, err := apld.parse(reflect.TypeOf(r.function).In(0))
	if err != nil {
		return nil, fmt.Errorf("Failed to parse Arguments. %v", err.Error())
	}
	parameters = append(parameters, arguments)

	if r.countArguments() == 2 {
		identity, err = ipld.parse(reflect.TypeOf(r.function).In(1))
		if err != nil {
			return nil, fmt.Errorf("Failed to parse Identity. %v", err.Error())
		}
		parameters = append(parameters, identity)
	} else if r.countArguments() == 3 {
		source, err = spld.parse(reflect.TypeOf(r.function).In(1))
		if err != nil {
			return nil, fmt.Errorf("Failed to parse Source. %v", err.Error())
		}
		parameters = append(parameters, source)
		identity, err = ipld.parse(reflect.TypeOf(r.function).In(2))
		if err != nil {
			return nil, fmt.Errorf("Failed to parse Identity. %v", err.Error())
		}
		parameters = append(parameters, identity)
	}

	// Call the handler and retrieve the response
	returnValues = reflect.ValueOf(r.function).Call(parameters)

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
