package paginator

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"reflect"
	"time"
)

type CursorDecoder interface {
	Decode(cursor string) []interface{}
}

func NewCursorDecoder(ref interface{}, keys ...string) (CursorDecoder, error) {
	decoder := &cursorDecoder{keys: keys}
	err := decoder.initKeyKinds(ref)
	if err != nil {
		return nil, err
	}
	return decoder, nil
}

var (
	ErrInvalidDecodeReference = errors.New("decode reference should be struct")
	ErrInvalidField           = errors.New("invalid field")
	ErrInvalidFieldType       = errors.New("invalid field type")
)

type kind uint

const (
	kindInvalid kind = iota
	kindBool
	kindInt
	kindUint
	kindFloat
	kindString
	kindTime
)

func toKind(rt reflect.Type) kind {
	for rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	if rt.ConvertibleTo(reflect.TypeOf(time.Time{})) {
		return kindTime
	}
	switch rt.Kind() {
	case reflect.Bool:
		return kindBool
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return kindInt
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return kindUint
	case reflect.Float32, reflect.Float64:
		return kindFloat
	default:
		return kindString
	}
}

type cursorDecoder struct {
	keys     []string
	keyKinds []kind
}

func (d *cursorDecoder) Decode(cursor string) []interface{} {
	b, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return nil
	}
	var fields []interface{}
	err = json.Unmarshal(b, &fields)
	if err != nil {
		return nil
	}
	return d.castJSONFields(fields)
}

func (d *cursorDecoder) initKeyKinds(ref interface{}) error {
	//	@TODO:	zero value error
	rt := toReflectValue(ref).Type()
	// reduce reflect type to underlying struct
	for rt.Kind() == reflect.Slice || rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	if rt.Kind() != reflect.Struct {
		// element of out must be struct, if not, just pass it to gorm to handle the error
		return ErrInvalidDecodeReference
	}
	d.keyKinds = make([]kind, len(d.keys))
	for i, key := range d.keys {
		field, ok := rt.FieldByName(key)
		if !ok {
			return ErrInvalidField
		}
		d.keyKinds[i] = toKind(field.Type)
	}
	return nil
}

func (d *cursorDecoder) castJSONFields(fields []interface{}) []interface{} {
	var result []interface{}
	for i, field := range fields {
		kind := d.keyKinds[i]
		switch f := field.(type) {
		case bool:
			bv, err := castJSONBool(f, kind)
			if err != nil {
				return nil
			}
			result = append(result, bv)
		case float64:
			fv, err := castJSONFloat(f, kind)
			if err != nil {
				return nil
			}
			result = append(result, fv)
		case string:
			sv, err := castJSONString(f, kind)
			if err != nil {
				return nil
			}
			result = append(result, sv)
		default:
			return nil
		}
	}
	return result
}

func castJSONBool(value bool, kind kind) (interface{}, error) {
	if kind != kindBool {
		return nil, ErrInvalidFieldType
	}
	return value, nil
}

func castJSONFloat(value float64, kind kind) (interface{}, error) {
	switch kind {
	case kindInt:
		return int(value), nil
	case kindUint:
		return uint(value), nil
	case kindFloat:
		return value, nil
	}
	return nil, ErrInvalidFieldType
}

func castJSONString(value string, kind kind) (interface{}, error) {
	if kind != kindString && kind != kindTime {
		return nil, ErrInvalidFieldType
	}
	if kind == kindString {
		return value, nil
	}
	tv, err := time.Parse(time.RFC3339Nano, value)
	if err != nil {
		return nil, ErrInvalidFieldType
	}
	return tv, nil
}
