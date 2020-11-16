package pair

import (
)

type pair struct {
	first	interface{}
	second	interface{}
}

// First returns the first element of the pair
func (p pair) First() interface{} {
	return p.first
}

// Second returns the second element of the pair
func (p pair) Second() interface{} {
	return p.second
}

// MakePair initialises the pair and return 
func MakePair(first, second interface{}) pair {
	return pair{first: first, second:second}
}

// CopyPair copies the pair into another
func CopyPair(p pair) pair {
	return p
}