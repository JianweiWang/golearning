package datafile

import (
	"reflect"
	"testing"
)

func TestGetString(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"test1", args{"/tmp/hello.txt"}, []string{"hello"}, false},
		{"test2", args{"/tmp/notexistfile.txt"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetString(tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}
