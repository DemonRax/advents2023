package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"

	"advents2023/util"
)

func Test_code(t *testing.T) {
	for _, test := range []struct {
		in   string
		want int
	}{
		{
			in: `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`,
			want: 30,
		},
	} {
		t.Run(test.in, func(t *testing.T) {
			if got := scratchCards(util.ReadString(test.in)); got != test.want {
				t.Errorf("got = %d, want %d", got, test.want)
			}
		})
	}
}

func Test_scanLine(t *testing.T) {
	for _, test := range []struct {
		in                 string
		wantWins, wantNums []int
	}{
		{
			in:       `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53`,
			wantWins: []int{41, 48, 83, 86, 17},
			wantNums: []int{83, 86, 6, 31, 17, 9, 48, 53},
		},
	} {
		t.Run(test.in, func(t *testing.T) {
			gotWins, gotNums := scanLine(test.in)
			if diff := cmp.Diff(test.wantWins, gotWins); diff != "" {
				t.Errorf("did not get expected wins (-want/+got):\n%s", diff)
			}
			if diff := cmp.Diff(test.wantNums, gotNums); diff != "" {
				t.Errorf("did not get expected nums (-want/+got):\n%s", diff)
			}
		})
	}
}
