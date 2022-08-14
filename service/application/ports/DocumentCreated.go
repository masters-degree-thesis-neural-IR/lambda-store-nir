package ports

import "lambda-store-nir/service/application/domain"

type DocumentEvent interface {
	Created(document domain.Document) error
}
