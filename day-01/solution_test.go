package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestFileOpenNotExist(t *testing.T) {
  _, _, err := LoadInput("wrong.txt")
  assert.True(t, err != nil, "Expected failure")
}

func TestFileContainsAlphaNumeric(t *testing.T) {
  _, _, err := LoadInput("input_test1.txt")
  assert.True(t, err != nil, "Expected failure")
}


func TestComputeDistance(t *testing.T) {
  val := ComputeDistane([]int{1, 2}, []int{2, 3})

  assert.Equal(t, 2, val)
} 
