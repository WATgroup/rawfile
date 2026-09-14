// SPDX-FileCopyrightText: © 2026 W-A-T EU Operations Oü
// SPDX-License-Identifier: MPL-2.0
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

// Package rawfile provides a very-barebones (just io.ReadWriteCloser) fd wrapper,
// for use with low-level packages requiring only the file descriptor
package rawfile

type sysFdType = int

type rawFile struct {
	fd  sysFdType
	sem uint32
}

//revive:disable:unexported-return  It's a private type on purpose...

// New wraps a provided Fd in a "rawFile"
// Recommended usage: from basedir·OpenRaw()
func New(fd int) (ret *rawFile) {
	return newFile(fd)
}

// Read implements io.Reader
func (f *rawFile) Read(buf []byte) (n int, err error) {
	return f.doRead(buf)
}

// Write implements io.Writer
func (f *rawFile) Write(buf []byte) (n int, err error) {
	return f.doWrite(buf)
}

// Close implements io.Closer
func (f *rawFile) Close() error {
	return f.doClose()
}

/////////////////////////////////////////////////////////////////////////////////

func (f *rawFile) Fd() int {
	return f.fd
}

// FromOpen is an alternate constructor, designed to be used from an Open() call
// BEWARE Close() races.. (resp. ownership)
//
//go:noinline
func FromOpen(obj any, err error) *rawFile {
	if nil != err {
		return nil
	}

	// Check if passed object is an *os.File or equivalent
	if fder, ok := obj.(interface{ Fd() uintptr }); ok {
		return newFile(int(fder.Fd())) // from *os.File and compatibles
	}
	// ...or our own basedir.OpenRaw()
	if fder, ok := obj.(interface{ Fd() int }); ok {
		return newFile(fder.Fd()) // from *os.File and compatibles
	}

	// ELSE, fail
	return nil
}
