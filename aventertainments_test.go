package opendmm

import (
  "testing"
)

func TestAventertainments(t *testing.T) {
  queries := []string {
    "SKYHD-001",
    "MCB3DBD-25",
    "lafbd-10",
  }
  assertSearchable(t, queries, aveSearch)
}
