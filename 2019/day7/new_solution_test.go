package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestDay7Task1(t *testing.T) {
	type args struct {
		memory        string
		phaseSequence sort.IntSlice
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []int
	}{
		{
			name:  "Example 1",
			args:  args{memory: "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0", phaseSequence: []int{0, 1, 2, 3, 4}},
			want:  43210,
			want1: []int{4, 3, 2, 1, 0},
		},
		{
			name:  "Example 2",
			args:  args{memory: "3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0", phaseSequence: []int{0, 1, 2, 3, 4}},
			want:  54321,
			want1: []int{0, 1, 2, 3, 4},
		},
		{
			name:  "Example 3",
			args:  args{memory: "3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0", phaseSequence: []int{0, 1, 2, 3, 4}},
			want:  65210,
			want1: []int{1, 0, 4, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Day7Task1as(tt.args.memory, tt.args.phaseSequence)
			if got != tt.want {
				t.Errorf("Day7Task1() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Day7Task1() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
