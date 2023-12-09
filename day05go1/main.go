package main

import (
	"fmt"
	"sort"
	"strings"

	"advents2023/util"
)

func main() {
	input := util.ReadFile("./day05go1/input.txt")
	fmt.Println(lowestLocation(input))
}

func lowestLocation(ss []string) int {
	line := 0
	seeds := scanSeeds(ss[line])

	nextMap := mapper{}
	for {
		nextMap, line = scanMap(ss, line+2)
		if nextMap.empty() {
			break
		}
		for i := range seeds {
			seeds[i] = nextMap.convert(seeds[i])
		}
	}
	sort.Ints(seeds)
	return seeds[0]
}

func scanMap(ss []string, line int) (mapper, int) {
	res := mapper{}
	for ; line < len(ss); line++ {
		s := ss[line]
		if s == "" {
			break
		}
		if strings.Contains(s, "map") {
			continue
		}
		res.maps = append(res.maps, scanMapChunk(s))
	}
	return res, line
}

func scanMapChunk(s string) mapChunk {
	str := strings.Split(s, " ")
	return mapChunk{
		dest:  util.ToInt(str[0]),
		src:   util.ToInt(str[1]),
		shift: util.ToInt(str[2]),
	}
}

func scanSeeds(s string) []int {
	prefix := strings.Split(s, ": ")
	return convertInts(strings.Split(prefix[1], " "))
}

func convertInts(in []string) []int {
	res := make([]int, len(in))
	for i := range in {
		res[i] = util.ToInt(in[i])
	}
	return res
}

type mapper struct {
	maps []mapChunk
}

func (m *mapper) empty() bool {
	return len(m.maps) == 0
}

func (m *mapper) convert(n int) int {
	for _, mc := range m.maps {
		nn := mc.convert(n)
		if nn != n {
			return nn
		}
	}
	return n
}

type mapChunk struct {
	dest, src, shift int
}

func (m *mapChunk) convert(n int) int {
	if n < m.src || n >= m.src+m.shift {
		return n
	}
	return n + m.dest - m.src
}
