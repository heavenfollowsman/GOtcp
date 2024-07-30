package message

import (
	"io"
)

type Encoder interface {
	Encode([]byte, io.Writer) error
}
