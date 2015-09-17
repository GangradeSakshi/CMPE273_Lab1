package main

import "testing"

func fibTest(t * testing.T) {
 var v int
 v = fibonacci()
 if v != 55 {
  t.Error("Expected 55, got ", v)
 }

}
