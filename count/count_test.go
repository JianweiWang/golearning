package count

import (
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	type args struct {
		names []string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]int
		wantErr bool
	}{
		{"test1", args{[]string{"hello", "world", "hello"}}, map[string]int{"hello": 2, "world": 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Count(tt.args.names)
			if (err != nil) != tt.wantErr {
				t.Errorf("Count() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}
