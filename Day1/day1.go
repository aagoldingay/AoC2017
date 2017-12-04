package main

import "fmt"

func ReviewSequenceP1(s string) int {
  lastNum, result := 0, 0
  for i, c := range s {
    if int(c)%48 == lastNum {
      result += lastNum
    }
    lastNum = (int(c) % 48)
    if ((i == len(s)-1) && ((int(c)%48) == int(s[0])%48)) {
      result += int(s[0])%48
    }
  }
  return result
}

func ReviewSequenceP2(s string) int {
  half, result := len(s)/2, 0
  //figure out mod
  result = half
  return result
}

func main() {
  fmt.Println("placeholder")
}