package file

import (
	"fmt"
	"io"
	"minmaxfilereplacer/domain"
	"os"
	"sync"
)

func CopyFiles(fi1, fi2 *domain.FileInfo) error {
	err := appendSourceContentToTarget(fi1, fi2)
	if err != nil {
		return err
	}
	err = appendSourceContentToTarget(fi2, fi1)
	if err != nil {
		return err
	}
	return nil
}

func appendSourceContentToTarget(source *domain.FileInfo, target *domain.FileInfo) error {
	s, err := os.OpenFile(source.Name, os.O_RDONLY, os.ModeAppend)
	if err != nil {
		return fmt.Errorf("cant open source file %v: %v", source.Name, err)
	}
	defer s.Close()
	t, err := os.OpenFile(target.Name, os.O_RDWR|os.O_APPEND, os.ModeAppend)
	if err != nil {
		return fmt.Errorf("cant open target file %v: %v", target.Name, err)
	}
	defer t.Close()

	err = copyContentWithLimit(s, t, source.Size)
	if err != nil {
		return err
	}
	return nil
}

func copyContentWithLimit(s, t *os.File, maxBytes int64) error {
	ch := make(chan byte)
	var writeErr error

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup, s *os.File, ch <-chan byte) {
		defer wg.Done()
		writeErr = write(t, ch)
	}(wg, t, ch)

	var readErr error
	readErr = read(s, int(maxBytes), ch)

	wg.Wait()

	if readErr != nil {
		return readErr
	}

	if writeErr != nil {
		return writeErr
	}

	return t.Sync()
}

func read(s *os.File, maxBytes int, ch chan<- byte) error {
	defer close(ch)
	arr := make([]byte, 1, 1)
	var read int
	for {
		n, err := s.Read(arr)
		if err == io.EOF {
			return nil
		}
		ch <- arr[0]
		read += n
		if read >= maxBytes {
			return err
		}
	}
}

func write(f *os.File, ch <-chan byte) error {
	arr := make([]byte, 1, 1)
	for b := range ch {
		arr[0] = b
		_, err := f.Write(arr)
		if err != nil {
			return err
		}
	}
	return nil
}
