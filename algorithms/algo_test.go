package algorithms_test

import (
	"fmt"
	"go_dsa/algorithms"
	"go_dsa/helper"
	"testing"
)

func TestValidLinearSearch(t *testing.T) {
	arr := []int{1, 2, 2, 5, 7, 6578, 876, 342, 2}
	item := 876
	wantedResult := 6
	result := algorithms.LinearSearch(item, arr)
	if result != wantedResult {
		t.Fatalf("Linear searched failed with result of: %v", result)
	}
}
func TestValidBinarySearch(t *testing.T) {
	arr := []int{1, 2, 2, 5, 7, 876}
	item := 876
	wantedResult := 5
	result := algorithms.BinarySearch(item, arr)
	if result != wantedResult {
		t.Fatalf("Binary searched failed with the result of %v", result)
	}
}
func TestCrystal(t *testing.T) {
	arr := []bool{false, false, false, false, true, true, true, true, true}
	wantedResult := 4
	result := algorithms.FindTheFirstFloorThatBreaksTheCrystalBalls(arr)
	if result != wantedResult {
		t.Fatalf("Crystal search failed with the result of %v", result)
	}
}
func TestBubbleSort(t *testing.T) {
	arr := []int{1, 5, 3, 2, 6, 7}
	wantedResult := []int{1, 2, 3, 5, 6, 7}
	algorithms.BubbleSort(arr)
	if !helper.CompareTwoUniDimensionalArrays(arr, wantedResult) {
		t.Fatalf("Bubble sort failed!")
	}
}
func TestRecursion(t *testing.T) {
	maze := []string{
		"xxxxxxxxxx x",
		"x        x x",
		"x        x x",
		"x xxxxxxxx x",
		"x          x",
		"x xxxxxxxxxx",
	}
	path := algorithms.MazeSolver(maze, 'x', algorithms.Point{10, 0}, algorithms.Point{1, 5})
	fmt.Printf("%v\n", path)
}
func TestQuickSort(t *testing.T) {
	arr := []int{1, 5, 3, 2, 6, 7}
	wantedResult := []int{1, 2, 3, 5, 6, 7}
	algorithms.QuickSort(arr)
	if !helper.CompareTwoUniDimensionalArrays(arr, wantedResult) {
		t.Fatalf("Bubble sort failed!")
	}
}
