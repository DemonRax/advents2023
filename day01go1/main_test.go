package main

import (
	"testing"

	"advents2023/util"
)

func Test_code(t *testing.T) {
	for _, test := range []struct {
		in   string
		want int
	}{
		{
			in: `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`,
			want: 142,
		},
	} {
		t.Run(test.in, func(t *testing.T) {
			if got := calibrate(util.ReadString(test.in)); got != test.want {
				t.Errorf("got = %d, want %d", got, test.want)
			}
		})
	}
}
