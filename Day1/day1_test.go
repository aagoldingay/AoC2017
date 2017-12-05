package main

import (
	"fmt"
	"testing"
)

var testsp1 = []string {"1122", "1234", "1111", "91212129"}
var resultsp1 = []int {3, 0, 4, 9}

var testsp2 = []string {"1212", "1221", "123425", "123123", "12131415"}
var resultsp2 = []int {6, 0, 4, 12, 4}

func Test_ReviewSequenceP1(t *testing.T) {
	fmt.Println("TEST PART 1")
	for i, s := range testsp1 {
		fmt.Println("testing : ", s)
		n := ReviewSequenceP1(s)
		fmt.Println("result :", n)
		if n != resultsp1[i] {
			t.Errorf("ERR - expected: %d, got: %d", resultsp1[i], n)
		}
	}
}

func Test_ReviewSequenceP2(t *testing.T) {
	fmt.Println("TEST PART 2")
	for i, s := range testsp2 {
		fmt.Println("testing : ", s)
		n := ReviewSequenceP2(s)
		fmt.Println("result : ", n)
		if n != resultsp2[i] {
			t.Errorf("ERR - expected: %d, got %d", resultsp2[i], n)
		}
	}
}