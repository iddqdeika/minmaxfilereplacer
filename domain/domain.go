package domain

import (
	"errors"
)

type FileProvider interface {
	GetFileList() ([]*FileInfo, error)
}

type FileFilter func(list []*FileInfo) []*FileInfo

type FileInfo struct {
	OrderValue int
	Name       string
	Path       string
	Size       int64
}

type SwapPairSelector func(list []*FileInfo) (*FileInfo, *FileInfo)

type FileSwapper func(f1, f2 *FileInfo) error

var ErrPathIsNotDirectory = errors.New("given path is not directory")

var ErrNoRelevantFiles = errors.New("there is no relevant files in given directory")

var ErrOnlyOneRelevantFile = errors.New("there is only one relevant file")
