package handlers

import (
	"go-swagger-sample/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"go-swagger-sample/common"
	"go-swagger-sample/models"
	"go-swagger-sample/restapi/operations/consumers"
	"fmt"
)

func init() {
	addHandler(func(api *operations.ATodoListApplicationAPI) {
			db, _ := common.InitializeDb()

		api.ConsumersFindConsumerHandler = consumers.FindConsumerHandlerFunc(func(params consumers.FindConsumerParams) middleware.Responder {
			var consumer models.Consumer

			db.Where("id = ?", 1).First(&consumer)

			fmt.Println(consumer)

			return consumers.NewFindConsumerOK().WithPayload(&consumer)
		})
	})
}
