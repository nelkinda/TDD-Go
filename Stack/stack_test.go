package Stack

import (
	"github.com/christianhujer/assert"
	"testing"
)

func Test_stack(t *testing.T) {
	stack := Stack{}
	_ = assert.True(t, stack.IsEmpty())
	_ = assert.Equals(t, 0, stack.GetSize())
	stack.Push(5)
	_ = assert.False(t, stack.IsEmpty())
	_ = assert.Equals(t, 1, stack.GetSize())
	_ = assert.Equals(t, 5, stack.Peek())
	stack.Push(10)
	_ = assert.False(t, stack.IsEmpty())
	_ = assert.Equals(t, 2, stack.GetSize())
	_ = assert.Equals(t, 10, stack.Peek())
	_ = assert.Equals(t, 10, stack.Pop())
	_ = assert.False(t, stack.IsEmpty())
	_ = assert.Equals(t, 1, stack.GetSize())
	_ = assert.Equals(t, 5, stack.Pop())
	_ = assert.True(t, stack.IsEmpty())
	_ = assert.Equals(t, 0, stack.GetSize())
}
