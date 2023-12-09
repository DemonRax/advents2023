package main

import (
	"github.com/google/go-cmp/cmp"
	"strconv"
	"testing"

	"advents2023/util"
)

func Test_code(t *testing.T) {
	for _, test := range []struct {
		in   string
		want int
	}{
		{
			in: `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......*755
...$......
.664.598..`,
			want: 467*35 + 592*755,
		},
	} {
		t.Run(test.in, func(t *testing.T) {
			if got := partNumbers(util.ReadString(test.in)); got != test.want {
				t.Errorf("got = %d, want %d", got, test.want)
			}
		})
	}
}

func Test_scanLine(t *testing.T) {
	for _, test := range []struct {
		in   string
		want [][]int
	}{
		{"467..114..", [][]int{{0, 1, 2}, {5, 6, 7}}},
		{"...*......", [][]int{}},
		{"..35..633.", [][]int{{2, 3}, {6, 7, 8}}},
		{".......755", [][]int{{7, 8, 9}}},
	} {
		t.Run(test.in, func(t *testing.T) {
			got := scanLine(test.in)
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("did not get expected result (-want/+got):\n%s", diff)
			}
		})
	}
}

func Test_convertNum(t *testing.T) {
	for _, test := range []struct {
		in     string
		coords []int
		want   int
	}{
		{"467..114..", []int{0, 1, 2}, 467},
		{"467..114..", []int{5, 6, 7}, 114},
	} {
		t.Run(test.in, func(t *testing.T) {
			got := convertNum(test.in, test.coords)
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("did not get expected result (-want/+got):\n%s", diff)
			}
		})
	}
}

func Test_gearsTouching(t *testing.T) {
	gears := findGears(util.ReadString(`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......*755
...$.*....
.664.598..`))
	for _, test := range []struct {
		coords       []int
		i            int
		wantX, wantY []int
	}{
		{[]int{0, 1, 2}, 0, []int{1}, []int{3}},
		{[]int{0, 1, 2}, 4, []int{4}, []int{3}},
		{[]int{7, 8, 9}, 7, []int{7}, []int{6}},
	} {
		t.Run(strconv.Itoa(test.i), func(t *testing.T) {
			gotX, gotY := gearsTouching(gears, test.coords, test.i)
			if diff := cmp.Diff(test.wantX, gotX); diff != "" {
				t.Errorf("did not get expected X (-want/+got):\n%s", diff)
			}
			if diff := cmp.Diff(test.wantY, gotY); diff != "" {
				t.Errorf("did not get expected Y (-want/+got):\n%s", diff)
			}
		})
	}
}
