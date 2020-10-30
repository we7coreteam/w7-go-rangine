package http

import (
	"github.com/gin-gonic/gin"
)

func NewRequest(context *gin.Context) *Request {
	return &Request{
		ginContext: context,
	}
}

type Request struct {
	ginContext *gin.Context
}

type File struct {
	OriginalName string
	MimeType     string
}

func (mySelf *Request) File(formName string) (*File, error) {
	file, error := mySelf.ginContext.FormFile(formName)
	if error != nil {
		return nil, error
	}

	return &File{
		OriginalName: file.Filename,
		MimeType:     file.Header["Content-Type"][0],
	}, nil
}
