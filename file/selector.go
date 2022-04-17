package file

import "minmaxfilereplacer/domain"

func MinMaxPairSelector(list []*domain.FileInfo) (*domain.FileInfo, *domain.FileInfo) {
	var min *domain.FileInfo
	var max *domain.FileInfo

	for _, f := range list {
		if min == nil || min.OrderValue > f.OrderValue {
			min = f
		}
		if max == nil || max.OrderValue < f.OrderValue {
			max = f
		}
	}
	return min, max
}
