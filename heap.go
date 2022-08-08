// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd

import "io"

// Heap provides access to metadata heaps as defined in §II.24.2.
type Heap struct {
	// Embed ReaderAt for ReadAt method.
	// Do not embed SectionReader directly
	// to avoid having Read and Seek.
	// If a client wants Read and Seek it must use
	// Open() to avoid fighting over the seek offset
	// with other clients.
	io.ReaderAt
	Size uint32
	sr   *io.SectionReader
	name string
}

// Data reads and returns the contents of the stream s.
func (s *Heap) Data() ([]byte, error) {
	dat := make([]byte, s.sr.Size())
	n, err := s.sr.ReadAt(dat, 0)
	if n == len(dat) {
		err = nil
	}
	return dat[0:n], err
}

// Open returns a new ReadSeeker reading the stream s.
func (s *Heap) Open() io.ReadSeeker {
	return io.NewSectionReader(s.sr, 0, 1<<63-1)
}
