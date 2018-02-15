// WARNING: This file has automatically been generated on Fri, 16 Feb 2018 00:12:48 EET.
// By https://git.io/c-for-go. DO NOT EDIT.

package imgui

/*
#include "cimgui/cimgui/cimgui.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"runtime"
	"sync"
	"unsafe"
)

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

func (x TextEditCallback) PassRef() (ref *C.ImGuiTextEditCallback, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if textEditCallback4D1820F7Func == nil {
		textEditCallback4D1820F7Func = x
	}
	return (*C.ImGuiTextEditCallback)(C.ImGuiTextEditCallback_4d1820f7), nil
}

func (x TextEditCallback) PassValue() (ref C.ImGuiTextEditCallback, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if textEditCallback4D1820F7Func == nil {
		textEditCallback4D1820F7Func = x
	}
	return (C.ImGuiTextEditCallback)(C.ImGuiTextEditCallback_4d1820f7), nil
}

func NewTextEditCallbackRef(ref unsafe.Pointer) *TextEditCallback {
	return (*TextEditCallback)(ref)
}

//export textEditCallback4D1820F7
func textEditCallback4D1820F7(cdata *C.struct_ImGuiTextEditCallbackData) C.int {
	if textEditCallback4D1820F7Func != nil {
		data4d1820f7 := NewTextEditCallbackDataRef(unsafe.Pointer(cdata))
		ret4d1820f7 := textEditCallback4D1820F7Func(data4d1820f7)
		ret, _ := (C.int)(ret4d1820f7), cgoAllocsUnknown
		return ret
	}
	panic("callback func has not been set (race?)")
}

var textEditCallback4D1820F7Func TextEditCallback

func (x SizeConstraintCallback) PassRef() (ref *C.ImGuiSizeConstraintCallback, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if sizeConstraintCallbackF2F47F64Func == nil {
		sizeConstraintCallbackF2F47F64Func = x
	}
	return (*C.ImGuiSizeConstraintCallback)(C.ImGuiSizeConstraintCallback_f2f47f64), nil
}

func (x SizeConstraintCallback) PassValue() (ref C.ImGuiSizeConstraintCallback, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if sizeConstraintCallbackF2F47F64Func == nil {
		sizeConstraintCallbackF2F47F64Func = x
	}
	return (C.ImGuiSizeConstraintCallback)(C.ImGuiSizeConstraintCallback_f2f47f64), nil
}

func NewSizeConstraintCallbackRef(ref unsafe.Pointer) *SizeConstraintCallback {
	return (*SizeConstraintCallback)(ref)
}

//export sizeConstraintCallbackF2F47F64
func sizeConstraintCallbackF2F47F64(cdata *C.struct_ImGuiSizeConstraintCallbackData) {
	if sizeConstraintCallbackF2F47F64Func != nil {
		dataf2f47f64 := NewSizeConstraintCallbackDataRef(unsafe.Pointer(cdata))
		sizeConstraintCallbackF2F47F64Func(dataf2f47f64)
		return
	}
	panic("callback func has not been set (race?)")
}

var sizeConstraintCallbackF2F47F64Func SizeConstraintCallback

func (x DrawCallback) PassRef() (ref *C.ImDrawCallback, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if drawCallbackADEAA4A0Func == nil {
		drawCallbackADEAA4A0Func = x
	}
	return (*C.ImDrawCallback)(C.ImDrawCallback_adeaa4a0), nil
}

func (x DrawCallback) PassValue() (ref C.ImDrawCallback, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if drawCallbackADEAA4A0Func == nil {
		drawCallbackADEAA4A0Func = x
	}
	return (C.ImDrawCallback)(C.ImDrawCallback_adeaa4a0), nil
}

func NewDrawCallbackRef(ref unsafe.Pointer) *DrawCallback {
	return (*DrawCallback)(ref)
}

//export drawCallbackADEAA4A0
func drawCallbackADEAA4A0(cparentList *C.struct_ImDrawList, ccmd *C.struct_ImDrawCmd) {
	if drawCallbackADEAA4A0Func != nil {
		parentListadeaa4a0 := (*DrawList)(unsafe.Pointer(cparentList))
		cmdadeaa4a0 := NewDrawCmdRef(unsafe.Pointer(ccmd))
		drawCallbackADEAA4A0Func(parentListadeaa4a0, cmdadeaa4a0)
		return
	}
	panic("callback func has not been set (race?)")
}

var drawCallbackADEAA4A0Func DrawCallback

// allocStructImDrawCmdMemory allocates memory for type C.struct_ImDrawCmd in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStructImDrawCmdMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStructImDrawCmdValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStructImDrawCmdValue = unsafe.Sizeof([1]C.struct_ImDrawCmd{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *DrawCmd) Ref() *C.struct_ImDrawCmd {
	if x == nil {
		return nil
	}
	return x.ref111e6f2b
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *DrawCmd) Free() {
	if x != nil && x.allocs111e6f2b != nil {
		x.allocs111e6f2b.(*cgoAllocMap).Free()
		x.ref111e6f2b = nil
	}
}

// NewDrawCmdRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewDrawCmdRef(ref unsafe.Pointer) *DrawCmd {
	if ref == nil {
		return nil
	}
	obj := new(DrawCmd)
	obj.ref111e6f2b = (*C.struct_ImDrawCmd)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *DrawCmd) PassRef() (*C.struct_ImDrawCmd, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref111e6f2b != nil {
		return x.ref111e6f2b, nil
	}
	mem111e6f2b := allocStructImDrawCmdMemory(1)
	ref111e6f2b := (*C.struct_ImDrawCmd)(mem111e6f2b)
	allocs111e6f2b := new(cgoAllocMap)
	allocs111e6f2b.Add(mem111e6f2b)

	var cElemCount_allocs *cgoAllocMap
	ref111e6f2b.ElemCount, cElemCount_allocs = (C.uint)(x.ElemCount), cgoAllocsUnknown
	allocs111e6f2b.Borrow(cElemCount_allocs)

	var cClipRect_allocs *cgoAllocMap
	ref111e6f2b.ClipRect, cClipRect_allocs = x.ClipRect.PassValue()
	allocs111e6f2b.Borrow(cClipRect_allocs)

	var cTextureId_allocs *cgoAllocMap
	ref111e6f2b.TextureId, cTextureId_allocs = *(**C.ImTextureID)(unsafe.Pointer(&x.TextureId)), cgoAllocsUnknown
	allocs111e6f2b.Borrow(cTextureId_allocs)

	var cUserCallback_allocs *cgoAllocMap
	ref111e6f2b.UserCallback, cUserCallback_allocs = x.UserCallback.PassValue()
	allocs111e6f2b.Borrow(cUserCallback_allocs)

	var cUserCallbackData_allocs *cgoAllocMap
	ref111e6f2b.UserCallbackData, cUserCallbackData_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.UserCallbackData)), cgoAllocsUnknown
	allocs111e6f2b.Borrow(cUserCallbackData_allocs)

	x.ref111e6f2b = ref111e6f2b
	x.allocs111e6f2b = allocs111e6f2b
	return ref111e6f2b, allocs111e6f2b

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x DrawCmd) PassValue() (C.struct_ImDrawCmd, *cgoAllocMap) {
	if x.ref111e6f2b != nil {
		return *x.ref111e6f2b, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *DrawCmd) Deref() {
	if x.ref111e6f2b == nil {
		return
	}
	x.ElemCount = (uint32)(x.ref111e6f2b.ElemCount)
	x.ClipRect = *NewVec4Ref(unsafe.Pointer(&x.ref111e6f2b.ClipRect))
	x.TextureId = (*TextureID)(unsafe.Pointer(x.ref111e6f2b.TextureId))
	x.UserCallback = *NewDrawCallbackRef(unsafe.Pointer(&x.ref111e6f2b.UserCallback))
	x.UserCallbackData = (unsafe.Pointer)(unsafe.Pointer(x.ref111e6f2b.UserCallbackData))
}

// allocStructImDrawDataMemory allocates memory for type C.struct_ImDrawData in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStructImDrawDataMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStructImDrawDataValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStructImDrawDataValue = unsafe.Sizeof([1]C.struct_ImDrawData{})

type sliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

// allocPStructImDrawListMemory allocates memory for type *C.struct_ImDrawList in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPStructImDrawListMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPStructImDrawListValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPStructImDrawListValue = unsafe.Sizeof([1]*C.struct_ImDrawList{})

const sizeOfPtr = unsafe.Sizeof(&struct{}{})

// unpackSSDrawList transforms a sliced Go data structure into plain C format.
func unpackSSDrawList(x [][]DrawList) (unpacked **C.struct_ImDrawList, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.struct_ImDrawList) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPStructImDrawListMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.struct_ImDrawList)(unsafe.Pointer(h0))
	for i0 := range x {
		h := (*sliceHeader)(unsafe.Pointer(&x[i0]))
		v0[i0] = (*C.struct_ImDrawList)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.struct_ImDrawList)(unsafe.Pointer(h.Data))
	return
}

// packSSDrawList reads sliced Go data structure out from plain C format.
func packSSDrawList(v [][]DrawList, ptr0 **C.struct_ImDrawList) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.struct_ImDrawList)(unsafe.Pointer(ptr0)))[i0]
		hxfc4425b := (*sliceHeader)(unsafe.Pointer(&v[i0]))
		hxfc4425b.Data = uintptr(unsafe.Pointer(ptr1))
		hxfc4425b.Cap = 0x7fffffff
		// hxfc4425b.Len = ?
	}
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *DrawData) Ref() *C.struct_ImDrawData {
	if x == nil {
		return nil
	}
	return x.ref9a158ae0
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *DrawData) Free() {
	if x != nil && x.allocs9a158ae0 != nil {
		x.allocs9a158ae0.(*cgoAllocMap).Free()
		x.ref9a158ae0 = nil
	}
}

// NewDrawDataRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewDrawDataRef(ref unsafe.Pointer) *DrawData {
	if ref == nil {
		return nil
	}
	obj := new(DrawData)
	obj.ref9a158ae0 = (*C.struct_ImDrawData)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *DrawData) PassRef() (*C.struct_ImDrawData, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref9a158ae0 != nil {
		return x.ref9a158ae0, nil
	}
	mem9a158ae0 := allocStructImDrawDataMemory(1)
	ref9a158ae0 := (*C.struct_ImDrawData)(mem9a158ae0)
	allocs9a158ae0 := new(cgoAllocMap)
	allocs9a158ae0.Add(mem9a158ae0)

	var cValid_allocs *cgoAllocMap
	ref9a158ae0.Valid, cValid_allocs = (C._Bool)(x.Valid), cgoAllocsUnknown
	allocs9a158ae0.Borrow(cValid_allocs)

	var cCmdLists_allocs *cgoAllocMap
	ref9a158ae0.CmdLists, cCmdLists_allocs = unpackSSDrawList(x.CmdLists)
	allocs9a158ae0.Borrow(cCmdLists_allocs)

	var cCmdListsCount_allocs *cgoAllocMap
	ref9a158ae0.CmdListsCount, cCmdListsCount_allocs = (C.int)(x.CmdListsCount), cgoAllocsUnknown
	allocs9a158ae0.Borrow(cCmdListsCount_allocs)

	var cTotalVtxCount_allocs *cgoAllocMap
	ref9a158ae0.TotalVtxCount, cTotalVtxCount_allocs = (C.int)(x.TotalVtxCount), cgoAllocsUnknown
	allocs9a158ae0.Borrow(cTotalVtxCount_allocs)

	var cTotalIdxCount_allocs *cgoAllocMap
	ref9a158ae0.TotalIdxCount, cTotalIdxCount_allocs = (C.int)(x.TotalIdxCount), cgoAllocsUnknown
	allocs9a158ae0.Borrow(cTotalIdxCount_allocs)

	x.ref9a158ae0 = ref9a158ae0
	x.allocs9a158ae0 = allocs9a158ae0
	return ref9a158ae0, allocs9a158ae0

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x DrawData) PassValue() (C.struct_ImDrawData, *cgoAllocMap) {
	if x.ref9a158ae0 != nil {
		return *x.ref9a158ae0, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *DrawData) Deref() {
	if x.ref9a158ae0 == nil {
		return
	}
	x.Valid = (bool)(x.ref9a158ae0.Valid)
	packSSDrawList(x.CmdLists, x.ref9a158ae0.CmdLists)
	x.CmdListsCount = (int32)(x.ref9a158ae0.CmdListsCount)
	x.TotalVtxCount = (int32)(x.ref9a158ae0.TotalVtxCount)
	x.TotalIdxCount = (int32)(x.ref9a158ae0.TotalIdxCount)
}

// Ref returns a reference to C object as it is.
func (x *DrawList) Ref() *C.struct_ImDrawList {
	if x == nil {
		return nil
	}
	return (*C.struct_ImDrawList)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *DrawList) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewDrawListRef converts the C object reference into a raw struct reference without wrapping.
func NewDrawListRef(ref unsafe.Pointer) *DrawList {
	return (*DrawList)(ref)
}

// NewDrawList allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewDrawList() *DrawList {
	return (*DrawList)(allocStructImDrawListMemory(1))
}

// allocStructImDrawListMemory allocates memory for type C.struct_ImDrawList in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStructImDrawListMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStructImDrawListValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStructImDrawListValue = unsafe.Sizeof([1]C.struct_ImDrawList{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *DrawList) PassRef() *C.struct_ImDrawList {
	if x == nil {
		x = (*DrawList)(allocStructImDrawListMemory(1))
	}
	return (*C.struct_ImDrawList)(unsafe.Pointer(x))
}

// allocStructImDrawVertMemory allocates memory for type C.struct_ImDrawVert in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStructImDrawVertMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStructImDrawVertValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStructImDrawVertValue = unsafe.Sizeof([1]C.struct_ImDrawVert{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *DrawVert) Ref() *C.struct_ImDrawVert {
	if x == nil {
		return nil
	}
	return x.ref5c8bfe45
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *DrawVert) Free() {
	if x != nil && x.allocs5c8bfe45 != nil {
		x.allocs5c8bfe45.(*cgoAllocMap).Free()
		x.ref5c8bfe45 = nil
	}
}

// NewDrawVertRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewDrawVertRef(ref unsafe.Pointer) *DrawVert {
	if ref == nil {
		return nil
	}
	obj := new(DrawVert)
	obj.ref5c8bfe45 = (*C.struct_ImDrawVert)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *DrawVert) PassRef() (*C.struct_ImDrawVert, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref5c8bfe45 != nil {
		return x.ref5c8bfe45, nil
	}
	mem5c8bfe45 := allocStructImDrawVertMemory(1)
	ref5c8bfe45 := (*C.struct_ImDrawVert)(mem5c8bfe45)
	allocs5c8bfe45 := new(cgoAllocMap)
	allocs5c8bfe45.Add(mem5c8bfe45)

	var cpos_allocs *cgoAllocMap
	ref5c8bfe45.pos, cpos_allocs = x.Pos.PassValue()
	allocs5c8bfe45.Borrow(cpos_allocs)

	var cuv_allocs *cgoAllocMap
	ref5c8bfe45.uv, cuv_allocs = x.Uv.PassValue()
	allocs5c8bfe45.Borrow(cuv_allocs)

	var ccol_allocs *cgoAllocMap
	ref5c8bfe45.col, ccol_allocs = (C.ImU32)(x.Col), cgoAllocsUnknown
	allocs5c8bfe45.Borrow(ccol_allocs)

	x.ref5c8bfe45 = ref5c8bfe45
	x.allocs5c8bfe45 = allocs5c8bfe45
	return ref5c8bfe45, allocs5c8bfe45

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x DrawVert) PassValue() (C.struct_ImDrawVert, *cgoAllocMap) {
	if x.ref5c8bfe45 != nil {
		return *x.ref5c8bfe45, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *DrawVert) Deref() {
	if x.ref5c8bfe45 == nil {
		return
	}
	x.Pos = *NewVec2Ref(unsafe.Pointer(&x.ref5c8bfe45.pos))
	x.Uv = *NewVec2Ref(unsafe.Pointer(&x.ref5c8bfe45.uv))
	x.Col = (U32)(x.ref5c8bfe45.col)
}

// Ref returns a reference to C object as it is.
func (x *Font) Ref() *C.struct_ImFont {
	if x == nil {
		return nil
	}
	return (*C.struct_ImFont)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *Font) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewFontRef converts the C object reference into a raw struct reference without wrapping.
func NewFontRef(ref unsafe.Pointer) *Font {
	return (*Font)(ref)
}

// NewFont allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewFont() *Font {
	return (*Font)(allocStructImFontMemory(1))
}

// allocStructImFontMemory allocates memory for type C.struct_ImFont in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStructImFontMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStructImFontValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStructImFontValue = unsafe.Sizeof([1]C.struct_ImFont{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *Font) PassRef() *C.struct_ImFont {
	if x == nil {
		x = (*Font)(allocStructImFontMemory(1))
	}
	return (*C.struct_ImFont)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *FontAtlas) Ref() *C.struct_ImFontAtlas {
	if x == nil {
		return nil
	}
	return (*C.struct_ImFontAtlas)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *FontAtlas) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewFontAtlasRef converts the C object reference into a raw struct reference without wrapping.
func NewFontAtlasRef(ref unsafe.Pointer) *FontAtlas {
	return (*FontAtlas)(ref)
}

// NewFontAtlas allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewFontAtlas() *FontAtlas {
	return (*FontAtlas)(allocStructImFontAtlasMemory(1))
}

// allocStructImFontAtlasMemory allocates memory for type C.struct_ImFontAtlas in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStructImFontAtlasMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStructImFontAtlasValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStructImFontAtlasValue = unsafe.Sizeof([1]C.struct_ImFontAtlas{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *FontAtlas) PassRef() *C.struct_ImFontAtlas {
	if x == nil {
		x = (*FontAtlas)(allocStructImFontAtlasMemory(1))
	}
	return (*C.struct_ImFontAtlas)(unsafe.Pointer(x))
}

// allocStructImFontConfigMemory allocates memory for type C.struct_ImFontConfig in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStructImFontConfigMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStructImFontConfigValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStructImFontConfigValue = unsafe.Sizeof([1]C.struct_ImFontConfig{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *FontConfig) Ref() *C.struct_ImFontConfig {
	if x == nil {
		return nil
	}
	return x.ref5c4f67a0
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *FontConfig) Free() {
	if x != nil && x.allocs5c4f67a0 != nil {
		x.allocs5c4f67a0.(*cgoAllocMap).Free()
		x.ref5c4f67a0 = nil
	}
}

// NewFontConfigRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewFontConfigRef(ref unsafe.Pointer) *FontConfig {
	if ref == nil {
		return nil
	}
	obj := new(FontConfig)
	obj.ref5c4f67a0 = (*C.struct_ImFontConfig)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *FontConfig) PassRef() (*C.struct_ImFontConfig, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref5c4f67a0 != nil {
		return x.ref5c4f67a0, nil
	}
	mem5c4f67a0 := allocStructImFontConfigMemory(1)
	ref5c4f67a0 := (*C.struct_ImFontConfig)(mem5c4f67a0)
	allocs5c4f67a0 := new(cgoAllocMap)
	allocs5c4f67a0.Add(mem5c4f67a0)

	var cFontData_allocs *cgoAllocMap
	ref5c4f67a0.FontData, cFontData_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.FontData)), cgoAllocsUnknown
	allocs5c4f67a0.Borrow(cFontData_allocs)

	var cFontDataSize_allocs *cgoAllocMap
	ref5c4f67a0.FontDataSize, cFontDataSize_allocs = (C.int)(x.FontDataSize), cgoAllocsUnknown
	allocs5c4f67a0.Borrow(cFontDataSize_allocs)

	var cFontDataOwnedByAtlas_allocs *cgoAllocMap
	ref5c4f67a0.FontDataOwnedByAtlas, cFontDataOwnedByAtlas_allocs = (C._Bool)(x.FontDataOwnedByAtlas), cgoAllocsUnknown
	allocs5c4f67a0.Borrow(cFontDataOwnedByAtlas_allocs)

	var cFontNo_allocs *cgoAllocMap
	ref5c4f67a0.FontNo, cFontNo_allocs = (C.int)(x.FontNo), cgoAllocsUnknown
	allocs5c4f67a0.Borrow(cFontNo_allocs)

	var cSizePixels_allocs *cgoAllocMap
	ref5c4f67a0.SizePixels, cSizePixels_allocs = (C.float)(x.SizePixels), cgoAllocsUnknown
	allocs5c4f67a0.Borrow(cSizePixels_allocs)

	var cOversampleH_allocs *cgoAllocMap
	ref5c4f67a0.OversampleH, cOversampleH_allocs = (C.int)(x.OversampleH), cgoAllocsUnknown
	allocs5c4f67a0.Borrow(cOversampleH_allocs)

	var cOversampleV_allocs *cgoAllocMap
	ref5c4f67a0.OversampleV, cOversampleV_allocs = (C.int)(x.OversampleV), cgoAllocsUnknown
	allocs5c4f67a0.Borrow(cOversampleV_allocs)

	var cPixelSnapH_allocs *cgoAllocMap
	ref5c4f67a0.PixelSnapH, cPixelSnapH_allocs = (C._Bool)(x.PixelSnapH), cgoAllocsUnknown
	allocs5c4f67a0.Borrow(cPixelSnapH_allocs)

	var cGlyphExtraSpacing_allocs *cgoAllocMap
	ref5c4f67a0.GlyphExtraSpacing, cGlyphExtraSpacing_allocs = x.GlyphExtraSpacing.PassValue()
	allocs5c4f67a0.Borrow(cGlyphExtraSpacing_allocs)

	var cGlyphOffset_allocs *cgoAllocMap
	ref5c4f67a0.GlyphOffset, cGlyphOffset_allocs = x.GlyphOffset.PassValue()
	allocs5c4f67a0.Borrow(cGlyphOffset_allocs)

	var cGlyphRanges_allocs *cgoAllocMap
	ref5c4f67a0.GlyphRanges, cGlyphRanges_allocs = (*C.ImWchar)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.GlyphRanges)).Data)), cgoAllocsUnknown
	allocs5c4f67a0.Borrow(cGlyphRanges_allocs)

	var cMergeMode_allocs *cgoAllocMap
	ref5c4f67a0.MergeMode, cMergeMode_allocs = (C._Bool)(x.MergeMode), cgoAllocsUnknown
	allocs5c4f67a0.Borrow(cMergeMode_allocs)

	var cRasterizerFlags_allocs *cgoAllocMap
	ref5c4f67a0.RasterizerFlags, cRasterizerFlags_allocs = (C.uint)(x.RasterizerFlags), cgoAllocsUnknown
	allocs5c4f67a0.Borrow(cRasterizerFlags_allocs)

	var cRasterizerMultiply_allocs *cgoAllocMap
	ref5c4f67a0.RasterizerMultiply, cRasterizerMultiply_allocs = (C.float)(x.RasterizerMultiply), cgoAllocsUnknown
	allocs5c4f67a0.Borrow(cRasterizerMultiply_allocs)

	var cName_allocs *cgoAllocMap
	ref5c4f67a0.Name, cName_allocs = *(*[32]C.char)(unsafe.Pointer(&x.Name)), cgoAllocsUnknown
	allocs5c4f67a0.Borrow(cName_allocs)

	var cDstFont_allocs *cgoAllocMap
	ref5c4f67a0.DstFont, cDstFont_allocs = (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.DstFont)).Data)), cgoAllocsUnknown
	allocs5c4f67a0.Borrow(cDstFont_allocs)

	x.ref5c4f67a0 = ref5c4f67a0
	x.allocs5c4f67a0 = allocs5c4f67a0
	return ref5c4f67a0, allocs5c4f67a0

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x FontConfig) PassValue() (C.struct_ImFontConfig, *cgoAllocMap) {
	if x.ref5c4f67a0 != nil {
		return *x.ref5c4f67a0, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *FontConfig) Deref() {
	if x.ref5c4f67a0 == nil {
		return
	}
	x.FontData = (unsafe.Pointer)(unsafe.Pointer(x.ref5c4f67a0.FontData))
	x.FontDataSize = (int32)(x.ref5c4f67a0.FontDataSize)
	x.FontDataOwnedByAtlas = (bool)(x.ref5c4f67a0.FontDataOwnedByAtlas)
	x.FontNo = (int32)(x.ref5c4f67a0.FontNo)
	x.SizePixels = (float32)(x.ref5c4f67a0.SizePixels)
	x.OversampleH = (int32)(x.ref5c4f67a0.OversampleH)
	x.OversampleV = (int32)(x.ref5c4f67a0.OversampleV)
	x.PixelSnapH = (bool)(x.ref5c4f67a0.PixelSnapH)
	x.GlyphExtraSpacing = *NewVec2Ref(unsafe.Pointer(&x.ref5c4f67a0.GlyphExtraSpacing))
	x.GlyphOffset = *NewVec2Ref(unsafe.Pointer(&x.ref5c4f67a0.GlyphOffset))
	hxf95e7c8 := (*sliceHeader)(unsafe.Pointer(&x.GlyphRanges))
	hxf95e7c8.Data = uintptr(unsafe.Pointer(x.ref5c4f67a0.GlyphRanges))
	hxf95e7c8.Cap = 0x7fffffff
	// hxf95e7c8.Len = ?

	x.MergeMode = (bool)(x.ref5c4f67a0.MergeMode)
	x.RasterizerFlags = (uint32)(x.ref5c4f67a0.RasterizerFlags)
	x.RasterizerMultiply = (float32)(x.ref5c4f67a0.RasterizerMultiply)
	x.Name = *(*[32]byte)(unsafe.Pointer(&x.ref5c4f67a0.Name))
	hxff2234b := (*sliceHeader)(unsafe.Pointer(&x.DstFont))
	hxff2234b.Data = uintptr(unsafe.Pointer(x.ref5c4f67a0.DstFont))
	hxff2234b.Cap = 0x7fffffff
	// hxff2234b.Len = ?

}

// allocStructImGuiIOMemory allocates memory for type C.struct_ImGuiIO in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStructImGuiIOMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStructImGuiIOValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStructImGuiIOValue = unsafe.Sizeof([1]C.struct_ImGuiIO{})

// unpackPCharString represents the data from Go string as *C.char and avoids copying.
func unpackPCharString(str string) (*C.char, *cgoAllocMap) {
	str = safeString(str)
	h := (*stringHeader)(unsafe.Pointer(&str))
	return (*C.char)(unsafe.Pointer(h.Data)), cgoAllocsUnknown
}

type stringHeader struct {
	Data uintptr
	Len  int
}

// safeString ensures that the string is NULL-terminated, a NULL-terminated copy is created otherwise.
func safeString(str string) string {
	if len(str) > 0 && str[len(str)-1] != '\x00' {
		str = str + "\x00"
	} else if len(str) == 0 {
		str = "\x00"
	}
	return str
}

// allocA5StructImVec2Memory allocates memory for type [5]C.struct_ImVec2 in C.
// The caller is responsible for freeing the this memory via C.free.
func allocA5StructImVec2Memory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfA5StructImVec2Value))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfA5StructImVec2Value = unsafe.Sizeof([1][5]C.struct_ImVec2{})

// unpackA5Vec2 transforms a sliced Go data structure into plain C format.
func unpackA5Vec2(x [5]Vec2) (unpacked [5]C.struct_ImVec2, allocs *cgoAllocMap) {
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(*[5]C.struct_ImVec2) {
		go allocs.Free()
	})

	mem0 := allocA5StructImVec2Memory(1)
	allocs.Add(mem0)
	v0 := (*[5]C.struct_ImVec2)(mem0)
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	unpacked = *(*[5]C.struct_ImVec2)(mem0)
	return
}

// packPCharString creates a Go string backed by *C.char and avoids copying.
func packPCharString(p *C.char) (raw string) {
	if p != nil && *p != 0 {
		h := (*stringHeader)(unsafe.Pointer(&raw))
		h.Data = uintptr(unsafe.Pointer(p))
		for *p != 0 {
			p = (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)) // p++
		}
		h.Len = int(uintptr(unsafe.Pointer(p)) - h.Data)
	}
	return
}

// RawString reperesents a string backed by data on the C side.
type RawString string

// Copy returns a Go-managed copy of raw string.
func (raw RawString) Copy() string {
	if len(raw) == 0 {
		return ""
	}
	h := (*stringHeader)(unsafe.Pointer(&raw))
	return C.GoStringN((*C.char)(unsafe.Pointer(h.Data)), C.int(h.Len))
}

// packA5Vec2 reads sliced Go data structure out from plain C format.
func packA5Vec2(v *[5]Vec2, ptr0 *[5]C.struct_ImVec2) {
	for i0 := range v {
		ptr1 := ptr0[i0]
		v[i0] = *NewVec2Ref(unsafe.Pointer(&ptr1))
	}
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *IO) Ref() *C.struct_ImGuiIO {
	if x == nil {
		return nil
	}
	return x.ref4700f756
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *IO) Free() {
	if x != nil && x.allocs4700f756 != nil {
		x.allocs4700f756.(*cgoAllocMap).Free()
		x.ref4700f756 = nil
	}
}

// NewIORef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewIORef(ref unsafe.Pointer) *IO {
	if ref == nil {
		return nil
	}
	obj := new(IO)
	obj.ref4700f756 = (*C.struct_ImGuiIO)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *IO) PassRef() (*C.struct_ImGuiIO, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref4700f756 != nil {
		return x.ref4700f756, nil
	}
	mem4700f756 := allocStructImGuiIOMemory(1)
	ref4700f756 := (*C.struct_ImGuiIO)(mem4700f756)
	allocs4700f756 := new(cgoAllocMap)
	allocs4700f756.Add(mem4700f756)

	var cDisplaySize_allocs *cgoAllocMap
	ref4700f756.DisplaySize, cDisplaySize_allocs = x.DisplaySize.PassValue()
	allocs4700f756.Borrow(cDisplaySize_allocs)

	var cDeltaTime_allocs *cgoAllocMap
	ref4700f756.DeltaTime, cDeltaTime_allocs = (C.float)(x.DeltaTime), cgoAllocsUnknown
	allocs4700f756.Borrow(cDeltaTime_allocs)

	var cIniSavingRate_allocs *cgoAllocMap
	ref4700f756.IniSavingRate, cIniSavingRate_allocs = (C.float)(x.IniSavingRate), cgoAllocsUnknown
	allocs4700f756.Borrow(cIniSavingRate_allocs)

	var cIniFilename_allocs *cgoAllocMap
	ref4700f756.IniFilename, cIniFilename_allocs = unpackPCharString(x.IniFilename)
	allocs4700f756.Borrow(cIniFilename_allocs)

	var cLogFilename_allocs *cgoAllocMap
	ref4700f756.LogFilename, cLogFilename_allocs = unpackPCharString(x.LogFilename)
	allocs4700f756.Borrow(cLogFilename_allocs)

	var cMouseDoubleClickTime_allocs *cgoAllocMap
	ref4700f756.MouseDoubleClickTime, cMouseDoubleClickTime_allocs = (C.float)(x.MouseDoubleClickTime), cgoAllocsUnknown
	allocs4700f756.Borrow(cMouseDoubleClickTime_allocs)

	var cMouseDoubleClickMaxDist_allocs *cgoAllocMap
	ref4700f756.MouseDoubleClickMaxDist, cMouseDoubleClickMaxDist_allocs = (C.float)(x.MouseDoubleClickMaxDist), cgoAllocsUnknown
	allocs4700f756.Borrow(cMouseDoubleClickMaxDist_allocs)

	var cMouseDragThreshold_allocs *cgoAllocMap
	ref4700f756.MouseDragThreshold, cMouseDragThreshold_allocs = (C.float)(x.MouseDragThreshold), cgoAllocsUnknown
	allocs4700f756.Borrow(cMouseDragThreshold_allocs)

	var cKeyMap_allocs *cgoAllocMap
	ref4700f756.KeyMap, cKeyMap_allocs = *(*[19]C.int)(unsafe.Pointer(&x.KeyMap)), cgoAllocsUnknown
	allocs4700f756.Borrow(cKeyMap_allocs)

	var cKeyRepeatDelay_allocs *cgoAllocMap
	ref4700f756.KeyRepeatDelay, cKeyRepeatDelay_allocs = (C.float)(x.KeyRepeatDelay), cgoAllocsUnknown
	allocs4700f756.Borrow(cKeyRepeatDelay_allocs)

	var cKeyRepeatRate_allocs *cgoAllocMap
	ref4700f756.KeyRepeatRate, cKeyRepeatRate_allocs = (C.float)(x.KeyRepeatRate), cgoAllocsUnknown
	allocs4700f756.Borrow(cKeyRepeatRate_allocs)

	var cUserData_allocs *cgoAllocMap
	ref4700f756.UserData, cUserData_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.UserData)), cgoAllocsUnknown
	allocs4700f756.Borrow(cUserData_allocs)

	var cFonts_allocs *cgoAllocMap
	ref4700f756.Fonts, cFonts_allocs = (*C.struct_ImFontAtlas)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.Fonts)).Data)), cgoAllocsUnknown
	allocs4700f756.Borrow(cFonts_allocs)

	var cFontGlobalScale_allocs *cgoAllocMap
	ref4700f756.FontGlobalScale, cFontGlobalScale_allocs = (C.float)(x.FontGlobalScale), cgoAllocsUnknown
	allocs4700f756.Borrow(cFontGlobalScale_allocs)

	var cFontAllowUserScaling_allocs *cgoAllocMap
	ref4700f756.FontAllowUserScaling, cFontAllowUserScaling_allocs = (C._Bool)(x.FontAllowUserScaling), cgoAllocsUnknown
	allocs4700f756.Borrow(cFontAllowUserScaling_allocs)

	var cFontDefault_allocs *cgoAllocMap
	ref4700f756.FontDefault, cFontDefault_allocs = (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.FontDefault)).Data)), cgoAllocsUnknown
	allocs4700f756.Borrow(cFontDefault_allocs)

	var cDisplayFramebufferScale_allocs *cgoAllocMap
	ref4700f756.DisplayFramebufferScale, cDisplayFramebufferScale_allocs = x.DisplayFramebufferScale.PassValue()
	allocs4700f756.Borrow(cDisplayFramebufferScale_allocs)

	var cDisplayVisibleMin_allocs *cgoAllocMap
	ref4700f756.DisplayVisibleMin, cDisplayVisibleMin_allocs = x.DisplayVisibleMin.PassValue()
	allocs4700f756.Borrow(cDisplayVisibleMin_allocs)

	var cDisplayVisibleMax_allocs *cgoAllocMap
	ref4700f756.DisplayVisibleMax, cDisplayVisibleMax_allocs = x.DisplayVisibleMax.PassValue()
	allocs4700f756.Borrow(cDisplayVisibleMax_allocs)

	var cOptMacOSXBehaviors_allocs *cgoAllocMap
	ref4700f756.OptMacOSXBehaviors, cOptMacOSXBehaviors_allocs = (C._Bool)(x.OptMacOSXBehaviors), cgoAllocsUnknown
	allocs4700f756.Borrow(cOptMacOSXBehaviors_allocs)

	var cOptCursorBlink_allocs *cgoAllocMap
	ref4700f756.OptCursorBlink, cOptCursorBlink_allocs = (C._Bool)(x.OptCursorBlink), cgoAllocsUnknown
	allocs4700f756.Borrow(cOptCursorBlink_allocs)

	var cRenderDrawListsFn_allocs *cgoAllocMap
	ref4700f756.RenderDrawListsFn, cRenderDrawListsFn_allocs = x.RenderDrawListsFn.PassRef()
	allocs4700f756.Borrow(cRenderDrawListsFn_allocs)

	var cGetClipboardTextFn_allocs *cgoAllocMap
	ref4700f756.GetClipboardTextFn, cGetClipboardTextFn_allocs = x.GetClipboardTextFn.PassRef()
	allocs4700f756.Borrow(cGetClipboardTextFn_allocs)

	var cSetClipboardTextFn_allocs *cgoAllocMap
	ref4700f756.SetClipboardTextFn, cSetClipboardTextFn_allocs = x.SetClipboardTextFn.PassRef()
	allocs4700f756.Borrow(cSetClipboardTextFn_allocs)

	var cClipboardUserData_allocs *cgoAllocMap
	ref4700f756.ClipboardUserData, cClipboardUserData_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.ClipboardUserData)), cgoAllocsUnknown
	allocs4700f756.Borrow(cClipboardUserData_allocs)

	var cMemAllocFn_allocs *cgoAllocMap
	ref4700f756.MemAllocFn, cMemAllocFn_allocs = x.MemAllocFn.PassRef()
	allocs4700f756.Borrow(cMemAllocFn_allocs)

	var cMemFreeFn_allocs *cgoAllocMap
	ref4700f756.MemFreeFn, cMemFreeFn_allocs = x.MemFreeFn.PassRef()
	allocs4700f756.Borrow(cMemFreeFn_allocs)

	var cImeSetInputScreenPosFn_allocs *cgoAllocMap
	ref4700f756.ImeSetInputScreenPosFn, cImeSetInputScreenPosFn_allocs = x.ESetInputScreenPosFn.PassRef()
	allocs4700f756.Borrow(cImeSetInputScreenPosFn_allocs)

	var cImeWindowHandle_allocs *cgoAllocMap
	ref4700f756.ImeWindowHandle, cImeWindowHandle_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.EWindowHandle)), cgoAllocsUnknown
	allocs4700f756.Borrow(cImeWindowHandle_allocs)

	var cMousePos_allocs *cgoAllocMap
	ref4700f756.MousePos, cMousePos_allocs = x.MousePos.PassValue()
	allocs4700f756.Borrow(cMousePos_allocs)

	var cMouseDown_allocs *cgoAllocMap
	ref4700f756.MouseDown, cMouseDown_allocs = *(*[5]C._Bool)(unsafe.Pointer(&x.MouseDown)), cgoAllocsUnknown
	allocs4700f756.Borrow(cMouseDown_allocs)

	var cMouseWheel_allocs *cgoAllocMap
	ref4700f756.MouseWheel, cMouseWheel_allocs = (C.float)(x.MouseWheel), cgoAllocsUnknown
	allocs4700f756.Borrow(cMouseWheel_allocs)

	var cMouseDrawCursor_allocs *cgoAllocMap
	ref4700f756.MouseDrawCursor, cMouseDrawCursor_allocs = (C._Bool)(x.MouseDrawCursor), cgoAllocsUnknown
	allocs4700f756.Borrow(cMouseDrawCursor_allocs)

	var cKeyCtrl_allocs *cgoAllocMap
	ref4700f756.KeyCtrl, cKeyCtrl_allocs = (C._Bool)(x.KeyCtrl), cgoAllocsUnknown
	allocs4700f756.Borrow(cKeyCtrl_allocs)

	var cKeyShift_allocs *cgoAllocMap
	ref4700f756.KeyShift, cKeyShift_allocs = (C._Bool)(x.KeyShift), cgoAllocsUnknown
	allocs4700f756.Borrow(cKeyShift_allocs)

	var cKeyAlt_allocs *cgoAllocMap
	ref4700f756.KeyAlt, cKeyAlt_allocs = (C._Bool)(x.KeyAlt), cgoAllocsUnknown
	allocs4700f756.Borrow(cKeyAlt_allocs)

	var cKeySuper_allocs *cgoAllocMap
	ref4700f756.KeySuper, cKeySuper_allocs = (C._Bool)(x.KeySuper), cgoAllocsUnknown
	allocs4700f756.Borrow(cKeySuper_allocs)

	var cKeysDown_allocs *cgoAllocMap
	ref4700f756.KeysDown, cKeysDown_allocs = *(*[512]C._Bool)(unsafe.Pointer(&x.KeysDown)), cgoAllocsUnknown
	allocs4700f756.Borrow(cKeysDown_allocs)

	var cInputCharacters_allocs *cgoAllocMap
	ref4700f756.InputCharacters, cInputCharacters_allocs = *(*[17]C.ImWchar)(unsafe.Pointer(&x.InputCharacters)), cgoAllocsUnknown
	allocs4700f756.Borrow(cInputCharacters_allocs)

	var cWantCaptureMouse_allocs *cgoAllocMap
	ref4700f756.WantCaptureMouse, cWantCaptureMouse_allocs = (C._Bool)(x.WantCaptureMouse), cgoAllocsUnknown
	allocs4700f756.Borrow(cWantCaptureMouse_allocs)

	var cWantCaptureKeyboard_allocs *cgoAllocMap
	ref4700f756.WantCaptureKeyboard, cWantCaptureKeyboard_allocs = (C._Bool)(x.WantCaptureKeyboard), cgoAllocsUnknown
	allocs4700f756.Borrow(cWantCaptureKeyboard_allocs)

	var cWantTextInput_allocs *cgoAllocMap
	ref4700f756.WantTextInput, cWantTextInput_allocs = (C._Bool)(x.WantTextInput), cgoAllocsUnknown
	allocs4700f756.Borrow(cWantTextInput_allocs)

	var cFramerate_allocs *cgoAllocMap
	ref4700f756.Framerate, cFramerate_allocs = (C.float)(x.Framerate), cgoAllocsUnknown
	allocs4700f756.Borrow(cFramerate_allocs)

	var cMetricsAllocs_allocs *cgoAllocMap
	ref4700f756.MetricsAllocs, cMetricsAllocs_allocs = (C.int)(x.MetricsAllocs), cgoAllocsUnknown
	allocs4700f756.Borrow(cMetricsAllocs_allocs)

	var cMetricsRenderVertices_allocs *cgoAllocMap
	ref4700f756.MetricsRenderVertices, cMetricsRenderVertices_allocs = (C.int)(x.MetricsRenderVertices), cgoAllocsUnknown
	allocs4700f756.Borrow(cMetricsRenderVertices_allocs)

	var cMetricsRenderIndices_allocs *cgoAllocMap
	ref4700f756.MetricsRenderIndices, cMetricsRenderIndices_allocs = (C.int)(x.MetricsRenderIndices), cgoAllocsUnknown
	allocs4700f756.Borrow(cMetricsRenderIndices_allocs)

	var cMetricsActiveWindows_allocs *cgoAllocMap
	ref4700f756.MetricsActiveWindows, cMetricsActiveWindows_allocs = (C.int)(x.MetricsActiveWindows), cgoAllocsUnknown
	allocs4700f756.Borrow(cMetricsActiveWindows_allocs)

	var cMouseDelta_allocs *cgoAllocMap
	ref4700f756.MouseDelta, cMouseDelta_allocs = x.MouseDelta.PassValue()
	allocs4700f756.Borrow(cMouseDelta_allocs)

	var cMousePosPrev_allocs *cgoAllocMap
	ref4700f756.MousePosPrev, cMousePosPrev_allocs = x.MousePosPrev.PassValue()
	allocs4700f756.Borrow(cMousePosPrev_allocs)

	var cMouseClicked_allocs *cgoAllocMap
	ref4700f756.MouseClicked, cMouseClicked_allocs = *(*[5]C._Bool)(unsafe.Pointer(&x.MouseClicked)), cgoAllocsUnknown
	allocs4700f756.Borrow(cMouseClicked_allocs)

	var cMouseClickedPos_allocs *cgoAllocMap
	ref4700f756.MouseClickedPos, cMouseClickedPos_allocs = unpackA5Vec2(x.MouseClickedPos)
	allocs4700f756.Borrow(cMouseClickedPos_allocs)

	var cMouseClickedTime_allocs *cgoAllocMap
	ref4700f756.MouseClickedTime, cMouseClickedTime_allocs = *(*[5]C.float)(unsafe.Pointer(&x.MouseClickedTime)), cgoAllocsUnknown
	allocs4700f756.Borrow(cMouseClickedTime_allocs)

	var cMouseDoubleClicked_allocs *cgoAllocMap
	ref4700f756.MouseDoubleClicked, cMouseDoubleClicked_allocs = *(*[5]C._Bool)(unsafe.Pointer(&x.MouseDoubleClicked)), cgoAllocsUnknown
	allocs4700f756.Borrow(cMouseDoubleClicked_allocs)

	var cMouseReleased_allocs *cgoAllocMap
	ref4700f756.MouseReleased, cMouseReleased_allocs = *(*[5]C._Bool)(unsafe.Pointer(&x.MouseReleased)), cgoAllocsUnknown
	allocs4700f756.Borrow(cMouseReleased_allocs)

	var cMouseDownOwned_allocs *cgoAllocMap
	ref4700f756.MouseDownOwned, cMouseDownOwned_allocs = *(*[5]C._Bool)(unsafe.Pointer(&x.MouseDownOwned)), cgoAllocsUnknown
	allocs4700f756.Borrow(cMouseDownOwned_allocs)

	var cMouseDownDuration_allocs *cgoAllocMap
	ref4700f756.MouseDownDuration, cMouseDownDuration_allocs = *(*[5]C.float)(unsafe.Pointer(&x.MouseDownDuration)), cgoAllocsUnknown
	allocs4700f756.Borrow(cMouseDownDuration_allocs)

	var cMouseDownDurationPrev_allocs *cgoAllocMap
	ref4700f756.MouseDownDurationPrev, cMouseDownDurationPrev_allocs = *(*[5]C.float)(unsafe.Pointer(&x.MouseDownDurationPrev)), cgoAllocsUnknown
	allocs4700f756.Borrow(cMouseDownDurationPrev_allocs)

	var cMouseDragMaxDistanceAbs_allocs *cgoAllocMap
	ref4700f756.MouseDragMaxDistanceAbs, cMouseDragMaxDistanceAbs_allocs = unpackA5Vec2(x.MouseDragMaxDistanceAbs)
	allocs4700f756.Borrow(cMouseDragMaxDistanceAbs_allocs)

	var cMouseDragMaxDistanceSqr_allocs *cgoAllocMap
	ref4700f756.MouseDragMaxDistanceSqr, cMouseDragMaxDistanceSqr_allocs = *(*[5]C.float)(unsafe.Pointer(&x.MouseDragMaxDistanceSqr)), cgoAllocsUnknown
	allocs4700f756.Borrow(cMouseDragMaxDistanceSqr_allocs)

	var cKeysDownDuration_allocs *cgoAllocMap
	ref4700f756.KeysDownDuration, cKeysDownDuration_allocs = *(*[512]C.float)(unsafe.Pointer(&x.KeysDownDuration)), cgoAllocsUnknown
	allocs4700f756.Borrow(cKeysDownDuration_allocs)

	var cKeysDownDurationPrev_allocs *cgoAllocMap
	ref4700f756.KeysDownDurationPrev, cKeysDownDurationPrev_allocs = *(*[512]C.float)(unsafe.Pointer(&x.KeysDownDurationPrev)), cgoAllocsUnknown
	allocs4700f756.Borrow(cKeysDownDurationPrev_allocs)

	x.ref4700f756 = ref4700f756
	x.allocs4700f756 = allocs4700f756
	return ref4700f756, allocs4700f756

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x IO) PassValue() (C.struct_ImGuiIO, *cgoAllocMap) {
	if x.ref4700f756 != nil {
		return *x.ref4700f756, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *IO) Deref() {
	if x.ref4700f756 == nil {
		return
	}
	x.DisplaySize = *NewVec2Ref(unsafe.Pointer(&x.ref4700f756.DisplaySize))
	x.DeltaTime = (float32)(x.ref4700f756.DeltaTime)
	x.IniSavingRate = (float32)(x.ref4700f756.IniSavingRate)
	x.IniFilename = packPCharString(x.ref4700f756.IniFilename)
	x.LogFilename = packPCharString(x.ref4700f756.LogFilename)
	x.MouseDoubleClickTime = (float32)(x.ref4700f756.MouseDoubleClickTime)
	x.MouseDoubleClickMaxDist = (float32)(x.ref4700f756.MouseDoubleClickMaxDist)
	x.MouseDragThreshold = (float32)(x.ref4700f756.MouseDragThreshold)
	x.KeyMap = *(*[19]int32)(unsafe.Pointer(&x.ref4700f756.KeyMap))
	x.KeyRepeatDelay = (float32)(x.ref4700f756.KeyRepeatDelay)
	x.KeyRepeatRate = (float32)(x.ref4700f756.KeyRepeatRate)
	x.UserData = (unsafe.Pointer)(unsafe.Pointer(x.ref4700f756.UserData))
	hxff73280 := (*sliceHeader)(unsafe.Pointer(&x.Fonts))
	hxff73280.Data = uintptr(unsafe.Pointer(x.ref4700f756.Fonts))
	hxff73280.Cap = 0x7fffffff
	// hxff73280.Len = ?

	x.FontGlobalScale = (float32)(x.ref4700f756.FontGlobalScale)
	x.FontAllowUserScaling = (bool)(x.ref4700f756.FontAllowUserScaling)
	hxfa9955c := (*sliceHeader)(unsafe.Pointer(&x.FontDefault))
	hxfa9955c.Data = uintptr(unsafe.Pointer(x.ref4700f756.FontDefault))
	hxfa9955c.Cap = 0x7fffffff
	// hxfa9955c.Len = ?

	x.DisplayFramebufferScale = *NewVec2Ref(unsafe.Pointer(&x.ref4700f756.DisplayFramebufferScale))
	x.DisplayVisibleMin = *NewVec2Ref(unsafe.Pointer(&x.ref4700f756.DisplayVisibleMin))
	x.DisplayVisibleMax = *NewVec2Ref(unsafe.Pointer(&x.ref4700f756.DisplayVisibleMax))
	x.OptMacOSXBehaviors = (bool)(x.ref4700f756.OptMacOSXBehaviors)
	x.OptCursorBlink = (bool)(x.ref4700f756.OptCursorBlink)
	x.RenderDrawListsFn = NewRef(unsafe.Pointer(x.ref4700f756.RenderDrawListsFn))
	x.GetClipboardTextFn = NewRef(unsafe.Pointer(x.ref4700f756.GetClipboardTextFn))
	x.SetClipboardTextFn = NewRef(unsafe.Pointer(x.ref4700f756.SetClipboardTextFn))
	x.ClipboardUserData = (unsafe.Pointer)(unsafe.Pointer(x.ref4700f756.ClipboardUserData))
	x.MemAllocFn = NewRef(unsafe.Pointer(x.ref4700f756.MemAllocFn))
	x.MemFreeFn = NewRef(unsafe.Pointer(x.ref4700f756.MemFreeFn))
	x.ESetInputScreenPosFn = NewRef(unsafe.Pointer(x.ref4700f756.ImeSetInputScreenPosFn))
	x.EWindowHandle = (unsafe.Pointer)(unsafe.Pointer(x.ref4700f756.ImeWindowHandle))
	x.MousePos = *NewVec2Ref(unsafe.Pointer(&x.ref4700f756.MousePos))
	x.MouseDown = *(*[5]bool)(unsafe.Pointer(&x.ref4700f756.MouseDown))
	x.MouseWheel = (float32)(x.ref4700f756.MouseWheel)
	x.MouseDrawCursor = (bool)(x.ref4700f756.MouseDrawCursor)
	x.KeyCtrl = (bool)(x.ref4700f756.KeyCtrl)
	x.KeyShift = (bool)(x.ref4700f756.KeyShift)
	x.KeyAlt = (bool)(x.ref4700f756.KeyAlt)
	x.KeySuper = (bool)(x.ref4700f756.KeySuper)
	x.KeysDown = *(*[512]bool)(unsafe.Pointer(&x.ref4700f756.KeysDown))
	x.InputCharacters = *(*[17]Wchar)(unsafe.Pointer(&x.ref4700f756.InputCharacters))
	x.WantCaptureMouse = (bool)(x.ref4700f756.WantCaptureMouse)
	x.WantCaptureKeyboard = (bool)(x.ref4700f756.WantCaptureKeyboard)
	x.WantTextInput = (bool)(x.ref4700f756.WantTextInput)
	x.Framerate = (float32)(x.ref4700f756.Framerate)
	x.MetricsAllocs = (int32)(x.ref4700f756.MetricsAllocs)
	x.MetricsRenderVertices = (int32)(x.ref4700f756.MetricsRenderVertices)
	x.MetricsRenderIndices = (int32)(x.ref4700f756.MetricsRenderIndices)
	x.MetricsActiveWindows = (int32)(x.ref4700f756.MetricsActiveWindows)
	x.MouseDelta = *NewVec2Ref(unsafe.Pointer(&x.ref4700f756.MouseDelta))
	x.MousePosPrev = *NewVec2Ref(unsafe.Pointer(&x.ref4700f756.MousePosPrev))
	x.MouseClicked = *(*[5]bool)(unsafe.Pointer(&x.ref4700f756.MouseClicked))
	packA5Vec2(&x.MouseClickedPos, (*[5]C.struct_ImVec2)(unsafe.Pointer(&x.ref4700f756.MouseClickedPos)))
	x.MouseClickedTime = *(*[5]float32)(unsafe.Pointer(&x.ref4700f756.MouseClickedTime))
	x.MouseDoubleClicked = *(*[5]bool)(unsafe.Pointer(&x.ref4700f756.MouseDoubleClicked))
	x.MouseReleased = *(*[5]bool)(unsafe.Pointer(&x.ref4700f756.MouseReleased))
	x.MouseDownOwned = *(*[5]bool)(unsafe.Pointer(&x.ref4700f756.MouseDownOwned))
	x.MouseDownDuration = *(*[5]float32)(unsafe.Pointer(&x.ref4700f756.MouseDownDuration))
	x.MouseDownDurationPrev = *(*[5]float32)(unsafe.Pointer(&x.ref4700f756.MouseDownDurationPrev))
	packA5Vec2(&x.MouseDragMaxDistanceAbs, (*[5]C.struct_ImVec2)(unsafe.Pointer(&x.ref4700f756.MouseDragMaxDistanceAbs)))
	x.MouseDragMaxDistanceSqr = *(*[5]float32)(unsafe.Pointer(&x.ref4700f756.MouseDragMaxDistanceSqr))
	x.KeysDownDuration = *(*[512]float32)(unsafe.Pointer(&x.ref4700f756.KeysDownDuration))
	x.KeysDownDurationPrev = *(*[512]float32)(unsafe.Pointer(&x.ref4700f756.KeysDownDurationPrev))
}

// allocStructImGuiListClipperMemory allocates memory for type C.struct_ImGuiListClipper in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStructImGuiListClipperMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStructImGuiListClipperValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStructImGuiListClipperValue = unsafe.Sizeof([1]C.struct_ImGuiListClipper{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *ListClipper) Ref() *C.struct_ImGuiListClipper {
	if x == nil {
		return nil
	}
	return x.refd52a46bd
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *ListClipper) Free() {
	if x != nil && x.allocsd52a46bd != nil {
		x.allocsd52a46bd.(*cgoAllocMap).Free()
		x.refd52a46bd = nil
	}
}

// NewListClipperRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewListClipperRef(ref unsafe.Pointer) *ListClipper {
	if ref == nil {
		return nil
	}
	obj := new(ListClipper)
	obj.refd52a46bd = (*C.struct_ImGuiListClipper)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *ListClipper) PassRef() (*C.struct_ImGuiListClipper, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refd52a46bd != nil {
		return x.refd52a46bd, nil
	}
	memd52a46bd := allocStructImGuiListClipperMemory(1)
	refd52a46bd := (*C.struct_ImGuiListClipper)(memd52a46bd)
	allocsd52a46bd := new(cgoAllocMap)
	allocsd52a46bd.Add(memd52a46bd)

	var cStartPosY_allocs *cgoAllocMap
	refd52a46bd.StartPosY, cStartPosY_allocs = (C.float)(x.StartPosY), cgoAllocsUnknown
	allocsd52a46bd.Borrow(cStartPosY_allocs)

	var cItemsHeight_allocs *cgoAllocMap
	refd52a46bd.ItemsHeight, cItemsHeight_allocs = (C.float)(x.ItemsHeight), cgoAllocsUnknown
	allocsd52a46bd.Borrow(cItemsHeight_allocs)

	var cItemsCount_allocs *cgoAllocMap
	refd52a46bd.ItemsCount, cItemsCount_allocs = (C.int)(x.ItemsCount), cgoAllocsUnknown
	allocsd52a46bd.Borrow(cItemsCount_allocs)

	var cStepNo_allocs *cgoAllocMap
	refd52a46bd.StepNo, cStepNo_allocs = (C.int)(x.StepNo), cgoAllocsUnknown
	allocsd52a46bd.Borrow(cStepNo_allocs)

	var cDisplayStart_allocs *cgoAllocMap
	refd52a46bd.DisplayStart, cDisplayStart_allocs = (C.int)(x.DisplayStart), cgoAllocsUnknown
	allocsd52a46bd.Borrow(cDisplayStart_allocs)

	var cDisplayEnd_allocs *cgoAllocMap
	refd52a46bd.DisplayEnd, cDisplayEnd_allocs = (C.int)(x.DisplayEnd), cgoAllocsUnknown
	allocsd52a46bd.Borrow(cDisplayEnd_allocs)

	x.refd52a46bd = refd52a46bd
	x.allocsd52a46bd = allocsd52a46bd
	return refd52a46bd, allocsd52a46bd

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x ListClipper) PassValue() (C.struct_ImGuiListClipper, *cgoAllocMap) {
	if x.refd52a46bd != nil {
		return *x.refd52a46bd, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *ListClipper) Deref() {
	if x.refd52a46bd == nil {
		return
	}
	x.StartPosY = (float32)(x.refd52a46bd.StartPosY)
	x.ItemsHeight = (float32)(x.refd52a46bd.ItemsHeight)
	x.ItemsCount = (int32)(x.refd52a46bd.ItemsCount)
	x.StepNo = (int32)(x.refd52a46bd.StepNo)
	x.DisplayStart = (int32)(x.refd52a46bd.DisplayStart)
	x.DisplayEnd = (int32)(x.refd52a46bd.DisplayEnd)
}

// allocStructImGuiPayloadMemory allocates memory for type C.struct_ImGuiPayload in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStructImGuiPayloadMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStructImGuiPayloadValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStructImGuiPayloadValue = unsafe.Sizeof([1]C.struct_ImGuiPayload{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *Payload) Ref() *C.struct_ImGuiPayload {
	if x == nil {
		return nil
	}
	return x.refea983e4e
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *Payload) Free() {
	if x != nil && x.allocsea983e4e != nil {
		x.allocsea983e4e.(*cgoAllocMap).Free()
		x.refea983e4e = nil
	}
}

// NewPayloadRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPayloadRef(ref unsafe.Pointer) *Payload {
	if ref == nil {
		return nil
	}
	obj := new(Payload)
	obj.refea983e4e = (*C.struct_ImGuiPayload)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *Payload) PassRef() (*C.struct_ImGuiPayload, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refea983e4e != nil {
		return x.refea983e4e, nil
	}
	memea983e4e := allocStructImGuiPayloadMemory(1)
	refea983e4e := (*C.struct_ImGuiPayload)(memea983e4e)
	allocsea983e4e := new(cgoAllocMap)
	allocsea983e4e.Add(memea983e4e)

	var cData_allocs *cgoAllocMap
	refea983e4e.Data, cData_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.Data)), cgoAllocsUnknown
	allocsea983e4e.Borrow(cData_allocs)

	var cDataSize_allocs *cgoAllocMap
	refea983e4e.DataSize, cDataSize_allocs = (C.int)(x.DataSize), cgoAllocsUnknown
	allocsea983e4e.Borrow(cDataSize_allocs)

	var cSourceId_allocs *cgoAllocMap
	refea983e4e.SourceId, cSourceId_allocs = (C.ImGuiID)(x.SourceId), cgoAllocsUnknown
	allocsea983e4e.Borrow(cSourceId_allocs)

	var cSourceParentId_allocs *cgoAllocMap
	refea983e4e.SourceParentId, cSourceParentId_allocs = (C.ImGuiID)(x.SourceParentId), cgoAllocsUnknown
	allocsea983e4e.Borrow(cSourceParentId_allocs)

	var cDataFrameCount_allocs *cgoAllocMap
	refea983e4e.DataFrameCount, cDataFrameCount_allocs = (C.int)(x.DataFrameCount), cgoAllocsUnknown
	allocsea983e4e.Borrow(cDataFrameCount_allocs)

	var cDataType_allocs *cgoAllocMap
	refea983e4e.DataType, cDataType_allocs = *(*[9]C.char)(unsafe.Pointer(&x.DataType)), cgoAllocsUnknown
	allocsea983e4e.Borrow(cDataType_allocs)

	var cPreview_allocs *cgoAllocMap
	refea983e4e.Preview, cPreview_allocs = (C._Bool)(x.Preview), cgoAllocsUnknown
	allocsea983e4e.Borrow(cPreview_allocs)

	var cDelivery_allocs *cgoAllocMap
	refea983e4e.Delivery, cDelivery_allocs = (C._Bool)(x.Delivery), cgoAllocsUnknown
	allocsea983e4e.Borrow(cDelivery_allocs)

	x.refea983e4e = refea983e4e
	x.allocsea983e4e = allocsea983e4e
	return refea983e4e, allocsea983e4e

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x Payload) PassValue() (C.struct_ImGuiPayload, *cgoAllocMap) {
	if x.refea983e4e != nil {
		return *x.refea983e4e, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *Payload) Deref() {
	if x.refea983e4e == nil {
		return
	}
	x.Data = (unsafe.Pointer)(unsafe.Pointer(x.refea983e4e.Data))
	x.DataSize = (int32)(x.refea983e4e.DataSize)
	x.SourceId = (ID)(x.refea983e4e.SourceId)
	x.SourceParentId = (ID)(x.refea983e4e.SourceParentId)
	x.DataFrameCount = (int32)(x.refea983e4e.DataFrameCount)
	x.DataType = *(*[9]byte)(unsafe.Pointer(&x.refea983e4e.DataType))
	x.Preview = (bool)(x.refea983e4e.Preview)
	x.Delivery = (bool)(x.refea983e4e.Delivery)
}

// allocStructImGuiSizeConstraintCallbackDataMemory allocates memory for type C.struct_ImGuiSizeConstraintCallbackData in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStructImGuiSizeConstraintCallbackDataMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStructImGuiSizeConstraintCallbackDataValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStructImGuiSizeConstraintCallbackDataValue = unsafe.Sizeof([1]C.struct_ImGuiSizeConstraintCallbackData{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *SizeConstraintCallbackData) Ref() *C.struct_ImGuiSizeConstraintCallbackData {
	if x == nil {
		return nil
	}
	return x.ref24c5ad70
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *SizeConstraintCallbackData) Free() {
	if x != nil && x.allocs24c5ad70 != nil {
		x.allocs24c5ad70.(*cgoAllocMap).Free()
		x.ref24c5ad70 = nil
	}
}

// NewSizeConstraintCallbackDataRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewSizeConstraintCallbackDataRef(ref unsafe.Pointer) *SizeConstraintCallbackData {
	if ref == nil {
		return nil
	}
	obj := new(SizeConstraintCallbackData)
	obj.ref24c5ad70 = (*C.struct_ImGuiSizeConstraintCallbackData)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *SizeConstraintCallbackData) PassRef() (*C.struct_ImGuiSizeConstraintCallbackData, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref24c5ad70 != nil {
		return x.ref24c5ad70, nil
	}
	mem24c5ad70 := allocStructImGuiSizeConstraintCallbackDataMemory(1)
	ref24c5ad70 := (*C.struct_ImGuiSizeConstraintCallbackData)(mem24c5ad70)
	allocs24c5ad70 := new(cgoAllocMap)
	allocs24c5ad70.Add(mem24c5ad70)

	var cUserData_allocs *cgoAllocMap
	ref24c5ad70.UserData, cUserData_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.UserData)), cgoAllocsUnknown
	allocs24c5ad70.Borrow(cUserData_allocs)

	var cPos_allocs *cgoAllocMap
	ref24c5ad70.Pos, cPos_allocs = x.Pos.PassValue()
	allocs24c5ad70.Borrow(cPos_allocs)

	var cCurrentSize_allocs *cgoAllocMap
	ref24c5ad70.CurrentSize, cCurrentSize_allocs = x.CurrentSize.PassValue()
	allocs24c5ad70.Borrow(cCurrentSize_allocs)

	var cDesiredSize_allocs *cgoAllocMap
	ref24c5ad70.DesiredSize, cDesiredSize_allocs = x.DesiredSize.PassValue()
	allocs24c5ad70.Borrow(cDesiredSize_allocs)

	x.ref24c5ad70 = ref24c5ad70
	x.allocs24c5ad70 = allocs24c5ad70
	return ref24c5ad70, allocs24c5ad70

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x SizeConstraintCallbackData) PassValue() (C.struct_ImGuiSizeConstraintCallbackData, *cgoAllocMap) {
	if x.ref24c5ad70 != nil {
		return *x.ref24c5ad70, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *SizeConstraintCallbackData) Deref() {
	if x.ref24c5ad70 == nil {
		return
	}
	x.UserData = (unsafe.Pointer)(unsafe.Pointer(x.ref24c5ad70.UserData))
	x.Pos = *NewVec2Ref(unsafe.Pointer(&x.ref24c5ad70.Pos))
	x.CurrentSize = *NewVec2Ref(unsafe.Pointer(&x.ref24c5ad70.CurrentSize))
	x.DesiredSize = *NewVec2Ref(unsafe.Pointer(&x.ref24c5ad70.DesiredSize))
}

// Ref returns a reference to C object as it is.
func (x *Storage) Ref() *C.struct_ImGuiStorage {
	if x == nil {
		return nil
	}
	return (*C.struct_ImGuiStorage)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *Storage) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewStorageRef converts the C object reference into a raw struct reference without wrapping.
func NewStorageRef(ref unsafe.Pointer) *Storage {
	return (*Storage)(ref)
}

// NewStorage allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewStorage() *Storage {
	return (*Storage)(allocStructImGuiStorageMemory(1))
}

// allocStructImGuiStorageMemory allocates memory for type C.struct_ImGuiStorage in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStructImGuiStorageMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStructImGuiStorageValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStructImGuiStorageValue = unsafe.Sizeof([1]C.struct_ImGuiStorage{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *Storage) PassRef() *C.struct_ImGuiStorage {
	if x == nil {
		x = (*Storage)(allocStructImGuiStorageMemory(1))
	}
	return (*C.struct_ImGuiStorage)(unsafe.Pointer(x))
}

// allocStructImGuiStyleMemory allocates memory for type C.struct_ImGuiStyle in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStructImGuiStyleMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStructImGuiStyleValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStructImGuiStyleValue = unsafe.Sizeof([1]C.struct_ImGuiStyle{})

// allocA43StructImVec4Memory allocates memory for type [43]C.struct_ImVec4 in C.
// The caller is responsible for freeing the this memory via C.free.
func allocA43StructImVec4Memory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfA43StructImVec4Value))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfA43StructImVec4Value = unsafe.Sizeof([1][43]C.struct_ImVec4{})

// unpackA43Vec4 transforms a sliced Go data structure into plain C format.
func unpackA43Vec4(x [43]Vec4) (unpacked [43]C.struct_ImVec4, allocs *cgoAllocMap) {
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(*[43]C.struct_ImVec4) {
		go allocs.Free()
	})

	mem0 := allocA43StructImVec4Memory(1)
	allocs.Add(mem0)
	v0 := (*[43]C.struct_ImVec4)(mem0)
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	unpacked = *(*[43]C.struct_ImVec4)(mem0)
	return
}

// packA43Vec4 reads sliced Go data structure out from plain C format.
func packA43Vec4(v *[43]Vec4, ptr0 *[43]C.struct_ImVec4) {
	for i0 := range v {
		ptr1 := ptr0[i0]
		v[i0] = *NewVec4Ref(unsafe.Pointer(&ptr1))
	}
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *Style) Ref() *C.struct_ImGuiStyle {
	if x == nil {
		return nil
	}
	return x.ref6e47ee0d
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *Style) Free() {
	if x != nil && x.allocs6e47ee0d != nil {
		x.allocs6e47ee0d.(*cgoAllocMap).Free()
		x.ref6e47ee0d = nil
	}
}

// NewStyleRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewStyleRef(ref unsafe.Pointer) *Style {
	if ref == nil {
		return nil
	}
	obj := new(Style)
	obj.ref6e47ee0d = (*C.struct_ImGuiStyle)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *Style) PassRef() (*C.struct_ImGuiStyle, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref6e47ee0d != nil {
		return x.ref6e47ee0d, nil
	}
	mem6e47ee0d := allocStructImGuiStyleMemory(1)
	ref6e47ee0d := (*C.struct_ImGuiStyle)(mem6e47ee0d)
	allocs6e47ee0d := new(cgoAllocMap)
	allocs6e47ee0d.Add(mem6e47ee0d)

	var cAlpha_allocs *cgoAllocMap
	ref6e47ee0d.Alpha, cAlpha_allocs = (C.float)(x.Alpha), cgoAllocsUnknown
	allocs6e47ee0d.Borrow(cAlpha_allocs)

	var cWindowPadding_allocs *cgoAllocMap
	ref6e47ee0d.WindowPadding, cWindowPadding_allocs = x.WindowPadding.PassValue()
	allocs6e47ee0d.Borrow(cWindowPadding_allocs)

	var cWindowRounding_allocs *cgoAllocMap
	ref6e47ee0d.WindowRounding, cWindowRounding_allocs = (C.float)(x.WindowRounding), cgoAllocsUnknown
	allocs6e47ee0d.Borrow(cWindowRounding_allocs)

	var cWindowBorderSize_allocs *cgoAllocMap
	ref6e47ee0d.WindowBorderSize, cWindowBorderSize_allocs = (C.float)(x.WindowBorderSize), cgoAllocsUnknown
	allocs6e47ee0d.Borrow(cWindowBorderSize_allocs)

	var cWindowMinSize_allocs *cgoAllocMap
	ref6e47ee0d.WindowMinSize, cWindowMinSize_allocs = x.WindowMinSize.PassValue()
	allocs6e47ee0d.Borrow(cWindowMinSize_allocs)

	var cWindowTitleAlign_allocs *cgoAllocMap
	ref6e47ee0d.WindowTitleAlign, cWindowTitleAlign_allocs = x.WindowTitleAlign.PassValue()
	allocs6e47ee0d.Borrow(cWindowTitleAlign_allocs)

	var cChildRounding_allocs *cgoAllocMap
	ref6e47ee0d.ChildRounding, cChildRounding_allocs = (C.float)(x.ChildRounding), cgoAllocsUnknown
	allocs6e47ee0d.Borrow(cChildRounding_allocs)

	var cChildBorderSize_allocs *cgoAllocMap
	ref6e47ee0d.ChildBorderSize, cChildBorderSize_allocs = (C.float)(x.ChildBorderSize), cgoAllocsUnknown
	allocs6e47ee0d.Borrow(cChildBorderSize_allocs)

	var cPopupRounding_allocs *cgoAllocMap
	ref6e47ee0d.PopupRounding, cPopupRounding_allocs = (C.float)(x.PopupRounding), cgoAllocsUnknown
	allocs6e47ee0d.Borrow(cPopupRounding_allocs)

	var cPopupBorderSize_allocs *cgoAllocMap
	ref6e47ee0d.PopupBorderSize, cPopupBorderSize_allocs = (C.float)(x.PopupBorderSize), cgoAllocsUnknown
	allocs6e47ee0d.Borrow(cPopupBorderSize_allocs)

	var cFramePadding_allocs *cgoAllocMap
	ref6e47ee0d.FramePadding, cFramePadding_allocs = x.FramePadding.PassValue()
	allocs6e47ee0d.Borrow(cFramePadding_allocs)

	var cFrameRounding_allocs *cgoAllocMap
	ref6e47ee0d.FrameRounding, cFrameRounding_allocs = (C.float)(x.FrameRounding), cgoAllocsUnknown
	allocs6e47ee0d.Borrow(cFrameRounding_allocs)

	var cFrameBorderSize_allocs *cgoAllocMap
	ref6e47ee0d.FrameBorderSize, cFrameBorderSize_allocs = (C.float)(x.FrameBorderSize), cgoAllocsUnknown
	allocs6e47ee0d.Borrow(cFrameBorderSize_allocs)

	var cItemSpacing_allocs *cgoAllocMap
	ref6e47ee0d.ItemSpacing, cItemSpacing_allocs = x.ItemSpacing.PassValue()
	allocs6e47ee0d.Borrow(cItemSpacing_allocs)

	var cItemInnerSpacing_allocs *cgoAllocMap
	ref6e47ee0d.ItemInnerSpacing, cItemInnerSpacing_allocs = x.ItemInnerSpacing.PassValue()
	allocs6e47ee0d.Borrow(cItemInnerSpacing_allocs)

	var cTouchExtraPadding_allocs *cgoAllocMap
	ref6e47ee0d.TouchExtraPadding, cTouchExtraPadding_allocs = x.TouchExtraPadding.PassValue()
	allocs6e47ee0d.Borrow(cTouchExtraPadding_allocs)

	var cIndentSpacing_allocs *cgoAllocMap
	ref6e47ee0d.IndentSpacing, cIndentSpacing_allocs = (C.float)(x.IndentSpacing), cgoAllocsUnknown
	allocs6e47ee0d.Borrow(cIndentSpacing_allocs)

	var cColumnsMinSpacing_allocs *cgoAllocMap
	ref6e47ee0d.ColumnsMinSpacing, cColumnsMinSpacing_allocs = (C.float)(x.ColumnsMinSpacing), cgoAllocsUnknown
	allocs6e47ee0d.Borrow(cColumnsMinSpacing_allocs)

	var cScrollbarSize_allocs *cgoAllocMap
	ref6e47ee0d.ScrollbarSize, cScrollbarSize_allocs = (C.float)(x.ScrollbarSize), cgoAllocsUnknown
	allocs6e47ee0d.Borrow(cScrollbarSize_allocs)

	var cScrollbarRounding_allocs *cgoAllocMap
	ref6e47ee0d.ScrollbarRounding, cScrollbarRounding_allocs = (C.float)(x.ScrollbarRounding), cgoAllocsUnknown
	allocs6e47ee0d.Borrow(cScrollbarRounding_allocs)

	var cGrabMinSize_allocs *cgoAllocMap
	ref6e47ee0d.GrabMinSize, cGrabMinSize_allocs = (C.float)(x.GrabMinSize), cgoAllocsUnknown
	allocs6e47ee0d.Borrow(cGrabMinSize_allocs)

	var cGrabRounding_allocs *cgoAllocMap
	ref6e47ee0d.GrabRounding, cGrabRounding_allocs = (C.float)(x.GrabRounding), cgoAllocsUnknown
	allocs6e47ee0d.Borrow(cGrabRounding_allocs)

	var cButtonTextAlign_allocs *cgoAllocMap
	ref6e47ee0d.ButtonTextAlign, cButtonTextAlign_allocs = x.ButtonTextAlign.PassValue()
	allocs6e47ee0d.Borrow(cButtonTextAlign_allocs)

	var cDisplayWindowPadding_allocs *cgoAllocMap
	ref6e47ee0d.DisplayWindowPadding, cDisplayWindowPadding_allocs = x.DisplayWindowPadding.PassValue()
	allocs6e47ee0d.Borrow(cDisplayWindowPadding_allocs)

	var cDisplaySafeAreaPadding_allocs *cgoAllocMap
	ref6e47ee0d.DisplaySafeAreaPadding, cDisplaySafeAreaPadding_allocs = x.DisplaySafeAreaPadding.PassValue()
	allocs6e47ee0d.Borrow(cDisplaySafeAreaPadding_allocs)

	var cAntiAliasedLines_allocs *cgoAllocMap
	ref6e47ee0d.AntiAliasedLines, cAntiAliasedLines_allocs = (C._Bool)(x.AntiAliasedLines), cgoAllocsUnknown
	allocs6e47ee0d.Borrow(cAntiAliasedLines_allocs)

	var cAntiAliasedFill_allocs *cgoAllocMap
	ref6e47ee0d.AntiAliasedFill, cAntiAliasedFill_allocs = (C._Bool)(x.AntiAliasedFill), cgoAllocsUnknown
	allocs6e47ee0d.Borrow(cAntiAliasedFill_allocs)

	var cCurveTessellationTol_allocs *cgoAllocMap
	ref6e47ee0d.CurveTessellationTol, cCurveTessellationTol_allocs = (C.float)(x.CurveTessellationTol), cgoAllocsUnknown
	allocs6e47ee0d.Borrow(cCurveTessellationTol_allocs)

	var cColors_allocs *cgoAllocMap
	ref6e47ee0d.Colors, cColors_allocs = unpackA43Vec4(x.Colors)
	allocs6e47ee0d.Borrow(cColors_allocs)

	x.ref6e47ee0d = ref6e47ee0d
	x.allocs6e47ee0d = allocs6e47ee0d
	return ref6e47ee0d, allocs6e47ee0d

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x Style) PassValue() (C.struct_ImGuiStyle, *cgoAllocMap) {
	if x.ref6e47ee0d != nil {
		return *x.ref6e47ee0d, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *Style) Deref() {
	if x.ref6e47ee0d == nil {
		return
	}
	x.Alpha = (float32)(x.ref6e47ee0d.Alpha)
	x.WindowPadding = *NewVec2Ref(unsafe.Pointer(&x.ref6e47ee0d.WindowPadding))
	x.WindowRounding = (float32)(x.ref6e47ee0d.WindowRounding)
	x.WindowBorderSize = (float32)(x.ref6e47ee0d.WindowBorderSize)
	x.WindowMinSize = *NewVec2Ref(unsafe.Pointer(&x.ref6e47ee0d.WindowMinSize))
	x.WindowTitleAlign = *NewVec2Ref(unsafe.Pointer(&x.ref6e47ee0d.WindowTitleAlign))
	x.ChildRounding = (float32)(x.ref6e47ee0d.ChildRounding)
	x.ChildBorderSize = (float32)(x.ref6e47ee0d.ChildBorderSize)
	x.PopupRounding = (float32)(x.ref6e47ee0d.PopupRounding)
	x.PopupBorderSize = (float32)(x.ref6e47ee0d.PopupBorderSize)
	x.FramePadding = *NewVec2Ref(unsafe.Pointer(&x.ref6e47ee0d.FramePadding))
	x.FrameRounding = (float32)(x.ref6e47ee0d.FrameRounding)
	x.FrameBorderSize = (float32)(x.ref6e47ee0d.FrameBorderSize)
	x.ItemSpacing = *NewVec2Ref(unsafe.Pointer(&x.ref6e47ee0d.ItemSpacing))
	x.ItemInnerSpacing = *NewVec2Ref(unsafe.Pointer(&x.ref6e47ee0d.ItemInnerSpacing))
	x.TouchExtraPadding = *NewVec2Ref(unsafe.Pointer(&x.ref6e47ee0d.TouchExtraPadding))
	x.IndentSpacing = (float32)(x.ref6e47ee0d.IndentSpacing)
	x.ColumnsMinSpacing = (float32)(x.ref6e47ee0d.ColumnsMinSpacing)
	x.ScrollbarSize = (float32)(x.ref6e47ee0d.ScrollbarSize)
	x.ScrollbarRounding = (float32)(x.ref6e47ee0d.ScrollbarRounding)
	x.GrabMinSize = (float32)(x.ref6e47ee0d.GrabMinSize)
	x.GrabRounding = (float32)(x.ref6e47ee0d.GrabRounding)
	x.ButtonTextAlign = *NewVec2Ref(unsafe.Pointer(&x.ref6e47ee0d.ButtonTextAlign))
	x.DisplayWindowPadding = *NewVec2Ref(unsafe.Pointer(&x.ref6e47ee0d.DisplayWindowPadding))
	x.DisplaySafeAreaPadding = *NewVec2Ref(unsafe.Pointer(&x.ref6e47ee0d.DisplaySafeAreaPadding))
	x.AntiAliasedLines = (bool)(x.ref6e47ee0d.AntiAliasedLines)
	x.AntiAliasedFill = (bool)(x.ref6e47ee0d.AntiAliasedFill)
	x.CurveTessellationTol = (float32)(x.ref6e47ee0d.CurveTessellationTol)
	packA43Vec4(&x.Colors, (*[43]C.struct_ImVec4)(unsafe.Pointer(&x.ref6e47ee0d.Colors)))
}

// allocStructImGuiTextEditCallbackDataMemory allocates memory for type C.struct_ImGuiTextEditCallbackData in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStructImGuiTextEditCallbackDataMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStructImGuiTextEditCallbackDataValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStructImGuiTextEditCallbackDataValue = unsafe.Sizeof([1]C.struct_ImGuiTextEditCallbackData{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *TextEditCallbackData) Ref() *C.struct_ImGuiTextEditCallbackData {
	if x == nil {
		return nil
	}
	return x.ref61acec8
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *TextEditCallbackData) Free() {
	if x != nil && x.allocs61acec8 != nil {
		x.allocs61acec8.(*cgoAllocMap).Free()
		x.ref61acec8 = nil
	}
}

// NewTextEditCallbackDataRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewTextEditCallbackDataRef(ref unsafe.Pointer) *TextEditCallbackData {
	if ref == nil {
		return nil
	}
	obj := new(TextEditCallbackData)
	obj.ref61acec8 = (*C.struct_ImGuiTextEditCallbackData)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *TextEditCallbackData) PassRef() (*C.struct_ImGuiTextEditCallbackData, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref61acec8 != nil {
		return x.ref61acec8, nil
	}
	mem61acec8 := allocStructImGuiTextEditCallbackDataMemory(1)
	ref61acec8 := (*C.struct_ImGuiTextEditCallbackData)(mem61acec8)
	allocs61acec8 := new(cgoAllocMap)
	allocs61acec8.Add(mem61acec8)

	var cEventFlag_allocs *cgoAllocMap
	ref61acec8.EventFlag, cEventFlag_allocs = (C.ImGuiInputTextFlags)(x.EventFlag), cgoAllocsUnknown
	allocs61acec8.Borrow(cEventFlag_allocs)

	var cFlags_allocs *cgoAllocMap
	ref61acec8.Flags, cFlags_allocs = (C.ImGuiInputTextFlags)(x.Flags), cgoAllocsUnknown
	allocs61acec8.Borrow(cFlags_allocs)

	var cUserData_allocs *cgoAllocMap
	ref61acec8.UserData, cUserData_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.UserData)), cgoAllocsUnknown
	allocs61acec8.Borrow(cUserData_allocs)

	var cReadOnly_allocs *cgoAllocMap
	ref61acec8.ReadOnly, cReadOnly_allocs = (C._Bool)(x.ReadOnly), cgoAllocsUnknown
	allocs61acec8.Borrow(cReadOnly_allocs)

	var cEventChar_allocs *cgoAllocMap
	ref61acec8.EventChar, cEventChar_allocs = (C.ImWchar)(x.EventChar), cgoAllocsUnknown
	allocs61acec8.Borrow(cEventChar_allocs)

	var cEventKey_allocs *cgoAllocMap
	ref61acec8.EventKey, cEventKey_allocs = (C.ImGuiKey)(x.EventKey), cgoAllocsUnknown
	allocs61acec8.Borrow(cEventKey_allocs)

	var cBuf_allocs *cgoAllocMap
	ref61acec8.Buf, cBuf_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.Buf)).Data)), cgoAllocsUnknown
	allocs61acec8.Borrow(cBuf_allocs)

	var cBufTextLen_allocs *cgoAllocMap
	ref61acec8.BufTextLen, cBufTextLen_allocs = (C.int)(x.BufTextLen), cgoAllocsUnknown
	allocs61acec8.Borrow(cBufTextLen_allocs)

	var cBufSize_allocs *cgoAllocMap
	ref61acec8.BufSize, cBufSize_allocs = (C.int)(x.BufSize), cgoAllocsUnknown
	allocs61acec8.Borrow(cBufSize_allocs)

	var cBufDirty_allocs *cgoAllocMap
	ref61acec8.BufDirty, cBufDirty_allocs = (C._Bool)(x.BufDirty), cgoAllocsUnknown
	allocs61acec8.Borrow(cBufDirty_allocs)

	var cCursorPos_allocs *cgoAllocMap
	ref61acec8.CursorPos, cCursorPos_allocs = (C.int)(x.CursorPos), cgoAllocsUnknown
	allocs61acec8.Borrow(cCursorPos_allocs)

	var cSelectionStart_allocs *cgoAllocMap
	ref61acec8.SelectionStart, cSelectionStart_allocs = (C.int)(x.SelectionStart), cgoAllocsUnknown
	allocs61acec8.Borrow(cSelectionStart_allocs)

	var cSelectionEnd_allocs *cgoAllocMap
	ref61acec8.SelectionEnd, cSelectionEnd_allocs = (C.int)(x.SelectionEnd), cgoAllocsUnknown
	allocs61acec8.Borrow(cSelectionEnd_allocs)

	x.ref61acec8 = ref61acec8
	x.allocs61acec8 = allocs61acec8
	return ref61acec8, allocs61acec8

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x TextEditCallbackData) PassValue() (C.struct_ImGuiTextEditCallbackData, *cgoAllocMap) {
	if x.ref61acec8 != nil {
		return *x.ref61acec8, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *TextEditCallbackData) Deref() {
	if x.ref61acec8 == nil {
		return
	}
	x.EventFlag = (InputTextFlags)(x.ref61acec8.EventFlag)
	x.Flags = (InputTextFlags)(x.ref61acec8.Flags)
	x.UserData = (unsafe.Pointer)(unsafe.Pointer(x.ref61acec8.UserData))
	x.ReadOnly = (bool)(x.ref61acec8.ReadOnly)
	x.EventChar = (Wchar)(x.ref61acec8.EventChar)
	x.EventKey = (Key)(x.ref61acec8.EventKey)
	hxfa3f05c := (*sliceHeader)(unsafe.Pointer(&x.Buf))
	hxfa3f05c.Data = uintptr(unsafe.Pointer(x.ref61acec8.Buf))
	hxfa3f05c.Cap = 0x7fffffff
	// hxfa3f05c.Len = ?

	x.BufTextLen = (int32)(x.ref61acec8.BufTextLen)
	x.BufSize = (int32)(x.ref61acec8.BufSize)
	x.BufDirty = (bool)(x.ref61acec8.BufDirty)
	x.CursorPos = (int32)(x.ref61acec8.CursorPos)
	x.SelectionStart = (int32)(x.ref61acec8.SelectionStart)
	x.SelectionEnd = (int32)(x.ref61acec8.SelectionEnd)
}

// Ref returns a reference to C object as it is.
func (x *TextFilter) Ref() *C.struct_ImGuiTextFilter {
	if x == nil {
		return nil
	}
	return (*C.struct_ImGuiTextFilter)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *TextFilter) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewTextFilterRef converts the C object reference into a raw struct reference without wrapping.
func NewTextFilterRef(ref unsafe.Pointer) *TextFilter {
	return (*TextFilter)(ref)
}

// NewTextFilter allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewTextFilter() *TextFilter {
	return (*TextFilter)(allocStructImGuiTextFilterMemory(1))
}

// allocStructImGuiTextFilterMemory allocates memory for type C.struct_ImGuiTextFilter in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStructImGuiTextFilterMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStructImGuiTextFilterValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStructImGuiTextFilterValue = unsafe.Sizeof([1]C.struct_ImGuiTextFilter{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *TextFilter) PassRef() *C.struct_ImGuiTextFilter {
	if x == nil {
		x = (*TextFilter)(allocStructImGuiTextFilterMemory(1))
	}
	return (*C.struct_ImGuiTextFilter)(unsafe.Pointer(x))
}

// allocStructImVec2Memory allocates memory for type C.struct_ImVec2 in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStructImVec2Memory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStructImVec2Value))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStructImVec2Value = unsafe.Sizeof([1]C.struct_ImVec2{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *Vec2) Ref() *C.struct_ImVec2 {
	if x == nil {
		return nil
	}
	return x.ref74e98a33
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *Vec2) Free() {
	if x != nil && x.allocs74e98a33 != nil {
		x.allocs74e98a33.(*cgoAllocMap).Free()
		x.ref74e98a33 = nil
	}
}

// NewVec2Ref creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewVec2Ref(ref unsafe.Pointer) *Vec2 {
	if ref == nil {
		return nil
	}
	obj := new(Vec2)
	obj.ref74e98a33 = (*C.struct_ImVec2)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *Vec2) PassRef() (*C.struct_ImVec2, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref74e98a33 != nil {
		return x.ref74e98a33, nil
	}
	mem74e98a33 := allocStructImVec2Memory(1)
	ref74e98a33 := (*C.struct_ImVec2)(mem74e98a33)
	allocs74e98a33 := new(cgoAllocMap)
	allocs74e98a33.Add(mem74e98a33)

	var cx_allocs *cgoAllocMap
	ref74e98a33.x, cx_allocs = (C.float)(x.X), cgoAllocsUnknown
	allocs74e98a33.Borrow(cx_allocs)

	var cy_allocs *cgoAllocMap
	ref74e98a33.y, cy_allocs = (C.float)(x.Y), cgoAllocsUnknown
	allocs74e98a33.Borrow(cy_allocs)

	x.ref74e98a33 = ref74e98a33
	x.allocs74e98a33 = allocs74e98a33
	return ref74e98a33, allocs74e98a33

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x Vec2) PassValue() (C.struct_ImVec2, *cgoAllocMap) {
	if x.ref74e98a33 != nil {
		return *x.ref74e98a33, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *Vec2) Deref() {
	if x.ref74e98a33 == nil {
		return
	}
	x.X = (float32)(x.ref74e98a33.x)
	x.Y = (float32)(x.ref74e98a33.y)
}

// allocStructImVec4Memory allocates memory for type C.struct_ImVec4 in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStructImVec4Memory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStructImVec4Value))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStructImVec4Value = unsafe.Sizeof([1]C.struct_ImVec4{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *Vec4) Ref() *C.struct_ImVec4 {
	if x == nil {
		return nil
	}
	return x.ref9d8a2f06
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *Vec4) Free() {
	if x != nil && x.allocs9d8a2f06 != nil {
		x.allocs9d8a2f06.(*cgoAllocMap).Free()
		x.ref9d8a2f06 = nil
	}
}

// NewVec4Ref creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewVec4Ref(ref unsafe.Pointer) *Vec4 {
	if ref == nil {
		return nil
	}
	obj := new(Vec4)
	obj.ref9d8a2f06 = (*C.struct_ImVec4)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *Vec4) PassRef() (*C.struct_ImVec4, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref9d8a2f06 != nil {
		return x.ref9d8a2f06, nil
	}
	mem9d8a2f06 := allocStructImVec4Memory(1)
	ref9d8a2f06 := (*C.struct_ImVec4)(mem9d8a2f06)
	allocs9d8a2f06 := new(cgoAllocMap)
	allocs9d8a2f06.Add(mem9d8a2f06)

	var cx_allocs *cgoAllocMap
	ref9d8a2f06.x, cx_allocs = (C.float)(x.X), cgoAllocsUnknown
	allocs9d8a2f06.Borrow(cx_allocs)

	var cy_allocs *cgoAllocMap
	ref9d8a2f06.y, cy_allocs = (C.float)(x.Y), cgoAllocsUnknown
	allocs9d8a2f06.Borrow(cy_allocs)

	var cz_allocs *cgoAllocMap
	ref9d8a2f06.z, cz_allocs = (C.float)(x.Z), cgoAllocsUnknown
	allocs9d8a2f06.Borrow(cz_allocs)

	var cw_allocs *cgoAllocMap
	ref9d8a2f06.w, cw_allocs = (C.float)(x.W), cgoAllocsUnknown
	allocs9d8a2f06.Borrow(cw_allocs)

	x.ref9d8a2f06 = ref9d8a2f06
	x.allocs9d8a2f06 = allocs9d8a2f06
	return ref9d8a2f06, allocs9d8a2f06

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x Vec4) PassValue() (C.struct_ImVec4, *cgoAllocMap) {
	if x.ref9d8a2f06 != nil {
		return *x.ref9d8a2f06, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *Vec4) Deref() {
	if x.ref9d8a2f06 == nil {
		return
	}
	x.X = (float32)(x.ref9d8a2f06.x)
	x.Y = (float32)(x.ref9d8a2f06.y)
	x.Z = (float32)(x.ref9d8a2f06.z)
	x.W = (float32)(x.ref9d8a2f06.w)
}

// allocPCharMemory allocates memory for type *C.char in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPCharMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPCharValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPCharValue = unsafe.Sizeof([1]*C.char{})

// unpackArgSString transforms a sliced Go data structure into plain C format.
func unpackArgSString(x []string) (unpacked **C.char, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.char) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPCharMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.char)(unsafe.Pointer(h0))
	for i0 := range x {
		v0[i0], _ = unpackPCharString(x[i0])
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.char)(unsafe.Pointer(h.Data))
	return
}

// packSString reads sliced Go data structure out from plain C format.
func packSString(v []string, ptr0 **C.char) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.char)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = packPCharString(ptr1)
	}
}

// unpackArgSStyle transforms a sliced Go data structure into plain C format.
func unpackArgSStyle(x []Style) (unpacked *C.struct_ImGuiStyle, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.struct_ImGuiStyle) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocStructImGuiStyleMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.struct_ImGuiStyle)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.struct_ImGuiStyle)(unsafe.Pointer(h.Data))
	return
}

// packSStyle reads sliced Go data structure out from plain C format.
func packSStyle(v []Style, ptr0 *C.struct_ImGuiStyle) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfStructImGuiStyleValue]C.struct_ImGuiStyle)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewStyleRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSVec2 transforms a sliced Go data structure into plain C format.
func unpackArgSVec2(x []Vec2) (unpacked *C.struct_ImVec2, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.struct_ImVec2) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocStructImVec2Memory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.struct_ImVec2)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.struct_ImVec2)(unsafe.Pointer(h.Data))
	return
}

// packSVec2 reads sliced Go data structure out from plain C format.
func packSVec2(v []Vec2, ptr0 *C.struct_ImVec2) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfStructImVec2Value]C.struct_ImVec2)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewVec2Ref(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSVec4 transforms a sliced Go data structure into plain C format.
func unpackArgSVec4(x []Vec4) (unpacked *C.struct_ImVec4, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.struct_ImVec4) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocStructImVec4Memory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.struct_ImVec4)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.struct_ImVec4)(unsafe.Pointer(h.Data))
	return
}

// packSVec4 reads sliced Go data structure out from plain C format.
func packSVec4(v []Vec4, ptr0 *C.struct_ImVec4) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfStructImVec4Value]C.struct_ImVec4)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewVec4Ref(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSFontConfig transforms a sliced Go data structure into plain C format.
func unpackArgSFontConfig(x []FontConfig) (unpacked *C.struct_ImFontConfig, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.struct_ImFontConfig) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocStructImFontConfigMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.struct_ImFontConfig)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.struct_ImFontConfig)(unsafe.Pointer(h.Data))
	return
}

// packSFontConfig reads sliced Go data structure out from plain C format.
func packSFontConfig(v []FontConfig, ptr0 *C.struct_ImFontConfig) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfStructImFontConfigValue]C.struct_ImFontConfig)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewFontConfigRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSListClipper transforms a sliced Go data structure into plain C format.
func unpackArgSListClipper(x []ListClipper) (unpacked *C.struct_ImGuiListClipper, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.struct_ImGuiListClipper) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocStructImGuiListClipperMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.struct_ImGuiListClipper)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.struct_ImGuiListClipper)(unsafe.Pointer(h.Data))
	return
}

// packSListClipper reads sliced Go data structure out from plain C format.
func packSListClipper(v []ListClipper, ptr0 *C.struct_ImGuiListClipper) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfStructImGuiListClipperValue]C.struct_ImGuiListClipper)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewListClipperRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSDrawData transforms a sliced Go data structure into plain C format.
func unpackArgSDrawData(x []DrawData) (unpacked *C.struct_ImDrawData, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.struct_ImDrawData) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocStructImDrawDataMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.struct_ImDrawData)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.struct_ImDrawData)(unsafe.Pointer(h.Data))
	return
}

// packSDrawData reads sliced Go data structure out from plain C format.
func packSDrawData(v []DrawData, ptr0 *C.struct_ImDrawData) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfStructImDrawDataValue]C.struct_ImDrawData)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewDrawDataRef(unsafe.Pointer(&ptr1))
	}
}

// allocPUcharMemory allocates memory for type *C.uchar in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPUcharMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPUcharValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPUcharValue = unsafe.Sizeof([1]*C.uchar{})

// unpackArgSSByte transforms a sliced Go data structure into plain C format.
func unpackArgSSByte(x [][]byte) (unpacked **C.uchar, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.uchar) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPUcharMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.uchar)(unsafe.Pointer(h0))
	for i0 := range x {
		h := (*sliceHeader)(unsafe.Pointer(&x[i0]))
		v0[i0] = (*C.uchar)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.uchar)(unsafe.Pointer(h.Data))
	return
}

// packSSByte reads sliced Go data structure out from plain C format.
func packSSByte(v [][]byte, ptr0 **C.uchar) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.uchar)(unsafe.Pointer(ptr0)))[i0]
		hxf0d18b7 := (*sliceHeader)(unsafe.Pointer(&v[i0]))
		hxf0d18b7.Data = uintptr(unsafe.Pointer(ptr1))
		hxf0d18b7.Cap = 0x7fffffff
		// hxf0d18b7.Len = ?
	}
}
