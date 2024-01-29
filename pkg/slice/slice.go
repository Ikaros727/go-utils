package slice

// Exist determines whether the target is in the s
func Exist[T comparable](s []T, target T) bool {
	return CustomizeExist(s, target, func(s, target T) bool { return s == target })
}

// CustomizeExist use customize function to determines the target is in the s
// customize: return true if s == target
func CustomizeExist[T any](s []T, target T, customize func(s, target T) bool) bool {
	for i := range s {
		if customize(s[i], target) {
			return true
		}
	}

	return false
}
