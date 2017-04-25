// Package semaphore provides a simple channel-based semaphore.
package semaphore

// Semaphore channel.
type Semaphore chan struct{}

// New semaphore with the given `concurrency`.
func New(concurrency int) Semaphore {
	return make(Semaphore, concurrency)
}

// Acquire resource.
func (s Semaphore) Acquire() {
	s <- struct{}{}
}

// Release resource.
func (s Semaphore) Release() {
	<-s
}

// Wait for completion.
func (s Semaphore) Wait() {
	for i := 0; i < cap(s); i++ {
		s <- struct{}{}
	}
}

// Run `fn` in a goroutine, acquiring then releasing after its return.
func (s Semaphore) Run(fn func()) {
	s.Acquire()
	go func() {
		defer s.Release()
		fn()
	}()
}
