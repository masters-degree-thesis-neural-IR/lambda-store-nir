package service

import (
	"lambda-store-nir/service/application/domain"
	"lambda-store-nir/service/application/exception"
	"lambda-store-nir/service/application/ports"
	"lambda-store-nir/service/application/repositories"
	"lambda-store-nir/service/application/usecases"
)

type DocumentService struct {
	DocumentEvent      ports.DocumentEvent
	DocumentRepository repositories.DocumentRepository
}

func NewDocumentService(documentEvent ports.DocumentEvent, documentRepository repositories.DocumentRepository) usecases.CreateDocumentUc {
	var c usecases.CreateDocumentUc = &DocumentService{
		DocumentEvent:      documentEvent,
		DocumentRepository: documentRepository,
	}
	return c
}

func (s *DocumentService) CreateDocument(id string, title string, body string) error {

	if id == "" {
		return *exception.ThrowValidationError("Invalid id from document")
	}

	if title == "" {
		return *exception.ThrowValidationError("Invalid title from document")
	}

	if body == "" {
		return *exception.ThrowValidationError("Invalid body from document")
	}

	document := domain.Document{
		Id:    id,
		Title: title,
		Body:  body,
	}

	var err = s.DocumentRepository.Save(document)

	if err != nil {
		return err
	}

	return s.DocumentEvent.Created(document)

}
