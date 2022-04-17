package file

import (
	"minmaxfilereplacer/domain"
	"testing"
)

func TestMinMaxSwapPairSelector(t *testing.T) {
	ds := []*domain.FileInfo{
		{
			OrderValue: 0,
		},
		{
			OrderValue: 1,
		},
		{
			OrderValue: 3,
		},
	}
	min, max := MinMaxPairSelector(ds)
	if min.OrderValue != 0 || max.OrderValue != 3 {
		t.Fatalf("incorrect min-max pair selection")
	}
}
