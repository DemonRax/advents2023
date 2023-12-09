package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"

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
...$.*....
.664.598..`,
			want: 4361,
		},
	} {
		t.Run(test.in, func(t *testing.T) {
			if got := partNumbers(util.ReadString(test.in)); got != test.want {
				t.Errorf("got = %d, want %d", got, test.want)
			}
		})
	}
}

func Test_convertInput(t *testing.T) {
	for _, test := range []struct {
		in   string
		want [][]int
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
...$.*....
.664.598..`,
			want: [][]int{
				{4, 6, 7, -1, -1, 1, 1, 4, -1, -1},
				{-1, -1, -1, -2, -1, -1, -1, -1, -1, -1},
				{-1, -1, 3, 5, -1, -1, 6, 3, 3, -1},
				{-1, -1, -1, -1, -1, -1, -2, -1, -1, -1},
				{6, 1, 7, -2, -1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -2, -1, 5, 8, -1},
				{-1, -1, 5, 9, 2, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1, -2, 7, 5, 5},
				{-1, -1, -1, -2, -1, -2, -1, -1, -1, -1},
				{-1, 6, 6, 4, -1, 5, 9, 8, -1, -1},
			},
		},
	} {
		t.Run(test.in, func(t *testing.T) {
			got := convertInput(util.ReadString(test.in))
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("did not get expected result (-want/+got):\n%s", diff)
			}
		})
	}
}

func Test_isTouching(t *testing.T) {
	src := [][]int{
		{4, 6, 7, -1, -1, 1, 1, 4, -1, -1},
		{-1, -1, -1, -2, -1, -1, -1, -1, -1, -1},
		{-1, -1, 3, 5, -1, -1, 6, 3, 3, -1},
		{-1, -1, -1, -1, -1, -1, -2, -1, -1, -1},
		{6, 1, 7, -2, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -2, -1, 5, 8, -1},
		{-1, -1, 5, 9, 2, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -2, 7, 5, 5},
		{-1, -1, -1, -2, -1, -2, -1, -1, -1, -1},
		{-1, 6, 6, 4, -1, 5, 9, 8, -1, -1},
	}
	for _, test := range []struct {
		i, j int
		want bool
	}{
		{0, 0, false},
		{0, 3, true},
		{2, 4, true},
		{2, 6, true},
		{2, 8, false},
		{7, 7, true},
	} {
		t.Run(fmt.Sprintf("%d %d", test.i, test.j), func(t *testing.T) {
			if got := isTouching(src, test.i, test.j); got != test.want {
				t.Errorf("want %t got %t", test.want, got)
			}
		})
	}
}

func Test_scanLine(t *testing.T) {
	src := [][]int{
		{4, 6, 7, -1, -1, 1, 1, 4, -1, -1},
		{-1, -1, -1, -2, -1, -1, -1, -1, -1, -1},
		{-1, -1, 3, 5, -1, -1, 6, 3, 3, -1},
		{-1, -1, -1, -1, -1, -1, -2, -1, -1, -1},
		{6, 1, 7, -2, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -2, -1, 5, 8, -1},
		{-1, -1, 5, 9, 2, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -2, 7, 5, 5},
		{-1, -1, -1, -2, -1, -2, -1, -1, -1, -1},
		{-1, 6, 6, 4, -1, 5, 9, 8, -1, -1},
	}
	for _, test := range []struct {
		i, want int
	}{
		{0, 467},
		{1, 0},
		{2, 35 + 633},
		{4, 617},
		{7, 755},
	} {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			if got := scanLine(src, test.i); got != test.want {
				t.Errorf("want %d got %d", test.want, got)
			}
		})
	}
}
