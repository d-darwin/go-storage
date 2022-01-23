package file

import "github.com/google/uuid"

type File struct {
	ID   uuid.UUID
	Name string
	Data []byte
}

func NewFile(filename string, blob []byte) File {
	fileId, err := uuid.NewUUID()

	return File{
		ID: ,
		Name: filename,
		Data: blob,
	}
}
