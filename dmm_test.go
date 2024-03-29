package opendmm

import (
	"testing"
)

func TestDmm(t *testing.T) {
	queries := []string{
		"MIDE-029",
		"mide-029",
		"XV-100",
		"XV-1001",
		"IPZ687",
		"WPE-037",
	}
	assertSearchable(t, queries, dmmSearch)
	blackhole := []string{
		"MCB3DBD-25",
		"S2MBD-048",
	}
	assertUnsearchable(t, blackhole, dmmSearch)
}
