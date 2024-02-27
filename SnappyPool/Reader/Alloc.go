package Reader

import (
	"io"
	"sync"

	"github.com/golang/snappy"
)

var Pool *sync.Pool = new(sync.Pool)

func Get(Reader io.Reader) *snappy.Reader {
	if Data := Pool.Get(); Data == nil {
		return snappy.NewReader(Reader)
	} else {
		Reader := Data.(*snappy.Reader)
		Reader.Reset(Reader)
		return Reader
	}
}
