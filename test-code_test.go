package mbiz

import (
	"testing"
)

func Test_case1(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{
			name:    "test1",
			args:    args{arr: []int{1, 2, 3}},
			wantRes: 6,
		},
		{
			name:    "test2",
			args:    args{arr: []int{4, 5, 6}},
			wantRes: 15,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := case1(tt.args.arr); gotRes != tt.wantRes {
				t.Errorf("case1() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_case2(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		{
			name:    "test1",
			args:    args{date: "Tue 16 Jul 2019"},
			wantRes: "2019-07-16",
		},
		{
			name:    "test2",
			args:    args{date: "Tue 17 Aug 2023"},
			wantRes: "2023-08-17",
		},
		{
			name:    "test3",
			args:    args{date: "Tue 32 Jul 2019"},
			wantRes: "invalid date format",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := case2(tt.args.date); gotRes != tt.wantRes {
				t.Errorf("case2() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_case3(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{
			name:    "test1",
			args:    args{num: 5},
			wantRes: 120,
		},
		{
			name:    "test2",
			args:    args{num: 10},
			wantRes: 3628800,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := case3(tt.args.num); gotRes != tt.wantRes {
				t.Errorf("case3() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_case4(t *testing.T) {
	type args struct {
		h int
		m int
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		{
			name:    "test1",
			args:    args{h: 5, m: 00},
			wantRes: "five o'clock",
		},
		{
			name:    "test2",
			args:    args{h: 5, m: 01},
			wantRes: "one minute past five",
		},
		{
			name:    "test3",
			args:    args{h: 5, m: 10},
			wantRes: "ten minutes past five",
		},
		{
			name:    "test4",
			args:    args{h: 5, m: 15},
			wantRes: "quarter past five",
		},
		{
			name:    "test5",
			args:    args{h: 5, m: 30},
			wantRes: "half past five",
		},
		{
			name:    "test6",
			args:    args{h: 5, m: 40},
			wantRes: "twenty minutes to six",
		},
		{
			name:    "test7",
			args:    args{h: 5, m: 45},
			wantRes: "quarter to six",
		},
		{
			name:    "test8",
			args:    args{h: 5, m: 47},
			wantRes: "thirteen minutes to six",
		},
		{
			name:    "test9",
			args:    args{h: 5, m: 28},
			wantRes: "twenty-eight minutes past five",
		},
		{
			name:    "test10",
			args:    args{h: 5, m: 60},
			wantRes: "six o'clock",
		},
		{
			name:    "test10",
			args:    args{h: 5, m: 61},
			wantRes: "invalid time format",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := case4(tt.args.h, tt.args.m); gotRes != tt.wantRes {
				t.Errorf("case4() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_case5(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{
			name: "test1",
			args: args{
				arr: []int{1, 2, 2, 3},
			},
			wantRes: 2,
		},
		{
			name: "test2",
			args: args{
				arr: []int{3, 3, 2, 1, 3},
			},
			wantRes: 2,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := case5(tt.args.arr); gotRes != tt.wantRes {
				t.Errorf("case5() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
