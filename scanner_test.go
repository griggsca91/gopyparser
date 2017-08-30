package gopyparser

import (
	"reflect"
	"testing"
)

func TestNewFromFile(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    *Scanner
		wantErr bool
	}{
		{"Check for valid file", args{"./pytestfiles/scanner1_test.py"}, New(`def main():
	print("hello world")
	return

main()`), false},
		{"Gets Invalid file", args{"./pytestfiles/scanner1_nonexistentfile.py"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFromFile(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
