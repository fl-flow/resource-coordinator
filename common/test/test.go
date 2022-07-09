package test

import (
  "testing"
)


func Assert(t *testing.T, a interface{}, b interface{}, description ...string) {
  if a != b {
    t.Errorf("%v: %v != %v", description, a, b)
  }
}
