package syncx

import (
	"errors"
)

type Semaphore struct {
	body chan struct{}
}

func NewSemaphore(capacity int) *Semaphore {
	sem := new(Semaphore)
	sem.body = make(chan struct{}, capacity)
	return sem
}

func (sem *Semaphore) Aquire() {
	entry := struct{}{}
	sem.body <- entry
}

func (sem *Semaphore) Release() {
	<-sem.body
}

func (sem *Semaphore) AquireN(n int) (err error) {
	if cap(sem.body) < n {
		err = errors.New("Capacity of semaphore is less than aquired value with AquireN.")
		return err
	}

	for i := 0; i < n; i++ {
		sem.Aquire()
	}

	return err
}

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

func (sem *Semaphore) AquireAll() {
	for i := 0; i < cap(sem.body); i++ {
		sem.Aquire()
	}
}

func (sem *Semaphore) ReleaseAll() {
	for i := 0; i < cap(sem.body); i++ {
		sem.Release()
	}
}
