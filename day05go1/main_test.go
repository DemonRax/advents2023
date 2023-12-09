package main

import (
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
			in: `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`,
			want: 35,
		},
	} {
		t.Run(test.in, func(t *testing.T) {
			if got := lowestLocation(util.ReadString(test.in)); got != test.want {
				t.Errorf("got = %d, want %d", got, test.want)
			}
		})
	}
}

func Test_mapChunkConvert(t *testing.T) {
	for _, test := range []struct {
		in   string
		n    int
		want int
	}{
		{"50 98 2", 50, 50},
		{"50 98 2", 98, 50},
		{"50 98 2", 99, 51},
		{"50 98 2", 100, 100},
		{"50 98 2", 52, 52},
		{"52 50 48", 49, 49},
		{"52 50 48", 50, 52},
		{"52 50 48", 97, 99},
		{"52 50 48", 98, 98},
	} {
		t.Run(strconv.Itoa(test.want), func(t *testing.T) {
			mc := scanMapChunk(test.in)
			if got := mc.convert(test.n); got != test.want {
				t.Errorf("got = %d, want %d", got, test.want)
			}
		})
	}
}

func Test_mapConvert(t *testing.T) {
	for _, test := range []struct {
		in   string
		n    int
		want int
	}{
		{`50 98 2
52 50 48`,
			49, 49},
		{`50 98 2
52 50 48`,
			50, 52},
		{`50 98 2
52 50 48`,
			79, 81},
		{`50 98 2
52 50 48`,
			100, 100},
		{`50 98 2
52 50 48`,
			14, 14},
	} {
		t.Run(strconv.Itoa(test.want), func(t *testing.T) {
			m, _ := scanMap(util.ReadString(test.in), 0)
			if got := m.convert(test.n); got != test.want {
				t.Errorf("got = %d, want %d", got, test.want)
			}
		})
	}
}
