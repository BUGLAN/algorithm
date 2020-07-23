package algorithm

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"空字符", args{s: ""}, ""},
		{"非空字符", args{s: "abc "}, " cba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseString(tt.args.s); got != tt.want {
				t.Errorf("ReverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFullArrange(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"正常的测试", args{"abc"}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FullArrange(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FullArrange() = %v, want %v", got, tt.want)
			}
		})
	}
}