package controller

import (
	"elastic-transfer/app/manage"
	pb "elastic-transfer/router"
)

type controller struct {
	pb.UnimplementedRouterServer
	manager *manage.ElasticManager
}

func New(manager *manage.ElasticManager) *controller {
	c := new(controller)
	c.manager = manager
	return c
}

func (c *controller) find(identity string) (information *pb.Information, err error) {
	option, err := c.manager.GetOption(identity)
	if err != nil {
		return
	}
	information = &pb.Information{
		Identity: option.Identity,
		Index:    option.Index,
		Validate: option.Validate,
		Topic:    option.Topic,
		Key:      option.Key,
	}
	return
}

func (c *controller) response(err error) (*pb.Response, error) {
	if err != nil {
		return &pb.Response{
			Error: 1,
			Msg:   err.Error(),
		}, nil
	} else {
		return &pb.Response{
			Error: 0,
			Msg:   "ok",
		}, nil
	}
}
