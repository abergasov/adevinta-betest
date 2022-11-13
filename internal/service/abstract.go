package service

import (
	"io"
)

type Decoder interface {
	Decode(reader io.Reader) (htmlBody string, err error)
}
