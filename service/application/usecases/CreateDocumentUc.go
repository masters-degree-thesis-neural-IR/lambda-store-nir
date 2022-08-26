package usecases

type CreateDocumentUc interface {
	CreateDocument(id string, title string, body string) error
}
