package repositories

import (
	"lambda-store-nir/service/application/domain"
)

type DocumentRepository interface {
	Save(document domain.Document) error
}
