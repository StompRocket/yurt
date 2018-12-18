package yurt

import "testing"

func TestBox(t *testing.T) {
	Box("Hello, World", "")
}

func TestColor(t *testing.T) {
	Box("Red", Red)
}