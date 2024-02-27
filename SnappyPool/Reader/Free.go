package Reader

import "github.com/golang/snappy"

func Put(Reader *snappy.Reader) {
	Reader.Reset(nil)
	Pool.Put(Reader)
}
