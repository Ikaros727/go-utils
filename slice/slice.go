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

// NoExist determines whether the target is not in the s
func NoExist[T comparable](s []T, target T) bool {
	for i := range s {
		if s[i] == target {
			return false
		}
	}

	return true
}
