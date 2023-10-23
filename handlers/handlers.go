package handlers

import "doApp/services"

type Handlers struct {
	Services services.IServices
}

func NewHandlers(services services.IServices) *Handlers {
	return &Handlers{
		Services: services,
	}
}
