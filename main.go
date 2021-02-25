package main

import "fmt"

type Iterator struct {
	slice []int
}

type Iterators struct {
	list []*Iterator
	minIdx int
	minVal int
	lastValues map[int]int
}

const maxNext = 9999999

func main() {
	iterators := Iterators{list: []*Iterator{
		{slice: []int{1, 6, 9, 12, 13}},
		{slice: []int{-5, 0, 3, 8, 14}},
		{slice: []int{-19, -15, -12, -7, 7}},
	}}
	iterators.getLastValues()

	for i := 0; i < 15; i++ {
		n := iterators.Next()
		fmt.Printf("n: %d\n", n)
	}

	_ = iterators.Next()
}

func (i *Iterator) Next() (next int) {
	if len(i.slice) == 0 {
		next = maxNext
	} else {
		next = i.slice[0]
		i.slice = i.slice[1:]
	}
	return
}

func (is *Iterators) getLastValues() {
	is.lastValues = make(map[int]int)

	for j, iterator := range is.list {
		is.lastValues[j] = iterator.Next()
	}
}

func (is *Iterators) getMinIndexAndValue() {
	is.minVal = is.lastValues[0]
	is.minIdx = 0

	for idx, val := range is.lastValues {
		if val < is.minVal {
			is.minIdx = idx
			is.minVal = val
		}
	}

	return
}

func (is *Iterators) Next() (next int) {
	is.getMinIndexAndValue()
	is.lastValues[is.minIdx] = is.list[is.minIdx].Next()

	return is.minVal
}
