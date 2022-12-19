package main

import (
	"testing"
)

func TestBaseCase(t *testing.T) {
	point1 := [3]int{1, 1, 1}
	point2 := [3]int{2, 1, 1}
	points := [][3]int{point1, point2}
	int := Solution(points)
	if int != 10 {
		t.Fatal("wrong solution for base test")
	}
}

func TestGetSides(t *testing.T) {
	point1 := [3]int{1, 1, 1}
	s := GetAllSides(point1)
	if len(s) != 6 {
		t.Fatal("wrong number of sides")
	}
	point2 := [3]int{2, 1, 1}
	var point2Found bool
	for _, side := range(s) {
		if side == point2 {
			point2Found = true
		}
	}
	if !point2Found {
		t.Fatal("point 2 not found")
	}
}
