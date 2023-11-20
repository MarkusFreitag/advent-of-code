package setutil

var exists = struct{}{}

type Set[E comparable] map[E]struct{}

func NewSet[E comparable](items ...E) Set[E] {
	set := make(Set[E])
	for _, item := range items {
		set[item] = exists
	}
	return set
}

func SliceToSet[S ~[]E, E comparable](slice S) Set[E] {
	set := make(Set[E])
	for _, item := range slice {
		set[item] = exists
	}
	return set
}

func SetToSlice[E comparable](set Set[E]) []E {
	slice := make([]E, len(set))
	var idx int
	for item := range set {
		slice[idx] = item
		idx++
	}
	return slice
}

func Difference[S ~[]E, E comparable](sliceA, sliceB S) (S, S) {
	diffA := make(S, 0)
	diffB := make(S, 0)

	setA, setB := SliceToSet(sliceA), SliceToSet(sliceB)

	for a := range setA {
		if _, ok := setB[a]; !ok {
			diffA = append(diffA, a)
		}
	}
	for b := range setB {
		if _, ok := setA[b]; !ok {
			diffB = append(diffB, b)
		}
	}

	return diffA, diffB
}

func Intersect[S ~[]E, E comparable](sliceA, sliceB S) S {
	intersect := make(S, 0)

	setA, setB := SliceToSet(sliceA), SliceToSet(sliceB)

	for a := range setA {
		if _, ok := setB[a]; ok {
			intersect = append(intersect, a)
		}
	}

	return intersect
}

func Union[S ~[]E, E comparable](slices ...S) S {
	if len(slices) == 0 {
		return nil
	}
	if len(slices) == 1 {
		return slices[0]
	}

	union := make(Set[E])
	for _, slice := range slices {
		for _, item := range slice {
			union[item] = exists
		}
	}

	return SetToSlice(union)
}

func UnionAll[S ~[]E, E comparable](slices ...S) S {
	if len(slices) == 0 {
		return nil
	}
	if len(slices) == 1 {
		return slices[0]
	}
	for _, slice := range slices[1:] {
		slices[0] = append(slices[0], slice...)
	}
	return slices[0]
}

func InnerJoin[S ~[]E, E comparable](left, right S) S {
	result := make(S, 0, len(left)+len(right))
	rightSet := SliceToSet(right)
	leftSet := make(Set[E], len(left))

	for _, v := range left {
		_, ok := rightSet[v]
		_, alreadyExists := leftSet[v]
		if ok && !alreadyExists {
			leftSet[v] = exists
			result = append(result, v)
		}
	}
	return result
}

func OuterJoin[S ~[]E, E comparable](left, right S) S {
	leftJoin := LeftJoin(left, right)
	rightJoin := RightJoin(left, right)

	result := make([]E, len(leftJoin)+len(rightJoin))
	copy(result, leftJoin)
	for i, v := range rightJoin {
		result[len(leftJoin)+i] = v
	}
	return result
}

func LeftJoin[S ~[]E, E comparable](left, right S) S {
	result := make([]E, 0, len(left))
	rightSet := SliceToSet(right)

	for _, v := range left {
		_, ok := rightSet[v]
		if !ok {
			result = append(result, v)
		}
	}
	return result
}

func RightJoin[S ~[]E, E comparable](left, right S) S {
	return LeftJoin(right, left)
}
