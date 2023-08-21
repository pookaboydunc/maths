package set

import (
	"fmt"
	"testing"
)

func Test_NewSet(t *testing.T) {
	A := NewSet(1, 1, 2, 3, 4, 5)
	if A.Cardinality() != 5 {
		t.Errorf("A should have 5 elements. Instead it has a cardinality of %f", A.Cardinality())
	}
	A.Remove(1)
	if A.Cardinality() != 4 {
		t.Errorf("A does should have 4 elements after one was removed. Instead it has a cardinality of %f", A.Cardinality())
	}
	B := NewSet()
	if B.Cardinality() != 0 {
		t.Errorf("B should not have any elements. Instead it has a cardinality of %f", B.Cardinality())
	}
	B.Add("20")
	if B.Cardinality() != 1 {
		t.Errorf("B should have 1 element. Instead it has a cardinality of %f", B.Cardinality())
	}
}

func Test_Contains(t *testing.T) {
	A := NewSet(1, 1, 2, 3, 4, 5)
	if A.Cardinality() != 5 {
		t.Errorf("A should have 5 elements. Instead it has a cardinality of %f", A.Cardinality())
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
	A := NewSet(1, 1, 2, 3, 4, 5)
	if A.Cardinality() != 5 {
		t.Errorf("A should have 5 elements. Instead it has a cardinality of %f", A.Cardinality())
	}
	A.Remove(1)
	B := NewSet(1, 1, 2, 3, 4, 5)
	B.Remove(2, 3, 4)
	C := A.Subset(B)
	if C.Cardinality() != 1 {
		t.Errorf("C has a cardinality of %f but was expecting a cardinality of %d", C.Cardinality(), 1)
	}
	D := NewSet(5, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
	F := C.Subset(D)
	if F.Cardinality() != 1 {
		t.Errorf("F has a cardinality of %f but was expecting a cardinality of %d", F.Cardinality(), 1)
	}
}

func Test_IsSubset(t *testing.T) {
	A := NewSet(1, 1, 2, 3, 4, 5)
	if A.Cardinality() != 5 {
		t.Errorf("A should have 5 elements. Instead it has a cardinality of %f", A.Cardinality())
	}
	A.Remove(1)
	B := NewSet(1, 1, 2, 3, 4, 5)
	B.Remove(2, 3, 4)
	C := A.Subset(B)
	if C.Cardinality() != 1 {
		t.Errorf("C has a cardinality of %f but was expecting a cardinality of %d", C.Cardinality(), 1)
	}
	if !C.IsSubset(A) {
		t.Error("C should be a subset of Set A")
	}
	D := NewSet("testing", true)
	if C.IsSubset(D) {
		t.Error("C should not be a subset of Set D")
	}
	E := NewSet(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
	if E.IsSubset(D) {
		t.Error("E should not be a subset of D")
	}
}

func Test_IsProperSubset(t *testing.T) {
	A := NewSet(1, 1, 2, 3, 4, 5)
	if A.Cardinality() != 5 {
		t.Errorf("A should have 5 elements. Instead it has a cardinality of %f", A.Cardinality())
	}
	A.Remove(1)
	B := NewSet(1, 1, 2, 3, 4, 5)
	B.Remove(2, 4)
	C := A.Subset(B)
	if C.Cardinality() != 2 {
		t.Errorf("C has a cardinality of %d but was expecting a cardinality of %f", 2, C.Cardinality())
	}
	if !C.IsSubset(A) {
		t.Error("C should be a subset of Set A")
	}
	if !C.IsProperSubset(A) {
		t.Error("C should be a proper subset of Set A")
	}
	D := NewSet(3, 5)
	if C.IsProperSubset(D) {
		t.Error("C should not be a proper subset of Set D as they are the same.")
	}
	C.Add(true)
	if C.IsProperSubset(D) {
		t.Error("C should be a proper subset of Set D as they are not the same.")
	}
}

func Test_IsEquivalent(t *testing.T) {
	A := NewSet(1, 2, 3, 4, 5)
	B := NewSet("1", "2", "3", "4", "5")
	C := NewSet("hello world")
	if !A.IsEquivalent(B) {
		t.Error("A and B should be equivalent")
	}
	if B.IsEquivalent(C) {
		t.Error("B and C should not be equivalent")
	}
}

func Test_Union(t *testing.T) {
	A := NewSet(1, 1, 2, 3, 4, 5)
	if A.Cardinality() != 5 {
		t.Errorf("A should have 5 elements. Instead it has a cardinality of %f", A.Cardinality())
	}
	A.Remove(1)
	B := NewSet(1, 1, 2, 3, 4, 5)
	C := A.Union(B)
	if C.Cardinality() != 5 {
		t.Errorf("C should have 5 elements. Instead it has a cardinality of %f", C.Cardinality())
	}
	for e := range C.E {
		if !A.Contains(e) {
			if !B.Contains(e) {
				t.Errorf("C should contain element %v but it does not", e)
			}
		}
	}
}

func Test_CommutativeUnion(t *testing.T) {
	A := NewSet(1, 1, 2, 3, 4, 5)
	if A.Cardinality() != 5 {
		t.Errorf("A should have 5 elements. Instead it has a cardinality of %f", A.Cardinality())
	}
	A.Remove(1)
	B := NewSet(1, 1, 2, 3, 4, 5)
	C := A.Union(B)
	if C.Cardinality() != 5 {
		t.Errorf("C should have 5 elements. Instead it has a cardinality of %f", C.Cardinality())
	}
	for e := range C.E {
		if !A.Contains(e) {
			if !B.Contains(e) {
				t.Errorf("C should contain element %v but it does not", e)
			}
		}
	}
	D := Union(B, A)

	if !D.IsEqual(C) {
		t.Errorf("Expected D (%v) to be the same as C (%v)", D, C)
	}
}

// A∪(B∪C)=(A∪B)∪C
func Test_AssociativityUnion(t *testing.T) {
	X := NewSet(1, 1, 2, 3, 4, 5)
	Y := NewSet(1, 2)
	Z := NewSet(3, 4, 5, 6)
	op1 := X.Union(Y.Union(Z))
	op2 := X.Union(Y).Union(Z)
	if op1.Cardinality() != op2.Cardinality() {
		t.Error("Set union operations do not appear to be associative")
	}
}

func Test_IsSuperset(t *testing.T) {
	A := NewSet(1, 1, 2, 3, 4, 5)
	B := NewSet(1, 1, 2, 3, 4, 5, 9, 10, 11)
	if !B.IsSuperset(A) {
		t.Error("Expecting B to be a superset of A")
	}
	if A.IsSuperset(B) {
		t.Error("Not expecting A to be a superset of B")
	}
}

func Test_IsProperSuperset(t *testing.T) {
	A := NewSet(1, 1, 2, 3, 4, 5)
	B := NewSet(1, 1, 2, 3, 4, 5, 9, 10, 11)
	if !B.IsProperSuperset(A) {
		t.Error("Expecting B to be a proper superset of A")
	}
	if A.IsProperSuperset(B) {
		t.Error("Not expecting A to be a proper superset of B")
	}
	C := B
	if C.IsProperSuperset(B) {
		t.Error("Not expecting C to be a proper superset of B as they are the same.")
	}
}

func Test_Equals(t *testing.T) {
	A := NewSet(1, 2, 3, 4, 5)
	B := NewSet(1, 2, 3, 4, 5, 9, 10, 11)
	if Equals(A, B) {
		t.Errorf("Expected A (%v) to be the same as B (%v)", A, B)
	}
	B.Remove(1, 2, 3)
	if Equals(A, B) {
		t.Errorf("Expected A (%v) to be the same as B (%v)", A, B)
	}
	C := NewSet(6, 7, 8, 9, 10)
	if Equals(A, C) {
		t.Errorf("Expected A (%v) to be the same as C (%v)", A, B)
	}
}

func Test_SuchThat(t *testing.T) {
	condition := func(x interface{}) bool {
		return x.(int)%2 == 0
	}
	A := SuchThat(condition, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	for e := range A.E {
		if e.(int)%2 != 0 {
			t.Errorf("Expecting all elements to be positive number yet found element %v", e)
		}
	}
}

func Test_Intersect(t *testing.T) {
	A := NewSet(1, 2, "3")
	B := NewSet("1", "2", "3", "4")
	C := A.Intersect(B)
	if C.Cardinality() != 1 && !C.Contains("3") {
		t.Errorf(`Expecting A intersect B to contain only "3" instead it contains %v`, C.E)
	}
	D := B.Intersect(A)
	if D.Cardinality() != 1 && !D.Contains("3") {
		t.Errorf(`Expecting B intersect A to contain only "3" instead it contains %v`, D.E)
	}
	X := NewSet(1, 1, 2, 3, 4, 5)
	Y := NewSet(1, 2)
	Z := NewSet(3, 4, 5, 6)
	op1 := X.Intersect(Y.Intersect(Z))
	op2 := X.Intersect(Y).Intersect(Z)
	if op1.Cardinality() != op2.Cardinality() {
		t.Error("Set intersection operations do not appear to be associative")
	}
}

func Test_IsDisjoint(t *testing.T) {
	A := NewSet(1, 2, "3")
	B := NewSet("1", "2", "3", "4")
	if A.IsDisjoint(B) {
		t.Error(`Expecting A & B to not be disjoint`)
	}
	X := NewSet(1, 2)
	Y := NewSet(3, 4, 5, 6)
	if !X.IsDisjoint(Y) {
		t.Error(`Expecting X & Y to be disjoint`)
	}
}

// A∩B=B∩A
func Test_CommutativeIntersect(t *testing.T) {
	A := NewSet(1, 1, 2, 3, 4, 5)
	if A.Cardinality() != 5 {
		t.Errorf("A should have 5 elements. Instead it has a cardinality of %f", A.Cardinality())
	}
	A.Remove(1)
	B := NewSet(1, 1, 2, 3, 4, 5)
	C := A.Intersect(B)
	if C.Cardinality() != 4 {
		t.Errorf("C should have 4 elements. Instead it has a cardinality of %f", C.Cardinality())
	}
	for e := range C.E {
		if !A.Contains(e) {
			if !B.Contains(e) {
				t.Errorf("C should contain element %v but it does not", e)
			}
		}
	}
	D := Intersect(B, A)

	if !D.IsEqual(C) {
		t.Errorf("Expected D (%v) to be the same as C (%v)", D, C)
	}
}
func Test_Difference(t *testing.T) {
	A := NewSet(2, "3")
	A.Add(1, 2, "3")
	B := NewSet("1", "2", "3", "4")
	C := Difference(A, B)
	if C.Cardinality() != 2 && !C.Contains(1, 2) {
		t.Errorf(`Expecting A difference B to contain only 1 & 2 instead it contains %v`, C.E)
	}
}

func Test_SymetricDifferencec(t *testing.T) {
	A := NewSet(1, 2, "3")
	B := NewSet("1", "2", "3", "4")
	C := SymetricDifferencec(A, B)
	if C.Cardinality() != 5 && !C.Contains(1, 2, "1", "2", "4") {
		t.Errorf(`Expecting A difference B to contain only 1, 2, "1", "2" & "4" instead it contains %v`, C.E)
	}
}

func Test_PowersetCardinality(t *testing.T) {
	A := NewSet(1, 2, 3)
	powCard := A.PowersetCardinality()
	if powCard != 8 {
		t.Errorf("Expecting a powerset cardinality of 4 but received %d", powCard)
	}
}

func Test_Complement(t *testing.T) {
	A := NewSet(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
	B := NewSet(1, 2, 3, 4, 5)
	C := NewSet(6, 7, 8, 9, 10)
	U := []*Set{B, C}
	D := Complement(A, U...)
	if D.Cardinality() != 5 {
		t.Errorf("Expecting cardinality of the complement to be 5 instead got %f", D.Cardinality())
	}
	t.Log(D.E)
}

func Test_Powerset(t *testing.T) {
	A := NewSet(1, 2, 3)
	P := A.Powerset()
	if P.Cardinality() != 8 {
		t.Errorf("Expecting a cardinality of %d instead got %f", 8, P.Cardinality())
	}
}

func Test_CartesianProduct(t *testing.T) {
	A := NewSet(1, 2, 3)
	B := NewSet(3, 4, 5, 6)
	C := CartesianProduct(A, B)
	if C.Cardinality() != (A.Cardinality() * B.Cardinality()) {
		t.Errorf("Expecting a cardinality of %f instead got %f", A.Cardinality()*B.Cardinality(), C.Cardinality())
	}
	t.Log(C.SetToSlice()...)
}

func Test_DisjointUnion(t *testing.T) {
	A := NewSet(1, 2)
	B := NewSet(3, 4, 5, 6)
	C := DisjointUnion(A, B)
	fmt.Println(C.String())
}
