package main

import "os"

// START OMIT

// Sizer describes the Size() method that gets the
// total size of an item.
type Sizer interface {
	Size() int64
}

func Fits(capacity int64, v Sizer) bool {
	return capacity > v.Size()
}

func IsEmailable(v Sizer) bool {
	return 1<<20 > v.Size()
}

// Size gets the size of a File.
func (f *File) Size() int64 {
	return f.info.Size()
}

// END OMIT

// TWO OMIT

type Sizers []Sizer

func (s Sizers) Size() int64 {
	var total int64
	for _, sizer := range s {
		total += sizer.Size()
	}
	return total
}

// TWO END OMIT

// THREE OMIT

type SizeFunc func() int64

func (s SizeFunc) Size() int64 {
	return s()
}

type Size int64

func (s Size) Size() int64 {
	return int64(s)
}

// THREE END OMIT

type File struct {
	info os.FileInfo
}
