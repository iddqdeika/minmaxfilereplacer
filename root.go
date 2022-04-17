package main

import (
	"fmt"
	"github.com/iddqdeika/rrr/helpful"
	"minmaxfilereplacer/domain"
	"minmaxfilereplacer/file"
)

func newRoot() *root {
	return &root{}
}

type root struct {
	p  domain.FileProvider
	f  domain.FileFilter
	ps domain.SwapPairSelector
	fs domain.FileSwapper
}

func (r *root) Register() error {
	cfg, err := helpful.NewJsonCfg("cfg.json")
	if err != nil {
		return err
	}
	path, err := cfg.GetString("path")
	if err != nil {
		return err
	}
	r.p, err = file.NewProvider(path)
	if err != nil {
		return err
	}
	r.f = file.LogFileFilter
	r.ps = file.MinMaxPairSelector
	r.fs = file.CopyFiles

	return nil
}

func (r *root) Resolve() error {
	list, err := r.p.GetFileList()
	if err != nil {
		return err
	}
	list = r.f(list)
	if err != nil {
		return fmt.Errorf("cant get files list: %v", err)
	}
	l := len(list)
	if l == 0 {
		return domain.ErrNoRelevantFiles
	}
	if l == 1 {
		return domain.ErrOnlyOneRelevantFile
	}
	return r.fs(r.ps(list))

}

func (r *root) Release() {

}
