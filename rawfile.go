// SPDX-FileCopyrightText: © 2026 W-A-T EU Operations Oü
// SPDX-License-Identifier: MPL-2.0
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

package rawfile

type SysFdType = int

type rawFile struct {
	fd  SysFdType
	sem uint32
}

// New wraps a provided Fd in a "rawFile"
// Recommended usage: from basedir·OpenRaw()
func New(fd int) (ret *rawFile) {
	return newFile(fd)
}

func (f *rawFile) Read(buf []byte) (n int, err error) {
	return f.doRead(buf)
}

func (f *rawFile) Write(buf []byte) (n int, err error) {
	return f.doWrite(buf)
}

func (f *rawFile) Close() error {
	return f.doClose()
}

func (f *rawFile) Fd() int {
	return f.fd
}
