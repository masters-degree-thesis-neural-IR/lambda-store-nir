package ports

import "lambda-store-nir/service/application/domain"

type Store interface {
	StoreDocument(document domain.Document) error
}
