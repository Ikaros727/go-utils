package slice

// Exist determines whether the target is in the s
func Exist[T comparable](s []T, target T) bool {
	for i := range s {
		if s[i] == target {
			return true
		}
	}

	return false
}

// Distinct make the element is unique in the slice
func Distinct[T comparable](slice []T) (final []T) {
	m := make(map[T]struct{})
	for _, e := range slice {
		if _, ok := m[e]; !ok {
			m[e] = struct{}{}
			final = append(final, e)
		}
	}

	return
}

// Union return the elements that exist in sliceA or sliceB
func Union[T comparable](sliceA, sliceB []T) []T {
	return Distinct(append(sliceA, sliceB...))
}

// Intersect return the elements that exist in both sliceA and sliceB
func Intersect[T comparable](sliceA, sliceB []T) (final []T) {
	m := make(map[T]struct{})
	for _, a := range sliceA {
		if _, ok := m[a]; !ok {
			m[a] = struct{}{}
		}
	}

	for _, b := range sliceB {
		if _, ok := m[b]; ok {
			final = append(final, b)
		}
	}

	return
}

// Subtract return the elements that exist in sliceIn but not exist in sliceNotIn
func Subtract[T comparable](sliceIn, sliceNotIn []T) (final []T) {
	mapNotIn := make(map[T]struct{})
	for _, notIn := range sliceNotIn {
		if _, ok := mapNotIn[notIn]; !ok {
			mapNotIn[notIn] = struct{}{}
		}
	}

	for _, in := range sliceIn {
		if _, ok := mapNotIn[in]; !ok {
			final = append(final, in)
		}
	}

	return
}

// ExclusiveOr return the elements that exist in sliceA but not in sliceB or in sliceB but not in sliceA
func ExclusiveOr[T comparable](sliceA, sliceB []T) (final []T) {
	mapA := make(map[T]struct{})
	mapB := make(map[T]struct{})

	for _, b := range sliceB {
		mapB[b] = struct{}{}
	}

	for _, a := range sliceA {
		mapA[a] = struct{}{}
		if _, ok := mapB[a]; !ok {
			final = append(final, a)
		}
	}

	for b := range mapB {
		if _, ok := mapA[b]; !ok {
			final = append(final, b)
		}
	}

	return
}
