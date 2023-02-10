package lock

type Lock struct {
	lock chan struct{}
}
