package main

import (
	SnappyPoolReader "TinyRDMCustomEncoderSnappy/SnappyPool/Reader"
	SnappyPoolWriter "TinyRDMCustomEncoderSnappy/SnappyPool/Writer"
	"bytes"
	"encoding/base64"
	"io"
	"os"

	SingBuffer "github.com/sagernet/sing/common/buf"
	"github.com/valyala/bytebufferpool"
)

func main() {
	if len(os.Args) == 3 {
		switch os.Args[1] {
		case "Encode", "Decode":
			if Bytes, Error := base64.StdEncoding.DecodeString(os.Args[2]); Error == nil {
				Buffer := bytebufferpool.Get()
				defer bytebufferpool.Put(Buffer)

				if os.Args[1] == "Encode" {
					Writer := SnappyPoolWriter.Get(Buffer)
					defer SnappyPoolWriter.Put(Writer)

					if _, Error := Writer.Write(Bytes); Error == nil {
						if Error := Writer.Flush(); Error != nil {
							os.Stdout.WriteString("[RDM-ERROR]")
							os.Exit(0)
						}
					} else {
						os.Stdout.WriteString("[RDM-ERROR]")
						os.Exit(0)
					}
				} else {
					DecodeBytes := SingBuffer.Get(64)
					defer SingBuffer.Put(DecodeBytes)

					Reader := SnappyPoolReader.Get(bytes.NewReader(Bytes))
					defer SnappyPoolReader.Put(Reader)

					for {
						if Length, Error := Reader.Read(DecodeBytes); Error == nil {
							if Length > 0 {
								if _, Error := Buffer.Write(DecodeBytes[:Length]); Error != nil {
									os.Stdout.WriteString("[RDM-ERROR]")
									os.Exit(0)
								}
							} else {
								break
							}
						} else {
							if Error != io.EOF {
								os.Stdout.WriteString("[RDM-ERROR]")
								os.Exit(0)
							}

							break
						}
					}
				}

				os.Stdout.WriteString(base64.StdEncoding.EncodeToString(Buffer.Bytes()))
			} else {
				os.Stdout.WriteString("[RDM-ERROR]")
			}
		default:
			os.Stdout.WriteString("[RDM-ERROR]")
		}
	} else {
		os.Stdout.WriteString("[RDM-ERROR]")
	}
}
