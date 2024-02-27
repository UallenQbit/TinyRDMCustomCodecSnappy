package Writer

import "github.com/golang/snappy"

func Put(Writer *snappy.Writer) {
	Writer.Reset(nil)
	Pool.Put(Writer)
}
