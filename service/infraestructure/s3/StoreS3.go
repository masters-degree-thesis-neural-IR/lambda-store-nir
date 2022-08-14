package s3

import (
	"lambda-store-nir/service/application/domain"
	"lambda-store-nir/service/application/ports"
)

type StoreS3 struct{}

func NewStoreS3() *ports.Store {
	var c ports.Store = &StoreS3{}
	return &c
}

func (receiver StoreS3) StoreDocument(document domain.Document) error {
	//TODO implement me
	panic("implement me")
}
