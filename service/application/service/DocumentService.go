package service

import (
	"lambda-store-nir/service/application/domain"
	"lambda-store-nir/service/application/exception"
	"lambda-store-nir/service/application/logger"
	"lambda-store-nir/service/application/ports"
	"lambda-store-nir/service/application/repositories"
	"lambda-store-nir/service/application/usecases"
)

type DocumentService struct {
	Logger             logger.Logger
	DocumentEvent      ports.DocumentEvent
	DocumentRepository repositories.DocumentRepository
}

func NewDocumentService(logger logger.Logger, documentEvent ports.DocumentEvent, documentRepository repositories.DocumentRepository) usecases.DocumentUc {

	return DocumentService{
		Logger:             logger,
		DocumentEvent:      documentEvent,
		DocumentRepository: documentRepository,
	}

}

func (s DocumentService) Create(id string, title string, body string) error {

	if id == "" {
		return exception.ThrowValidationError("Invalid id from document")
	}

	document := domain.Document{
		Id:    id,
		Title: title,
		Body:  body,
	}

	err := s.DocumentRepository.Save(document)

	if err != nil {
		return err
	}

	return s.DocumentEvent.Created(document)

}
