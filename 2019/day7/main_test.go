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
			got, got1 := Day7Task1(tt.args.memory, tt.args.phaseSequence)
			if got != tt.want {
				t.Errorf("Day7Task1() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Day7Task1() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDay7Task2(t *testing.T) {
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
			name: "Example 1", args: args{memory: "3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5", phaseSequence: []int{5, 6, 7, 8, 9}},
			want:  139629729,
			want1: []int{9, 8, 7, 6, 5},
		},
		{
			name: "Example 2", args: args{memory: "3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10", phaseSequence: []int{5, 6, 7, 8, 9}},
			want:  18216,
			want1: []int{9, 7, 8, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Day7Task2(tt.args.memory, tt.args.phaseSequence)
			if got != tt.want {
				t.Errorf("Day7Task2() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Day7Task2() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
