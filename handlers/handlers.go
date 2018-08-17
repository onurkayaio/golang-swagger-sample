package handlers

import "go-swagger-sample/restapi/operations"

var initializers []handlerInitializer

type handlerInitializer func(api *operations.ATodoListApplicationAPI)

// function to get handlers from configure app api.
func InitializeHandlers(api *operations.ATodoListApplicationAPI) {
	for _, initializer := range initializers {
		initializer(api)
	}
}

// add handlers to array to set configure app api.
func addHandler(h handlerInitializer) {
	if initializers == nil {
		initializers = make([]handlerInitializer, 0)
	}

	initializers = append(initializers, h)
}
