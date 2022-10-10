package set

// JaccardSimilarity
// Jaccard Index = (the number in both sets) / (the number in either set)
//
// The same formula in notation is:
// J(A,B) = |A∩B| / |A∪B|
func JaccardSimilarity(A, B *Set) float64 {
	D := A.Intersect(B)
	U := A.Union(B)
	return float64(D.Cardinality()) / float64(U.Cardinality())
}

// JaccardDistance
// Jaccard distance = 1 - JaccardSimilarity
func JaccardDistance(A, B *Set) float64 {
	return 1 - JaccardSimilarity(A, B)
}

// DSC
// Dice Similarity Coefficient
// The Sorensen Coefficient equals twice the number of elements common to both sets divided by the sum of the number of elements in each set.
func DSC(A, B *Set) float64 {
	commonElements := float64(A.Intersect(B).Cardinality())
	sumOfElements := float64(A.Cardinality() + B.Cardinality())
	return commonElements * 2 / sumOfElements
}

// OverlapCoefficient
// The Overlap Coefficient is defined as the size of the intersection divided by the size of the smaller of the two sets.
func OverlapCoefficient(A, B *Set) float64 {
	commonElements := float64(A.Intersect(B).Cardinality())
	min := A.Cardinality()
	if B.Cardinality() < A.Cardinality() {
		min = B.Cardinality()
	}
	return commonElements / float64(min)
}

// CosineSimilarity

// Tversky Index
