package file

import (
	"minmaxfilereplacer/domain"
	"strconv"
	"strings"
	"testing"
)

func TestLogFileFilter(t *testing.T) {
	ds := []*domain.FileInfo{
		{
			Name: "1.txt",
		},
		{
			Name: "2.log",
		},
		{
			Name: "3.log",
		},
		{
			Name: "2.txt",
		},
	}
	res := LogFileFilter(ds)
	if len(res) != 2 {
		t.Fatalf("incorrect filtered file list size")
	}
	for _, fi := range res {
		if !strings.HasPrefix(fi.Name, strconv.Itoa(fi.OrderValue)) {
			t.Fatalf("incorrect order (%v) for file with name %v", fi.OrderValue, fi.Name)
		}
	}
}
