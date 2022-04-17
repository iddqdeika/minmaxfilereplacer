package file

import (
	"minmaxfilereplacer/domain"
	"strconv"
	"strings"
)

// LogFileFilter фильтрует инфо о файлах, оставляя лишь нужные
// также обогощает инфо о файлах данными о порядке (число order)
func LogFileFilter(list []*domain.FileInfo) []*domain.FileInfo {

	res := make([]*domain.FileInfo, 0)
	for _, f := range list {
		fi, ok := tryConvert(f)
		if ok {
			res = append(res, fi)
		}
	}
	return res
}

func tryConvert(f *domain.FileInfo) (*domain.FileInfo, bool) {
	if !strings.HasSuffix(f.Name, ".log") {
		return nil, false
	}
	v := trimSuffix(f.Name)
	iv, err := strconv.Atoi(v)
	if err != nil {
		return nil, false
	}
	f.OrderValue = iv
	return f, true
}

func trimSuffix(f string) string {
	runes := []rune(f)
	return string(runes[:len(runes)-4])
}
