package tree

import (
	"testing"
)

func TestNew(t *testing.T) {
	root := NewBST()
	t.Log(root)
}
