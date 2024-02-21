package doublylinkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoublyLinkedList_Add(t *testing.T) {
	var dl = &DoublyLinkedList[string]{}
	assert.Nilf(t, dl.Add(0, "first item"), "Error is not expected.")
	assert.Nilf(t, dl.Add(0, "second item"), "Error is not expected.")

	result := dl.PrintForward()
	assert.NotEmptyf(t, result, "Empty is not expected.")
	assert.Equalf(t, result, "HEAD -> first item -> second item -> NULL", "Unexpected result.")
}

func TestDoublyLinkedList_AddElements(t *testing.T) {
	testCases := []struct {
		index int
		value string
	}{
		{index: 0, value: "C"},
		{index: 0, value: "A"},
		{index: 1, value: "B"},
		{index: 3, value: "D"},
	}
	dl := &DoublyLinkedList[string]{}
	assert.Nilf(t, dl.AddElements(testCases), "Error is not expected.")

	result := dl.PrintForward()
	assert.Equalf(t, result, "HEAD -> A -> B -> C -> D -> NULL", "Unexpected result.")
}
