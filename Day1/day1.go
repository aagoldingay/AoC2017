package main

import ()

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
  half, numAhead, result := len(s)/2, 0, 0

  for i, c := range s {
    numAhead = int(s[(i + half) % len(s)])
    if int(c) == numAhead {
      result += numAhead % 48
    }
  }
  return result
}

func main() {}