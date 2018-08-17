package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"go-swagger-sample/restapi/operations/todos"
	"go-swagger-sample/models"
	"go-swagger-sample/common"
	"go-swagger-sample/restapi/operations"
)

func init() {
	addHandler(func(api *operations.ATodoListApplicationAPI) {
		db, _ := common.InitializeDb()

		api.TodosFindTodosHandler = todos.FindTodosHandlerFunc(func(params todos.FindTodosParams) middleware.Responder {
			var item []*models.Item
			db.Find(&item)

			return todos.NewFindTodosOK().WithPayload(item)
		})
	})
}
