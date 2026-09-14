// SPDX-FileCopyrightText: © 2026 W-A-T EU Operations Oü
// SPDX-License-Identifier: MPL-2.0
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

//*** this is Linux-ONLY

package rawfile

import (
	"runtime"
	"sync/atomic"
	"syscall"
)

func newFile(fd int) (ret *rawFile) {
	syscall.CloseOnExec(fd)
	ret = &rawFile{
		fd:  fd,
		sem: 1,
	}
	// link finalizer
	runtime.SetFinalizer(ret, (*rawFile).close)
	return
}

func (f *rawFile) doRead(buf []byte) (int, error) {
	addRef(&f.sem)
	defer decRef(&f.sem)
	return syscall.Read(f.fd, buf)
}

func (f *rawFile) doWrite(buf []byte) (int, error) {
	addRef(&f.sem)
	defer decRef(&f.sem)
	return syscall.Write(f.fd, buf)
}

func (f *rawFile) doClose() (ret error) {
	if atomic.CompareAndSwapUint32(&f.sem, 1, 0) {
		// last one wins (+closes)
		ret = syscall.Close(f.fd)
		f.fd = -1 // not ours anymore
		if nil != ret {
			return ret
		}
		// unlink finalizer
		runtime.SetFinalizer(f, nil)
		return
	}
	decRef(&f.sem)
	return nil
}

func (f *rawFile) close() error {
	syscall.Close(f.fd)
	f.fd = -1
	return nil
}

func addRef(v *uint32) uint32 {
	return atomic.AddUint32(v, 1)
}

func decRef(v *uint32) uint32 {
	return atomic.AddUint32(v, ^uint32(0))
}
