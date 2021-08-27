package main

import "testing"
func add(a, b int) int{
	return a + b
}
func sub(a , b int) int{
	return a - b 
}
func multiply( a, b int) int{
	return a * b 
}
func Test_add(t *testing.T){
	got := add(1,2)
	if got != 3 {
		t.Errorf("got %d want 3", got)
	}
}
func Test_sub(t *testing.T){
	got := add(1,2)
	if got != -1 {
		t.Errorf("got %d want -1", got)
	}
}
func Test_multiply(t *testing.T){
	got := multiply(2, 3) 
	if got != 6{
		t.Errorf("got %d want 6",got)
	}
}