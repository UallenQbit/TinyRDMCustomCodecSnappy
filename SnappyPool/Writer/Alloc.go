package Writer

import (
	"io"
	"sync"

	"github.com/golang/snappy"
)

var Pool *sync.Pool = new(sync.Pool)

func Get(Writer io.Writer) *snappy.Writer {
	if Data := Pool.Get(); Data == nil {
		return snappy.NewBufferedWriter(Writer)
	} else {
		Writer := Data.(*snappy.Writer)
		Writer.Reset(Writer)
		return Writer
	}
}
