package service

import (
	"dev.ssouza/rest-api/domain"
	"dev.ssouza/rest-api/persistence"
)

func CreatEndpoint() domain.Endpoint {
	return persistence.CreateEndpoint(domain.Endpoint{})
}