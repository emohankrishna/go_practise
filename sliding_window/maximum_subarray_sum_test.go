package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	assert := assert.New(t)
	assert.ElementsMatch(MaxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3), []int{3, 3, 5, 5, 6, 7}, "3 should be returned")
}
func TestMaxSlidingWindow1(t *testing.T) {
	assert := assert.New(t)
	result := MaxSlidingWindow([]int{4, 3, 11}, 3)
	assert.ElementsMatch(result, []int{11}, "3 should be returned")
}
