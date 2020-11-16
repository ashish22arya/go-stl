package pair

import (
)

type pair struct {
	first	interface{}
	second	interface{}
}

func (p pair) First() interface{} {
	return p.first
}

func (p pair) Second() interface{} {
	return p.second
}

func MakePair(first, second interface{}) pair {
	return pair{first: first, second:second}
}

func CopyPair(p pair) pair {
	return p
}