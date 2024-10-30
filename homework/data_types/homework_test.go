package main

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

func ToLittleEndian[T uint64 | uint16 | uint32](number T) T {
	size := int(unsafe.Sizeof(number))
	pointer := unsafe.Pointer(&number)

	reversedBytes := make([]byte, 0, size)
	for i := size - 1; i >= 0; i-- {
		nextPointer := unsafe.Add(pointer, i)
		reversedBytes = append(reversedBytes, *(*byte)(nextPointer))
	}
	return *(*T)(unsafe.Pointer(unsafe.SliceData(reversedBytes)))
}

func TestSerializationProperties(t *testing.T) {
	tests := map[string]struct {
		number uint32
		result uint32
	}{
		"test case #1": {
			number: 0x00000000,
			result: 0x00000000,
		},
		"test case #2": {
			number: 0xFFFFFFFF,
			result: 0xFFFFFFFF,
		},
		"test case #3": {
			number: 0x00FF00FF,
			result: 0xFF00FF00,
		},
		"test case #4": {
			number: 0x0000FFFF,
			result: 0xFFFF0000,
		},
		"test case #5": {
			number: 0x01020304,
			result: 0x04030201,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}

func TestSerializationProperties_uint64(t *testing.T) {
	tests := map[string]struct {
		number uint64
		result uint64
	}{
		"test case #1": {
			number: 0x00000000,
			result: 0x00000000,
		},
		"test case #2": {
			number: 0xFFFFFFFF,
			result: 0xffffffff00000000,
		},
		"test case #3": {
			number: 0x00FF00FF,
			result: 0xff00ff0000000000,
		},
		"test case #4": {
			number: 0x0000FFFF,
			result: 0xffff000000000000,
		},
		"test case #5": {
			number: 0x01020304,
			result: 0x403020100000000,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}
