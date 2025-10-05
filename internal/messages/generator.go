package messages

import "math/rand"



func GeneraterandomMessage() ([]byte, string, error) {
	if rand.Intn(2) == 0 {
		data, err := GenerateJSONMesssage()
		return data, "application/json", err

	}
	data, err := GenerateXMLMessage()
	return data, "applicatio/xml", err
}