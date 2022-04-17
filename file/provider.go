package file

import (
	"fmt"
	"minmaxfilereplacer/domain"
	"os"
)

func NewProvider(p string) (*fileProvider, error) {
	if !validPath(p) {
		return nil, fmt.Errorf("given path (\"%v\") is incorrect", p)
	}
	return &fileProvider{
		fp: p,
	}, nil
}

func validPath(p string) bool {
	if p == "" {
		return false
	}
	return true
}

type fileProvider struct {
	fp string
}

func (p *fileProvider) GetFileList() ([]*domain.FileInfo, error) {
	f, err := os.Open(p.fp)
	if err != nil {
		return nil, fmt.Errorf("cant open path %v: %v", p.fp, err)
	}
	s, err := f.Stat()
	if err != nil {
		return nil, fmt.Errorf("cant get path %v stats: %v", p.fp, err)
	}
	if !s.IsDir() {
		return nil, domain.ErrPathIsNotDirectory
	}
	names, err := f.Readdir(0)
	if err != nil {
		return nil, err
	}
	return convertToFileInfo(p.fp, names), nil
}

func convertToFileInfo(path string, names []os.FileInfo) []*domain.FileInfo {
	var res []*domain.FileInfo
	for _, fi := range names {
		if !fi.IsDir() {
			res = append(res, fileInfo(path, fi))
		}
	}
	return res
}

func fileInfo(path string, fi os.FileInfo) *domain.FileInfo {
	return &domain.FileInfo{
		Name: fi.Name(),
		Path: path + string(os.PathSeparator) + fi.Name(),
		Size: fi.Size(),
	}
}
