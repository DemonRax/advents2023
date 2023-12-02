package main

import (
	"fmt"
	"strings"

	"advents2023/util"
)

func main() {
	input := util.ReadFile("./day02go1/input.txt")
	fmt.Println(minPossible(input))
}

const (
	r = "red"
	g = "green"
	b = "blue"
)

func minPossible(ss []string) int {
	var sum int
	for _, s := range ss {
		gm := scanGame(s)
		mr, mg, mb := gm.minPossible()
		sum += mr * mg * mb
	}
	return sum
}

type game struct {
	raw   string
	index int
	sets  []map[string]int // RGB
}

func scanGame(s string) game {
	sub := strings.Split(s, ":")
	title := strings.Split(sub[0], " ")
	res := game{
		raw:   s,
		index: util.ToInt(title[1]),
	}
	sets := strings.Split(sub[1], ";")
	res.sets = make([]map[string]int, len(sets))
	for i, set := range sets {
		res.sets[i] = scanSet(set)
	}
	return res
}

func scanSet(s string) map[string]int {
	tuples := strings.Split(strings.TrimSpace(s), ", ")
	res := make(map[string]int, 3)
	for _, t := range tuples {
		tuple := strings.Split(t, " ")
		res[tuple[1]] = util.ToInt(tuple[0])
	}
	return res
}

func (gm game) minPossible() (int, int, int) {
	mr, mg, mb := 1, 1, 1
	for _, set := range gm.sets {
		if set[r] > mr {
			mr = set[r]
		}
		if set[g] > mg {
			mg = set[g]
		}
		if set[b] > mb {
			mb = set[b]
		}
	}
	return mr, mg, mb
}
