package paginator

import (
	"encoding/base64"
	"encoding/json"
	"reflect"
)

type CursorEncoder interface {
	Encode(v interface{}) string
}

func NewCursorEncoder(keys ...string) CursorEncoder {
	return &cursorEncoder{keys}
}

type cursorEncoder struct {
	keys []string
}

func (e *cursorEncoder) Encode(v interface{}) string {
	return base64.StdEncoding.EncodeToString(e.marshalJSON(v))
}

func (e *cursorEncoder) marshalJSON(value interface{}) []byte {
	rv := toReflectValue(value)
	for rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	fields := make([]interface{}, len(e.keys))
	for i, key := range e.keys {
		fields[i] = rv.FieldByName(key).Interface()
	}
	b, err := json.Marshal(fields)
	if err != nil {
		return nil
	}
	return b
}
