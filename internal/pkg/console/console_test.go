package console

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrintf(t *testing.T) {
	writer := &TestableWriter{}
	SetOut(writer)
	Printf("a test\n")

	assert.Equal(t, "a test\n", string(writer.Data))
}
