// Code generated by command: go run blend_asm.go -pkg mem -out ../mem/blend_amd64.s -stubs ../mem/blend_amd64.go. DO NOT EDIT.

//go:build !purego

package mem

// Blend copies the one-bits of src to dst, returning the number of bytes written.
func Blend(dst []byte, src []byte) int
