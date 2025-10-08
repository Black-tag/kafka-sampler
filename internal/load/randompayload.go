package load

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"math/rand"
)








func RandomPayload(format string, minSize, MaxSize int) []byte {
	size := rand.Intn(MaxSize-minSize+1) + minSize
	data := map[string]interface{}{
		"id": rand.Intn(100000),
		"body": string(bytes.Repeat([]byte("x"), size-32)),
	}
	switch format {
	case "json":
		b, _ := json.Marshal(data)
		return b
	case "xml":
		b, _ := xml.Marshal(data)
		return b
	case "protobuf":
		return bytes.Repeat([]byte{0xA5}, size)
	case "avro":
		return bytes.Repeat([]byte{0xB2}, size)
	default:
		return bytes.Repeat([]byte("x"), size)
	}
}