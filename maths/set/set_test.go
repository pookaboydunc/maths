package set

import (
	"testing"
)

func Test_NewSet(t *testing.T) {
	elements := []interface{}{1, 1, 2, 3, 4, 5}
	A := NewSet(elements...)
	if A.Cardinality() != 5 {
		t.Errorf("A should have 5 elements. Instead it has a cardinality of %d", A.Cardinality())
	}
	A.Remove(1)
	if A.Cardinality() != 4 {
		t.Errorf("A does should have 4 elements after one was removed. Instead it has a cardinality of %d", A.Cardinality())
	}
	B := NewSet()
	if B.Cardinality() != 0 {
		t.Errorf("B should not have any elements. Instead it has a cardinality of %d", B.Cardinality())
	}
	B.Add("20")
	if B.Cardinality() != 1 {
		t.Errorf("B should have 1 element. Instead it has a cardinality of %d", B.Cardinality())
	}
}

func Test_Contains(t *testing.T) {
	elements := []interface{}{1, 1, 2, 3, 4, 5}
	A := NewSet(elements...)
	if A.Cardinality() != 5 {
		t.Errorf("A should have 5 elements. Instead it has a cardinality of %d", A.Cardinality())
	}
	if !A.Contains(1) {
		t.Error("A should contain the 1 element.")
	}
	if A.Contains("1") {
		t.Error("A should not contain the '1' element.")
	}
	if A.Contains("1", 2) {
		t.Error("A should not contain the '1' element.")
	}
	if !A.Contains(1, 2) {
		t.Error("A should contain elements 1 & 2.")
	}
}

func Test_SubsetOf(t *testing.T) {
	elements := []interface{}{1, 1, 2, 3, 4, 5}
	A := NewSet(elements...)
	if A.Cardinality() != 5 {
		t.Errorf("A should have 5 elements. Instead it has a cardinality of %d", A.Cardinality())
	}
	A.Remove(1)
	B := NewSet(elements...)
	B.Remove(2, 3, 4)
	C := Subset(&A, &B)
	if C.Cardinality() != 1 {
		t.Errorf("C has a cardinality of %d but was expecting a cardinality of %d", 1, C.Cardinality())
	}
	felements := []interface{}{5, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	D := NewSet(felements...)
	F := Subset(&C, &D)
	if F.Cardinality() != 1 {
		t.Errorf("F has a cardinality of %d but was expecting a cardinality of %d", 1, F.Cardinality())
	}
}

func Test_IsSubset(t *testing.T) {
	elements := []interface{}{1, 1, 2, 3, 4, 5}
	A := NewSet(elements...)
	if A.Cardinality() != 5 {
		t.Errorf("A should have 5 elements. Instead it has a cardinality of %d", A.Cardinality())
	}
	A.Remove(1)
	B := NewSet(elements...)
	B.Remove(2, 3, 4)
	C := Subset(&A, &B)
	if C.Cardinality() != 1 {
		t.Errorf("C has a cardinality of %d but was expecting a cardinality of %d", 1, C.Cardinality())
	}
	if !C.IsSubset(&A) {
		t.Error("C should be a subset of Set A")
	}
	D := NewSet([]interface{}{"testing", true, nil}...)
	if C.IsSubset(&D) {
		t.Error("C should not be a subset of Set D")
	}
	E := NewSet([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}...)
	if E.IsSubset(&D) {
		t.Error("E should not be a subset of D")
	}
}

func Test_IsProperSubset(t *testing.T) {
	elements := []interface{}{1, 1, 2, 3, 4, 5}
	A := NewSet(elements...)
	if A.Cardinality() != 5 {
		t.Errorf("A should have 5 elements. Instead it has a cardinality of %d", A.Cardinality())
	}
	A.Remove(1)
	B := NewSet(elements...)
	B.Remove(2, 4)
	C := Subset(&A, &B)
	if C.Cardinality() != 2 {
		t.Errorf("C has a cardinality of %d but was expecting a cardinality of %d", 2, C.Cardinality())
	}
	if !C.IsSubset(&A) {
		t.Error("C should be a subset of Set A")
	}
	if !C.IsProperSubset(&A) {
		t.Error("C should be a proper subset of Set A")
	}
	D := NewSet([]interface{}{3, 5}...)
	if C.IsProperSubset(&D) {
		t.Error("C should not be a proper subset of Set D as they are the same.")
	}
	C.Add(true)
	if C.IsProperSubset(&D) {
		t.Error("C should be a proper subset of Set D as they are not the same.")
	}
}

func Test_Equivalence(t *testing.T) {
	elements := []interface{}{1, 2, 3, 4, 5}
	elements2 := []interface{}{"1", "2", "3", "4", "5"}
	elements3 := []interface{}{"hello world"}
	A := NewSet(elements...)
	B := NewSet(elements2...)
	C := NewSet(elements3...)
	if !A.Equivalence(&B) {
		t.Error("A and B should be equivalent")
	}
	if B.Equivalence(&C) {
		t.Error("B and C should not be equivalent")
	}
}

func Test_Union(t *testing.T) {
	elements := []interface{}{1, 1, 2, 3, 4, 5}
	A := NewSet(elements...)
	if A.Cardinality() != 5 {
		t.Errorf("A should have 5 elements. Instead it has a cardinality of %d", A.Cardinality())
	}
	A.Remove(1)
	B := NewSet(elements...)
	C := A.Union(&B)
	if C.Cardinality() != 5 {
		t.Errorf("C should have 5 elements. Instead it has a cardinality of %d", C.Cardinality())
	}
	for e := range C.E {
		if !A.Contains(e) || !B.Contains(e) {
			t.Errorf("C should contain element %v but it does not", e)
		}
	}
}

func Test_IsSuperset(t *testing.T) {
	elements := []interface{}{1, 1, 2, 3, 4, 5}
	A := NewSet(elements...)
	elements = append(elements, 9, 10, 11)
	B := NewSet(elements...)
	if !B.IsSuperset(&A) {
		t.Error("Expecting B to be a superset of A")
	}
	if A.IsSuperset(&B) {
		t.Error("Not expecting A to be a superset of B")
	}
}

func Test_IsProperSuperset(t *testing.T) {
	elements := []interface{}{1, 1, 2, 3, 4, 5}
	A := NewSet(elements...)
	elements = append(elements, 9, 10, 11)
	B := NewSet(elements...)
	if !B.IsProperSuperset(&A) {
		t.Error("Expecting B to be a proper superset of A")
	}
	if A.IsProperSuperset(&B) {
		t.Error("Not expecting A to be a proper superset of B")
	}
	C := B
	if C.IsProperSuperset(&B) {
		t.Error("Not expecting C to be a proper superset of B as they are the same.")
	}
}
