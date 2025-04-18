package controllers

import (
	"github.com/cloudimpl/next-coder-sdk/api"
	"github.com/cloudimpl/next-coder-sdk/polycode"
	"portal/register/model"
)

func greeting(ctx polycode.WorkflowContext, input model.HelloRequest) (model.HelloResponse, error) {
	greetingService := ctx.Service("greeting-service").Get()

	var output model.HelloResponse
	res := greetingService.RequestReply(polycode.TaskOptions{}, "greeting", input)
	if err := res.Get(&output); err != nil {
		return model.HelloResponse{}, err
	}

	return output, nil
}

var Greeting = api.FromWorkflow(greeting)
