package slice

import "github.com/MarkusFreitag/advent-of-code/util/constraints"

func Difference[T constraints.Comparable](sliceA, sliceB []T) ([]T, []T) {
	diffA := make([]T, 0)
	diffB := make([]T, 0)

	mapA, mapB := SliceToSet(sliceA), SliceToSet(sliceB)

	for a := range mapA {
		if _, ok := mapB[a]; !ok {
			diffA = append(diffA, a)
		}
	}
	for b := range mapB {
		if _, ok := mapA[b]; !ok {
			diffB = append(diffB, b)
		}
	}

	return diffA, diffB
}

func Intersect[T constraints.Comparable](sliceA, sliceB []T) []T {
	intersect := make([]T, 0)

	mapA, mapB := SliceToSet(sliceA), SliceToSet(sliceB)

	for a := range mapA {
		if _, ok := mapB[a]; ok {
			intersect = append(intersect, a)
		}
	}

	return intersect
}

func Union[T constraints.Comparable](slices ...[]T) []T {
	if len(slices) == 0 {
		return nil
	}
	if len(slices) == 1 {
		return slices[0]
	}

	union := make(map[T]struct{})
	for _, slice := range slices {
		for _, item := range slice {
			union[item] = exists
		}
	}

	return SetToSlice(union)
}

func UnionAll[T constraints.Comparable](slices ...[]T) []T {
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

func InnerJoin[T constraints.Comparable](left, right []T) []T {
	result := make([]T, 0, len(left)+len(right))
	rightSet := SliceToSet(right)
	leftSet := make(map[T]struct{}, len(left))

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

func OuterJoin[T constraints.Comparable](left, right []T) []T {
	leftJoin := LeftJoin(left, right)
	rightJoin := RightJoin(left, right)

	result := make([]T, len(leftJoin)+len(rightJoin))
	copy(result, leftJoin)
	for i, v := range rightJoin {
		result[len(leftJoin)+i] = v
	}
	return result
}

func LeftJoin[T constraints.Comparable](left, right []T) []T {
	result := make([]T, 0, len(left))
	rightSet := SliceToSet(right)

	for _, v := range left {
		_, ok := rightSet[v]
		if !ok {
			result = append(result, v)
		}
	}
	return result
}

func RightJoin[T constraints.Comparable](left, right []T) []T {
	return LeftJoin(right, left)
}
