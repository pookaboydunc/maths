// Package set represents the mathematical set.
package set

import (
	"sync"
)

type nothing struct{}

type elements map[interface{}]nothing

// Set is the main structure used to denote a set.
//
// {}	set	used to define a set	S={1,2,3,4,…}
type Set struct {
	E     elements
	rLock sync.Mutex
}

// NewSet returns a new set (A) of all unique elements passed into the function call.
func NewSet(els ...interface{}) (A Set) {
	A.E = make(elements, len(els))
	for _, e := range els {
		A.Add(e)
	}
	return
}

// Add inserts one or more elements into A.
func (A *Set) Add(els ...interface{}) {
	for _, e := range els {
		A.E[e] = nothing{}
	}
}

// Remove deletes one or more existing elements from A.
func (A *Set) Remove(els ...interface{}) {
	A.rLock.Lock()
	defer A.rLock.Unlock()
	for _, e := range els {
		delete(A.E, e)
	}
}

// Contains checks if one or more elements are in A.
//
// ∈	in, element of	used to denote that an element is part of a set	1∈1,2,3
// ∉	not in, not an element of	used to denote than an element is not part of a set	4∉1,2,3
func (A *Set) Contains(els ...interface{}) bool {
	A.rLock.Lock()
	defer A.rLock.Unlock()
	for _, e := range els {
		if _, ok := A.E[e]; !ok {
			return false
		}
	}
	return true
}

// Cardinality returns the size of A defined as the number of unique elements within A.
//
// ∣S∣	cardinality	used to describe the size of a set (refers to the number of unique elements if A is finite)
// S={1,2,2,2,3,4,5,5}
// ∣S∣=5
func (A *Set) Cardinality() int {
	A.rLock.Lock()
	defer A.rLock.Unlock()
	return len(A.E)
}

// Subset returns a new set (C) that is the subset of A & B.
//
// ⊆	subset	set A is a subset of set B when each element in A is also an element in B
// A={1,2}
// B={2,1,4,3,5}
// A⊆B
func Subset(A, B *Set) (C Set) {
	l := A.Cardinality()
	if l < B.Cardinality() {
		l = B.Cardinality()
	}
	C.E = make(elements, l)
	for e := range A.E {
		if B.Contains(e) {
			C.Add(e)
		}
	}
	return
}

// IsSubset checks if A is a subset of B.
//
// ⊆	subset	set A is a subset of set B when each element in A is also an element in B
// A={1,2}
// B={2,1,4,3,5}
// A⊆B
func (A *Set) IsSubset(B *Set) bool {
	if A.Cardinality() > B.Cardinality() {
		return false
	}
	for e := range A.E {
		if !B.Contains(e) {
			return false
		}
	}
	return true
}

// IsProperSubset checks if the A is a proper subset of B.
// ⊂	proper subset	set A is a proper subset of set B when each element in A is also an element in B and A≠B
// A={1,2,3,4,5}
// B={2,1,4,3,5}
// A⊆B is true but A⊂B is not true
func (A *Set) IsProperSubset(B *Set) bool {
	return A.Cardinality() < B.Cardinality() && A.IsSubset(B)
}

// Equivalence checks if A & B have the same Cardinality.
//
// Sets are equivalent when their cardinality is the same. NOT to be mistaken with equality.
func (A *Set) Equivalence(B *Set) bool {
	return A.Cardinality() == B.Cardinality()
}

// IsSuperset checks if A is a superset of B.
// ⊇	superset	set A is a superset of set B when B is a subset of A
// A={2,4,6,7,8}
// B={2,4,8}
// A⊇B
func (A *Set) IsSuperset(B *Set) bool {
	return B.IsSubset(A)
}

// IsProperSuperset checks if A is a proper superset of B.
// ⊇	proper superset	set A is a proper superset of set B when B is a subset of A and A!=B
// A={2,4,6,7,8}
// B={2,4,8}
// A⊇B
func (A *Set) IsProperSuperset(B *Set) bool {
	return A.Cardinality() > B.Cardinality() && A.IsSuperset(B)
}

// Union creates a new set (C) from elements in A & B.
// ∪ 	union	a set with the elements in set A or in set B
// A={1,2}
// B={2,3,5}
// A∪B={1,2,3,5}
func (A *Set) Union(B *Set) (C Set) {
	C.E = A.E
	for e := range B.E {
		C.Add(e)
	}
	return
}

// Intersection creates a new set (C) from elements in A || B.
// ∩	intersection	a set with the elements in set A and in set B
// A={1,2}
// B={2,3,5}
// A∩B={2}
func Intersection(A, B *Set) (C Set) {
	if A.Cardinality() < B.Cardinality() {
		for e := range A.E {
			if B.Contains(e) {
				C.Add(e)
			}
		}
	} else {
		for e := range B.E {
			if A.Contains(e) {
				C.Add(e)
			}
		}
	}
	return
}

// Difference
// −, ∖	set difference	elements in set A that are not in B
// A={1,2,3,4}
// B={2,3,5,8}
// A−B={1,4}
// B−A={5,8}

// SymetricDifference

// CartesianProduct
// ×	Cartesian product	a set containing all possible combinations of one element from A and one element from B
// A={1,2}
// B={3,4}
// A×B={(1,3),(2,3),(1,4),(2,4)}
// B×A={(3,1),(3,2),(4,1),(4,2)}

// MappingFunction

// JaccardSimilarity
// Jaccard Index = (the number in both sets) / (the number in either set) * 100

// The same formula in notation is:
// J(A,B) = |A∩B| / |A∪B|

// SuchThat
// :, ∣	such that	used to denote a condition, usually in set-builder notation or in a mathematical definition
// {x2:x+3 is prime}
