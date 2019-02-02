package closer

import (
	"errors"
	"testing"
)

var mockErr = errors.New("failed to close")

type closerMock struct {
	Counter int
	FailOn  int
}

func (c *closerMock) Close() error {
	c.Counter += 1
	if c.Counter == c.FailOn {
		return mockErr
	}
	return nil
}

func TestMayClose(t *testing.T) {
	t.Run("should close successfully", func(t *testing.T) {
		c := new(closerMock)
		MayClose(c, c, c)

		if c.Counter != 3 {
			t.Fatalf("counter not equal 3")
		}
	})
	t.Run("should failed", func(t *testing.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Fatal("should not panic")
			}
		}()
		c := &closerMock{0, 2}
		MayClose(c, c, c)
	})
}

func TestTryClose(t *testing.T) {
	t.Run("should close successfully", func(t *testing.T) {
		c := new(closerMock)
		TryClose(c, c, c)

		if c.Counter != 3 {
			t.Fatalf("counter not equal 3")
		}
	})
	t.Run("should failed", func(t *testing.T) {
		c := &closerMock{0, 2}
		defer func() {
			if c.Counter != 3 {
				t.Fatalf("counter not equal 3, but equals %d", c.Counter)
			}
			if err := recover(); err == nil {
				t.Fatal("test not paniced")
			}
		}()
		TryClose(c, c, c)
	})
}

func TestMustClose(t *testing.T) {
	t.Run("should close successfully", func(t *testing.T) {
		c := new(closerMock)
		MustClose(c, c, c)

		if c.Counter != 3 {
			t.Fatalf("counter not equal 3")
		}
	})
	t.Run("should failed", func(t *testing.T) {
		c := &closerMock{0, 2}
		defer func() {
			if c.Counter != 2 {
				t.Fatalf("counter not equal 2, but equals %d", c.Counter)
			}
			if err := recover(); err == nil {
				t.Fatal("test not paniced")
			}
		}()
		MustClose(c, c, c)
	})
}