package gopyparser

import (
	"bufio"
	"reflect"
	"strings"
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
		// TODO: Add test cases.
		{"Check for valid file", args{"./pytestfiles/scanner1_test.py"}, New(strings.NewReader(`def main():
	print("hello world")
	return

main()`)), false},
		{"loading empty file", args{"./pytestfiles/scanner2_empty.py"}, New(strings.NewReader(``)), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFromFile(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			b := make([]byte, 100)
			b2 := make([]byte, 100)
			_, _ = got.r.Read(b)
			_, _ = tt.want.r.Read(b2)
			if string(b) != string(b2) {
				t.Errorf("NewFromFile() = %v, want %v", string(b), string(b2))
			}
		})
	}
}

func TestScanner_Read(t *testing.T) {
	type fields struct {
		r *bufio.Reader
	}
	tests := []struct {
		name   string
		fields fields
		want   rune
	}{
		// TODO: Add test cases.
		{"consume character", fields{r: bufio.NewReader(strings.NewReader("abcd"))}, rune('a')},
		{"empty string", fields{r: bufio.NewReader(strings.NewReader(""))}, EOF},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scanner{
				r: tt.fields.r,
			}
			if got := s.Read(); got != tt.want {
				t.Errorf("Scanner.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSpace(t *testing.T) {
	type args struct {
		c rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"Testing for space", args{rune(' ')}, true},
		{"Is not a space", args{rune('\t')}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSpace(tt.args.c); got != tt.want {
				t.Errorf("IsSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsTab(t *testing.T) {
	type args struct {
		c rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Testing for tab", args{rune('\t')}, true},
		{"Is not a tab", args{rune(' ')}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsTab(tt.args.c); got != tt.want {
				t.Errorf("IsTab() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_NextToken(t *testing.T) {
	tests := []struct {
		name string
		s    *Scanner
		want Token
	}{
		{"Gets correct token", NewFromString("="), NewToken(Equal, "=")},
		{"Gets correct token", NewFromString(")"), NewToken(RPar, ")")},
		{"Gets correct token", NewFromString("("), NewToken(LPar, "(")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.NextToken(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scanner.NextToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
