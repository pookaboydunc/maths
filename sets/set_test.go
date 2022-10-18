package set

import (
	"fmt"
	"math"
	"testing"
)

func Test_NewSet(t *testing.T) {
	A := NewSet(1, 1, 2, 3, 4, 5)
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
	A := NewSet(1, 1, 2, 3, 4, 5)
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
	A := NewSet(1, 1, 2, 3, 4, 5)
	if A.Cardinality() != 5 {
		t.Errorf("A should have 5 elements. Instead it has a cardinality of %d", A.Cardinality())
	}
	A.Remove(1)
	B := NewSet(1, 1, 2, 3, 4, 5)
	B.Remove(2, 3, 4)
	C := A.Subset(B)
	if C.Cardinality() != 1 {
		t.Errorf("C has a cardinality of %d but was expecting a cardinality of %d", 1, C.Cardinality())
	}
	D := NewSet(5, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
	F := C.Subset(D)
	if F.Cardinality() != 1 {
		t.Errorf("F has a cardinality of %d but was expecting a cardinality of %d", 1, F.Cardinality())
	}
}

func Test_IsSubset(t *testing.T) {
	A := NewSet(1, 1, 2, 3, 4, 5)
	if A.Cardinality() != 5 {
		t.Errorf("A should have 5 elements. Instead it has a cardinality of %d", A.Cardinality())
	}
	A.Remove(1)
	B := NewSet(1, 1, 2, 3, 4, 5)
	B.Remove(2, 3, 4)
	C := A.Subset(B)
	if C.Cardinality() != 1 {
		t.Errorf("C has a cardinality of %d but was expecting a cardinality of %d", 1, C.Cardinality())
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
		t.Errorf("A should have 5 elements. Instead it has a cardinality of %d", A.Cardinality())
	}
	A.Remove(1)
	B := NewSet(1, 1, 2, 3, 4, 5)
	B.Remove(2, 4)
	C := A.Subset(B)
	if C.Cardinality() != 2 {
		t.Errorf("C has a cardinality of %d but was expecting a cardinality of %d", 2, C.Cardinality())
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
		t.Errorf("A should have 5 elements. Instead it has a cardinality of %d", A.Cardinality())
	}
	A.Remove(1)
	B := NewSet(1, 1, 2, 3, 4, 5)
	C := A.Union(B)
	if C.Cardinality() != 5 {
		t.Errorf("C should have 5 elements. Instead it has a cardinality of %d", C.Cardinality())
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
		t.Errorf("A should have 5 elements. Instead it has a cardinality of %d", A.Cardinality())
	}
	A.Remove(1)
	B := NewSet(1, 1, 2, 3, 4, 5)
	C := A.Union(B)
	if C.Cardinality() != 5 {
		t.Errorf("C should have 5 elements. Instead it has a cardinality of %d", C.Cardinality())
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

func Test_JaccardSimilarity(t *testing.T) {
	A := NewSet(0, 1, 2, 5, 6, 8, 9)
	B := NewSet(0, 2, 3, 4, 5, 7, 9)
	index := JaccardSimilarity(A, B)
	if index != 0.4 {
		t.Errorf("expected a similarity index of 0.4 instead got %.2f", index)
	}
	C := NewSet(0, 1, 2, 3, 4, 5)
	D := NewSet(6, 7, 8, 9, 10)
	index2 := JaccardSimilarity(C, D)
	if index2 != 0 {
		t.Errorf("expected a similarity index of 0 instead got %.2f", index)
	}
	E := NewSet("cat", "dog", "hippo", "monkey")
	F := NewSet("monkey", "rhino", "ostrich", "salmon")
	index3 := JaccardSimilarity(E, F)
	shouldBe := 0.14
	tolerance := 0.01
	if diff := math.Abs(index3 - shouldBe); diff > tolerance {
		t.Errorf("expected a similarity index of 0.14 instead got %.6f", index3)
	}
}

func Test_JaccardDistance(t *testing.T) {
	A := NewSet(0, 1, 2, 5, 6, 8, 9)
	B := NewSet(0, 2, 3, 4, 5, 7, 9)
	index := JaccardSimilarity(A, B)
	distance := JaccardDistance(A, B)
	if index != 0.4 {
		t.Errorf("expected a similarity index of 0.4 instead got %.2f", index)
	}
	if distance != 0.6 {
		t.Errorf("expected a distance of 0.6 instead got %.2f", distance)
		t.Log(index, distance)
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

func Test_DSC(t *testing.T) {
	A := NewSet("ni", "ig", "gh", "ht")
	B := NewSet("na", "ac", "ch", "ht")
	dsc := DSC(A, B)
	if dsc != 0.25 {
		t.Errorf("Expecting 0.25 but got %f", dsc)
	}
	t.Log(dsc)
}

func Test_OverlapCoeffiecient(t *testing.T) {
	A := NewSet("ni", "ig", "gh", "ht")
	B := NewSet("br", "ri", "ig", "gh", "ht")
	overlapCo := OverlapCoefficient(A, B)
	if (overlapCo) != 0.75 {
		t.Errorf("Expecting 0.75 but got %f", overlapCo)
	}
	C := NewSet("br", "ri", "ig", "gh", "ht")
	D := NewSet("fr", "ry")
	overlapCo = OverlapCoefficient(C, D)
	if (overlapCo) != 0 {
		t.Errorf("Expecting 0 but got %f", overlapCo)
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
		t.Errorf("Expecting cardinality of the complement to be 5 instead got %d", D.Cardinality())
	}
	t.Log(D.E)
}

func Test_Powerset(t *testing.T) {
	A := NewSet(1, 2, 3)
	P := A.Powerset()
	if P.Cardinality() != 8 {
		t.Errorf("Expecting a cardinality of %d instead got %d", 8, P.Cardinality())
	}
}

func Test_Map(t *testing.T) {
	A := NewSet(1, 2, 3, 4)
	f := func(x interface{}) interface{} {
		return x.(int) + A.Cardinality()
	}
	B := A.Map(f)
	for e := range A.E {
		if !B.Contains((f(e))) {
			t.Error("The mapping function has not correctly mapped elements")
		}
	}
}

// A={1,2}
// B={3,4}
// A×B={(1,3),(2,3),(1,4),(2,4)}
// B×A={(3,1),(3,2),(4,1),(4,2)}

func Test_CartesianProduct(t *testing.T) {
	A := NewSet(1, 2)
	B := NewSet(3, 4)
	C := A.CartesianProduct(B)
	fmt.Println(C.String())
}
