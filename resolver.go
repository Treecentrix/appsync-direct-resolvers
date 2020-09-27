package resolvers

import (
	"encoding/json"
	"reflect"
)

type resolver struct {
	function interface{}
}

func (r *resolver) hasArguments() bool {
	return reflect.TypeOf(r.function).NumIn() >= 1
}

func (r *resolver) hasHeaders() bool {
	return reflect.TypeOf(r.function).NumIn() >= 2
}

func (r *resolver) hasIdentity() bool {
	return reflect.TypeOf(r.function).NumIn() >= 3
}

func (r *resolver) call(p json.RawMessage, h json.RawMessage, i *Identity) (interface{}, error) {
	var args []reflect.Value

	if r.hasArguments() {
		pld := payload{p}
		arguments, err := pld.parse(reflect.TypeOf(r.function).In(0))

		if err != nil {
			return nil, err
		}
		args = append(args, *arguments)
	}

	if r.hasHeaders() {
		pld := payload{h}
		headers, err := pld.parse(reflect.TypeOf(r.function).In(1))

		if err != nil {
			return nil, err
		}
		args = append(args, *headers)
	}

	if r.hasIdentity() && i != nil {
		args = append(args, reflect.ValueOf(i))
	}

	returnValues := reflect.ValueOf(r.function).Call(args)
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
