package handlers

import (
	"doApp/services"

	"cloud.google.com/go/storage"
)

type Handlers struct {
	Services      services.IServices
	StorageClient *storage.Client
}

func NewHandlers(services services.IServices) *Handlers {
	return &Handlers{
		Services: services,
	}
}
