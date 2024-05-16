package dateutils

import (
	"fmt"
	"testing"
	"time"
)

func TestGetStartTimestampOfDate(t *testing.T) {
	type args struct {
		dateStr  string
		location *time.Location
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				dateStr:  "02/05/2024",
				location: time.UTC,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetStartTimestampOfDate(tt.args.dateStr, tt.args.location)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStartTimestampOfDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			fmt.Println("==> got: ", got)
		})
	}
}

func TestGetEndTimestampOfDate(t *testing.T) {
	type args struct {
		dateStr  string
		location *time.Location
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				dateStr:  "08/05/2024",
				location: time.UTC,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetEndTimestampOfDate(tt.args.dateStr, tt.args.location)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEndTimestampOfDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			fmt.Println("==> got: ", got)
		})
	}
}

func TestParseDateFromTimestamp(t *testing.T) {
	type args struct {
		timestampMs int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Success",
			args: args{
				timestampMs: 1673740800000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseDateFromTimestamp(tt.args.timestampMs)
			fmt.Println("==> got: ", got)
		})
	}
}
