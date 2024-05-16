package iteratorlibrary

import (
	"reflect"
	"testing"

	jsonutils "project-v/pkg/json"
)

func TestNewJsoniterJsonUtils(t *testing.T) {
	tests := []struct {
		name string
		want jsonutils.IJson
	}{
		{
			name: "OK",
			want: NewJsoniterJsonUtils(),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := NewJsoniterJsonUtils(); !reflect.DeepEqual(
					got, tt.want,
				) {
					t.Errorf("NewJsonUtils() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestJsonUtilsMarshal(t *testing.T) {
	type args struct {
		req interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				req: struct {
					Name string `json:"name"`
				}{
					Name: "test",
				},
			},
			want:    []byte(`{"name":"test"}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				j := jsoniteratorJsonUtils{}
				got, err := j.Marshal(tt.args.req)
				if (err != nil) != tt.wantErr {
					t.Errorf(
						"Marshal() error = %v, wantErr %v", err, tt.wantErr,
					)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Marshal() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestJsonUtilsUnmarshal(t *testing.T) {
	type args struct {
		data []byte
		req  interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				data: []byte(`{"name": "test"}`),
				req: &struct {
					Name string `json:"name"`
				}{
					Name: "test",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				j := jsoniteratorJsonUtils{}
				if err := j.Unmarshal(
					tt.args.data, tt.args.req,
				); (err != nil) != tt.wantErr {
					t.Errorf(
						"Unmarshal() error = %v, wantErr %v", err, tt.wantErr,
					)
				}
			},
		)
	}
}
