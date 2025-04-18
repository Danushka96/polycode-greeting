package greeting_service

import (
	"github.com/cloudimpl/next-coder-sdk/polycode"
	"portal/register/model"
)

func Greeting(ctx polycode.ServiceContext, input model.HelloRequest) (model.HelloResponse, error) {
	return model.HelloResponse{
		Message: "Hello " + input.Name,
	}, nil
}
