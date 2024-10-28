package input_test

import (
	"sync"
	"testing"

	"github.com/GabzAraujo/go-speed-typer/internal/input"
)

func TestRingBufferConcurrency(t *testing.T) {
	rb := input.NewRingBuffer(100)
	var wg sync.WaitGroup

	t.Run("Concurrent Adds", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func(r rune) {
				defer wg.Done()
				rb.Add(r)
			}(rune(i%26 + 'a'))
		}
		wg.Wait()

		if rb.Count() != 100 {
			t.Errorf("Expected count 100, got %d", rb.Count())
		}
	})

	rb.Clear()

	t.Run("Concurrent Adds and Removes", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			wg.Add(2)
			go func() {
				defer wg.Done()
				rb.Add('a')
			}()
			go func() {
				defer wg.Done()
				rb.RemoveLast()
			}()
		}
		wg.Wait()

		if rb.Count() < 0 || rb.Count() > 100 {
			t.Errorf("Unexpected count: %d", rb.Count())
		}
	})

	rb.Clear()

	t.Run("Concurrent Reads", func(t *testing.T) {
		// Pre-populate the buffer
		for i := 0; i < 50; i++ {
			rb.Add(rune(i%26 + 'a'))
		}

		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				_ = rb.Get()
			}()
		}
		wg.Wait()

		if rb.Count() != 50 {
			t.Errorf("Expected count 50, got %d", rb.Count())
		}
	})
}
