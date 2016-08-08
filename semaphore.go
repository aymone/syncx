// Package syncx is intended to provide functionality additional to stdlib sync package.
package syncx

import (
	"errors"
)

// Simple semaphore not visible from outside of programm.
// It is implemented using golang channels. Functions related to it
// (like Acquire, Release, etc.) will simply block till they will be able to
// perform action on semaphore.
type Semaphore struct {
	body chan struct{}
}

// NewSemaphore function accepts single argument - expected capacity as integer and
// returns pointer to created semaphore.
func NewSemaphore(capacity int) *Semaphore {
	sem := new(Semaphore)
	sem.body = make(chan struct{}, capacity)
	return sem
}

// Acquire function increments semaphore.
func (sem *Semaphore) Acquire() {
	entry := struct{}{}
	sem.body <- entry
}

// Release function decrements semaphore.
func (sem *Semaphore) Release() {
	<-sem.body
}

// AcquireN function increments semaphore by N. It will return error if N > cap(sem).
func (sem *Semaphore) AcquireN(n int) (err error) {
	if cap(sem.body) < n {
		err = errors.New("Capacity of semaphore is less than acquired value with AcquireN.")
		return err
	}

	for i := 0; i < n; i++ {
		sem.Acquire()
	}

	return err
}

// AcquireNUnsafe function is the same as AcquireN except it won't return error if N > cap(sem).
// It will simply block until semaphore is incremented N times.
func (sem *Semaphore) AcquireNUnsafe(n int) {
	for i := 0; i < n; i++ {
		sem.Acquire()
	}
}

// ReleaseN function decrements semaphore by N. It will return error if N > cap(sem).
func (sem *Semaphore) ReleaseN(n int) (err error) {
	if cap(sem.body) < n {
		err = errors.New("Capacity of semaphore is less than requested value for release.")
		return err
	}

	for i := 0; i < n; i++ {
		sem.Release()
	}

	return err
}

// ReleaseNUnsafe function is the same as ReleaseN except it won't return error if N > cap(sem).
// It will simply block until semaphore is decremented N times.
func (sem *Semaphore) ReleaseNUnsafe(n int) {
	for i := 0; i < n; i++ {
		sem.Release()
	}
}

// AcquireAll function increments semaphore till its maximum capacity.
func (sem *Semaphore) AcquireAll() {
	for i := 0; i < cap(sem.body); i++ {
		sem.Acquire()
	}
}

// ReleaseAll function decrements semaphore till its maximum capacity.
func (sem *Semaphore) ReleaseAll() {
	for i := 0; i < cap(sem.body); i++ {
		sem.Release()
	}
}
