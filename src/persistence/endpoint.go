package persistence

import(
	"dev.ssouza/rest-api/domain"
)

func CreateEndpoint(endpoint domain.Endpoint) domain.Endpoint {
	newEndpoint := domain.Endpoint{

	}

	return newEndpoint
}

func ReadEndpoint(id int) domain.Endpoint {
	newEndpoint := domain.Endpoint{
		Id: id,
	}

	return newEndpoint
}

func UpdateEndpoint(endpoint domain.Endpoint) domain.Endpoint {
	newEndpoint := domain.Endpoint{
		Key: "update",
	}

	return newEndpoint
}

func DeleteEndpoint(id int) {
	
}
