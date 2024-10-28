package input

import (
	"sync"
)

type RingBuffer struct {
	buffer []rune
	size   int
	head   int
	tail   int
	count  int
	mutex  sync.RWMutex
}

func NewRingBuffer(size int) *RingBuffer {
	if size <= 0 {
		size = 1
	}
	return &RingBuffer{
		buffer: make([]rune, size),
		size:   size,
		head:   0,
		tail:   0,
		count:  0,
	}
}

func (rb *RingBuffer) Add(r rune) {
	rb.mutex.Lock()
	defer rb.mutex.Unlock()

	if rb.buffer == nil {
		panic("Buffer is nil")
	}

	if rb.count == rb.size {
		rb.head = (rb.head + 1) % rb.size
	} else {
		rb.count++
	}

	rb.buffer[rb.tail] = r
	rb.tail = (rb.tail + 1) % rb.size
}

func (rb *RingBuffer) Get() []rune {
	rb.mutex.Lock()
	defer rb.mutex.Unlock()

	if rb.count == 0 {
		return []rune{}
	}
	result := make([]rune, rb.count)
	for i := 0; i < rb.count; i++ {
		// The modulo operator makes sure that it will retrieve the data within the bound of the RingBuffer
		result[i] = rb.buffer[(rb.head+i)%rb.size]
	}

	return result
}

func (rb *RingBuffer) Clear() {
	rb.mutex.Lock()
	defer rb.mutex.Unlock()

	rb.head = 0
	rb.tail = 0
	rb.count = 0
}

func (rb *RingBuffer) IsFull() bool {
	rb.mutex.RLock()
	defer rb.mutex.RUnlock()

	return rb.count == rb.size
}

func (rb *RingBuffer) IsEmpty() bool {
	rb.mutex.RLock()
	defer rb.mutex.RUnlock()

	return rb.count == 0
}

func (rb *RingBuffer) Count() int {
	rb.mutex.RLock()
	defer rb.mutex.RUnlock()

	return rb.count
}

func (rb *RingBuffer) RemoveLast() {
	rb.mutex.Lock()
	defer rb.mutex.Unlock()

	if rb.count > 0 {
		rb.tail = (rb.tail - 1 + rb.size) % rb.size
		rb.count--
	}
}
