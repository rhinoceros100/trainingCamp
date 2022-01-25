package main

import (
	"reflect"
	"testing"
)

func TestGoimDecoder(t *testing.T) {
	msg1 := "test"
	header1 := &Header{
		PackageLen: int32(16 + len(msg1)),
		HeaderLen:  16,
		Ver:        2,
		Operation:  5,
		Seq:        333,
	}
	header2 := &Header{
		PackageLen: int32(16 + len(msg1) + 100),
		HeaderLen:  16,
		Ver:        2,
		Operation:  5,
		Seq:        333,
	}
	type args struct {
		ReceivedMsg *[]byte
	}
	tests := []struct {
		name       string
		args       args
		wantHeader *Header
		wantBody   string
		wantErr    bool
	}{
		{
			name: "common",
			args: args{
				ReceivedMsg: genPackage(header1, msg1),
			},
			wantHeader: header1,
			wantBody:   msg1,
			wantErr:    false,
		},
		{
			name: "wrong package len",
			args: args{
				ReceivedMsg: genPackage(header2, msg1),
			},
			wantHeader: nil,
			wantBody:   "",
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHeader, gotBody, err := GoimDecoder(tt.args.ReceivedMsg)
			if (err != nil) != tt.wantErr {
				t.Errorf("GoimDecoder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotHeader, tt.wantHeader) {
				t.Errorf("GoimDecoder() gotHeader = %v, want %v", gotHeader, tt.wantHeader)
			}
			if !reflect.DeepEqual(gotBody, tt.wantBody) {
				t.Errorf("GoimDecoder() gotBody = %v, want %v", gotBody, tt.wantBody)
			}
		})
	}
}
