package json

type IJson interface {
	Marshal(interface{}) ([]byte, error)
	Unmarshal(data []byte, req interface{}) error
}
