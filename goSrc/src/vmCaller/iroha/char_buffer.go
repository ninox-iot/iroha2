package iroha

// #cgo CFLAGS: -I ../../../../irohad
// #cgo LDFLAGS: -Wl,-unresolved-symbols=ignore-all
// #include <stdlib.h>
// #include "ametsuchi/impl/burrow_storage.h"
import "C"
import (
	"encoding/hex"
	"unsafe"

	"github.com/hyperledger/burrow/binary"
)

func MakeIrohaCharBuffer(data string) *C.struct_Iroha_CharBuffer {
	return &C.struct_Iroha_CharBuffer{
		data: C.CString(data),
		size: C.ulonglong(len(data)),
	}
}

func (buf *C.struct_Iroha_CharBuffer) free() {
	C.free(unsafe.Pointer(buf.data))
}

type Iroha_CharBufferArray_Wrapper struct {
	charBuffers []C.struct_Iroha_CharBuffer
	cArray      *C.struct_Iroha_CharBufferArray
}

func MakeIrohaCharBufferArray(data []binary.Word256) *Iroha_CharBufferArray_Wrapper {
	array := make([]C.struct_Iroha_CharBuffer, len(data))
	for idx, el := range data {
		array[idx] = *MakeIrohaCharBuffer(hex.EncodeToString(el.Bytes()))
	}
	return &Iroha_CharBufferArray_Wrapper{
		array,
		&C.struct_Iroha_CharBufferArray{
			data: &array[0],
			size: C.ulonglong(len(data)),
		},
	}
}

func (arr *Iroha_CharBufferArray_Wrapper) free() {
	for _, el := range arr.charBuffers {
		C.free(unsafe.Pointer(el.data))
	}
}
