package services

import (
	"context"
	"github.com/imulab/go-scim/core/prop"
	"github.com/imulab/go-scim/protocol/crud"
	"github.com/imulab/go-scim/protocol/db"
	"github.com/imulab/go-scim/protocol/log"
)

type (
	GetRequest struct {
		*crud.Projection
		ResourceID         string
	}
	GetResponse struct {
		Resource *prop.Resource
		Location string
		Version  string
	}
	GetService struct {
		Logger   log.Logger
		Database db.DB
	}
)

func (s *GetService) GetResource(ctx context.Context, request *GetRequest) (*GetResponse, error) {
	resource, err := s.Database.Get(ctx, request.ResourceID, request.Projection)
	if err != nil {
		s.Logger.Error("failed to get resource from persistence", log.Args{
			"resourceId": request.ResourceID,
			"error": err,
		})
		return nil, err
	}

	return &GetResponse{
		Resource: resource,
		Location: resource.Location(),
		Version:  resource.Version(),
	}, nil
}