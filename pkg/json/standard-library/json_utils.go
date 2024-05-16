package iteratorlibrary

import (
	jsoniter "github.com/json-iterator/go"

	jsonutils "project-v/pkg/json"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

type jsoniteratorJsonUtils struct{}

// Marshal This function converts a given struct to JSON string
func (jju jsoniteratorJsonUtils) Marshal(req interface{}) ([]byte, error) {
	return json.Marshal(req)
}

// Unmarshal This function converts a given JSON string to struct
func (jju jsoniteratorJsonUtils) Unmarshal(data []byte, req interface{}) error {
	return json.Unmarshal(data, req)
}

// NewJsoniterJsonUtils This function returns the instance of MarshalUtils
func NewJsoniterJsonUtils() jsonutils.IJson {
	return jsoniteratorJsonUtils{}
}
