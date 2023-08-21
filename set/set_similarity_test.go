package set

import (
	"math"
	"testing"
)

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
