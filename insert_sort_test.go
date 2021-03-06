package algorithm

import (
	"reflect"
	"testing"
)

func TestInsertSort(t *testing.T) {
type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"empty", args{list: []int{}}, []int{}},
		{"value", args{list: []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}}, []int{4, 12, 27, 37, 80, 81, 84, 85, 93, 93}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertSort(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertSort() = %v, want %v", got, tt.want)
			}
		})
	}
}