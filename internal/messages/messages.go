package messages

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)






type Message struct {
    ID        string `json:"id" xml:"id"`
    Format    string `json:"format" xml:"format"` // json, xml, avro, etc.
    Event     string `json:"event" xml:"event"`
    UserID    int    `json:"user_id" xml:"user_id"`
    Payload   string `json:"payload" xml:"payload"`
    Timestamp int64  `json:"timestamp" xml:"timestamp"`
}



func GenerateRandomMessage() ([]byte, string, error) {
    formats := []string{"json", "xml"}
    format := formats[rand.Intn(len(formats))]

    msg := Message{
        ID:        uuid.New().String(),
        Format:    format,
        Event:     "UserSignedIn",
        UserID:    rand.Intn(10000),
        Payload:   fmt.Sprintf("Payload-%d", rand.Intn(999999)),
        Timestamp: time.Now().UnixMilli(),
    }

    switch format {
    case "json":
        data, err := json.Marshal(msg)
        return data, format, err
    case "xml":
        data, err := xml.Marshal(msg)
        return data, format, err
    default:
        return nil, "", fmt.Errorf("unknown format: %s", format)
    }
}
