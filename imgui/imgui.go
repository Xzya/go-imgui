// WARNING: This file has automatically been generated on Thu, 15 Feb 2018 23:59:51 EET.
// By https://git.io/c-for-go. DO NOT EDIT.

package imgui

/*
#include "cimgui/cimgui.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"runtime"
	"unsafe"
)

// GetIO function as declared in cimgui/cimgui.h:547
func GetIO() *IO {
	__ret := C.igGetIO()
	__v := NewIORef(unsafe.Pointer(__ret))
	return __v
}

// GetStyle function as declared in cimgui/cimgui.h:548
func GetStyle() *Style {
	__ret := C.igGetStyle()
	__v := NewStyleRef(unsafe.Pointer(__ret))
	return __v
}

// GetDrawData function as declared in cimgui/cimgui.h:549
func GetDrawData() *DrawData {
	__ret := C.igGetDrawData()
	__v := NewDrawDataRef(unsafe.Pointer(__ret))
	return __v
}

// NewFrame function as declared in cimgui/cimgui.h:550
func NewFrame() {
	C.igNewFrame()
}

// Render function as declared in cimgui/cimgui.h:551
func Render() {
	C.igRender()
}

// EndFrame function as declared in cimgui/cimgui.h:552
func EndFrame() {
	C.igEndFrame()
}

// Shutdown function as declared in cimgui/cimgui.h:553
func Shutdown() {
	C.igShutdown()
}

// ShowDemoWindow function as declared in cimgui/cimgui.h:556
func ShowDemoWindow(opened *bool) {
	copened, _ := (*C._Bool)(unsafe.Pointer(opened)), cgoAllocsUnknown
	C.igShowDemoWindow(copened)
}

// ShowMetricsWindow function as declared in cimgui/cimgui.h:557
func ShowMetricsWindow(opened *bool) {
	copened, _ := (*C._Bool)(unsafe.Pointer(opened)), cgoAllocsUnknown
	C.igShowMetricsWindow(copened)
}

// ShowStyleEditor function as declared in cimgui/cimgui.h:558
func ShowStyleEditor(ref *Style) {
	cref, _ := ref.PassRef()
	C.igShowStyleEditor(cref)
}

// ShowStyleSelector function as declared in cimgui/cimgui.h:559
func ShowStyleSelector(label string) {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	C.igShowStyleSelector(clabel)
	runtime.KeepAlive(label)
}

// ShowFontSelector function as declared in cimgui/cimgui.h:560
func ShowFontSelector(label string) {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	C.igShowFontSelector(clabel)
	runtime.KeepAlive(label)
}

// ShowUserGuide function as declared in cimgui/cimgui.h:561
func ShowUserGuide() {
	C.igShowUserGuide()
}

// Begin function as declared in cimgui/cimgui.h:564
func Begin(name string, pOpen *bool, flags WindowFlags) bool {
	name = safeString(name)
	cname, _ := unpackPCharString(name)
	cpOpen, _ := (*C._Bool)(unsafe.Pointer(pOpen)), cgoAllocsUnknown
	cflags, _ := (C.ImGuiWindowFlags)(flags), cgoAllocsUnknown
	__ret := C.igBegin(cname, cpOpen, cflags)
	runtime.KeepAlive(name)
	__v := (bool)(__ret)
	return __v
}

// End function as declared in cimgui/cimgui.h:567
func End() {
	C.igEnd()
}

// BeginChild function as declared in cimgui/cimgui.h:568
func BeginChild(strId string, size Vec2, border bool, extraFlags WindowFlags) bool {
	strId = safeString(strId)
	cstrId, _ := unpackPCharString(strId)
	csize, _ := size.PassValue()
	cborder, _ := (C._Bool)(border), cgoAllocsUnknown
	cextraFlags, _ := (C.ImGuiWindowFlags)(extraFlags), cgoAllocsUnknown
	__ret := C.igBeginChild(cstrId, csize, cborder, cextraFlags)
	runtime.KeepAlive(strId)
	__v := (bool)(__ret)
	return __v
}

// BeginChildEx function as declared in cimgui/cimgui.h:569
func BeginChildEx(id ID, size Vec2, border bool, extraFlags WindowFlags) bool {
	cid, _ := (C.ImGuiID)(id), cgoAllocsUnknown
	csize, _ := size.PassValue()
	cborder, _ := (C._Bool)(border), cgoAllocsUnknown
	cextraFlags, _ := (C.ImGuiWindowFlags)(extraFlags), cgoAllocsUnknown
	__ret := C.igBeginChildEx(cid, csize, cborder, cextraFlags)
	__v := (bool)(__ret)
	return __v
}

// EndChild function as declared in cimgui/cimgui.h:570
func EndChild() {
	C.igEndChild()
}

// GetContentRegionMax function as declared in cimgui/cimgui.h:571
func GetContentRegionMax(out *Vec2) {
	cout, _ := out.PassRef()
	C.igGetContentRegionMax(cout)
}

// GetContentRegionAvail function as declared in cimgui/cimgui.h:572
func GetContentRegionAvail(out *Vec2) {
	cout, _ := out.PassRef()
	C.igGetContentRegionAvail(cout)
}

// GetContentRegionAvailWidth function as declared in cimgui/cimgui.h:573
func GetContentRegionAvailWidth() float32 {
	__ret := C.igGetContentRegionAvailWidth()
	__v := (float32)(__ret)
	return __v
}

// GetWindowContentRegionMin function as declared in cimgui/cimgui.h:574
func GetWindowContentRegionMin(out *Vec2) {
	cout, _ := out.PassRef()
	C.igGetWindowContentRegionMin(cout)
}

// GetWindowContentRegionMax function as declared in cimgui/cimgui.h:575
func GetWindowContentRegionMax(out *Vec2) {
	cout, _ := out.PassRef()
	C.igGetWindowContentRegionMax(cout)
}

// GetWindowContentRegionWidth function as declared in cimgui/cimgui.h:576
func GetWindowContentRegionWidth() float32 {
	__ret := C.igGetWindowContentRegionWidth()
	__v := (float32)(__ret)
	return __v
}

// GetWindowDrawList function as declared in cimgui/cimgui.h:577
func GetWindowDrawList() *DrawList {
	__ret := C.igGetWindowDrawList()
	__v := *(**DrawList)(unsafe.Pointer(&__ret))
	return __v
}

// GetWindowPos function as declared in cimgui/cimgui.h:578
func GetWindowPos(out *Vec2) {
	cout, _ := out.PassRef()
	C.igGetWindowPos(cout)
}

// GetWindowSize function as declared in cimgui/cimgui.h:579
func GetWindowSize(out *Vec2) {
	cout, _ := out.PassRef()
	C.igGetWindowSize(cout)
}

// GetWindowWidth function as declared in cimgui/cimgui.h:580
func GetWindowWidth() float32 {
	__ret := C.igGetWindowWidth()
	__v := (float32)(__ret)
	return __v
}

// GetWindowHeight function as declared in cimgui/cimgui.h:581
func GetWindowHeight() float32 {
	__ret := C.igGetWindowHeight()
	__v := (float32)(__ret)
	return __v
}

// IsWindowCollapsed function as declared in cimgui/cimgui.h:582
func IsWindowCollapsed() bool {
	__ret := C.igIsWindowCollapsed()
	__v := (bool)(__ret)
	return __v
}

// IsWindowAppearing function as declared in cimgui/cimgui.h:583
func IsWindowAppearing() bool {
	__ret := C.igIsWindowAppearing()
	__v := (bool)(__ret)
	return __v
}

// SetWindowFontScale function as declared in cimgui/cimgui.h:584
func SetWindowFontScale(scale float32) {
	cscale, _ := (C.float)(scale), cgoAllocsUnknown
	C.igSetWindowFontScale(cscale)
}

// SetNextWindowPos function as declared in cimgui/cimgui.h:586
func SetNextWindowPos(pos Vec2, cond Cond, pivot Vec2) {
	cpos, _ := pos.PassValue()
	ccond, _ := (C.ImGuiCond)(cond), cgoAllocsUnknown
	cpivot, _ := pivot.PassValue()
	C.igSetNextWindowPos(cpos, ccond, cpivot)
}

// SetNextWindowSize function as declared in cimgui/cimgui.h:587
func SetNextWindowSize(size Vec2, cond Cond) {
	csize, _ := size.PassValue()
	ccond, _ := (C.ImGuiCond)(cond), cgoAllocsUnknown
	C.igSetNextWindowSize(csize, ccond)
}

// SetNextWindowSizeConstraints function as declared in cimgui/cimgui.h:588
func SetNextWindowSizeConstraints(sizeMin Vec2, sizeMax Vec2, customCallback SizeConstraintCallback, customCallbackData unsafe.Pointer) {
	csizeMin, _ := sizeMin.PassValue()
	csizeMax, _ := sizeMax.PassValue()
	ccustomCallback, _ := customCallback.PassValue()
	ccustomCallbackData, _ := customCallbackData, cgoAllocsUnknown
	C.igSetNextWindowSizeConstraints(csizeMin, csizeMax, ccustomCallback, ccustomCallbackData)
}

// SetNextWindowContentSize function as declared in cimgui/cimgui.h:589
func SetNextWindowContentSize(size Vec2) {
	csize, _ := size.PassValue()
	C.igSetNextWindowContentSize(csize)
}

// SetNextWindowCollapsed function as declared in cimgui/cimgui.h:590
func SetNextWindowCollapsed(collapsed bool, cond Cond) {
	ccollapsed, _ := (C._Bool)(collapsed), cgoAllocsUnknown
	ccond, _ := (C.ImGuiCond)(cond), cgoAllocsUnknown
	C.igSetNextWindowCollapsed(ccollapsed, ccond)
}

// SetNextWindowFocus function as declared in cimgui/cimgui.h:591
func SetNextWindowFocus() {
	C.igSetNextWindowFocus()
}

// SetWindowPos function as declared in cimgui/cimgui.h:592
func SetWindowPos(pos Vec2, cond Cond) {
	cpos, _ := pos.PassValue()
	ccond, _ := (C.ImGuiCond)(cond), cgoAllocsUnknown
	C.igSetWindowPos(cpos, ccond)
}

// SetWindowSize function as declared in cimgui/cimgui.h:593
func SetWindowSize(size Vec2, cond Cond) {
	csize, _ := size.PassValue()
	ccond, _ := (C.ImGuiCond)(cond), cgoAllocsUnknown
	C.igSetWindowSize(csize, ccond)
}

// SetWindowCollapsed function as declared in cimgui/cimgui.h:594
func SetWindowCollapsed(collapsed bool, cond Cond) {
	ccollapsed, _ := (C._Bool)(collapsed), cgoAllocsUnknown
	ccond, _ := (C.ImGuiCond)(cond), cgoAllocsUnknown
	C.igSetWindowCollapsed(ccollapsed, ccond)
}

// SetWindowFocus function as declared in cimgui/cimgui.h:595
func SetWindowFocus() {
	C.igSetWindowFocus()
}

// SetWindowPosByName function as declared in cimgui/cimgui.h:596
func SetWindowPosByName(name string, pos Vec2, cond Cond) {
	name = safeString(name)
	cname, _ := unpackPCharString(name)
	cpos, _ := pos.PassValue()
	ccond, _ := (C.ImGuiCond)(cond), cgoAllocsUnknown
	C.igSetWindowPosByName(cname, cpos, ccond)
	runtime.KeepAlive(name)
}

// SetWindowSize2 function as declared in cimgui/cimgui.h:597
func SetWindowSize2(name string, size Vec2, cond Cond) {
	name = safeString(name)
	cname, _ := unpackPCharString(name)
	csize, _ := size.PassValue()
	ccond, _ := (C.ImGuiCond)(cond), cgoAllocsUnknown
	C.igSetWindowSize2(cname, csize, ccond)
	runtime.KeepAlive(name)
}

// SetWindowCollapsed2 function as declared in cimgui/cimgui.h:598
func SetWindowCollapsed2(name string, collapsed bool, cond Cond) {
	name = safeString(name)
	cname, _ := unpackPCharString(name)
	ccollapsed, _ := (C._Bool)(collapsed), cgoAllocsUnknown
	ccond, _ := (C.ImGuiCond)(cond), cgoAllocsUnknown
	C.igSetWindowCollapsed2(cname, ccollapsed, ccond)
	runtime.KeepAlive(name)
}

// SetWindowFocus2 function as declared in cimgui/cimgui.h:599
func SetWindowFocus2(name string) {
	name = safeString(name)
	cname, _ := unpackPCharString(name)
	C.igSetWindowFocus2(cname)
	runtime.KeepAlive(name)
}

// GetScrollX function as declared in cimgui/cimgui.h:601
func GetScrollX() float32 {
	__ret := C.igGetScrollX()
	__v := (float32)(__ret)
	return __v
}

// GetScrollY function as declared in cimgui/cimgui.h:602
func GetScrollY() float32 {
	__ret := C.igGetScrollY()
	__v := (float32)(__ret)
	return __v
}

// GetScrollMaxX function as declared in cimgui/cimgui.h:603
func GetScrollMaxX() float32 {
	__ret := C.igGetScrollMaxX()
	__v := (float32)(__ret)
	return __v
}

// GetScrollMaxY function as declared in cimgui/cimgui.h:604
func GetScrollMaxY() float32 {
	__ret := C.igGetScrollMaxY()
	__v := (float32)(__ret)
	return __v
}

// SetScrollX function as declared in cimgui/cimgui.h:605
func SetScrollX(scrollX float32) {
	cscrollX, _ := (C.float)(scrollX), cgoAllocsUnknown
	C.igSetScrollX(cscrollX)
}

// SetScrollY function as declared in cimgui/cimgui.h:606
func SetScrollY(scrollY float32) {
	cscrollY, _ := (C.float)(scrollY), cgoAllocsUnknown
	C.igSetScrollY(cscrollY)
}

// SetScrollHere function as declared in cimgui/cimgui.h:607
func SetScrollHere(centerYRatio float32) {
	ccenterYRatio, _ := (C.float)(centerYRatio), cgoAllocsUnknown
	C.igSetScrollHere(ccenterYRatio)
}

// SetScrollFromPosY function as declared in cimgui/cimgui.h:608
func SetScrollFromPosY(posY float32, centerYRatio float32) {
	cposY, _ := (C.float)(posY), cgoAllocsUnknown
	ccenterYRatio, _ := (C.float)(centerYRatio), cgoAllocsUnknown
	C.igSetScrollFromPosY(cposY, ccenterYRatio)
}

// SetStateStorage function as declared in cimgui/cimgui.h:609
func SetStateStorage(tree *Storage) {
	ctree, _ := (*C.struct_ImGuiStorage)(unsafe.Pointer(tree)), cgoAllocsUnknown
	C.igSetStateStorage(ctree)
}

// GetStateStorage function as declared in cimgui/cimgui.h:610
func GetStateStorage() *Storage {
	__ret := C.igGetStateStorage()
	__v := *(**Storage)(unsafe.Pointer(&__ret))
	return __v
}

// PushFont function as declared in cimgui/cimgui.h:613
func PushFont(font *Font) {
	cfont, _ := (*C.struct_ImFont)(unsafe.Pointer(font)), cgoAllocsUnknown
	C.igPushFont(cfont)
}

// PopFont function as declared in cimgui/cimgui.h:614
func PopFont() {
	C.igPopFont()
}

// PushStyleColorU32 function as declared in cimgui/cimgui.h:615
func PushStyleColorU32(idx Col, col U32) {
	cidx, _ := (C.ImGuiCol)(idx), cgoAllocsUnknown
	ccol, _ := (C.ImU32)(col), cgoAllocsUnknown
	C.igPushStyleColorU32(cidx, ccol)
}

// PushStyleColor function as declared in cimgui/cimgui.h:616
func PushStyleColor(idx Col, col Vec4) {
	cidx, _ := (C.ImGuiCol)(idx), cgoAllocsUnknown
	ccol, _ := col.PassValue()
	C.igPushStyleColor(cidx, ccol)
}

// PopStyleColor function as declared in cimgui/cimgui.h:617
func PopStyleColor(count int32) {
	ccount, _ := (C.int)(count), cgoAllocsUnknown
	C.igPopStyleColor(ccount)
}

// PushStyleVar function as declared in cimgui/cimgui.h:618
func PushStyleVar(idx StyleVar, val float32) {
	cidx, _ := (C.ImGuiStyleVar)(idx), cgoAllocsUnknown
	cval, _ := (C.float)(val), cgoAllocsUnknown
	C.igPushStyleVar(cidx, cval)
}

// PushStyleVarVec function as declared in cimgui/cimgui.h:619
func PushStyleVarVec(idx StyleVar, val Vec2) {
	cidx, _ := (C.ImGuiStyleVar)(idx), cgoAllocsUnknown
	cval, _ := val.PassValue()
	C.igPushStyleVarVec(cidx, cval)
}

// PopStyleVar function as declared in cimgui/cimgui.h:620
func PopStyleVar(count int32) {
	ccount, _ := (C.int)(count), cgoAllocsUnknown
	C.igPopStyleVar(ccount)
}

// GetStyleColorVec4 function as declared in cimgui/cimgui.h:621
func GetStyleColorVec4(pOut *Vec4, idx Col) {
	cpOut, _ := pOut.PassRef()
	cidx, _ := (C.ImGuiCol)(idx), cgoAllocsUnknown
	C.igGetStyleColorVec4(cpOut, cidx)
}

// GetFont function as declared in cimgui/cimgui.h:622
func GetFont() *Font {
	__ret := C.igGetFont()
	__v := *(**Font)(unsafe.Pointer(&__ret))
	return __v
}

// GetFontSize function as declared in cimgui/cimgui.h:623
func GetFontSize() float32 {
	__ret := C.igGetFontSize()
	__v := (float32)(__ret)
	return __v
}

// GetFontTexUvWhitePixel function as declared in cimgui/cimgui.h:624
func GetFontTexUvWhitePixel(pOut *Vec2) {
	cpOut, _ := pOut.PassRef()
	C.igGetFontTexUvWhitePixel(cpOut)
}

// GetColorU32 function as declared in cimgui/cimgui.h:625
func GetColorU32(idx Col, alphaMul float32) U32 {
	cidx, _ := (C.ImGuiCol)(idx), cgoAllocsUnknown
	calphaMul, _ := (C.float)(alphaMul), cgoAllocsUnknown
	__ret := C.igGetColorU32(cidx, calphaMul)
	__v := (U32)(__ret)
	return __v
}

// GetColorU32Vec function as declared in cimgui/cimgui.h:626
func GetColorU32Vec(col *Vec4) U32 {
	ccol, _ := col.PassRef()
	__ret := C.igGetColorU32Vec(ccol)
	__v := (U32)(__ret)
	return __v
}

// GetColorU32U32 function as declared in cimgui/cimgui.h:627
func GetColorU32U32(col U32) U32 {
	ccol, _ := (C.ImU32)(col), cgoAllocsUnknown
	__ret := C.igGetColorU32U32(ccol)
	__v := (U32)(__ret)
	return __v
}

// PushItemWidth function as declared in cimgui/cimgui.h:630
func PushItemWidth(itemWidth float32) {
	citemWidth, _ := (C.float)(itemWidth), cgoAllocsUnknown
	C.igPushItemWidth(citemWidth)
}

// PopItemWidth function as declared in cimgui/cimgui.h:631
func PopItemWidth() {
	C.igPopItemWidth()
}

// CalcItemWidth function as declared in cimgui/cimgui.h:632
func CalcItemWidth() float32 {
	__ret := C.igCalcItemWidth()
	__v := (float32)(__ret)
	return __v
}

// PushTextWrapPos function as declared in cimgui/cimgui.h:633
func PushTextWrapPos(wrapPosX float32) {
	cwrapPosX, _ := (C.float)(wrapPosX), cgoAllocsUnknown
	C.igPushTextWrapPos(cwrapPosX)
}

// PopTextWrapPos function as declared in cimgui/cimgui.h:634
func PopTextWrapPos() {
	C.igPopTextWrapPos()
}

// PushAllowKeyboardFocus function as declared in cimgui/cimgui.h:635
func PushAllowKeyboardFocus(v bool) {
	cv, _ := (C._Bool)(v), cgoAllocsUnknown
	C.igPushAllowKeyboardFocus(cv)
}

// PopAllowKeyboardFocus function as declared in cimgui/cimgui.h:636
func PopAllowKeyboardFocus() {
	C.igPopAllowKeyboardFocus()
}

// PushButtonRepeat function as declared in cimgui/cimgui.h:637
func PushButtonRepeat(repeat bool) {
	crepeat, _ := (C._Bool)(repeat), cgoAllocsUnknown
	C.igPushButtonRepeat(crepeat)
}

// PopButtonRepeat function as declared in cimgui/cimgui.h:638
func PopButtonRepeat() {
	C.igPopButtonRepeat()
}

// Separator function as declared in cimgui/cimgui.h:641
func Separator() {
	C.igSeparator()
}

// SameLine function as declared in cimgui/cimgui.h:642
func SameLine(posX float32, spacingW float32) {
	cposX, _ := (C.float)(posX), cgoAllocsUnknown
	cspacingW, _ := (C.float)(spacingW), cgoAllocsUnknown
	C.igSameLine(cposX, cspacingW)
}

// NewLine function as declared in cimgui/cimgui.h:643
func NewLine() {
	C.igNewLine()
}

// Spacing function as declared in cimgui/cimgui.h:644
func Spacing() {
	C.igSpacing()
}

// Dummy function as declared in cimgui/cimgui.h:645
func Dummy(size *Vec2) {
	csize, _ := size.PassRef()
	C.igDummy(csize)
}

// Indent function as declared in cimgui/cimgui.h:646
func Indent(indentW float32) {
	cindentW, _ := (C.float)(indentW), cgoAllocsUnknown
	C.igIndent(cindentW)
}

// Unindent function as declared in cimgui/cimgui.h:647
func Unindent(indentW float32) {
	cindentW, _ := (C.float)(indentW), cgoAllocsUnknown
	C.igUnindent(cindentW)
}

// BeginGroup function as declared in cimgui/cimgui.h:648
func BeginGroup() {
	C.igBeginGroup()
}

// EndGroup function as declared in cimgui/cimgui.h:649
func EndGroup() {
	C.igEndGroup()
}

// GetCursorPos function as declared in cimgui/cimgui.h:650
func GetCursorPos(pOut *Vec2) {
	cpOut, _ := pOut.PassRef()
	C.igGetCursorPos(cpOut)
}

// GetCursorPosX function as declared in cimgui/cimgui.h:651
func GetCursorPosX() float32 {
	__ret := C.igGetCursorPosX()
	__v := (float32)(__ret)
	return __v
}

// GetCursorPosY function as declared in cimgui/cimgui.h:652
func GetCursorPosY() float32 {
	__ret := C.igGetCursorPosY()
	__v := (float32)(__ret)
	return __v
}

// SetCursorPos function as declared in cimgui/cimgui.h:653
func SetCursorPos(localPos Vec2) {
	clocalPos, _ := localPos.PassValue()
	C.igSetCursorPos(clocalPos)
}

// SetCursorPosX function as declared in cimgui/cimgui.h:654
func SetCursorPosX(x float32) {
	cx, _ := (C.float)(x), cgoAllocsUnknown
	C.igSetCursorPosX(cx)
}

// SetCursorPosY function as declared in cimgui/cimgui.h:655
func SetCursorPosY(y float32) {
	cy, _ := (C.float)(y), cgoAllocsUnknown
	C.igSetCursorPosY(cy)
}

// GetCursorStartPos function as declared in cimgui/cimgui.h:656
func GetCursorStartPos(pOut *Vec2) {
	cpOut, _ := pOut.PassRef()
	C.igGetCursorStartPos(cpOut)
}

// GetCursorScreenPos function as declared in cimgui/cimgui.h:657
func GetCursorScreenPos(pOut *Vec2) {
	cpOut, _ := pOut.PassRef()
	C.igGetCursorScreenPos(cpOut)
}

// SetCursorScreenPos function as declared in cimgui/cimgui.h:658
func SetCursorScreenPos(pos Vec2) {
	cpos, _ := pos.PassValue()
	C.igSetCursorScreenPos(cpos)
}

// AlignTextToFramePadding function as declared in cimgui/cimgui.h:659
func AlignTextToFramePadding() {
	C.igAlignTextToFramePadding()
}

// GetTextLineHeight function as declared in cimgui/cimgui.h:660
func GetTextLineHeight() float32 {
	__ret := C.igGetTextLineHeight()
	__v := (float32)(__ret)
	return __v
}

// GetTextLineHeightWithSpacing function as declared in cimgui/cimgui.h:661
func GetTextLineHeightWithSpacing() float32 {
	__ret := C.igGetTextLineHeightWithSpacing()
	__v := (float32)(__ret)
	return __v
}

// GetFrameHeight function as declared in cimgui/cimgui.h:662
func GetFrameHeight() float32 {
	__ret := C.igGetFrameHeight()
	__v := (float32)(__ret)
	return __v
}

// GetFrameHeightWithSpacing function as declared in cimgui/cimgui.h:663
func GetFrameHeightWithSpacing() float32 {
	__ret := C.igGetFrameHeightWithSpacing()
	__v := (float32)(__ret)
	return __v
}

// Columns function as declared in cimgui/cimgui.h:666
func Columns(count int32, id string, border bool) {
	ccount, _ := (C.int)(count), cgoAllocsUnknown
	id = safeString(id)
	cid, _ := unpackPCharString(id)
	cborder, _ := (C._Bool)(border), cgoAllocsUnknown
	C.igColumns(ccount, cid, cborder)
	runtime.KeepAlive(id)
}

// NextColumn function as declared in cimgui/cimgui.h:667
func NextColumn() {
	C.igNextColumn()
}

// GetColumnIndex function as declared in cimgui/cimgui.h:668
func GetColumnIndex() int32 {
	__ret := C.igGetColumnIndex()
	__v := (int32)(__ret)
	return __v
}

// GetColumnWidth function as declared in cimgui/cimgui.h:669
func GetColumnWidth(columnIndex int32) float32 {
	ccolumnIndex, _ := (C.int)(columnIndex), cgoAllocsUnknown
	__ret := C.igGetColumnWidth(ccolumnIndex)
	__v := (float32)(__ret)
	return __v
}

// SetColumnWidth function as declared in cimgui/cimgui.h:670
func SetColumnWidth(columnIndex int32, width float32) {
	ccolumnIndex, _ := (C.int)(columnIndex), cgoAllocsUnknown
	cwidth, _ := (C.float)(width), cgoAllocsUnknown
	C.igSetColumnWidth(ccolumnIndex, cwidth)
}

// GetColumnOffset function as declared in cimgui/cimgui.h:671
func GetColumnOffset(columnIndex int32) float32 {
	ccolumnIndex, _ := (C.int)(columnIndex), cgoAllocsUnknown
	__ret := C.igGetColumnOffset(ccolumnIndex)
	__v := (float32)(__ret)
	return __v
}

// SetColumnOffset function as declared in cimgui/cimgui.h:672
func SetColumnOffset(columnIndex int32, offsetX float32) {
	ccolumnIndex, _ := (C.int)(columnIndex), cgoAllocsUnknown
	coffsetX, _ := (C.float)(offsetX), cgoAllocsUnknown
	C.igSetColumnOffset(ccolumnIndex, coffsetX)
}

// GetColumnsCount function as declared in cimgui/cimgui.h:673
func GetColumnsCount() int32 {
	__ret := C.igGetColumnsCount()
	__v := (int32)(__ret)
	return __v
}

// PushIDStr function as declared in cimgui/cimgui.h:678
func PushIDStr(strId string) {
	strId = safeString(strId)
	cstrId, _ := unpackPCharString(strId)
	C.igPushIDStr(cstrId)
	runtime.KeepAlive(strId)
}

// PushIDStrRange function as declared in cimgui/cimgui.h:679
func PushIDStrRange(strBegin string, strEnd string) {
	strBegin = safeString(strBegin)
	cstrBegin, _ := unpackPCharString(strBegin)
	strEnd = safeString(strEnd)
	cstrEnd, _ := unpackPCharString(strEnd)
	C.igPushIDStrRange(cstrBegin, cstrEnd)
	runtime.KeepAlive(strEnd)
	runtime.KeepAlive(strBegin)
}

// PushIDPtr function as declared in cimgui/cimgui.h:680
func PushIDPtr(ptrId unsafe.Pointer) {
	cptrId, _ := ptrId, cgoAllocsUnknown
	C.igPushIDPtr(cptrId)
}

// PushIDInt function as declared in cimgui/cimgui.h:681
func PushIDInt(intId int32) {
	cintId, _ := (C.int)(intId), cgoAllocsUnknown
	C.igPushIDInt(cintId)
}

// PopID function as declared in cimgui/cimgui.h:682
func PopID() {
	C.igPopID()
}

// GetIDStr function as declared in cimgui/cimgui.h:683
func GetIDStr(strId string) ID {
	strId = safeString(strId)
	cstrId, _ := unpackPCharString(strId)
	__ret := C.igGetIDStr(cstrId)
	runtime.KeepAlive(strId)
	__v := (ID)(__ret)
	return __v
}

// GetIDStrRange function as declared in cimgui/cimgui.h:684
func GetIDStrRange(strBegin string, strEnd string) ID {
	strBegin = safeString(strBegin)
	cstrBegin, _ := unpackPCharString(strBegin)
	strEnd = safeString(strEnd)
	cstrEnd, _ := unpackPCharString(strEnd)
	__ret := C.igGetIDStrRange(cstrBegin, cstrEnd)
	runtime.KeepAlive(strEnd)
	runtime.KeepAlive(strBegin)
	__v := (ID)(__ret)
	return __v
}

// GetIDPtr function as declared in cimgui/cimgui.h:685
func GetIDPtr(ptrId unsafe.Pointer) ID {
	cptrId, _ := ptrId, cgoAllocsUnknown
	__ret := C.igGetIDPtr(cptrId)
	__v := (ID)(__ret)
	return __v
}

// Bullet function as declared in cimgui/cimgui.h:701
func Bullet() {
	C.igBullet()
}

// Button function as declared in cimgui/cimgui.h:704
func Button(label string, size Vec2) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	csize, _ := size.PassValue()
	__ret := C.igButton(clabel, csize)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// SmallButton function as declared in cimgui/cimgui.h:705
func SmallButton(label string) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	__ret := C.igSmallButton(clabel)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// InvisibleButton function as declared in cimgui/cimgui.h:706
func InvisibleButton(strId string, size Vec2) bool {
	strId = safeString(strId)
	cstrId, _ := unpackPCharString(strId)
	csize, _ := size.PassValue()
	__ret := C.igInvisibleButton(cstrId, csize)
	runtime.KeepAlive(strId)
	__v := (bool)(__ret)
	return __v
}

// Image function as declared in cimgui/cimgui.h:707
func Image(userTextureId *TextureID, size Vec2, uv0 Vec2, uv1 Vec2, tintCol Vec4, borderCol Vec4) {
	cuserTextureId, _ := (C.ImTextureID)(unsafe.Pointer(userTextureId)), cgoAllocsUnknown
	csize, _ := size.PassValue()
	cuv0, _ := uv0.PassValue()
	cuv1, _ := uv1.PassValue()
	ctintCol, _ := tintCol.PassValue()
	cborderCol, _ := borderCol.PassValue()
	C.igImage(cuserTextureId, csize, cuv0, cuv1, ctintCol, cborderCol)
}

// ImageButton function as declared in cimgui/cimgui.h:708
func ImageButton(userTextureId *TextureID, size Vec2, uv0 Vec2, uv1 Vec2, framePadding int32, bgCol Vec4, tintCol Vec4) bool {
	cuserTextureId, _ := (C.ImTextureID)(unsafe.Pointer(userTextureId)), cgoAllocsUnknown
	csize, _ := size.PassValue()
	cuv0, _ := uv0.PassValue()
	cuv1, _ := uv1.PassValue()
	cframePadding, _ := (C.int)(framePadding), cgoAllocsUnknown
	cbgCol, _ := bgCol.PassValue()
	ctintCol, _ := tintCol.PassValue()
	__ret := C.igImageButton(cuserTextureId, csize, cuv0, cuv1, cframePadding, cbgCol, ctintCol)
	__v := (bool)(__ret)
	return __v
}

// Checkbox function as declared in cimgui/cimgui.h:709
func Checkbox(label string, v *bool) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cv, _ := (*C._Bool)(unsafe.Pointer(v)), cgoAllocsUnknown
	__ret := C.igCheckbox(clabel, cv)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// CheckboxFlags function as declared in cimgui/cimgui.h:710
func CheckboxFlags(label string, flags *uint32, flagsValue uint32) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cflags, _ := (*C.uint)(unsafe.Pointer(flags)), cgoAllocsUnknown
	cflagsValue, _ := (C.uint)(flagsValue), cgoAllocsUnknown
	__ret := C.igCheckboxFlags(clabel, cflags, cflagsValue)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// RadioButtonBool function as declared in cimgui/cimgui.h:711
func RadioButtonBool(label string, active bool) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cactive, _ := (C._Bool)(active), cgoAllocsUnknown
	__ret := C.igRadioButtonBool(clabel, cactive)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// RadioButton function as declared in cimgui/cimgui.h:712
func RadioButton(label string, v *int32, vButton int32) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cv, _ := (*C.int)(unsafe.Pointer(v)), cgoAllocsUnknown
	cvButton, _ := (C.int)(vButton), cgoAllocsUnknown
	__ret := C.igRadioButton(clabel, cv, cvButton)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// PlotLines function as declared in cimgui/cimgui.h:713
func PlotLines(label string, values *float32, valuesCount int32, valuesOffset int32, overlayText string, scaleMin float32, scaleMax float32, graphSize Vec2, stride int32) {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cvalues, _ := (*C.float)(unsafe.Pointer(values)), cgoAllocsUnknown
	cvaluesCount, _ := (C.int)(valuesCount), cgoAllocsUnknown
	cvaluesOffset, _ := (C.int)(valuesOffset), cgoAllocsUnknown
	overlayText = safeString(overlayText)
	coverlayText, _ := unpackPCharString(overlayText)
	cscaleMin, _ := (C.float)(scaleMin), cgoAllocsUnknown
	cscaleMax, _ := (C.float)(scaleMax), cgoAllocsUnknown
	cgraphSize, _ := graphSize.PassValue()
	cstride, _ := (C.int)(stride), cgoAllocsUnknown
	C.igPlotLines(clabel, cvalues, cvaluesCount, cvaluesOffset, coverlayText, cscaleMin, cscaleMax, cgraphSize, cstride)
	runtime.KeepAlive(overlayText)
	runtime.KeepAlive(label)
}

// PlotHistogram function as declared in cimgui/cimgui.h:715
func PlotHistogram(label string, values *float32, valuesCount int32, valuesOffset int32, overlayText string, scaleMin float32, scaleMax float32, graphSize Vec2, stride int32) {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cvalues, _ := (*C.float)(unsafe.Pointer(values)), cgoAllocsUnknown
	cvaluesCount, _ := (C.int)(valuesCount), cgoAllocsUnknown
	cvaluesOffset, _ := (C.int)(valuesOffset), cgoAllocsUnknown
	overlayText = safeString(overlayText)
	coverlayText, _ := unpackPCharString(overlayText)
	cscaleMin, _ := (C.float)(scaleMin), cgoAllocsUnknown
	cscaleMax, _ := (C.float)(scaleMax), cgoAllocsUnknown
	cgraphSize, _ := graphSize.PassValue()
	cstride, _ := (C.int)(stride), cgoAllocsUnknown
	C.igPlotHistogram(clabel, cvalues, cvaluesCount, cvaluesOffset, coverlayText, cscaleMin, cscaleMax, cgraphSize, cstride)
	runtime.KeepAlive(overlayText)
	runtime.KeepAlive(label)
}

// ProgressBar function as declared in cimgui/cimgui.h:717
func ProgressBar(fraction float32, sizeArg *Vec2, overlay string) {
	cfraction, _ := (C.float)(fraction), cgoAllocsUnknown
	csizeArg, _ := sizeArg.PassRef()
	overlay = safeString(overlay)
	coverlay, _ := unpackPCharString(overlay)
	C.igProgressBar(cfraction, csizeArg, coverlay)
	runtime.KeepAlive(overlay)
}

// BeginCombo function as declared in cimgui/cimgui.h:719
func BeginCombo(label string, previewValue string, flags int32) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	previewValue = safeString(previewValue)
	cpreviewValue, _ := unpackPCharString(previewValue)
	cflags, _ := (C.ImGuiComboFlags)(flags), cgoAllocsUnknown
	__ret := C.igBeginCombo(clabel, cpreviewValue, cflags)
	runtime.KeepAlive(previewValue)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// EndCombo function as declared in cimgui/cimgui.h:720
func EndCombo() {
	C.igEndCombo()
}

// Combo function as declared in cimgui/cimgui.h:721
func Combo(label string, currentItem *int32, items []string, itemsCount int32, popupMaxHeightInItems int32) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	ccurrentItem, _ := (*C.int)(unsafe.Pointer(currentItem)), cgoAllocsUnknown
	citems, _ := unpackArgSString(items)
	citemsCount, _ := (C.int)(itemsCount), cgoAllocsUnknown
	cpopupMaxHeightInItems, _ := (C.int)(popupMaxHeightInItems), cgoAllocsUnknown
	__ret := C.igCombo(clabel, ccurrentItem, citems, citemsCount, cpopupMaxHeightInItems)
	packSString(items, citems)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// Combo2 function as declared in cimgui/cimgui.h:722
func Combo2(label string, currentItem *int32, itemsSeparatedByZeros string, popupMaxHeightInItems int32) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	ccurrentItem, _ := (*C.int)(unsafe.Pointer(currentItem)), cgoAllocsUnknown
	itemsSeparatedByZeros = safeString(itemsSeparatedByZeros)
	citemsSeparatedByZeros, _ := unpackPCharString(itemsSeparatedByZeros)
	cpopupMaxHeightInItems, _ := (C.int)(popupMaxHeightInItems), cgoAllocsUnknown
	__ret := C.igCombo2(clabel, ccurrentItem, citemsSeparatedByZeros, cpopupMaxHeightInItems)
	runtime.KeepAlive(itemsSeparatedByZeros)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// DragFloat function as declared in cimgui/cimgui.h:727
func DragFloat(label string, v *float32, vSpeed float32, vMin float32, vMax float32, displayFormat string, power float32) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cv, _ := (*C.float)(unsafe.Pointer(v)), cgoAllocsUnknown
	cvSpeed, _ := (C.float)(vSpeed), cgoAllocsUnknown
	cvMin, _ := (C.float)(vMin), cgoAllocsUnknown
	cvMax, _ := (C.float)(vMax), cgoAllocsUnknown
	displayFormat = safeString(displayFormat)
	cdisplayFormat, _ := unpackPCharString(displayFormat)
	cpower, _ := (C.float)(power), cgoAllocsUnknown
	__ret := C.igDragFloat(clabel, cv, cvSpeed, cvMin, cvMax, cdisplayFormat, cpower)
	runtime.KeepAlive(displayFormat)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// DragFloat2 function as declared in cimgui/cimgui.h:728
func DragFloat2(label string, v *[2]float32, vSpeed float32, vMin float32, vMax float32, displayFormat string, power float32) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cv, _ := *(**C.float)(unsafe.Pointer(&v)), cgoAllocsUnknown
	cvSpeed, _ := (C.float)(vSpeed), cgoAllocsUnknown
	cvMin, _ := (C.float)(vMin), cgoAllocsUnknown
	cvMax, _ := (C.float)(vMax), cgoAllocsUnknown
	displayFormat = safeString(displayFormat)
	cdisplayFormat, _ := unpackPCharString(displayFormat)
	cpower, _ := (C.float)(power), cgoAllocsUnknown
	__ret := C.igDragFloat2(clabel, cv, cvSpeed, cvMin, cvMax, cdisplayFormat, cpower)
	runtime.KeepAlive(displayFormat)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// DragFloat3 function as declared in cimgui/cimgui.h:729
func DragFloat3(label string, v *[3]float32, vSpeed float32, vMin float32, vMax float32, displayFormat string, power float32) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cv, _ := *(**C.float)(unsafe.Pointer(&v)), cgoAllocsUnknown
	cvSpeed, _ := (C.float)(vSpeed), cgoAllocsUnknown
	cvMin, _ := (C.float)(vMin), cgoAllocsUnknown
	cvMax, _ := (C.float)(vMax), cgoAllocsUnknown
	displayFormat = safeString(displayFormat)
	cdisplayFormat, _ := unpackPCharString(displayFormat)
	cpower, _ := (C.float)(power), cgoAllocsUnknown
	__ret := C.igDragFloat3(clabel, cv, cvSpeed, cvMin, cvMax, cdisplayFormat, cpower)
	runtime.KeepAlive(displayFormat)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// DragFloat4 function as declared in cimgui/cimgui.h:730
func DragFloat4(label string, v *[4]float32, vSpeed float32, vMin float32, vMax float32, displayFormat string, power float32) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cv, _ := *(**C.float)(unsafe.Pointer(&v)), cgoAllocsUnknown
	cvSpeed, _ := (C.float)(vSpeed), cgoAllocsUnknown
	cvMin, _ := (C.float)(vMin), cgoAllocsUnknown
	cvMax, _ := (C.float)(vMax), cgoAllocsUnknown
	displayFormat = safeString(displayFormat)
	cdisplayFormat, _ := unpackPCharString(displayFormat)
	cpower, _ := (C.float)(power), cgoAllocsUnknown
	__ret := C.igDragFloat4(clabel, cv, cvSpeed, cvMin, cvMax, cdisplayFormat, cpower)
	runtime.KeepAlive(displayFormat)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// DragFloatRange2 function as declared in cimgui/cimgui.h:731
func DragFloatRange2(label string, vCurrentMin *float32, vCurrentMax *float32, vSpeed float32, vMin float32, vMax float32, displayFormat string, displayFormatMax string, power float32) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cvCurrentMin, _ := (*C.float)(unsafe.Pointer(vCurrentMin)), cgoAllocsUnknown
	cvCurrentMax, _ := (*C.float)(unsafe.Pointer(vCurrentMax)), cgoAllocsUnknown
	cvSpeed, _ := (C.float)(vSpeed), cgoAllocsUnknown
	cvMin, _ := (C.float)(vMin), cgoAllocsUnknown
	cvMax, _ := (C.float)(vMax), cgoAllocsUnknown
	displayFormat = safeString(displayFormat)
	cdisplayFormat, _ := unpackPCharString(displayFormat)
	displayFormatMax = safeString(displayFormatMax)
	cdisplayFormatMax, _ := unpackPCharString(displayFormatMax)
	cpower, _ := (C.float)(power), cgoAllocsUnknown
	__ret := C.igDragFloatRange2(clabel, cvCurrentMin, cvCurrentMax, cvSpeed, cvMin, cvMax, cdisplayFormat, cdisplayFormatMax, cpower)
	runtime.KeepAlive(displayFormatMax)
	runtime.KeepAlive(displayFormat)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// DragInt function as declared in cimgui/cimgui.h:732
func DragInt(label string, v *int32, vSpeed float32, vMin int32, vMax int32, displayFormat string) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cv, _ := (*C.int)(unsafe.Pointer(v)), cgoAllocsUnknown
	cvSpeed, _ := (C.float)(vSpeed), cgoAllocsUnknown
	cvMin, _ := (C.int)(vMin), cgoAllocsUnknown
	cvMax, _ := (C.int)(vMax), cgoAllocsUnknown
	displayFormat = safeString(displayFormat)
	cdisplayFormat, _ := unpackPCharString(displayFormat)
	__ret := C.igDragInt(clabel, cv, cvSpeed, cvMin, cvMax, cdisplayFormat)
	runtime.KeepAlive(displayFormat)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// DragInt2 function as declared in cimgui/cimgui.h:733
func DragInt2(label string, v *[2]int32, vSpeed float32, vMin int32, vMax int32, displayFormat string) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cv, _ := *(**C.int)(unsafe.Pointer(&v)), cgoAllocsUnknown
	cvSpeed, _ := (C.float)(vSpeed), cgoAllocsUnknown
	cvMin, _ := (C.int)(vMin), cgoAllocsUnknown
	cvMax, _ := (C.int)(vMax), cgoAllocsUnknown
	displayFormat = safeString(displayFormat)
	cdisplayFormat, _ := unpackPCharString(displayFormat)
	__ret := C.igDragInt2(clabel, cv, cvSpeed, cvMin, cvMax, cdisplayFormat)
	runtime.KeepAlive(displayFormat)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// DragInt3 function as declared in cimgui/cimgui.h:734
func DragInt3(label string, v *[3]int32, vSpeed float32, vMin int32, vMax int32, displayFormat string) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cv, _ := *(**C.int)(unsafe.Pointer(&v)), cgoAllocsUnknown
	cvSpeed, _ := (C.float)(vSpeed), cgoAllocsUnknown
	cvMin, _ := (C.int)(vMin), cgoAllocsUnknown
	cvMax, _ := (C.int)(vMax), cgoAllocsUnknown
	displayFormat = safeString(displayFormat)
	cdisplayFormat, _ := unpackPCharString(displayFormat)
	__ret := C.igDragInt3(clabel, cv, cvSpeed, cvMin, cvMax, cdisplayFormat)
	runtime.KeepAlive(displayFormat)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// DragInt4 function as declared in cimgui/cimgui.h:735
func DragInt4(label string, v *[4]int32, vSpeed float32, vMin int32, vMax int32, displayFormat string) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cv, _ := *(**C.int)(unsafe.Pointer(&v)), cgoAllocsUnknown
	cvSpeed, _ := (C.float)(vSpeed), cgoAllocsUnknown
	cvMin, _ := (C.int)(vMin), cgoAllocsUnknown
	cvMax, _ := (C.int)(vMax), cgoAllocsUnknown
	displayFormat = safeString(displayFormat)
	cdisplayFormat, _ := unpackPCharString(displayFormat)
	__ret := C.igDragInt4(clabel, cv, cvSpeed, cvMin, cvMax, cdisplayFormat)
	runtime.KeepAlive(displayFormat)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// DragIntRange2 function as declared in cimgui/cimgui.h:736
func DragIntRange2(label string, vCurrentMin *int32, vCurrentMax *int32, vSpeed float32, vMin int32, vMax int32, displayFormat string, displayFormatMax string) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cvCurrentMin, _ := (*C.int)(unsafe.Pointer(vCurrentMin)), cgoAllocsUnknown
	cvCurrentMax, _ := (*C.int)(unsafe.Pointer(vCurrentMax)), cgoAllocsUnknown
	cvSpeed, _ := (C.float)(vSpeed), cgoAllocsUnknown
	cvMin, _ := (C.int)(vMin), cgoAllocsUnknown
	cvMax, _ := (C.int)(vMax), cgoAllocsUnknown
	displayFormat = safeString(displayFormat)
	cdisplayFormat, _ := unpackPCharString(displayFormat)
	displayFormatMax = safeString(displayFormatMax)
	cdisplayFormatMax, _ := unpackPCharString(displayFormatMax)
	__ret := C.igDragIntRange2(clabel, cvCurrentMin, cvCurrentMax, cvSpeed, cvMin, cvMax, cdisplayFormat, cdisplayFormatMax)
	runtime.KeepAlive(displayFormatMax)
	runtime.KeepAlive(displayFormat)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// InputText function as declared in cimgui/cimgui.h:739
func InputText(label string, buf *byte, bufSize uint, flags InputTextFlags, callback TextEditCallback, userData unsafe.Pointer) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cbuf, _ := (*C.char)(unsafe.Pointer(buf)), cgoAllocsUnknown
	cbufSize, _ := (C.size_t)(bufSize), cgoAllocsUnknown
	cflags, _ := (C.ImGuiInputTextFlags)(flags), cgoAllocsUnknown
	ccallback, _ := callback.PassValue()
	cuserData, _ := userData, cgoAllocsUnknown
	__ret := C.igInputText(clabel, cbuf, cbufSize, cflags, ccallback, cuserData)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// InputTextMultiline function as declared in cimgui/cimgui.h:740
func InputTextMultiline(label string, buf *byte, bufSize uint, size Vec2, flags InputTextFlags, callback TextEditCallback, userData unsafe.Pointer) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cbuf, _ := (*C.char)(unsafe.Pointer(buf)), cgoAllocsUnknown
	cbufSize, _ := (C.size_t)(bufSize), cgoAllocsUnknown
	csize, _ := size.PassValue()
	cflags, _ := (C.ImGuiInputTextFlags)(flags), cgoAllocsUnknown
	ccallback, _ := callback.PassValue()
	cuserData, _ := userData, cgoAllocsUnknown
	__ret := C.igInputTextMultiline(clabel, cbuf, cbufSize, csize, cflags, ccallback, cuserData)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// InputFloat function as declared in cimgui/cimgui.h:741
func InputFloat(label string, v *float32, step float32, stepFast float32, decimalPrecision int32, extraFlags InputTextFlags) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cv, _ := (*C.float)(unsafe.Pointer(v)), cgoAllocsUnknown
	cstep, _ := (C.float)(step), cgoAllocsUnknown
	cstepFast, _ := (C.float)(stepFast), cgoAllocsUnknown
	cdecimalPrecision, _ := (C.int)(decimalPrecision), cgoAllocsUnknown
	cextraFlags, _ := (C.ImGuiInputTextFlags)(extraFlags), cgoAllocsUnknown
	__ret := C.igInputFloat(clabel, cv, cstep, cstepFast, cdecimalPrecision, cextraFlags)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// InputFloat2 function as declared in cimgui/cimgui.h:742
func InputFloat2(label string, v *[2]float32, decimalPrecision int32, extraFlags InputTextFlags) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cv, _ := *(**C.float)(unsafe.Pointer(&v)), cgoAllocsUnknown
	cdecimalPrecision, _ := (C.int)(decimalPrecision), cgoAllocsUnknown
	cextraFlags, _ := (C.ImGuiInputTextFlags)(extraFlags), cgoAllocsUnknown
	__ret := C.igInputFloat2(clabel, cv, cdecimalPrecision, cextraFlags)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// InputFloat3 function as declared in cimgui/cimgui.h:743
func InputFloat3(label string, v *[3]float32, decimalPrecision int32, extraFlags InputTextFlags) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cv, _ := *(**C.float)(unsafe.Pointer(&v)), cgoAllocsUnknown
	cdecimalPrecision, _ := (C.int)(decimalPrecision), cgoAllocsUnknown
	cextraFlags, _ := (C.ImGuiInputTextFlags)(extraFlags), cgoAllocsUnknown
	__ret := C.igInputFloat3(clabel, cv, cdecimalPrecision, cextraFlags)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// InputFloat4 function as declared in cimgui/cimgui.h:744
func InputFloat4(label string, v *[4]float32, decimalPrecision int32, extraFlags InputTextFlags) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cv, _ := *(**C.float)(unsafe.Pointer(&v)), cgoAllocsUnknown
	cdecimalPrecision, _ := (C.int)(decimalPrecision), cgoAllocsUnknown
	cextraFlags, _ := (C.ImGuiInputTextFlags)(extraFlags), cgoAllocsUnknown
	__ret := C.igInputFloat4(clabel, cv, cdecimalPrecision, cextraFlags)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// InputInt function as declared in cimgui/cimgui.h:745
func InputInt(label string, v *int32, step int32, stepFast int32, extraFlags InputTextFlags) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cv, _ := (*C.int)(unsafe.Pointer(v)), cgoAllocsUnknown
	cstep, _ := (C.int)(step), cgoAllocsUnknown
	cstepFast, _ := (C.int)(stepFast), cgoAllocsUnknown
	cextraFlags, _ := (C.ImGuiInputTextFlags)(extraFlags), cgoAllocsUnknown
	__ret := C.igInputInt(clabel, cv, cstep, cstepFast, cextraFlags)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// InputInt2 function as declared in cimgui/cimgui.h:746
func InputInt2(label string, v *[2]int32, extraFlags InputTextFlags) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cv, _ := *(**C.int)(unsafe.Pointer(&v)), cgoAllocsUnknown
	cextraFlags, _ := (C.ImGuiInputTextFlags)(extraFlags), cgoAllocsUnknown
	__ret := C.igInputInt2(clabel, cv, cextraFlags)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// InputInt3 function as declared in cimgui/cimgui.h:747
func InputInt3(label string, v *[3]int32, extraFlags InputTextFlags) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cv, _ := *(**C.int)(unsafe.Pointer(&v)), cgoAllocsUnknown
	cextraFlags, _ := (C.ImGuiInputTextFlags)(extraFlags), cgoAllocsUnknown
	__ret := C.igInputInt3(clabel, cv, cextraFlags)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// InputInt4 function as declared in cimgui/cimgui.h:748
func InputInt4(label string, v *[4]int32, extraFlags InputTextFlags) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cv, _ := *(**C.int)(unsafe.Pointer(&v)), cgoAllocsUnknown
	cextraFlags, _ := (C.ImGuiInputTextFlags)(extraFlags), cgoAllocsUnknown
	__ret := C.igInputInt4(clabel, cv, cextraFlags)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// SliderFloat function as declared in cimgui/cimgui.h:751
func SliderFloat(label string, v *float32, vMin float32, vMax float32, displayFormat string, power float32) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cv, _ := (*C.float)(unsafe.Pointer(v)), cgoAllocsUnknown
	cvMin, _ := (C.float)(vMin), cgoAllocsUnknown
	cvMax, _ := (C.float)(vMax), cgoAllocsUnknown
	displayFormat = safeString(displayFormat)
	cdisplayFormat, _ := unpackPCharString(displayFormat)
	cpower, _ := (C.float)(power), cgoAllocsUnknown
	__ret := C.igSliderFloat(clabel, cv, cvMin, cvMax, cdisplayFormat, cpower)
	runtime.KeepAlive(displayFormat)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// SliderFloat2 function as declared in cimgui/cimgui.h:752
func SliderFloat2(label string, v *[2]float32, vMin float32, vMax float32, displayFormat string, power float32) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cv, _ := *(**C.float)(unsafe.Pointer(&v)), cgoAllocsUnknown
	cvMin, _ := (C.float)(vMin), cgoAllocsUnknown
	cvMax, _ := (C.float)(vMax), cgoAllocsUnknown
	displayFormat = safeString(displayFormat)
	cdisplayFormat, _ := unpackPCharString(displayFormat)
	cpower, _ := (C.float)(power), cgoAllocsUnknown
	__ret := C.igSliderFloat2(clabel, cv, cvMin, cvMax, cdisplayFormat, cpower)
	runtime.KeepAlive(displayFormat)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// SliderFloat3 function as declared in cimgui/cimgui.h:753
func SliderFloat3(label string, v *[3]float32, vMin float32, vMax float32, displayFormat string, power float32) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cv, _ := *(**C.float)(unsafe.Pointer(&v)), cgoAllocsUnknown
	cvMin, _ := (C.float)(vMin), cgoAllocsUnknown
	cvMax, _ := (C.float)(vMax), cgoAllocsUnknown
	displayFormat = safeString(displayFormat)
	cdisplayFormat, _ := unpackPCharString(displayFormat)
	cpower, _ := (C.float)(power), cgoAllocsUnknown
	__ret := C.igSliderFloat3(clabel, cv, cvMin, cvMax, cdisplayFormat, cpower)
	runtime.KeepAlive(displayFormat)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// SliderFloat4 function as declared in cimgui/cimgui.h:754
func SliderFloat4(label string, v *[4]float32, vMin float32, vMax float32, displayFormat string, power float32) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cv, _ := *(**C.float)(unsafe.Pointer(&v)), cgoAllocsUnknown
	cvMin, _ := (C.float)(vMin), cgoAllocsUnknown
	cvMax, _ := (C.float)(vMax), cgoAllocsUnknown
	displayFormat = safeString(displayFormat)
	cdisplayFormat, _ := unpackPCharString(displayFormat)
	cpower, _ := (C.float)(power), cgoAllocsUnknown
	__ret := C.igSliderFloat4(clabel, cv, cvMin, cvMax, cdisplayFormat, cpower)
	runtime.KeepAlive(displayFormat)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// SliderAngle function as declared in cimgui/cimgui.h:755
func SliderAngle(label string, vRad *float32, vDegreesMin float32, vDegreesMax float32) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cvRad, _ := (*C.float)(unsafe.Pointer(vRad)), cgoAllocsUnknown
	cvDegreesMin, _ := (C.float)(vDegreesMin), cgoAllocsUnknown
	cvDegreesMax, _ := (C.float)(vDegreesMax), cgoAllocsUnknown
	__ret := C.igSliderAngle(clabel, cvRad, cvDegreesMin, cvDegreesMax)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// SliderInt function as declared in cimgui/cimgui.h:756
func SliderInt(label string, v *int32, vMin int32, vMax int32, displayFormat string) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cv, _ := (*C.int)(unsafe.Pointer(v)), cgoAllocsUnknown
	cvMin, _ := (C.int)(vMin), cgoAllocsUnknown
	cvMax, _ := (C.int)(vMax), cgoAllocsUnknown
	displayFormat = safeString(displayFormat)
	cdisplayFormat, _ := unpackPCharString(displayFormat)
	__ret := C.igSliderInt(clabel, cv, cvMin, cvMax, cdisplayFormat)
	runtime.KeepAlive(displayFormat)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// SliderInt2 function as declared in cimgui/cimgui.h:757
func SliderInt2(label string, v *[2]int32, vMin int32, vMax int32, displayFormat string) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cv, _ := *(**C.int)(unsafe.Pointer(&v)), cgoAllocsUnknown
	cvMin, _ := (C.int)(vMin), cgoAllocsUnknown
	cvMax, _ := (C.int)(vMax), cgoAllocsUnknown
	displayFormat = safeString(displayFormat)
	cdisplayFormat, _ := unpackPCharString(displayFormat)
	__ret := C.igSliderInt2(clabel, cv, cvMin, cvMax, cdisplayFormat)
	runtime.KeepAlive(displayFormat)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// SliderInt3 function as declared in cimgui/cimgui.h:758
func SliderInt3(label string, v *[3]int32, vMin int32, vMax int32, displayFormat string) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cv, _ := *(**C.int)(unsafe.Pointer(&v)), cgoAllocsUnknown
	cvMin, _ := (C.int)(vMin), cgoAllocsUnknown
	cvMax, _ := (C.int)(vMax), cgoAllocsUnknown
	displayFormat = safeString(displayFormat)
	cdisplayFormat, _ := unpackPCharString(displayFormat)
	__ret := C.igSliderInt3(clabel, cv, cvMin, cvMax, cdisplayFormat)
	runtime.KeepAlive(displayFormat)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// SliderInt4 function as declared in cimgui/cimgui.h:759
func SliderInt4(label string, v *[4]int32, vMin int32, vMax int32, displayFormat string) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cv, _ := *(**C.int)(unsafe.Pointer(&v)), cgoAllocsUnknown
	cvMin, _ := (C.int)(vMin), cgoAllocsUnknown
	cvMax, _ := (C.int)(vMax), cgoAllocsUnknown
	displayFormat = safeString(displayFormat)
	cdisplayFormat, _ := unpackPCharString(displayFormat)
	__ret := C.igSliderInt4(clabel, cv, cvMin, cvMax, cdisplayFormat)
	runtime.KeepAlive(displayFormat)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// VSliderFloat function as declared in cimgui/cimgui.h:760
func VSliderFloat(label string, size Vec2, v *float32, vMin float32, vMax float32, displayFormat string, power float32) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	csize, _ := size.PassValue()
	cv, _ := (*C.float)(unsafe.Pointer(v)), cgoAllocsUnknown
	cvMin, _ := (C.float)(vMin), cgoAllocsUnknown
	cvMax, _ := (C.float)(vMax), cgoAllocsUnknown
	displayFormat = safeString(displayFormat)
	cdisplayFormat, _ := unpackPCharString(displayFormat)
	cpower, _ := (C.float)(power), cgoAllocsUnknown
	__ret := C.igVSliderFloat(clabel, csize, cv, cvMin, cvMax, cdisplayFormat, cpower)
	runtime.KeepAlive(displayFormat)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// VSliderInt function as declared in cimgui/cimgui.h:761
func VSliderInt(label string, size Vec2, v *int32, vMin int32, vMax int32, displayFormat string) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	csize, _ := size.PassValue()
	cv, _ := (*C.int)(unsafe.Pointer(v)), cgoAllocsUnknown
	cvMin, _ := (C.int)(vMin), cgoAllocsUnknown
	cvMax, _ := (C.int)(vMax), cgoAllocsUnknown
	displayFormat = safeString(displayFormat)
	cdisplayFormat, _ := unpackPCharString(displayFormat)
	__ret := C.igVSliderInt(clabel, csize, cv, cvMin, cvMax, cdisplayFormat)
	runtime.KeepAlive(displayFormat)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// ColorEdit3 function as declared in cimgui/cimgui.h:765
func ColorEdit3(label string, col *[3]float32, flags ColorEditFlags) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	ccol, _ := *(**C.float)(unsafe.Pointer(&col)), cgoAllocsUnknown
	cflags, _ := (C.ImGuiColorEditFlags)(flags), cgoAllocsUnknown
	__ret := C.igColorEdit3(clabel, ccol, cflags)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// ColorEdit4 function as declared in cimgui/cimgui.h:766
func ColorEdit4(label string, col *[4]float32, flags ColorEditFlags) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	ccol, _ := *(**C.float)(unsafe.Pointer(&col)), cgoAllocsUnknown
	cflags, _ := (C.ImGuiColorEditFlags)(flags), cgoAllocsUnknown
	__ret := C.igColorEdit4(clabel, ccol, cflags)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// ColorPicker3 function as declared in cimgui/cimgui.h:767
func ColorPicker3(label string, col *[3]float32, flags ColorEditFlags) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	ccol, _ := *(**C.float)(unsafe.Pointer(&col)), cgoAllocsUnknown
	cflags, _ := (C.ImGuiColorEditFlags)(flags), cgoAllocsUnknown
	__ret := C.igColorPicker3(clabel, ccol, cflags)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// ColorPicker4 function as declared in cimgui/cimgui.h:768
func ColorPicker4(label string, col *[4]float32, flags ColorEditFlags, refCol []float32) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	ccol, _ := *(**C.float)(unsafe.Pointer(&col)), cgoAllocsUnknown
	cflags, _ := (C.ImGuiColorEditFlags)(flags), cgoAllocsUnknown
	crefCol, _ := (*C.float)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&refCol)).Data)), cgoAllocsUnknown
	__ret := C.igColorPicker4(clabel, ccol, cflags, crefCol)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// ColorButton function as declared in cimgui/cimgui.h:769
func ColorButton(descId string, col Vec4, flags ColorEditFlags, size Vec2) bool {
	descId = safeString(descId)
	cdescId, _ := unpackPCharString(descId)
	ccol, _ := col.PassValue()
	cflags, _ := (C.ImGuiColorEditFlags)(flags), cgoAllocsUnknown
	csize, _ := size.PassValue()
	__ret := C.igColorButton(cdescId, ccol, cflags, csize)
	runtime.KeepAlive(descId)
	__v := (bool)(__ret)
	return __v
}

// SetColorEditOptions function as declared in cimgui/cimgui.h:770
func SetColorEditOptions(flags ColorEditFlags) {
	cflags, _ := (C.ImGuiColorEditFlags)(flags), cgoAllocsUnknown
	C.igSetColorEditOptions(cflags)
}

// TreeNode function as declared in cimgui/cimgui.h:773
func TreeNode(label string) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	__ret := C.igTreeNode(clabel)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// TreeNodeEx function as declared in cimgui/cimgui.h:778
func TreeNodeEx(label string, flags TreeNodeFlags) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cflags, _ := (C.ImGuiTreeNodeFlags)(flags), cgoAllocsUnknown
	__ret := C.igTreeNodeEx(clabel, cflags)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// TreeNodeExV function as declared in cimgui/cimgui.h:781
func TreeNodeExV(strId string, flags TreeNodeFlags, fmt string, args unsafe.Pointer) bool {
	strId = safeString(strId)
	cstrId, _ := unpackPCharString(strId)
	cflags, _ := (C.ImGuiTreeNodeFlags)(flags), cgoAllocsUnknown
	fmt = safeString(fmt)
	cfmt, _ := unpackPCharString(fmt)
	cargs, _ := args, cgoAllocsUnknown
	__ret := C.igTreeNodeExV(cstrId, cflags, cfmt, cargs)
	runtime.KeepAlive(fmt)
	runtime.KeepAlive(strId)
	__v := (bool)(__ret)
	return __v
}

// TreeNodeExVPtr function as declared in cimgui/cimgui.h:782
func TreeNodeExVPtr(ptrId unsafe.Pointer, flags TreeNodeFlags, fmt string, args unsafe.Pointer) bool {
	cptrId, _ := ptrId, cgoAllocsUnknown
	cflags, _ := (C.ImGuiTreeNodeFlags)(flags), cgoAllocsUnknown
	fmt = safeString(fmt)
	cfmt, _ := unpackPCharString(fmt)
	cargs, _ := args, cgoAllocsUnknown
	__ret := C.igTreeNodeExVPtr(cptrId, cflags, cfmt, cargs)
	runtime.KeepAlive(fmt)
	__v := (bool)(__ret)
	return __v
}

// TreePushStr function as declared in cimgui/cimgui.h:783
func TreePushStr(strId string) {
	strId = safeString(strId)
	cstrId, _ := unpackPCharString(strId)
	C.igTreePushStr(cstrId)
	runtime.KeepAlive(strId)
}

// TreePushPtr function as declared in cimgui/cimgui.h:784
func TreePushPtr(ptrId unsafe.Pointer) {
	cptrId, _ := ptrId, cgoAllocsUnknown
	C.igTreePushPtr(cptrId)
}

// TreePop function as declared in cimgui/cimgui.h:785
func TreePop() {
	C.igTreePop()
}

// TreeAdvanceToLabelPos function as declared in cimgui/cimgui.h:786
func TreeAdvanceToLabelPos() {
	C.igTreeAdvanceToLabelPos()
}

// GetTreeNodeToLabelSpacing function as declared in cimgui/cimgui.h:787
func GetTreeNodeToLabelSpacing() float32 {
	__ret := C.igGetTreeNodeToLabelSpacing()
	__v := (float32)(__ret)
	return __v
}

// SetNextTreeNodeOpen function as declared in cimgui/cimgui.h:788
func SetNextTreeNodeOpen(opened bool, cond Cond) {
	copened, _ := (C._Bool)(opened), cgoAllocsUnknown
	ccond, _ := (C.ImGuiCond)(cond), cgoAllocsUnknown
	C.igSetNextTreeNodeOpen(copened, ccond)
}

// CollapsingHeader function as declared in cimgui/cimgui.h:789
func CollapsingHeader(label string, flags TreeNodeFlags) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cflags, _ := (C.ImGuiTreeNodeFlags)(flags), cgoAllocsUnknown
	__ret := C.igCollapsingHeader(clabel, cflags)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// CollapsingHeaderEx function as declared in cimgui/cimgui.h:790
func CollapsingHeaderEx(label string, pOpen []bool, flags TreeNodeFlags) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cpOpen, _ := (*C._Bool)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&pOpen)).Data)), cgoAllocsUnknown
	cflags, _ := (C.ImGuiTreeNodeFlags)(flags), cgoAllocsUnknown
	__ret := C.igCollapsingHeaderEx(clabel, cpOpen, cflags)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// Selectable function as declared in cimgui/cimgui.h:793
func Selectable(label string, selected bool, flags SelectableFlags, size Vec2) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cselected, _ := (C._Bool)(selected), cgoAllocsUnknown
	cflags, _ := (C.ImGuiSelectableFlags)(flags), cgoAllocsUnknown
	csize, _ := size.PassValue()
	__ret := C.igSelectable(clabel, cselected, cflags, csize)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// SelectableEx function as declared in cimgui/cimgui.h:794
func SelectableEx(label string, pSelected []bool, flags SelectableFlags, size Vec2) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cpSelected, _ := (*C._Bool)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&pSelected)).Data)), cgoAllocsUnknown
	cflags, _ := (C.ImGuiSelectableFlags)(flags), cgoAllocsUnknown
	csize, _ := size.PassValue()
	__ret := C.igSelectableEx(clabel, cpSelected, cflags, csize)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// ListBox function as declared in cimgui/cimgui.h:795
func ListBox(label string, currentItem []int32, items []string, itemsCount int32, heightInItems int32) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	ccurrentItem, _ := (*C.int)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&currentItem)).Data)), cgoAllocsUnknown
	citems, _ := unpackArgSString(items)
	citemsCount, _ := (C.int)(itemsCount), cgoAllocsUnknown
	cheightInItems, _ := (C.int)(heightInItems), cgoAllocsUnknown
	__ret := C.igListBox(clabel, ccurrentItem, citems, citemsCount, cheightInItems)
	packSString(items, citems)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// ListBoxHeader function as declared in cimgui/cimgui.h:797
func ListBoxHeader(label string, size Vec2) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	csize, _ := size.PassValue()
	__ret := C.igListBoxHeader(clabel, csize)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// ListBoxHeader2 function as declared in cimgui/cimgui.h:798
func ListBoxHeader2(label string, itemsCount int32, heightInItems int32) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	citemsCount, _ := (C.int)(itemsCount), cgoAllocsUnknown
	cheightInItems, _ := (C.int)(heightInItems), cgoAllocsUnknown
	__ret := C.igListBoxHeader2(clabel, citemsCount, cheightInItems)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// ListBoxFooter function as declared in cimgui/cimgui.h:799
func ListBoxFooter() {
	C.igListBoxFooter()
}

// ValueBool function as declared in cimgui/cimgui.h:802
func ValueBool(prefix string, b bool) {
	prefix = safeString(prefix)
	cprefix, _ := unpackPCharString(prefix)
	cb, _ := (C._Bool)(b), cgoAllocsUnknown
	C.igValueBool(cprefix, cb)
	runtime.KeepAlive(prefix)
}

// ValueInt function as declared in cimgui/cimgui.h:803
func ValueInt(prefix string, v int32) {
	prefix = safeString(prefix)
	cprefix, _ := unpackPCharString(prefix)
	cv, _ := (C.int)(v), cgoAllocsUnknown
	C.igValueInt(cprefix, cv)
	runtime.KeepAlive(prefix)
}

// ValueUInt function as declared in cimgui/cimgui.h:804
func ValueUInt(prefix string, v uint32) {
	prefix = safeString(prefix)
	cprefix, _ := unpackPCharString(prefix)
	cv, _ := (C.uint)(v), cgoAllocsUnknown
	C.igValueUInt(cprefix, cv)
	runtime.KeepAlive(prefix)
}

// ValueFloat function as declared in cimgui/cimgui.h:805
func ValueFloat(prefix string, v float32, floatFormat string) {
	prefix = safeString(prefix)
	cprefix, _ := unpackPCharString(prefix)
	cv, _ := (C.float)(v), cgoAllocsUnknown
	floatFormat = safeString(floatFormat)
	cfloatFormat, _ := unpackPCharString(floatFormat)
	C.igValueFloat(cprefix, cv, cfloatFormat)
	runtime.KeepAlive(floatFormat)
	runtime.KeepAlive(prefix)
}

// BeginTooltip function as declared in cimgui/cimgui.h:810
func BeginTooltip() {
	C.igBeginTooltip()
}

// EndTooltip function as declared in cimgui/cimgui.h:811
func EndTooltip() {
	C.igEndTooltip()
}

// BeginMainMenuBar function as declared in cimgui/cimgui.h:814
func BeginMainMenuBar() bool {
	__ret := C.igBeginMainMenuBar()
	__v := (bool)(__ret)
	return __v
}

// EndMainMenuBar function as declared in cimgui/cimgui.h:815
func EndMainMenuBar() {
	C.igEndMainMenuBar()
}

// BeginMenuBar function as declared in cimgui/cimgui.h:816
func BeginMenuBar() bool {
	__ret := C.igBeginMenuBar()
	__v := (bool)(__ret)
	return __v
}

// EndMenuBar function as declared in cimgui/cimgui.h:817
func EndMenuBar() {
	C.igEndMenuBar()
}

// BeginMenu function as declared in cimgui/cimgui.h:818
func BeginMenu(label string, enabled bool) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cenabled, _ := (C._Bool)(enabled), cgoAllocsUnknown
	__ret := C.igBeginMenu(clabel, cenabled)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// EndMenu function as declared in cimgui/cimgui.h:819
func EndMenu() {
	C.igEndMenu()
}

// MenuItem function as declared in cimgui/cimgui.h:820
func MenuItem(label string, shortcut string, selected bool, enabled bool) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	shortcut = safeString(shortcut)
	cshortcut, _ := unpackPCharString(shortcut)
	cselected, _ := (C._Bool)(selected), cgoAllocsUnknown
	cenabled, _ := (C._Bool)(enabled), cgoAllocsUnknown
	__ret := C.igMenuItem(clabel, cshortcut, cselected, cenabled)
	runtime.KeepAlive(shortcut)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// MenuItemPtr function as declared in cimgui/cimgui.h:821
func MenuItemPtr(label string, shortcut string, pSelected []bool, enabled bool) bool {
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	shortcut = safeString(shortcut)
	cshortcut, _ := unpackPCharString(shortcut)
	cpSelected, _ := (*C._Bool)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&pSelected)).Data)), cgoAllocsUnknown
	cenabled, _ := (C._Bool)(enabled), cgoAllocsUnknown
	__ret := C.igMenuItemPtr(clabel, cshortcut, cpSelected, cenabled)
	runtime.KeepAlive(shortcut)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// OpenPopup function as declared in cimgui/cimgui.h:824
func OpenPopup(strId string) {
	strId = safeString(strId)
	cstrId, _ := unpackPCharString(strId)
	C.igOpenPopup(cstrId)
	runtime.KeepAlive(strId)
}

// OpenPopupOnItemClick function as declared in cimgui/cimgui.h:825
func OpenPopupOnItemClick(strId string, mouseButton int32) bool {
	strId = safeString(strId)
	cstrId, _ := unpackPCharString(strId)
	cmouseButton, _ := (C.int)(mouseButton), cgoAllocsUnknown
	__ret := C.igOpenPopupOnItemClick(cstrId, cmouseButton)
	runtime.KeepAlive(strId)
	__v := (bool)(__ret)
	return __v
}

// BeginPopup function as declared in cimgui/cimgui.h:826
func BeginPopup(strId string) bool {
	strId = safeString(strId)
	cstrId, _ := unpackPCharString(strId)
	__ret := C.igBeginPopup(cstrId)
	runtime.KeepAlive(strId)
	__v := (bool)(__ret)
	return __v
}

// BeginPopupModal function as declared in cimgui/cimgui.h:827
func BeginPopupModal(name string, pOpen *bool, extraFlags WindowFlags) bool {
	name = safeString(name)
	cname, _ := unpackPCharString(name)
	cpOpen, _ := (*C._Bool)(unsafe.Pointer(pOpen)), cgoAllocsUnknown
	cextraFlags, _ := (C.ImGuiWindowFlags)(extraFlags), cgoAllocsUnknown
	__ret := C.igBeginPopupModal(cname, cpOpen, cextraFlags)
	runtime.KeepAlive(name)
	__v := (bool)(__ret)
	return __v
}

// BeginPopupContextItem function as declared in cimgui/cimgui.h:828
func BeginPopupContextItem(strId string, mouseButton int32) bool {
	strId = safeString(strId)
	cstrId, _ := unpackPCharString(strId)
	cmouseButton, _ := (C.int)(mouseButton), cgoAllocsUnknown
	__ret := C.igBeginPopupContextItem(cstrId, cmouseButton)
	runtime.KeepAlive(strId)
	__v := (bool)(__ret)
	return __v
}

// BeginPopupContextWindow function as declared in cimgui/cimgui.h:829
func BeginPopupContextWindow(strId string, mouseButton int32, alsoOverItems bool) bool {
	strId = safeString(strId)
	cstrId, _ := unpackPCharString(strId)
	cmouseButton, _ := (C.int)(mouseButton), cgoAllocsUnknown
	calsoOverItems, _ := (C._Bool)(alsoOverItems), cgoAllocsUnknown
	__ret := C.igBeginPopupContextWindow(cstrId, cmouseButton, calsoOverItems)
	runtime.KeepAlive(strId)
	__v := (bool)(__ret)
	return __v
}

// BeginPopupContextVoid function as declared in cimgui/cimgui.h:830
func BeginPopupContextVoid(strId string, mouseButton int32) bool {
	strId = safeString(strId)
	cstrId, _ := unpackPCharString(strId)
	cmouseButton, _ := (C.int)(mouseButton), cgoAllocsUnknown
	__ret := C.igBeginPopupContextVoid(cstrId, cmouseButton)
	runtime.KeepAlive(strId)
	__v := (bool)(__ret)
	return __v
}

// EndPopup function as declared in cimgui/cimgui.h:831
func EndPopup() {
	C.igEndPopup()
}

// IsPopupOpen function as declared in cimgui/cimgui.h:832
func IsPopupOpen(strId string) bool {
	strId = safeString(strId)
	cstrId, _ := unpackPCharString(strId)
	__ret := C.igIsPopupOpen(cstrId)
	runtime.KeepAlive(strId)
	__v := (bool)(__ret)
	return __v
}

// CloseCurrentPopup function as declared in cimgui/cimgui.h:833
func CloseCurrentPopup() {
	C.igCloseCurrentPopup()
}

// LogToTTY function as declared in cimgui/cimgui.h:836
func LogToTTY(maxDepth int32) {
	cmaxDepth, _ := (C.int)(maxDepth), cgoAllocsUnknown
	C.igLogToTTY(cmaxDepth)
}

// LogToFile function as declared in cimgui/cimgui.h:837
func LogToFile(maxDepth int32, filename string) {
	cmaxDepth, _ := (C.int)(maxDepth), cgoAllocsUnknown
	filename = safeString(filename)
	cfilename, _ := unpackPCharString(filename)
	C.igLogToFile(cmaxDepth, cfilename)
	runtime.KeepAlive(filename)
}

// LogToClipboard function as declared in cimgui/cimgui.h:838
func LogToClipboard(maxDepth int32) {
	cmaxDepth, _ := (C.int)(maxDepth), cgoAllocsUnknown
	C.igLogToClipboard(cmaxDepth)
}

// LogFinish function as declared in cimgui/cimgui.h:839
func LogFinish() {
	C.igLogFinish()
}

// LogButtons function as declared in cimgui/cimgui.h:840
func LogButtons() {
	C.igLogButtons()
}

// BeginDragDropSource function as declared in cimgui/cimgui.h:843
func BeginDragDropSource(flags int32, mouseButton int32) bool {
	cflags, _ := (C.ImGuiDragDropFlags)(flags), cgoAllocsUnknown
	cmouseButton, _ := (C.int)(mouseButton), cgoAllocsUnknown
	__ret := C.igBeginDragDropSource(cflags, cmouseButton)
	__v := (bool)(__ret)
	return __v
}

// SetDragDropPayload function as declared in cimgui/cimgui.h:844
func SetDragDropPayload(kind string, data unsafe.Pointer, size uint, cond Cond) bool {
	kind = safeString(kind)
	ckind, _ := unpackPCharString(kind)
	cdata, _ := data, cgoAllocsUnknown
	csize, _ := (C.size_t)(size), cgoAllocsUnknown
	ccond, _ := (C.ImGuiCond)(cond), cgoAllocsUnknown
	__ret := C.igSetDragDropPayload(ckind, cdata, csize, ccond)
	runtime.KeepAlive(kind)
	__v := (bool)(__ret)
	return __v
}

// EndDragDropSource function as declared in cimgui/cimgui.h:845
func EndDragDropSource() {
	C.igEndDragDropSource()
}

// BeginDragDropTarget function as declared in cimgui/cimgui.h:846
func BeginDragDropTarget() bool {
	__ret := C.igBeginDragDropTarget()
	__v := (bool)(__ret)
	return __v
}

// AcceptDragDropPayload function as declared in cimgui/cimgui.h:847
func AcceptDragDropPayload(kind string, flags int32) *Payload {
	kind = safeString(kind)
	ckind, _ := unpackPCharString(kind)
	cflags, _ := (C.ImGuiDragDropFlags)(flags), cgoAllocsUnknown
	__ret := C.igAcceptDragDropPayload(ckind, cflags)
	runtime.KeepAlive(kind)
	__v := NewPayloadRef(unsafe.Pointer(__ret))
	return __v
}

// EndDragDropTarget function as declared in cimgui/cimgui.h:848
func EndDragDropTarget() {
	C.igEndDragDropTarget()
}

// PushClipRect function as declared in cimgui/cimgui.h:851
func PushClipRect(clipRectMin Vec2, clipRectMax Vec2, intersectWithCurrentClipRect bool) {
	cclipRectMin, _ := clipRectMin.PassValue()
	cclipRectMax, _ := clipRectMax.PassValue()
	cintersectWithCurrentClipRect, _ := (C._Bool)(intersectWithCurrentClipRect), cgoAllocsUnknown
	C.igPushClipRect(cclipRectMin, cclipRectMax, cintersectWithCurrentClipRect)
}

// PopClipRect function as declared in cimgui/cimgui.h:852
func PopClipRect() {
	C.igPopClipRect()
}

// StyleColorsClassic function as declared in cimgui/cimgui.h:855
func StyleColorsClassic(dst []Style) {
	cdst, _ := unpackArgSStyle(dst)
	C.igStyleColorsClassic(cdst)
	packSStyle(dst, cdst)
}

// StyleColorsDark function as declared in cimgui/cimgui.h:856
func StyleColorsDark(dst []Style) {
	cdst, _ := unpackArgSStyle(dst)
	C.igStyleColorsDark(cdst)
	packSStyle(dst, cdst)
}

// StyleColorsLight function as declared in cimgui/cimgui.h:857
func StyleColorsLight(dst []Style) {
	cdst, _ := unpackArgSStyle(dst)
	C.igStyleColorsLight(cdst)
	packSStyle(dst, cdst)
}

// SetItemDefaultFocus function as declared in cimgui/cimgui.h:859
func SetItemDefaultFocus() {
	C.igSetItemDefaultFocus()
}

// SetKeyboardFocusHere function as declared in cimgui/cimgui.h:860
func SetKeyboardFocusHere(offset int32) {
	coffset, _ := (C.int)(offset), cgoAllocsUnknown
	C.igSetKeyboardFocusHere(coffset)
}

// IsItemHovered function as declared in cimgui/cimgui.h:863
func IsItemHovered(flags int32) bool {
	cflags, _ := (C.ImGuiHoveredFlags)(flags), cgoAllocsUnknown
	__ret := C.igIsItemHovered(cflags)
	__v := (bool)(__ret)
	return __v
}

// IsItemActive function as declared in cimgui/cimgui.h:864
func IsItemActive() bool {
	__ret := C.igIsItemActive()
	__v := (bool)(__ret)
	return __v
}

// IsItemClicked function as declared in cimgui/cimgui.h:865
func IsItemClicked(mouseButton int32) bool {
	cmouseButton, _ := (C.int)(mouseButton), cgoAllocsUnknown
	__ret := C.igIsItemClicked(cmouseButton)
	__v := (bool)(__ret)
	return __v
}

// IsItemVisible function as declared in cimgui/cimgui.h:866
func IsItemVisible() bool {
	__ret := C.igIsItemVisible()
	__v := (bool)(__ret)
	return __v
}

// IsAnyItemHovered function as declared in cimgui/cimgui.h:867
func IsAnyItemHovered() bool {
	__ret := C.igIsAnyItemHovered()
	__v := (bool)(__ret)
	return __v
}

// IsAnyItemActive function as declared in cimgui/cimgui.h:868
func IsAnyItemActive() bool {
	__ret := C.igIsAnyItemActive()
	__v := (bool)(__ret)
	return __v
}

// GetItemRectMin function as declared in cimgui/cimgui.h:869
func GetItemRectMin(pOut []Vec2) {
	cpOut, _ := unpackArgSVec2(pOut)
	C.igGetItemRectMin(cpOut)
	packSVec2(pOut, cpOut)
}

// GetItemRectMax function as declared in cimgui/cimgui.h:870
func GetItemRectMax(pOut []Vec2) {
	cpOut, _ := unpackArgSVec2(pOut)
	C.igGetItemRectMax(cpOut)
	packSVec2(pOut, cpOut)
}

// GetItemRectSize function as declared in cimgui/cimgui.h:871
func GetItemRectSize(pOut []Vec2) {
	cpOut, _ := unpackArgSVec2(pOut)
	C.igGetItemRectSize(cpOut)
	packSVec2(pOut, cpOut)
}

// SetItemAllowOverlap function as declared in cimgui/cimgui.h:872
func SetItemAllowOverlap() {
	C.igSetItemAllowOverlap()
}

// IsWindowFocused function as declared in cimgui/cimgui.h:873
func IsWindowFocused(flags int32) bool {
	cflags, _ := (C.ImGuiFocusedFlags)(flags), cgoAllocsUnknown
	__ret := C.igIsWindowFocused(cflags)
	__v := (bool)(__ret)
	return __v
}

// IsWindowHovered function as declared in cimgui/cimgui.h:874
func IsWindowHovered(falgs int32) bool {
	cfalgs, _ := (C.ImGuiHoveredFlags)(falgs), cgoAllocsUnknown
	__ret := C.igIsWindowHovered(cfalgs)
	__v := (bool)(__ret)
	return __v
}

// IsAnyWindowFocused function as declared in cimgui/cimgui.h:875
func IsAnyWindowFocused() bool {
	__ret := C.igIsAnyWindowFocused()
	__v := (bool)(__ret)
	return __v
}

// IsAnyWindowHovered function as declared in cimgui/cimgui.h:876
func IsAnyWindowHovered() bool {
	__ret := C.igIsAnyWindowHovered()
	__v := (bool)(__ret)
	return __v
}

// IsRectVisible function as declared in cimgui/cimgui.h:877
func IsRectVisible(itemSize Vec2) bool {
	citemSize, _ := itemSize.PassValue()
	__ret := C.igIsRectVisible(citemSize)
	__v := (bool)(__ret)
	return __v
}

// IsRectVisible2 function as declared in cimgui/cimgui.h:878
func IsRectVisible2(rectMin []Vec2, rectMax []Vec2) bool {
	crectMin, _ := unpackArgSVec2(rectMin)
	crectMax, _ := unpackArgSVec2(rectMax)
	__ret := C.igIsRectVisible2(crectMin, crectMax)
	packSVec2(rectMax, crectMax)
	packSVec2(rectMin, crectMin)
	__v := (bool)(__ret)
	return __v
}

// GetTime function as declared in cimgui/cimgui.h:879
func GetTime() float32 {
	__ret := C.igGetTime()
	__v := (float32)(__ret)
	return __v
}

// GetFrameCount function as declared in cimgui/cimgui.h:880
func GetFrameCount() int32 {
	__ret := C.igGetFrameCount()
	__v := (int32)(__ret)
	return __v
}

// GetOverlayDrawList function as declared in cimgui/cimgui.h:882
func GetOverlayDrawList() *DrawList {
	__ret := C.igGetOverlayDrawList()
	__v := *(**DrawList)(unsafe.Pointer(&__ret))
	return __v
}

// GetDrawListSharedData function as declared in cimgui/cimgui.h:883
func GetDrawListSharedData() *DrawListSharedData {
	__ret := C.igGetDrawListSharedData()
	__v := *(**DrawListSharedData)(unsafe.Pointer(&__ret))
	return __v
}

// GetStyleColorName function as declared in cimgui/cimgui.h:885
func GetStyleColorName(idx Col) string {
	cidx, _ := (C.ImGuiCol)(idx), cgoAllocsUnknown
	__ret := C.igGetStyleColorName(cidx)
	__v := packPCharString(__ret)
	return __v
}

// CalcItemRectClosestPoint function as declared in cimgui/cimgui.h:886
func CalcItemRectClosestPoint(pOut []Vec2, pos Vec2, onEdge bool, outward float32) {
	cpOut, _ := unpackArgSVec2(pOut)
	cpos, _ := pos.PassValue()
	conEdge, _ := (C._Bool)(onEdge), cgoAllocsUnknown
	coutward, _ := (C.float)(outward), cgoAllocsUnknown
	C.igCalcItemRectClosestPoint(cpOut, cpos, conEdge, coutward)
	packSVec2(pOut, cpOut)
}

// CalcTextSize function as declared in cimgui/cimgui.h:887
func CalcTextSize(pOut []Vec2, text string, textEnd string, hideTextAfterDoubleHash bool, wrapWidth float32) {
	cpOut, _ := unpackArgSVec2(pOut)
	text = safeString(text)
	ctext, _ := unpackPCharString(text)
	textEnd = safeString(textEnd)
	ctextEnd, _ := unpackPCharString(textEnd)
	chideTextAfterDoubleHash, _ := (C._Bool)(hideTextAfterDoubleHash), cgoAllocsUnknown
	cwrapWidth, _ := (C.float)(wrapWidth), cgoAllocsUnknown
	C.igCalcTextSize(cpOut, ctext, ctextEnd, chideTextAfterDoubleHash, cwrapWidth)
	runtime.KeepAlive(textEnd)
	runtime.KeepAlive(text)
	packSVec2(pOut, cpOut)
}

// CalcListClipping function as declared in cimgui/cimgui.h:888
func CalcListClipping(itemsCount int32, itemsHeight float32, outItemsDisplayStart []int32, outItemsDisplayEnd []int32) {
	citemsCount, _ := (C.int)(itemsCount), cgoAllocsUnknown
	citemsHeight, _ := (C.float)(itemsHeight), cgoAllocsUnknown
	coutItemsDisplayStart, _ := (*C.int)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&outItemsDisplayStart)).Data)), cgoAllocsUnknown
	coutItemsDisplayEnd, _ := (*C.int)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&outItemsDisplayEnd)).Data)), cgoAllocsUnknown
	C.igCalcListClipping(citemsCount, citemsHeight, coutItemsDisplayStart, coutItemsDisplayEnd)
}

// BeginChildFrame function as declared in cimgui/cimgui.h:890
func BeginChildFrame(id ID, size Vec2, extraFlags WindowFlags) bool {
	cid, _ := (C.ImGuiID)(id), cgoAllocsUnknown
	csize, _ := size.PassValue()
	cextraFlags, _ := (C.ImGuiWindowFlags)(extraFlags), cgoAllocsUnknown
	__ret := C.igBeginChildFrame(cid, csize, cextraFlags)
	__v := (bool)(__ret)
	return __v
}

// EndChildFrame function as declared in cimgui/cimgui.h:891
func EndChildFrame() {
	C.igEndChildFrame()
}

// ColorConvertU32ToFloat4 function as declared in cimgui/cimgui.h:893
func ColorConvertU32ToFloat4(pOut []Vec4, in U32) {
	cpOut, _ := unpackArgSVec4(pOut)
	cin, _ := (C.ImU32)(in), cgoAllocsUnknown
	C.igColorConvertU32ToFloat4(cpOut, cin)
	packSVec4(pOut, cpOut)
}

// ColorConvertFloat4ToU32 function as declared in cimgui/cimgui.h:894
func ColorConvertFloat4ToU32(in Vec4) U32 {
	cin, _ := in.PassValue()
	__ret := C.igColorConvertFloat4ToU32(cin)
	__v := (U32)(__ret)
	return __v
}

// ColorConvertRGBtoHSV function as declared in cimgui/cimgui.h:895
func ColorConvertRGBtoHSV(r float32, g float32, b float32, outH []float32, outS []float32, outV []float32) {
	cr, _ := (C.float)(r), cgoAllocsUnknown
	cg, _ := (C.float)(g), cgoAllocsUnknown
	cb, _ := (C.float)(b), cgoAllocsUnknown
	coutH, _ := (*C.float)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&outH)).Data)), cgoAllocsUnknown
	coutS, _ := (*C.float)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&outS)).Data)), cgoAllocsUnknown
	coutV, _ := (*C.float)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&outV)).Data)), cgoAllocsUnknown
	C.igColorConvertRGBtoHSV(cr, cg, cb, coutH, coutS, coutV)
}

// ColorConvertHSVtoRGB function as declared in cimgui/cimgui.h:896
func ColorConvertHSVtoRGB(h float32, s float32, v float32, outR []float32, outG []float32, outB []float32) {
	ch, _ := (C.float)(h), cgoAllocsUnknown
	cs, _ := (C.float)(s), cgoAllocsUnknown
	cv, _ := (C.float)(v), cgoAllocsUnknown
	coutR, _ := (*C.float)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&outR)).Data)), cgoAllocsUnknown
	coutG, _ := (*C.float)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&outG)).Data)), cgoAllocsUnknown
	coutB, _ := (*C.float)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&outB)).Data)), cgoAllocsUnknown
	C.igColorConvertHSVtoRGB(ch, cs, cv, coutR, coutG, coutB)
}

// GetKeyIndex function as declared in cimgui/cimgui.h:899
func GetKeyIndex(imguiKey Key) int32 {
	cimguiKey, _ := (C.ImGuiKey)(imguiKey), cgoAllocsUnknown
	__ret := C.igGetKeyIndex(cimguiKey)
	__v := (int32)(__ret)
	return __v
}

// IsKeyDown function as declared in cimgui/cimgui.h:900
func IsKeyDown(userKeyIndex int32) bool {
	cuserKeyIndex, _ := (C.int)(userKeyIndex), cgoAllocsUnknown
	__ret := C.igIsKeyDown(cuserKeyIndex)
	__v := (bool)(__ret)
	return __v
}

// IsKeyPressed function as declared in cimgui/cimgui.h:901
func IsKeyPressed(userKeyIndex int32, repeat bool) bool {
	cuserKeyIndex, _ := (C.int)(userKeyIndex), cgoAllocsUnknown
	crepeat, _ := (C._Bool)(repeat), cgoAllocsUnknown
	__ret := C.igIsKeyPressed(cuserKeyIndex, crepeat)
	__v := (bool)(__ret)
	return __v
}

// IsKeyReleased function as declared in cimgui/cimgui.h:902
func IsKeyReleased(userKeyIndex int32) bool {
	cuserKeyIndex, _ := (C.int)(userKeyIndex), cgoAllocsUnknown
	__ret := C.igIsKeyReleased(cuserKeyIndex)
	__v := (bool)(__ret)
	return __v
}

// GetKeyPressedAmount function as declared in cimgui/cimgui.h:903
func GetKeyPressedAmount(keyIndex int32, repeatDelay float32, rate float32) int32 {
	ckeyIndex, _ := (C.int)(keyIndex), cgoAllocsUnknown
	crepeatDelay, _ := (C.float)(repeatDelay), cgoAllocsUnknown
	crate, _ := (C.float)(rate), cgoAllocsUnknown
	__ret := C.igGetKeyPressedAmount(ckeyIndex, crepeatDelay, crate)
	__v := (int32)(__ret)
	return __v
}

// IsMouseDown function as declared in cimgui/cimgui.h:904
func IsMouseDown(button int32) bool {
	cbutton, _ := (C.int)(button), cgoAllocsUnknown
	__ret := C.igIsMouseDown(cbutton)
	__v := (bool)(__ret)
	return __v
}

// IsMouseClicked function as declared in cimgui/cimgui.h:905
func IsMouseClicked(button int32, repeat bool) bool {
	cbutton, _ := (C.int)(button), cgoAllocsUnknown
	crepeat, _ := (C._Bool)(repeat), cgoAllocsUnknown
	__ret := C.igIsMouseClicked(cbutton, crepeat)
	__v := (bool)(__ret)
	return __v
}

// IsMouseDoubleClicked function as declared in cimgui/cimgui.h:906
func IsMouseDoubleClicked(button int32) bool {
	cbutton, _ := (C.int)(button), cgoAllocsUnknown
	__ret := C.igIsMouseDoubleClicked(cbutton)
	__v := (bool)(__ret)
	return __v
}

// IsMouseReleased function as declared in cimgui/cimgui.h:907
func IsMouseReleased(button int32) bool {
	cbutton, _ := (C.int)(button), cgoAllocsUnknown
	__ret := C.igIsMouseReleased(cbutton)
	__v := (bool)(__ret)
	return __v
}

// IsMouseDragging function as declared in cimgui/cimgui.h:908
func IsMouseDragging(button int32, lockThreshold float32) bool {
	cbutton, _ := (C.int)(button), cgoAllocsUnknown
	clockThreshold, _ := (C.float)(lockThreshold), cgoAllocsUnknown
	__ret := C.igIsMouseDragging(cbutton, clockThreshold)
	__v := (bool)(__ret)
	return __v
}

// IsMouseHoveringRect function as declared in cimgui/cimgui.h:909
func IsMouseHoveringRect(rMin Vec2, rMax Vec2, clip bool) bool {
	crMin, _ := rMin.PassValue()
	crMax, _ := rMax.PassValue()
	cclip, _ := (C._Bool)(clip), cgoAllocsUnknown
	__ret := C.igIsMouseHoveringRect(crMin, crMax, cclip)
	__v := (bool)(__ret)
	return __v
}

// IsMousePosValid function as declared in cimgui/cimgui.h:910
func IsMousePosValid(mousePos []Vec2) bool {
	cmousePos, _ := unpackArgSVec2(mousePos)
	__ret := C.igIsMousePosValid(cmousePos)
	packSVec2(mousePos, cmousePos)
	__v := (bool)(__ret)
	return __v
}

// GetMousePos function as declared in cimgui/cimgui.h:912
func GetMousePos(pOut []Vec2) {
	cpOut, _ := unpackArgSVec2(pOut)
	C.igGetMousePos(cpOut)
	packSVec2(pOut, cpOut)
}

// GetMousePosOnOpeningCurrentPopup function as declared in cimgui/cimgui.h:913
func GetMousePosOnOpeningCurrentPopup(pOut []Vec2) {
	cpOut, _ := unpackArgSVec2(pOut)
	C.igGetMousePosOnOpeningCurrentPopup(cpOut)
	packSVec2(pOut, cpOut)
}

// GetMouseDragDelta function as declared in cimgui/cimgui.h:914
func GetMouseDragDelta(pOut []Vec2, button int32, lockThreshold float32) {
	cpOut, _ := unpackArgSVec2(pOut)
	cbutton, _ := (C.int)(button), cgoAllocsUnknown
	clockThreshold, _ := (C.float)(lockThreshold), cgoAllocsUnknown
	C.igGetMouseDragDelta(cpOut, cbutton, clockThreshold)
	packSVec2(pOut, cpOut)
}

// ResetMouseDragDelta function as declared in cimgui/cimgui.h:915
func ResetMouseDragDelta(button int32) {
	cbutton, _ := (C.int)(button), cgoAllocsUnknown
	C.igResetMouseDragDelta(cbutton)
}

// GetMouseCursor function as declared in cimgui/cimgui.h:916
func GetMouseCursor() MouseCursor {
	__ret := C.igGetMouseCursor()
	__v := (MouseCursor)(__ret)
	return __v
}

// SetMouseCursor function as declared in cimgui/cimgui.h:917
func SetMouseCursor(kind MouseCursor) {
	ckind, _ := (C.ImGuiMouseCursor)(kind), cgoAllocsUnknown
	C.igSetMouseCursor(ckind)
}

// CaptureKeyboardFromApp function as declared in cimgui/cimgui.h:918
func CaptureKeyboardFromApp(capture bool) {
	ccapture, _ := (C._Bool)(capture), cgoAllocsUnknown
	C.igCaptureKeyboardFromApp(ccapture)
}

// CaptureMouseFromApp function as declared in cimgui/cimgui.h:919
func CaptureMouseFromApp(capture bool) {
	ccapture, _ := (C._Bool)(capture), cgoAllocsUnknown
	C.igCaptureMouseFromApp(ccapture)
}

// MemAlloc function as declared in cimgui/cimgui.h:922
func MemAlloc(sz uint) unsafe.Pointer {
	csz, _ := (C.size_t)(sz), cgoAllocsUnknown
	__ret := C.igMemAlloc(csz)
	__v := *(*unsafe.Pointer)(unsafe.Pointer(&__ret))
	return __v
}

// MemFree function as declared in cimgui/cimgui.h:923
func MemFree(ptr unsafe.Pointer) {
	cptr, _ := ptr, cgoAllocsUnknown
	C.igMemFree(cptr)
}

// GetClipboardText function as declared in cimgui/cimgui.h:924
func GetClipboardText() string {
	__ret := C.igGetClipboardText()
	__v := packPCharString(__ret)
	return __v
}

// SetClipboardText function as declared in cimgui/cimgui.h:925
func SetClipboardText(text string) {
	text = safeString(text)
	ctext, _ := unpackPCharString(text)
	C.igSetClipboardText(ctext)
	runtime.KeepAlive(text)
}

// GetVersion function as declared in cimgui/cimgui.h:928
func GetVersion() string {
	__ret := C.igGetVersion()
	__v := packPCharString(__ret)
	return __v
}

// DestroyContext function as declared in cimgui/cimgui.h:930
func DestroyContext(ctx []Context) {
	cctx, _ := (*C.struct_ImGuiContext)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&ctx)).Data)), cgoAllocsUnknown
	C.igDestroyContext(cctx)
}

// GetCurrentContext function as declared in cimgui/cimgui.h:931
func GetCurrentContext() *Context {
	__ret := C.igGetCurrentContext()
	__v := *(**Context)(unsafe.Pointer(&__ret))
	return __v
}

// SetCurrentContext function as declared in cimgui/cimgui.h:932
func SetCurrentContext(ctx []Context) {
	cctx, _ := (*C.struct_ImGuiContext)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&ctx)).Data)), cgoAllocsUnknown
	C.igSetCurrentContext(cctx)
}

// FontConfigDefaultConstructor function as declared in cimgui/cimgui.h:934
func FontConfigDefaultConstructor(config []FontConfig) {
	cconfig, _ := unpackArgSFontConfig(config)
	C.ImFontConfig_DefaultConstructor(cconfig)
	packSFontConfig(config, cconfig)
}

// IOAddInputCharacter function as declared in cimgui/cimgui.h:937
func IOAddInputCharacter(c uint16) {
	cc, _ := (C.ushort)(c), cgoAllocsUnknown
	C.ImGuiIO_AddInputCharacter(cc)
}

// IOAddInputCharactersUTF8 function as declared in cimgui/cimgui.h:938
func IOAddInputCharactersUTF8(utf8Chars string) {
	utf8Chars = safeString(utf8Chars)
	cutf8Chars, _ := unpackPCharString(utf8Chars)
	C.ImGuiIO_AddInputCharactersUTF8(cutf8Chars)
	runtime.KeepAlive(utf8Chars)
}

// IOClearInputCharacters function as declared in cimgui/cimgui.h:939
func IOClearInputCharacters() {
	C.ImGuiIO_ClearInputCharacters()
}

// TextFilterCreate function as declared in cimgui/cimgui.h:942
func TextFilterCreate(defaultFilter string) *TextFilter {
	defaultFilter = safeString(defaultFilter)
	cdefaultFilter, _ := unpackPCharString(defaultFilter)
	__ret := C.ImGuiTextFilter_Create(cdefaultFilter)
	runtime.KeepAlive(defaultFilter)
	__v := *(**TextFilter)(unsafe.Pointer(&__ret))
	return __v
}

// TextFilterDestroy function as declared in cimgui/cimgui.h:943
func TextFilterDestroy(filter []TextFilter) {
	cfilter, _ := (*C.struct_ImGuiTextFilter)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&filter)).Data)), cgoAllocsUnknown
	C.ImGuiTextFilter_Destroy(cfilter)
}

// TextFilterClear function as declared in cimgui/cimgui.h:944
func TextFilterClear(filter []TextFilter) {
	cfilter, _ := (*C.struct_ImGuiTextFilter)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&filter)).Data)), cgoAllocsUnknown
	C.ImGuiTextFilter_Clear(cfilter)
}

// TextFilterDraw function as declared in cimgui/cimgui.h:945
func TextFilterDraw(filter []TextFilter, label string, width float32) bool {
	cfilter, _ := (*C.struct_ImGuiTextFilter)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&filter)).Data)), cgoAllocsUnknown
	label = safeString(label)
	clabel, _ := unpackPCharString(label)
	cwidth, _ := (C.float)(width), cgoAllocsUnknown
	__ret := C.ImGuiTextFilter_Draw(cfilter, clabel, cwidth)
	runtime.KeepAlive(label)
	__v := (bool)(__ret)
	return __v
}

// TextFilterPassFilter function as declared in cimgui/cimgui.h:946
func TextFilterPassFilter(filter []TextFilter, text string, textEnd string) bool {
	cfilter, _ := (*C.struct_ImGuiTextFilter)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&filter)).Data)), cgoAllocsUnknown
	text = safeString(text)
	ctext, _ := unpackPCharString(text)
	textEnd = safeString(textEnd)
	ctextEnd, _ := unpackPCharString(textEnd)
	__ret := C.ImGuiTextFilter_PassFilter(cfilter, ctext, ctextEnd)
	runtime.KeepAlive(textEnd)
	runtime.KeepAlive(text)
	__v := (bool)(__ret)
	return __v
}

// TextFilterIsActive function as declared in cimgui/cimgui.h:947
func TextFilterIsActive(filter []TextFilter) bool {
	cfilter, _ := (*C.struct_ImGuiTextFilter)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&filter)).Data)), cgoAllocsUnknown
	__ret := C.ImGuiTextFilter_IsActive(cfilter)
	__v := (bool)(__ret)
	return __v
}

// TextFilterBuild function as declared in cimgui/cimgui.h:948
func TextFilterBuild(filter []TextFilter) {
	cfilter, _ := (*C.struct_ImGuiTextFilter)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&filter)).Data)), cgoAllocsUnknown
	C.ImGuiTextFilter_Build(cfilter)
}

// TextFilterGetInputBuf function as declared in cimgui/cimgui.h:949
func TextFilterGetInputBuf(filter []TextFilter) string {
	cfilter, _ := (*C.struct_ImGuiTextFilter)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&filter)).Data)), cgoAllocsUnknown
	__ret := C.ImGuiTextFilter_GetInputBuf(cfilter)
	__v := packPCharString(__ret)
	return __v
}

// TextBufferCreate function as declared in cimgui/cimgui.h:952
func TextBufferCreate() *TextBuffer {
	__ret := C.ImGuiTextBuffer_Create()
	__v := *(**TextBuffer)(unsafe.Pointer(&__ret))
	return __v
}

// TextBufferDestroy function as declared in cimgui/cimgui.h:953
func TextBufferDestroy(buffer []TextBuffer) {
	cbuffer, _ := (*C.struct_ImGuiTextBuffer)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&buffer)).Data)), cgoAllocsUnknown
	C.ImGuiTextBuffer_Destroy(cbuffer)
}

// TextBufferIndex function as declared in cimgui/cimgui.h:954
func TextBufferIndex(buffer []TextBuffer, i int32) byte {
	cbuffer, _ := (*C.struct_ImGuiTextBuffer)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&buffer)).Data)), cgoAllocsUnknown
	ci, _ := (C.int)(i), cgoAllocsUnknown
	__ret := C.ImGuiTextBuffer_index(cbuffer, ci)
	__v := (byte)(__ret)
	return __v
}

// TextBufferBegin function as declared in cimgui/cimgui.h:955
func TextBufferBegin(buffer []TextBuffer) string {
	cbuffer, _ := (*C.struct_ImGuiTextBuffer)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&buffer)).Data)), cgoAllocsUnknown
	__ret := C.ImGuiTextBuffer_begin(cbuffer)
	__v := packPCharString(__ret)
	return __v
}

// TextBufferEnd function as declared in cimgui/cimgui.h:956
func TextBufferEnd(buffer []TextBuffer) string {
	cbuffer, _ := (*C.struct_ImGuiTextBuffer)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&buffer)).Data)), cgoAllocsUnknown
	__ret := C.ImGuiTextBuffer_end(cbuffer)
	__v := packPCharString(__ret)
	return __v
}

// TextBufferSize function as declared in cimgui/cimgui.h:957
func TextBufferSize(buffer []TextBuffer) int32 {
	cbuffer, _ := (*C.struct_ImGuiTextBuffer)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&buffer)).Data)), cgoAllocsUnknown
	__ret := C.ImGuiTextBuffer_size(cbuffer)
	__v := (int32)(__ret)
	return __v
}

// TextBufferEmpty function as declared in cimgui/cimgui.h:958
func TextBufferEmpty(buffer []TextBuffer) bool {
	cbuffer, _ := (*C.struct_ImGuiTextBuffer)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&buffer)).Data)), cgoAllocsUnknown
	__ret := C.ImGuiTextBuffer_empty(cbuffer)
	__v := (bool)(__ret)
	return __v
}

// TextBufferClear function as declared in cimgui/cimgui.h:959
func TextBufferClear(buffer []TextBuffer) {
	cbuffer, _ := (*C.struct_ImGuiTextBuffer)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&buffer)).Data)), cgoAllocsUnknown
	C.ImGuiTextBuffer_clear(cbuffer)
}

// TextBufferCStr function as declared in cimgui/cimgui.h:960
func TextBufferCStr(buffer []TextBuffer) string {
	cbuffer, _ := (*C.struct_ImGuiTextBuffer)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&buffer)).Data)), cgoAllocsUnknown
	__ret := C.ImGuiTextBuffer_c_str(cbuffer)
	__v := packPCharString(__ret)
	return __v
}

// StorageCreate function as declared in cimgui/cimgui.h:965
func StorageCreate() *Storage {
	__ret := C.ImGuiStorage_Create()
	__v := *(**Storage)(unsafe.Pointer(&__ret))
	return __v
}

// StorageDestroy function as declared in cimgui/cimgui.h:966
func StorageDestroy(storage []Storage) {
	cstorage, _ := (*C.struct_ImGuiStorage)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&storage)).Data)), cgoAllocsUnknown
	C.ImGuiStorage_Destroy(cstorage)
}

// StorageGetInt function as declared in cimgui/cimgui.h:967
func StorageGetInt(storage []Storage, key ID, defaultVal int32) int32 {
	cstorage, _ := (*C.struct_ImGuiStorage)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&storage)).Data)), cgoAllocsUnknown
	ckey, _ := (C.ImGuiID)(key), cgoAllocsUnknown
	cdefaultVal, _ := (C.int)(defaultVal), cgoAllocsUnknown
	__ret := C.ImGuiStorage_GetInt(cstorage, ckey, cdefaultVal)
	__v := (int32)(__ret)
	return __v
}

// StorageSetInt function as declared in cimgui/cimgui.h:968
func StorageSetInt(storage []Storage, key ID, val int32) {
	cstorage, _ := (*C.struct_ImGuiStorage)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&storage)).Data)), cgoAllocsUnknown
	ckey, _ := (C.ImGuiID)(key), cgoAllocsUnknown
	cval, _ := (C.int)(val), cgoAllocsUnknown
	C.ImGuiStorage_SetInt(cstorage, ckey, cval)
}

// StorageGetBool function as declared in cimgui/cimgui.h:969
func StorageGetBool(storage []Storage, key ID, defaultVal bool) bool {
	cstorage, _ := (*C.struct_ImGuiStorage)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&storage)).Data)), cgoAllocsUnknown
	ckey, _ := (C.ImGuiID)(key), cgoAllocsUnknown
	cdefaultVal, _ := (C._Bool)(defaultVal), cgoAllocsUnknown
	__ret := C.ImGuiStorage_GetBool(cstorage, ckey, cdefaultVal)
	__v := (bool)(__ret)
	return __v
}

// StorageSetBool function as declared in cimgui/cimgui.h:970
func StorageSetBool(storage []Storage, key ID, val bool) {
	cstorage, _ := (*C.struct_ImGuiStorage)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&storage)).Data)), cgoAllocsUnknown
	ckey, _ := (C.ImGuiID)(key), cgoAllocsUnknown
	cval, _ := (C._Bool)(val), cgoAllocsUnknown
	C.ImGuiStorage_SetBool(cstorage, ckey, cval)
}

// StorageGetFloat function as declared in cimgui/cimgui.h:971
func StorageGetFloat(storage []Storage, key ID, defaultVal float32) float32 {
	cstorage, _ := (*C.struct_ImGuiStorage)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&storage)).Data)), cgoAllocsUnknown
	ckey, _ := (C.ImGuiID)(key), cgoAllocsUnknown
	cdefaultVal, _ := (C.float)(defaultVal), cgoAllocsUnknown
	__ret := C.ImGuiStorage_GetFloat(cstorage, ckey, cdefaultVal)
	__v := (float32)(__ret)
	return __v
}

// StorageSetFloat function as declared in cimgui/cimgui.h:972
func StorageSetFloat(storage []Storage, key ID, val float32) {
	cstorage, _ := (*C.struct_ImGuiStorage)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&storage)).Data)), cgoAllocsUnknown
	ckey, _ := (C.ImGuiID)(key), cgoAllocsUnknown
	cval, _ := (C.float)(val), cgoAllocsUnknown
	C.ImGuiStorage_SetFloat(cstorage, ckey, cval)
}

// StorageGetVoidPtr function as declared in cimgui/cimgui.h:973
func StorageGetVoidPtr(storage []Storage, key ID) unsafe.Pointer {
	cstorage, _ := (*C.struct_ImGuiStorage)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&storage)).Data)), cgoAllocsUnknown
	ckey, _ := (C.ImGuiID)(key), cgoAllocsUnknown
	__ret := C.ImGuiStorage_GetVoidPtr(cstorage, ckey)
	__v := *(*unsafe.Pointer)(unsafe.Pointer(&__ret))
	return __v
}

// StorageSetVoidPtr function as declared in cimgui/cimgui.h:974
func StorageSetVoidPtr(storage []Storage, key ID, val unsafe.Pointer) {
	cstorage, _ := (*C.struct_ImGuiStorage)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&storage)).Data)), cgoAllocsUnknown
	ckey, _ := (C.ImGuiID)(key), cgoAllocsUnknown
	cval, _ := val, cgoAllocsUnknown
	C.ImGuiStorage_SetVoidPtr(cstorage, ckey, cval)
}

// StorageGetIntRef function as declared in cimgui/cimgui.h:975
func StorageGetIntRef(storage []Storage, key ID, defaultVal int32) *int32 {
	cstorage, _ := (*C.struct_ImGuiStorage)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&storage)).Data)), cgoAllocsUnknown
	ckey, _ := (C.ImGuiID)(key), cgoAllocsUnknown
	cdefaultVal, _ := (C.int)(defaultVal), cgoAllocsUnknown
	__ret := C.ImGuiStorage_GetIntRef(cstorage, ckey, cdefaultVal)
	__v := *(**int32)(unsafe.Pointer(&__ret))
	return __v
}

// StorageGetBoolRef function as declared in cimgui/cimgui.h:976
func StorageGetBoolRef(storage []Storage, key ID, defaultVal bool) *bool {
	cstorage, _ := (*C.struct_ImGuiStorage)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&storage)).Data)), cgoAllocsUnknown
	ckey, _ := (C.ImGuiID)(key), cgoAllocsUnknown
	cdefaultVal, _ := (C._Bool)(defaultVal), cgoAllocsUnknown
	__ret := C.ImGuiStorage_GetBoolRef(cstorage, ckey, cdefaultVal)
	__v := *(**bool)(unsafe.Pointer(&__ret))
	return __v
}

// StorageGetFloatRef function as declared in cimgui/cimgui.h:977
func StorageGetFloatRef(storage []Storage, key ID, defaultVal float32) *float32 {
	cstorage, _ := (*C.struct_ImGuiStorage)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&storage)).Data)), cgoAllocsUnknown
	ckey, _ := (C.ImGuiID)(key), cgoAllocsUnknown
	cdefaultVal, _ := (C.float)(defaultVal), cgoAllocsUnknown
	__ret := C.ImGuiStorage_GetFloatRef(cstorage, ckey, cdefaultVal)
	__v := *(**float32)(unsafe.Pointer(&__ret))
	return __v
}

// StorageGetVoidPtrRef function as declared in cimgui/cimgui.h:978
func StorageGetVoidPtrRef(storage []Storage, key ID, defaultVal unsafe.Pointer) *unsafe.Pointer {
	cstorage, _ := (*C.struct_ImGuiStorage)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&storage)).Data)), cgoAllocsUnknown
	ckey, _ := (C.ImGuiID)(key), cgoAllocsUnknown
	cdefaultVal, _ := defaultVal, cgoAllocsUnknown
	__ret := C.ImGuiStorage_GetVoidPtrRef(cstorage, ckey, cdefaultVal)
	__v := *(**unsafe.Pointer)(unsafe.Pointer(&__ret))
	return __v
}

// StorageSetAllInt function as declared in cimgui/cimgui.h:979
func StorageSetAllInt(storage []Storage, val int32) {
	cstorage, _ := (*C.struct_ImGuiStorage)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&storage)).Data)), cgoAllocsUnknown
	cval, _ := (C.int)(val), cgoAllocsUnknown
	C.ImGuiStorage_SetAllInt(cstorage, cval)
}

// TextEditCallbackDataDeleteChars function as declared in cimgui/cimgui.h:982
func TextEditCallbackDataDeleteChars(data *TextEditCallbackData, pos int32, bytesCount int32) {
	cdata, _ := data.PassRef()
	cpos, _ := (C.int)(pos), cgoAllocsUnknown
	cbytesCount, _ := (C.int)(bytesCount), cgoAllocsUnknown
	C.ImGuiTextEditCallbackData_DeleteChars(cdata, cpos, cbytesCount)
}

// TextEditCallbackDataInsertChars function as declared in cimgui/cimgui.h:983
func TextEditCallbackDataInsertChars(data *TextEditCallbackData, pos int32, text string, textEnd string) {
	cdata, _ := data.PassRef()
	cpos, _ := (C.int)(pos), cgoAllocsUnknown
	text = safeString(text)
	ctext, _ := unpackPCharString(text)
	textEnd = safeString(textEnd)
	ctextEnd, _ := unpackPCharString(textEnd)
	C.ImGuiTextEditCallbackData_InsertChars(cdata, cpos, ctext, ctextEnd)
	runtime.KeepAlive(textEnd)
	runtime.KeepAlive(text)
}

// TextEditCallbackDataHasSelection function as declared in cimgui/cimgui.h:984
func TextEditCallbackDataHasSelection(data *TextEditCallbackData) bool {
	cdata, _ := data.PassRef()
	__ret := C.ImGuiTextEditCallbackData_HasSelection(cdata)
	__v := (bool)(__ret)
	return __v
}

// ListClipperStep function as declared in cimgui/cimgui.h:987
func ListClipperStep(clipper []ListClipper) bool {
	cclipper, _ := unpackArgSListClipper(clipper)
	__ret := C.ImGuiListClipper_Step(cclipper)
	packSListClipper(clipper, cclipper)
	__v := (bool)(__ret)
	return __v
}

// ListClipperBegin function as declared in cimgui/cimgui.h:988
func ListClipperBegin(clipper []ListClipper, count int32, itemsHeight float32) {
	cclipper, _ := unpackArgSListClipper(clipper)
	ccount, _ := (C.int)(count), cgoAllocsUnknown
	citemsHeight, _ := (C.float)(itemsHeight), cgoAllocsUnknown
	C.ImGuiListClipper_Begin(cclipper, ccount, citemsHeight)
	packSListClipper(clipper, cclipper)
}

// ListClipperEnd function as declared in cimgui/cimgui.h:989
func ListClipperEnd(clipper []ListClipper) {
	cclipper, _ := unpackArgSListClipper(clipper)
	C.ImGuiListClipper_End(cclipper)
	packSListClipper(clipper, cclipper)
}

// ListClipperGetDisplayStart function as declared in cimgui/cimgui.h:990
func ListClipperGetDisplayStart(clipper []ListClipper) int32 {
	cclipper, _ := unpackArgSListClipper(clipper)
	__ret := C.ImGuiListClipper_GetDisplayStart(cclipper)
	packSListClipper(clipper, cclipper)
	__v := (int32)(__ret)
	return __v
}

// ListClipperGetDisplayEnd function as declared in cimgui/cimgui.h:991
func ListClipperGetDisplayEnd(clipper []ListClipper) int32 {
	cclipper, _ := unpackArgSListClipper(clipper)
	__ret := C.ImGuiListClipper_GetDisplayEnd(cclipper)
	packSListClipper(clipper, cclipper)
	__v := (int32)(__ret)
	return __v
}

// DrawListGetVertexBufferSize function as declared in cimgui/cimgui.h:994
func DrawListGetVertexBufferSize(list []DrawList) int32 {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	__ret := C.ImDrawList_GetVertexBufferSize(clist)
	__v := (int32)(__ret)
	return __v
}

// DrawListGetVertexPtr function as declared in cimgui/cimgui.h:995
func DrawListGetVertexPtr(list []DrawList, n int32) *DrawVert {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	cn, _ := (C.int)(n), cgoAllocsUnknown
	__ret := C.ImDrawList_GetVertexPtr(clist, cn)
	__v := NewDrawVertRef(unsafe.Pointer(__ret))
	return __v
}

// DrawListGetIndexBufferSize function as declared in cimgui/cimgui.h:996
func DrawListGetIndexBufferSize(list []DrawList) int32 {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	__ret := C.ImDrawList_GetIndexBufferSize(clist)
	__v := (int32)(__ret)
	return __v
}

// DrawListGetIndexPtr function as declared in cimgui/cimgui.h:997
func DrawListGetIndexPtr(list []DrawList, n int32) *DrawIdx {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	cn, _ := (C.int)(n), cgoAllocsUnknown
	__ret := C.ImDrawList_GetIndexPtr(clist, cn)
	__v := *(**DrawIdx)(unsafe.Pointer(&__ret))
	return __v
}

// DrawListGetCmdSize function as declared in cimgui/cimgui.h:998
func DrawListGetCmdSize(list []DrawList) int32 {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	__ret := C.ImDrawList_GetCmdSize(clist)
	__v := (int32)(__ret)
	return __v
}

// DrawListGetCmdPtr function as declared in cimgui/cimgui.h:999
func DrawListGetCmdPtr(list []DrawList, n int32) *DrawCmd {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	cn, _ := (C.int)(n), cgoAllocsUnknown
	__ret := C.ImDrawList_GetCmdPtr(clist, cn)
	__v := NewDrawCmdRef(unsafe.Pointer(__ret))
	return __v
}

// DrawListClear function as declared in cimgui/cimgui.h:1001
func DrawListClear(list []DrawList) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	C.ImDrawList_Clear(clist)
}

// DrawListClearFreeMemory function as declared in cimgui/cimgui.h:1002
func DrawListClearFreeMemory(list []DrawList) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	C.ImDrawList_ClearFreeMemory(clist)
}

// DrawListPushClipRect function as declared in cimgui/cimgui.h:1003
func DrawListPushClipRect(list []DrawList, clipRectMin Vec2, clipRectMax Vec2, intersectWithCurrentClipRect bool) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	cclipRectMin, _ := clipRectMin.PassValue()
	cclipRectMax, _ := clipRectMax.PassValue()
	cintersectWithCurrentClipRect, _ := (C._Bool)(intersectWithCurrentClipRect), cgoAllocsUnknown
	C.ImDrawList_PushClipRect(clist, cclipRectMin, cclipRectMax, cintersectWithCurrentClipRect)
}

// DrawListPushClipRectFullScreen function as declared in cimgui/cimgui.h:1004
func DrawListPushClipRectFullScreen(list []DrawList) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	C.ImDrawList_PushClipRectFullScreen(clist)
}

// DrawListPopClipRect function as declared in cimgui/cimgui.h:1005
func DrawListPopClipRect(list []DrawList) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	C.ImDrawList_PopClipRect(clist)
}

// DrawListPushTextureID function as declared in cimgui/cimgui.h:1006
func DrawListPushTextureID(list []DrawList, textureId *TextureID) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	ctextureId, _ := (C.ImTextureID)(unsafe.Pointer(textureId)), cgoAllocsUnknown
	C.ImDrawList_PushTextureID(clist, ctextureId)
}

// DrawListPopTextureID function as declared in cimgui/cimgui.h:1007
func DrawListPopTextureID(list []DrawList) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	C.ImDrawList_PopTextureID(clist)
}

// DrawListGetClipRectMin function as declared in cimgui/cimgui.h:1008
func DrawListGetClipRectMin(pOut []Vec2, list []DrawList) {
	cpOut, _ := unpackArgSVec2(pOut)
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	C.ImDrawList_GetClipRectMin(cpOut, clist)
	packSVec2(pOut, cpOut)
}

// DrawListGetClipRectMax function as declared in cimgui/cimgui.h:1009
func DrawListGetClipRectMax(pOut []Vec2, list []DrawList) {
	cpOut, _ := unpackArgSVec2(pOut)
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	C.ImDrawList_GetClipRectMax(cpOut, clist)
	packSVec2(pOut, cpOut)
}

// DrawListAddLine function as declared in cimgui/cimgui.h:1012
func DrawListAddLine(list []DrawList, a Vec2, b Vec2, col U32, thickness float32) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	ca, _ := a.PassValue()
	cb, _ := b.PassValue()
	ccol, _ := (C.ImU32)(col), cgoAllocsUnknown
	cthickness, _ := (C.float)(thickness), cgoAllocsUnknown
	C.ImDrawList_AddLine(clist, ca, cb, ccol, cthickness)
}

// DrawListAddRect function as declared in cimgui/cimgui.h:1013
func DrawListAddRect(list []DrawList, a Vec2, b Vec2, col U32, rounding float32, roundingCornersFlags int32, thickness float32) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	ca, _ := a.PassValue()
	cb, _ := b.PassValue()
	ccol, _ := (C.ImU32)(col), cgoAllocsUnknown
	crounding, _ := (C.float)(rounding), cgoAllocsUnknown
	croundingCornersFlags, _ := (C.int)(roundingCornersFlags), cgoAllocsUnknown
	cthickness, _ := (C.float)(thickness), cgoAllocsUnknown
	C.ImDrawList_AddRect(clist, ca, cb, ccol, crounding, croundingCornersFlags, cthickness)
}

// DrawListAddRectFilled function as declared in cimgui/cimgui.h:1014
func DrawListAddRectFilled(list []DrawList, a Vec2, b Vec2, col U32, rounding float32, roundingCornersFlags int32) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	ca, _ := a.PassValue()
	cb, _ := b.PassValue()
	ccol, _ := (C.ImU32)(col), cgoAllocsUnknown
	crounding, _ := (C.float)(rounding), cgoAllocsUnknown
	croundingCornersFlags, _ := (C.int)(roundingCornersFlags), cgoAllocsUnknown
	C.ImDrawList_AddRectFilled(clist, ca, cb, ccol, crounding, croundingCornersFlags)
}

// DrawListAddRectFilledMultiColor function as declared in cimgui/cimgui.h:1015
func DrawListAddRectFilledMultiColor(list []DrawList, a Vec2, b Vec2, colUprLeft U32, colUprRight U32, colBotRight U32, colBotLeft U32) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	ca, _ := a.PassValue()
	cb, _ := b.PassValue()
	ccolUprLeft, _ := (C.ImU32)(colUprLeft), cgoAllocsUnknown
	ccolUprRight, _ := (C.ImU32)(colUprRight), cgoAllocsUnknown
	ccolBotRight, _ := (C.ImU32)(colBotRight), cgoAllocsUnknown
	ccolBotLeft, _ := (C.ImU32)(colBotLeft), cgoAllocsUnknown
	C.ImDrawList_AddRectFilledMultiColor(clist, ca, cb, ccolUprLeft, ccolUprRight, ccolBotRight, ccolBotLeft)
}

// DrawListAddQuad function as declared in cimgui/cimgui.h:1016
func DrawListAddQuad(list []DrawList, a Vec2, b Vec2, c Vec2, d Vec2, col U32, thickness float32) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	ca, _ := a.PassValue()
	cb, _ := b.PassValue()
	cc, _ := c.PassValue()
	cd, _ := d.PassValue()
	ccol, _ := (C.ImU32)(col), cgoAllocsUnknown
	cthickness, _ := (C.float)(thickness), cgoAllocsUnknown
	C.ImDrawList_AddQuad(clist, ca, cb, cc, cd, ccol, cthickness)
}

// DrawListAddQuadFilled function as declared in cimgui/cimgui.h:1017
func DrawListAddQuadFilled(list []DrawList, a Vec2, b Vec2, c Vec2, d Vec2, col U32) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	ca, _ := a.PassValue()
	cb, _ := b.PassValue()
	cc, _ := c.PassValue()
	cd, _ := d.PassValue()
	ccol, _ := (C.ImU32)(col), cgoAllocsUnknown
	C.ImDrawList_AddQuadFilled(clist, ca, cb, cc, cd, ccol)
}

// DrawListAddTriangle function as declared in cimgui/cimgui.h:1018
func DrawListAddTriangle(list []DrawList, a Vec2, b Vec2, c Vec2, col U32, thickness float32) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	ca, _ := a.PassValue()
	cb, _ := b.PassValue()
	cc, _ := c.PassValue()
	ccol, _ := (C.ImU32)(col), cgoAllocsUnknown
	cthickness, _ := (C.float)(thickness), cgoAllocsUnknown
	C.ImDrawList_AddTriangle(clist, ca, cb, cc, ccol, cthickness)
}

// DrawListAddTriangleFilled function as declared in cimgui/cimgui.h:1019
func DrawListAddTriangleFilled(list []DrawList, a Vec2, b Vec2, c Vec2, col U32) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	ca, _ := a.PassValue()
	cb, _ := b.PassValue()
	cc, _ := c.PassValue()
	ccol, _ := (C.ImU32)(col), cgoAllocsUnknown
	C.ImDrawList_AddTriangleFilled(clist, ca, cb, cc, ccol)
}

// DrawListAddCircle function as declared in cimgui/cimgui.h:1020
func DrawListAddCircle(list []DrawList, centre Vec2, radius float32, col U32, numSegments int32, thickness float32) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	ccentre, _ := centre.PassValue()
	cradius, _ := (C.float)(radius), cgoAllocsUnknown
	ccol, _ := (C.ImU32)(col), cgoAllocsUnknown
	cnumSegments, _ := (C.int)(numSegments), cgoAllocsUnknown
	cthickness, _ := (C.float)(thickness), cgoAllocsUnknown
	C.ImDrawList_AddCircle(clist, ccentre, cradius, ccol, cnumSegments, cthickness)
}

// DrawListAddCircleFilled function as declared in cimgui/cimgui.h:1021
func DrawListAddCircleFilled(list []DrawList, centre Vec2, radius float32, col U32, numSegments int32) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	ccentre, _ := centre.PassValue()
	cradius, _ := (C.float)(radius), cgoAllocsUnknown
	ccol, _ := (C.ImU32)(col), cgoAllocsUnknown
	cnumSegments, _ := (C.int)(numSegments), cgoAllocsUnknown
	C.ImDrawList_AddCircleFilled(clist, ccentre, cradius, ccol, cnumSegments)
}

// DrawListAddText function as declared in cimgui/cimgui.h:1022
func DrawListAddText(list []DrawList, pos Vec2, col U32, textBegin string, textEnd string) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	cpos, _ := pos.PassValue()
	ccol, _ := (C.ImU32)(col), cgoAllocsUnknown
	textBegin = safeString(textBegin)
	ctextBegin, _ := unpackPCharString(textBegin)
	textEnd = safeString(textEnd)
	ctextEnd, _ := unpackPCharString(textEnd)
	C.ImDrawList_AddText(clist, cpos, ccol, ctextBegin, ctextEnd)
	runtime.KeepAlive(textEnd)
	runtime.KeepAlive(textBegin)
}

// DrawListAddTextExt function as declared in cimgui/cimgui.h:1023
func DrawListAddTextExt(list []DrawList, font []Font, fontSize float32, pos Vec2, col U32, textBegin string, textEnd string, wrapWidth float32, cpuFineClipRect []Vec4) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	cfont, _ := (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&font)).Data)), cgoAllocsUnknown
	cfontSize, _ := (C.float)(fontSize), cgoAllocsUnknown
	cpos, _ := pos.PassValue()
	ccol, _ := (C.ImU32)(col), cgoAllocsUnknown
	textBegin = safeString(textBegin)
	ctextBegin, _ := unpackPCharString(textBegin)
	textEnd = safeString(textEnd)
	ctextEnd, _ := unpackPCharString(textEnd)
	cwrapWidth, _ := (C.float)(wrapWidth), cgoAllocsUnknown
	ccpuFineClipRect, _ := unpackArgSVec4(cpuFineClipRect)
	C.ImDrawList_AddTextExt(clist, cfont, cfontSize, cpos, ccol, ctextBegin, ctextEnd, cwrapWidth, ccpuFineClipRect)
	packSVec4(cpuFineClipRect, ccpuFineClipRect)
	runtime.KeepAlive(textEnd)
	runtime.KeepAlive(textBegin)
}

// DrawListAddImage function as declared in cimgui/cimgui.h:1024
func DrawListAddImage(list []DrawList, userTextureId *TextureID, a Vec2, b Vec2, uvA Vec2, uvB Vec2, col U32) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	cuserTextureId, _ := (C.ImTextureID)(unsafe.Pointer(userTextureId)), cgoAllocsUnknown
	ca, _ := a.PassValue()
	cb, _ := b.PassValue()
	cuvA, _ := uvA.PassValue()
	cuvB, _ := uvB.PassValue()
	ccol, _ := (C.ImU32)(col), cgoAllocsUnknown
	C.ImDrawList_AddImage(clist, cuserTextureId, ca, cb, cuvA, cuvB, ccol)
}

// DrawListAddImageQuad function as declared in cimgui/cimgui.h:1025
func DrawListAddImageQuad(list []DrawList, userTextureId *TextureID, a Vec2, b Vec2, c Vec2, d Vec2, uvA Vec2, uvB Vec2, uvC Vec2, uvD Vec2, col U32) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	cuserTextureId, _ := (C.ImTextureID)(unsafe.Pointer(userTextureId)), cgoAllocsUnknown
	ca, _ := a.PassValue()
	cb, _ := b.PassValue()
	cc, _ := c.PassValue()
	cd, _ := d.PassValue()
	cuvA, _ := uvA.PassValue()
	cuvB, _ := uvB.PassValue()
	cuvC, _ := uvC.PassValue()
	cuvD, _ := uvD.PassValue()
	ccol, _ := (C.ImU32)(col), cgoAllocsUnknown
	C.ImDrawList_AddImageQuad(clist, cuserTextureId, ca, cb, cc, cd, cuvA, cuvB, cuvC, cuvD, ccol)
}

// DrawListAddImageRounded function as declared in cimgui/cimgui.h:1026
func DrawListAddImageRounded(list []DrawList, userTextureId *TextureID, a Vec2, b Vec2, uvA Vec2, uvB Vec2, col U32, rounding float32, roundingCorners int32) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	cuserTextureId, _ := (C.ImTextureID)(unsafe.Pointer(userTextureId)), cgoAllocsUnknown
	ca, _ := a.PassValue()
	cb, _ := b.PassValue()
	cuvA, _ := uvA.PassValue()
	cuvB, _ := uvB.PassValue()
	ccol, _ := (C.ImU32)(col), cgoAllocsUnknown
	crounding, _ := (C.float)(rounding), cgoAllocsUnknown
	croundingCorners, _ := (C.int)(roundingCorners), cgoAllocsUnknown
	C.ImDrawList_AddImageRounded(clist, cuserTextureId, ca, cb, cuvA, cuvB, ccol, crounding, croundingCorners)
}

// DrawListAddPolyline function as declared in cimgui/cimgui.h:1027
func DrawListAddPolyline(list []DrawList, points []Vec2, numPoints int32, col U32, closed bool, thickness float32) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	cpoints, _ := unpackArgSVec2(points)
	cnumPoints, _ := (C.int)(numPoints), cgoAllocsUnknown
	ccol, _ := (C.ImU32)(col), cgoAllocsUnknown
	cclosed, _ := (C._Bool)(closed), cgoAllocsUnknown
	cthickness, _ := (C.float)(thickness), cgoAllocsUnknown
	C.ImDrawList_AddPolyline(clist, cpoints, cnumPoints, ccol, cclosed, cthickness)
	packSVec2(points, cpoints)
}

// DrawListAddConvexPolyFilled function as declared in cimgui/cimgui.h:1028
func DrawListAddConvexPolyFilled(list []DrawList, points []Vec2, numPoints int32, col U32) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	cpoints, _ := unpackArgSVec2(points)
	cnumPoints, _ := (C.int)(numPoints), cgoAllocsUnknown
	ccol, _ := (C.ImU32)(col), cgoAllocsUnknown
	C.ImDrawList_AddConvexPolyFilled(clist, cpoints, cnumPoints, ccol)
	packSVec2(points, cpoints)
}

// DrawListAddBezierCurve function as declared in cimgui/cimgui.h:1029
func DrawListAddBezierCurve(list []DrawList, pos0 Vec2, cp0 Vec2, cp1 Vec2, pos1 Vec2, col U32, thickness float32, numSegments int32) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	cpos0, _ := pos0.PassValue()
	ccp0, _ := cp0.PassValue()
	ccp1, _ := cp1.PassValue()
	cpos1, _ := pos1.PassValue()
	ccol, _ := (C.ImU32)(col), cgoAllocsUnknown
	cthickness, _ := (C.float)(thickness), cgoAllocsUnknown
	cnumSegments, _ := (C.int)(numSegments), cgoAllocsUnknown
	C.ImDrawList_AddBezierCurve(clist, cpos0, ccp0, ccp1, cpos1, ccol, cthickness, cnumSegments)
}

// DrawListPathClear function as declared in cimgui/cimgui.h:1032
func DrawListPathClear(list []DrawList) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	C.ImDrawList_PathClear(clist)
}

// DrawListPathLineTo function as declared in cimgui/cimgui.h:1033
func DrawListPathLineTo(list []DrawList, pos Vec2) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	cpos, _ := pos.PassValue()
	C.ImDrawList_PathLineTo(clist, cpos)
}

// DrawListPathLineToMergeDuplicate function as declared in cimgui/cimgui.h:1034
func DrawListPathLineToMergeDuplicate(list []DrawList, pos Vec2) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	cpos, _ := pos.PassValue()
	C.ImDrawList_PathLineToMergeDuplicate(clist, cpos)
}

// DrawListPathFillConvex function as declared in cimgui/cimgui.h:1035
func DrawListPathFillConvex(list []DrawList, col U32) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	ccol, _ := (C.ImU32)(col), cgoAllocsUnknown
	C.ImDrawList_PathFillConvex(clist, ccol)
}

// DrawListPathStroke function as declared in cimgui/cimgui.h:1036
func DrawListPathStroke(list []DrawList, col U32, closed bool, thickness float32) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	ccol, _ := (C.ImU32)(col), cgoAllocsUnknown
	cclosed, _ := (C._Bool)(closed), cgoAllocsUnknown
	cthickness, _ := (C.float)(thickness), cgoAllocsUnknown
	C.ImDrawList_PathStroke(clist, ccol, cclosed, cthickness)
}

// DrawListPathArcTo function as declared in cimgui/cimgui.h:1037
func DrawListPathArcTo(list []DrawList, centre Vec2, radius float32, aMin float32, aMax float32, numSegments int32) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	ccentre, _ := centre.PassValue()
	cradius, _ := (C.float)(radius), cgoAllocsUnknown
	caMin, _ := (C.float)(aMin), cgoAllocsUnknown
	caMax, _ := (C.float)(aMax), cgoAllocsUnknown
	cnumSegments, _ := (C.int)(numSegments), cgoAllocsUnknown
	C.ImDrawList_PathArcTo(clist, ccentre, cradius, caMin, caMax, cnumSegments)
}

// DrawListPathArcToFast function as declared in cimgui/cimgui.h:1038
func DrawListPathArcToFast(list []DrawList, centre Vec2, radius float32, aMinOf12 int32, aMaxOf12 int32) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	ccentre, _ := centre.PassValue()
	cradius, _ := (C.float)(radius), cgoAllocsUnknown
	caMinOf12, _ := (C.int)(aMinOf12), cgoAllocsUnknown
	caMaxOf12, _ := (C.int)(aMaxOf12), cgoAllocsUnknown
	C.ImDrawList_PathArcToFast(clist, ccentre, cradius, caMinOf12, caMaxOf12)
}

// DrawListPathBezierCurveTo function as declared in cimgui/cimgui.h:1039
func DrawListPathBezierCurveTo(list []DrawList, p1 Vec2, p2 Vec2, p3 Vec2, numSegments int32) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	cp1, _ := p1.PassValue()
	cp2, _ := p2.PassValue()
	cp3, _ := p3.PassValue()
	cnumSegments, _ := (C.int)(numSegments), cgoAllocsUnknown
	C.ImDrawList_PathBezierCurveTo(clist, cp1, cp2, cp3, cnumSegments)
}

// DrawListPathRect function as declared in cimgui/cimgui.h:1040
func DrawListPathRect(list []DrawList, rectMin Vec2, rectMax Vec2, rounding float32, roundingCornersFlags int32) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	crectMin, _ := rectMin.PassValue()
	crectMax, _ := rectMax.PassValue()
	crounding, _ := (C.float)(rounding), cgoAllocsUnknown
	croundingCornersFlags, _ := (C.int)(roundingCornersFlags), cgoAllocsUnknown
	C.ImDrawList_PathRect(clist, crectMin, crectMax, crounding, croundingCornersFlags)
}

// DrawListChannelsSplit function as declared in cimgui/cimgui.h:1043
func DrawListChannelsSplit(list []DrawList, channelsCount int32) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	cchannelsCount, _ := (C.int)(channelsCount), cgoAllocsUnknown
	C.ImDrawList_ChannelsSplit(clist, cchannelsCount)
}

// DrawListChannelsMerge function as declared in cimgui/cimgui.h:1044
func DrawListChannelsMerge(list []DrawList) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	C.ImDrawList_ChannelsMerge(clist)
}

// DrawListChannelsSetCurrent function as declared in cimgui/cimgui.h:1045
func DrawListChannelsSetCurrent(list []DrawList, channelIndex int32) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	cchannelIndex, _ := (C.int)(channelIndex), cgoAllocsUnknown
	C.ImDrawList_ChannelsSetCurrent(clist, cchannelIndex)
}

// DrawListAddCallback function as declared in cimgui/cimgui.h:1048
func DrawListAddCallback(list []DrawList, callback DrawCallback, callbackData unsafe.Pointer) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	ccallback, _ := callback.PassValue()
	ccallbackData, _ := callbackData, cgoAllocsUnknown
	C.ImDrawList_AddCallback(clist, ccallback, ccallbackData)
}

// DrawListAddDrawCmd function as declared in cimgui/cimgui.h:1049
func DrawListAddDrawCmd(list []DrawList) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	C.ImDrawList_AddDrawCmd(clist)
}

// DrawListPrimReserve function as declared in cimgui/cimgui.h:1052
func DrawListPrimReserve(list []DrawList, idxCount int32, vtxCount int32) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	cidxCount, _ := (C.int)(idxCount), cgoAllocsUnknown
	cvtxCount, _ := (C.int)(vtxCount), cgoAllocsUnknown
	C.ImDrawList_PrimReserve(clist, cidxCount, cvtxCount)
}

// DrawListPrimRect function as declared in cimgui/cimgui.h:1053
func DrawListPrimRect(list []DrawList, a Vec2, b Vec2, col U32) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	ca, _ := a.PassValue()
	cb, _ := b.PassValue()
	ccol, _ := (C.ImU32)(col), cgoAllocsUnknown
	C.ImDrawList_PrimRect(clist, ca, cb, ccol)
}

// DrawListPrimRectUV function as declared in cimgui/cimgui.h:1054
func DrawListPrimRectUV(list []DrawList, a Vec2, b Vec2, uvA Vec2, uvB Vec2, col U32) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	ca, _ := a.PassValue()
	cb, _ := b.PassValue()
	cuvA, _ := uvA.PassValue()
	cuvB, _ := uvB.PassValue()
	ccol, _ := (C.ImU32)(col), cgoAllocsUnknown
	C.ImDrawList_PrimRectUV(clist, ca, cb, cuvA, cuvB, ccol)
}

// DrawListPrimQuadUV function as declared in cimgui/cimgui.h:1055
func DrawListPrimQuadUV(list []DrawList, a Vec2, b Vec2, c Vec2, d Vec2, uvA Vec2, uvB Vec2, uvC Vec2, uvD Vec2, col U32) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	ca, _ := a.PassValue()
	cb, _ := b.PassValue()
	cc, _ := c.PassValue()
	cd, _ := d.PassValue()
	cuvA, _ := uvA.PassValue()
	cuvB, _ := uvB.PassValue()
	cuvC, _ := uvC.PassValue()
	cuvD, _ := uvD.PassValue()
	ccol, _ := (C.ImU32)(col), cgoAllocsUnknown
	C.ImDrawList_PrimQuadUV(clist, ca, cb, cc, cd, cuvA, cuvB, cuvC, cuvD, ccol)
}

// DrawListPrimWriteVtx function as declared in cimgui/cimgui.h:1056
func DrawListPrimWriteVtx(list []DrawList, pos Vec2, uv Vec2, col U32) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	cpos, _ := pos.PassValue()
	cuv, _ := uv.PassValue()
	ccol, _ := (C.ImU32)(col), cgoAllocsUnknown
	C.ImDrawList_PrimWriteVtx(clist, cpos, cuv, ccol)
}

// DrawListPrimWriteIdx function as declared in cimgui/cimgui.h:1057
func DrawListPrimWriteIdx(list []DrawList, idx DrawIdx) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	cidx, _ := (C.ImDrawIdx)(idx), cgoAllocsUnknown
	C.ImDrawList_PrimWriteIdx(clist, cidx)
}

// DrawListPrimVtx function as declared in cimgui/cimgui.h:1058
func DrawListPrimVtx(list []DrawList, pos Vec2, uv Vec2, col U32) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	cpos, _ := pos.PassValue()
	cuv, _ := uv.PassValue()
	ccol, _ := (C.ImU32)(col), cgoAllocsUnknown
	C.ImDrawList_PrimVtx(clist, cpos, cuv, ccol)
}

// DrawListUpdateClipRect function as declared in cimgui/cimgui.h:1059
func DrawListUpdateClipRect(list []DrawList) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	C.ImDrawList_UpdateClipRect(clist)
}

// DrawListUpdateTextureID function as declared in cimgui/cimgui.h:1060
func DrawListUpdateTextureID(list []DrawList) {
	clist, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&list)).Data)), cgoAllocsUnknown
	C.ImDrawList_UpdateTextureID(clist)
}

// DrawDataDeIndexAllBuffers function as declared in cimgui/cimgui.h:1063
func DrawDataDeIndexAllBuffers(drawData []DrawData) {
	cdrawData, _ := unpackArgSDrawData(drawData)
	C.ImDrawData_DeIndexAllBuffers(cdrawData)
	packSDrawData(drawData, cdrawData)
}

// DrawDataScaleClipRects function as declared in cimgui/cimgui.h:1064
func DrawDataScaleClipRects(drawData []DrawData, sc Vec2) {
	cdrawData, _ := unpackArgSDrawData(drawData)
	csc, _ := sc.PassValue()
	C.ImDrawData_ScaleClipRects(cdrawData, csc)
	packSDrawData(drawData, cdrawData)
}

// FontAtlasGetTexDataAsRGBA32 function as declared in cimgui/cimgui.h:1067
func FontAtlasGetTexDataAsRGBA32(atlas []FontAtlas, outPixels [][]byte, outWidth []int32, outHeight []int32, outBytesPerPixel []int32) {
	catlas, _ := (*C.struct_ImFontAtlas)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&atlas)).Data)), cgoAllocsUnknown
	coutPixels, _ := unpackArgSSByte(outPixels)
	coutWidth, _ := (*C.int)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&outWidth)).Data)), cgoAllocsUnknown
	coutHeight, _ := (*C.int)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&outHeight)).Data)), cgoAllocsUnknown
	coutBytesPerPixel, _ := (*C.int)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&outBytesPerPixel)).Data)), cgoAllocsUnknown
	C.ImFontAtlas_GetTexDataAsRGBA32(catlas, coutPixels, coutWidth, coutHeight, coutBytesPerPixel)
	packSSByte(outPixels, coutPixels)
}

// FontAtlasGetTexDataAsAlpha8 function as declared in cimgui/cimgui.h:1068
func FontAtlasGetTexDataAsAlpha8(atlas []FontAtlas, outPixels [][]byte, outWidth []int32, outHeight []int32, outBytesPerPixel []int32) {
	catlas, _ := (*C.struct_ImFontAtlas)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&atlas)).Data)), cgoAllocsUnknown
	coutPixels, _ := unpackArgSSByte(outPixels)
	coutWidth, _ := (*C.int)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&outWidth)).Data)), cgoAllocsUnknown
	coutHeight, _ := (*C.int)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&outHeight)).Data)), cgoAllocsUnknown
	coutBytesPerPixel, _ := (*C.int)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&outBytesPerPixel)).Data)), cgoAllocsUnknown
	C.ImFontAtlas_GetTexDataAsAlpha8(catlas, coutPixels, coutWidth, coutHeight, coutBytesPerPixel)
	packSSByte(outPixels, coutPixels)
}

// FontAtlasSetTexID function as declared in cimgui/cimgui.h:1069
func FontAtlasSetTexID(atlas []FontAtlas, id *TextureID) {
	catlas, _ := (*C.struct_ImFontAtlas)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&atlas)).Data)), cgoAllocsUnknown
	cid, _ := (C.ImTextureID)(unsafe.Pointer(id)), cgoAllocsUnknown
	C.ImFontAtlas_SetTexID(catlas, cid)
}

// FontAtlasAddFont function as declared in cimgui/cimgui.h:1070
func FontAtlasAddFont(atlas []FontAtlas, fontCfg []FontConfig) *Font {
	catlas, _ := (*C.struct_ImFontAtlas)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&atlas)).Data)), cgoAllocsUnknown
	cfontCfg, _ := unpackArgSFontConfig(fontCfg)
	__ret := C.ImFontAtlas_AddFont(catlas, cfontCfg)
	packSFontConfig(fontCfg, cfontCfg)
	__v := *(**Font)(unsafe.Pointer(&__ret))
	return __v
}

// FontAtlasAddFontDefault function as declared in cimgui/cimgui.h:1071
func FontAtlasAddFontDefault(atlas []FontAtlas, fontCfg []FontConfig) *Font {
	catlas, _ := (*C.struct_ImFontAtlas)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&atlas)).Data)), cgoAllocsUnknown
	cfontCfg, _ := unpackArgSFontConfig(fontCfg)
	__ret := C.ImFontAtlas_AddFontDefault(catlas, cfontCfg)
	packSFontConfig(fontCfg, cfontCfg)
	__v := *(**Font)(unsafe.Pointer(&__ret))
	return __v
}

// FontAtlasAddFontFromFileTTF function as declared in cimgui/cimgui.h:1072
func FontAtlasAddFontFromFileTTF(atlas []FontAtlas, filename string, sizePixels float32, fontCfg []FontConfig, glyphRanges []Wchar) *Font {
	catlas, _ := (*C.struct_ImFontAtlas)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&atlas)).Data)), cgoAllocsUnknown
	filename = safeString(filename)
	cfilename, _ := unpackPCharString(filename)
	csizePixels, _ := (C.float)(sizePixels), cgoAllocsUnknown
	cfontCfg, _ := unpackArgSFontConfig(fontCfg)
	cglyphRanges, _ := (*C.ImWchar)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&glyphRanges)).Data)), cgoAllocsUnknown
	__ret := C.ImFontAtlas_AddFontFromFileTTF(catlas, cfilename, csizePixels, cfontCfg, cglyphRanges)
	packSFontConfig(fontCfg, cfontCfg)
	runtime.KeepAlive(filename)
	__v := *(**Font)(unsafe.Pointer(&__ret))
	return __v
}

// FontAtlasAddFontFromMemoryTTF function as declared in cimgui/cimgui.h:1073
func FontAtlasAddFontFromMemoryTTF(atlas []FontAtlas, fontData unsafe.Pointer, fontSize int32, sizePixels float32, fontCfg []FontConfig, glyphRanges []Wchar) *Font {
	catlas, _ := (*C.struct_ImFontAtlas)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&atlas)).Data)), cgoAllocsUnknown
	cfontData, _ := fontData, cgoAllocsUnknown
	cfontSize, _ := (C.int)(fontSize), cgoAllocsUnknown
	csizePixels, _ := (C.float)(sizePixels), cgoAllocsUnknown
	cfontCfg, _ := unpackArgSFontConfig(fontCfg)
	cglyphRanges, _ := (*C.ImWchar)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&glyphRanges)).Data)), cgoAllocsUnknown
	__ret := C.ImFontAtlas_AddFontFromMemoryTTF(catlas, cfontData, cfontSize, csizePixels, cfontCfg, cglyphRanges)
	packSFontConfig(fontCfg, cfontCfg)
	__v := *(**Font)(unsafe.Pointer(&__ret))
	return __v
}

// FontAtlasAddFontFromMemoryCompressedTTF function as declared in cimgui/cimgui.h:1074
func FontAtlasAddFontFromMemoryCompressedTTF(atlas []FontAtlas, compressedFontData unsafe.Pointer, compressedFontSize int32, sizePixels float32, fontCfg []FontConfig, glyphRanges []Wchar) *Font {
	catlas, _ := (*C.struct_ImFontAtlas)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&atlas)).Data)), cgoAllocsUnknown
	ccompressedFontData, _ := compressedFontData, cgoAllocsUnknown
	ccompressedFontSize, _ := (C.int)(compressedFontSize), cgoAllocsUnknown
	csizePixels, _ := (C.float)(sizePixels), cgoAllocsUnknown
	cfontCfg, _ := unpackArgSFontConfig(fontCfg)
	cglyphRanges, _ := (*C.ImWchar)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&glyphRanges)).Data)), cgoAllocsUnknown
	__ret := C.ImFontAtlas_AddFontFromMemoryCompressedTTF(catlas, ccompressedFontData, ccompressedFontSize, csizePixels, cfontCfg, cglyphRanges)
	packSFontConfig(fontCfg, cfontCfg)
	__v := *(**Font)(unsafe.Pointer(&__ret))
	return __v
}

// FontAtlasAddFontFromMemoryCompressedBase85TTF function as declared in cimgui/cimgui.h:1075
func FontAtlasAddFontFromMemoryCompressedBase85TTF(atlas []FontAtlas, compressedFontDataBase85 string, sizePixels float32, fontCfg []FontConfig, glyphRanges []Wchar) *Font {
	catlas, _ := (*C.struct_ImFontAtlas)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&atlas)).Data)), cgoAllocsUnknown
	compressedFontDataBase85 = safeString(compressedFontDataBase85)
	ccompressedFontDataBase85, _ := unpackPCharString(compressedFontDataBase85)
	csizePixels, _ := (C.float)(sizePixels), cgoAllocsUnknown
	cfontCfg, _ := unpackArgSFontConfig(fontCfg)
	cglyphRanges, _ := (*C.ImWchar)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&glyphRanges)).Data)), cgoAllocsUnknown
	__ret := C.ImFontAtlas_AddFontFromMemoryCompressedBase85TTF(catlas, ccompressedFontDataBase85, csizePixels, cfontCfg, cglyphRanges)
	packSFontConfig(fontCfg, cfontCfg)
	runtime.KeepAlive(compressedFontDataBase85)
	__v := *(**Font)(unsafe.Pointer(&__ret))
	return __v
}

// FontAtlasClearTexData function as declared in cimgui/cimgui.h:1076
func FontAtlasClearTexData(atlas []FontAtlas) {
	catlas, _ := (*C.struct_ImFontAtlas)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&atlas)).Data)), cgoAllocsUnknown
	C.ImFontAtlas_ClearTexData(catlas)
}

// FontAtlasClear function as declared in cimgui/cimgui.h:1077
func FontAtlasClear(atlas []FontAtlas) {
	catlas, _ := (*C.struct_ImFontAtlas)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&atlas)).Data)), cgoAllocsUnknown
	C.ImFontAtlas_Clear(catlas)
}

// FontAtlasGetGlyphRangesDefault function as declared in cimgui/cimgui.h:1078
func FontAtlasGetGlyphRangesDefault(atlas []FontAtlas) *Wchar {
	catlas, _ := (*C.struct_ImFontAtlas)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&atlas)).Data)), cgoAllocsUnknown
	__ret := C.ImFontAtlas_GetGlyphRangesDefault(catlas)
	__v := *(**Wchar)(unsafe.Pointer(&__ret))
	return __v
}

// FontAtlasGetGlyphRangesKorean function as declared in cimgui/cimgui.h:1079
func FontAtlasGetGlyphRangesKorean(atlas []FontAtlas) *Wchar {
	catlas, _ := (*C.struct_ImFontAtlas)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&atlas)).Data)), cgoAllocsUnknown
	__ret := C.ImFontAtlas_GetGlyphRangesKorean(catlas)
	__v := *(**Wchar)(unsafe.Pointer(&__ret))
	return __v
}

// FontAtlasGetGlyphRangesJapanese function as declared in cimgui/cimgui.h:1080
func FontAtlasGetGlyphRangesJapanese(atlas []FontAtlas) *Wchar {
	catlas, _ := (*C.struct_ImFontAtlas)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&atlas)).Data)), cgoAllocsUnknown
	__ret := C.ImFontAtlas_GetGlyphRangesJapanese(catlas)
	__v := *(**Wchar)(unsafe.Pointer(&__ret))
	return __v
}

// FontAtlasGetGlyphRangesChinese function as declared in cimgui/cimgui.h:1081
func FontAtlasGetGlyphRangesChinese(atlas []FontAtlas) *Wchar {
	catlas, _ := (*C.struct_ImFontAtlas)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&atlas)).Data)), cgoAllocsUnknown
	__ret := C.ImFontAtlas_GetGlyphRangesChinese(catlas)
	__v := *(**Wchar)(unsafe.Pointer(&__ret))
	return __v
}

// FontAtlasGetGlyphRangesCyrillic function as declared in cimgui/cimgui.h:1082
func FontAtlasGetGlyphRangesCyrillic(atlas []FontAtlas) *Wchar {
	catlas, _ := (*C.struct_ImFontAtlas)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&atlas)).Data)), cgoAllocsUnknown
	__ret := C.ImFontAtlas_GetGlyphRangesCyrillic(catlas)
	__v := *(**Wchar)(unsafe.Pointer(&__ret))
	return __v
}

// FontAtlasGetGlyphRangesThai function as declared in cimgui/cimgui.h:1083
func FontAtlasGetGlyphRangesThai(atlas []FontAtlas) *Wchar {
	catlas, _ := (*C.struct_ImFontAtlas)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&atlas)).Data)), cgoAllocsUnknown
	__ret := C.ImFontAtlas_GetGlyphRangesThai(catlas)
	__v := *(**Wchar)(unsafe.Pointer(&__ret))
	return __v
}

// FontAtlasGetTexID function as declared in cimgui/cimgui.h:1085
func FontAtlasGetTexID(atlas []FontAtlas) *TextureID {
	catlas, _ := (*C.struct_ImFontAtlas)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&atlas)).Data)), cgoAllocsUnknown
	__ret := C.ImFontAtlas_GetTexID(catlas)
	__v := *(**TextureID)(unsafe.Pointer(&__ret))
	return __v
}

// FontAtlasGetTexPixelsAlpha8 function as declared in cimgui/cimgui.h:1086
func FontAtlasGetTexPixelsAlpha8(atlas []FontAtlas) *byte {
	catlas, _ := (*C.struct_ImFontAtlas)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&atlas)).Data)), cgoAllocsUnknown
	__ret := C.ImFontAtlas_GetTexPixelsAlpha8(catlas)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// FontAtlasGetTexPixelsRGBA32 function as declared in cimgui/cimgui.h:1087
func FontAtlasGetTexPixelsRGBA32(atlas []FontAtlas) *uint32 {
	catlas, _ := (*C.struct_ImFontAtlas)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&atlas)).Data)), cgoAllocsUnknown
	__ret := C.ImFontAtlas_GetTexPixelsRGBA32(catlas)
	__v := *(**uint32)(unsafe.Pointer(&__ret))
	return __v
}

// FontAtlasGetTexWidth function as declared in cimgui/cimgui.h:1088
func FontAtlasGetTexWidth(atlas []FontAtlas) int32 {
	catlas, _ := (*C.struct_ImFontAtlas)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&atlas)).Data)), cgoAllocsUnknown
	__ret := C.ImFontAtlas_GetTexWidth(catlas)
	__v := (int32)(__ret)
	return __v
}

// FontAtlasGetTexHeight function as declared in cimgui/cimgui.h:1089
func FontAtlasGetTexHeight(atlas []FontAtlas) int32 {
	catlas, _ := (*C.struct_ImFontAtlas)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&atlas)).Data)), cgoAllocsUnknown
	__ret := C.ImFontAtlas_GetTexHeight(catlas)
	__v := (int32)(__ret)
	return __v
}

// FontAtlasGetTexDesiredWidth function as declared in cimgui/cimgui.h:1090
func FontAtlasGetTexDesiredWidth(atlas []FontAtlas) int32 {
	catlas, _ := (*C.struct_ImFontAtlas)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&atlas)).Data)), cgoAllocsUnknown
	__ret := C.ImFontAtlas_GetTexDesiredWidth(catlas)
	__v := (int32)(__ret)
	return __v
}

// FontAtlasSetTexDesiredWidth function as declared in cimgui/cimgui.h:1091
func FontAtlasSetTexDesiredWidth(atlas []FontAtlas, texDesiredWidth int32) {
	catlas, _ := (*C.struct_ImFontAtlas)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&atlas)).Data)), cgoAllocsUnknown
	ctexDesiredWidth, _ := (C.int)(texDesiredWidth), cgoAllocsUnknown
	C.ImFontAtlas_SetTexDesiredWidth(catlas, ctexDesiredWidth)
}

// FontAtlasGetTexGlyphPadding function as declared in cimgui/cimgui.h:1092
func FontAtlasGetTexGlyphPadding(atlas []FontAtlas) int32 {
	catlas, _ := (*C.struct_ImFontAtlas)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&atlas)).Data)), cgoAllocsUnknown
	__ret := C.ImFontAtlas_GetTexGlyphPadding(catlas)
	__v := (int32)(__ret)
	return __v
}

// FontAtlasSetTexGlyphPadding function as declared in cimgui/cimgui.h:1093
func FontAtlasSetTexGlyphPadding(atlas []FontAtlas, texGlyphPadding int32) {
	catlas, _ := (*C.struct_ImFontAtlas)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&atlas)).Data)), cgoAllocsUnknown
	ctexGlyphPadding, _ := (C.int)(texGlyphPadding), cgoAllocsUnknown
	C.ImFontAtlas_SetTexGlyphPadding(catlas, ctexGlyphPadding)
}

// FontAtlasGetTexUvWhitePixel function as declared in cimgui/cimgui.h:1094
func FontAtlasGetTexUvWhitePixel(atlas []FontAtlas, pOut []Vec2) {
	catlas, _ := (*C.struct_ImFontAtlas)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&atlas)).Data)), cgoAllocsUnknown
	cpOut, _ := unpackArgSVec2(pOut)
	C.ImFontAtlas_GetTexUvWhitePixel(catlas, cpOut)
	packSVec2(pOut, cpOut)
}

// FontAtlasFontsSize function as declared in cimgui/cimgui.h:1097
func FontAtlasFontsSize(atlas []FontAtlas) int32 {
	catlas, _ := (*C.struct_ImFontAtlas)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&atlas)).Data)), cgoAllocsUnknown
	__ret := C.ImFontAtlas_Fonts_size(catlas)
	__v := (int32)(__ret)
	return __v
}

// FontAtlasFontsIndex function as declared in cimgui/cimgui.h:1098
func FontAtlasFontsIndex(atlas []FontAtlas, index int32) *Font {
	catlas, _ := (*C.struct_ImFontAtlas)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&atlas)).Data)), cgoAllocsUnknown
	cindex, _ := (C.int)(index), cgoAllocsUnknown
	__ret := C.ImFontAtlas_Fonts_index(catlas, cindex)
	__v := *(**Font)(unsafe.Pointer(&__ret))
	return __v
}

// FontGetFontSize function as declared in cimgui/cimgui.h:1101
func FontGetFontSize(font []Font) float32 {
	cfont, _ := (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&font)).Data)), cgoAllocsUnknown
	__ret := C.ImFont_GetFontSize(cfont)
	__v := (float32)(__ret)
	return __v
}

// FontSetFontSize function as declared in cimgui/cimgui.h:1102
func FontSetFontSize(font []Font, fontSize float32) {
	cfont, _ := (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&font)).Data)), cgoAllocsUnknown
	cfontSize, _ := (C.float)(fontSize), cgoAllocsUnknown
	C.ImFont_SetFontSize(cfont, cfontSize)
}

// FontGetScale function as declared in cimgui/cimgui.h:1103
func FontGetScale(font []Font) float32 {
	cfont, _ := (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&font)).Data)), cgoAllocsUnknown
	__ret := C.ImFont_GetScale(cfont)
	__v := (float32)(__ret)
	return __v
}

// FontSetScale function as declared in cimgui/cimgui.h:1104
func FontSetScale(font []Font, scale float32) {
	cfont, _ := (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&font)).Data)), cgoAllocsUnknown
	cscale, _ := (C.float)(scale), cgoAllocsUnknown
	C.ImFont_SetScale(cfont, cscale)
}

// FontGetDisplayOffset function as declared in cimgui/cimgui.h:1105
func FontGetDisplayOffset(font []Font, pOut []Vec2) {
	cfont, _ := (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&font)).Data)), cgoAllocsUnknown
	cpOut, _ := unpackArgSVec2(pOut)
	C.ImFont_GetDisplayOffset(cfont, cpOut)
	packSVec2(pOut, cpOut)
}

// FontGetFallbackGlyph function as declared in cimgui/cimgui.h:1106
func FontGetFallbackGlyph(font []Font) *Glyph {
	cfont, _ := (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&font)).Data)), cgoAllocsUnknown
	__ret := C.ImFont_GetFallbackGlyph(cfont)
	__v := *(**Glyph)(unsafe.Pointer(&__ret))
	return __v
}

// FontSetFallbackGlyph function as declared in cimgui/cimgui.h:1107
func FontSetFallbackGlyph(font []Font, fallbackGlyph []Glyph) {
	cfont, _ := (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&font)).Data)), cgoAllocsUnknown
	cfallbackGlyph, _ := (*C.struct_Glyph)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&fallbackGlyph)).Data)), cgoAllocsUnknown
	C.ImFont_SetFallbackGlyph(cfont, cfallbackGlyph)
}

// FontGetFallbackAdvanceX function as declared in cimgui/cimgui.h:1108
func FontGetFallbackAdvanceX(font []Font) float32 {
	cfont, _ := (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&font)).Data)), cgoAllocsUnknown
	__ret := C.ImFont_GetFallbackAdvanceX(cfont)
	__v := (float32)(__ret)
	return __v
}

// FontGetFallbackChar function as declared in cimgui/cimgui.h:1109
func FontGetFallbackChar(font []Font) Wchar {
	cfont, _ := (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&font)).Data)), cgoAllocsUnknown
	__ret := C.ImFont_GetFallbackChar(cfont)
	__v := (Wchar)(__ret)
	return __v
}

// FontGetConfigDataCount function as declared in cimgui/cimgui.h:1110
func FontGetConfigDataCount(font []Font) int16 {
	cfont, _ := (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&font)).Data)), cgoAllocsUnknown
	__ret := C.ImFont_GetConfigDataCount(cfont)
	__v := (int16)(__ret)
	return __v
}

// FontGetConfigData function as declared in cimgui/cimgui.h:1111
func FontGetConfigData(font []Font) *FontConfig {
	cfont, _ := (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&font)).Data)), cgoAllocsUnknown
	__ret := C.ImFont_GetConfigData(cfont)
	__v := NewFontConfigRef(unsafe.Pointer(__ret))
	return __v
}

// FontGetContainerAtlas function as declared in cimgui/cimgui.h:1112
func FontGetContainerAtlas(font []Font) *FontAtlas {
	cfont, _ := (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&font)).Data)), cgoAllocsUnknown
	__ret := C.ImFont_GetContainerAtlas(cfont)
	__v := *(**FontAtlas)(unsafe.Pointer(&__ret))
	return __v
}

// FontGetAscent function as declared in cimgui/cimgui.h:1113
func FontGetAscent(font []Font) float32 {
	cfont, _ := (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&font)).Data)), cgoAllocsUnknown
	__ret := C.ImFont_GetAscent(cfont)
	__v := (float32)(__ret)
	return __v
}

// FontGetDescent function as declared in cimgui/cimgui.h:1114
func FontGetDescent(font []Font) float32 {
	cfont, _ := (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&font)).Data)), cgoAllocsUnknown
	__ret := C.ImFont_GetDescent(cfont)
	__v := (float32)(__ret)
	return __v
}

// FontGetMetricsTotalSurface function as declared in cimgui/cimgui.h:1115
func FontGetMetricsTotalSurface(font []Font) int32 {
	cfont, _ := (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&font)).Data)), cgoAllocsUnknown
	__ret := C.ImFont_GetMetricsTotalSurface(cfont)
	__v := (int32)(__ret)
	return __v
}

// FontClearOutputData function as declared in cimgui/cimgui.h:1116
func FontClearOutputData(font []Font) {
	cfont, _ := (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&font)).Data)), cgoAllocsUnknown
	C.ImFont_ClearOutputData(cfont)
}

// FontBuildLookupTable function as declared in cimgui/cimgui.h:1117
func FontBuildLookupTable(font []Font) {
	cfont, _ := (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&font)).Data)), cgoAllocsUnknown
	C.ImFont_BuildLookupTable(cfont)
}

// FontFindGlyph function as declared in cimgui/cimgui.h:1118
func FontFindGlyph(font []Font, c Wchar) *Glyph {
	cfont, _ := (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&font)).Data)), cgoAllocsUnknown
	cc, _ := (C.ImWchar)(c), cgoAllocsUnknown
	__ret := C.ImFont_FindGlyph(cfont, cc)
	__v := *(**Glyph)(unsafe.Pointer(&__ret))
	return __v
}

// FontSetFallbackChar function as declared in cimgui/cimgui.h:1119
func FontSetFallbackChar(font []Font, c Wchar) {
	cfont, _ := (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&font)).Data)), cgoAllocsUnknown
	cc, _ := (C.ImWchar)(c), cgoAllocsUnknown
	C.ImFont_SetFallbackChar(cfont, cc)
}

// FontGetCharAdvance function as declared in cimgui/cimgui.h:1120
func FontGetCharAdvance(font []Font, c Wchar) float32 {
	cfont, _ := (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&font)).Data)), cgoAllocsUnknown
	cc, _ := (C.ImWchar)(c), cgoAllocsUnknown
	__ret := C.ImFont_GetCharAdvance(cfont, cc)
	__v := (float32)(__ret)
	return __v
}

// FontIsLoaded function as declared in cimgui/cimgui.h:1121
func FontIsLoaded(font []Font) bool {
	cfont, _ := (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&font)).Data)), cgoAllocsUnknown
	__ret := C.ImFont_IsLoaded(cfont)
	__v := (bool)(__ret)
	return __v
}

// FontGetDebugName function as declared in cimgui/cimgui.h:1122
func FontGetDebugName(font []Font) string {
	cfont, _ := (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&font)).Data)), cgoAllocsUnknown
	__ret := C.ImFont_GetDebugName(cfont)
	__v := packPCharString(__ret)
	return __v
}

// FontCalcTextSizeA function as declared in cimgui/cimgui.h:1123
func FontCalcTextSizeA(font []Font, pOut []Vec2, size float32, maxWidth float32, wrapWidth float32, textBegin string, textEnd string, remaining []string) {
	cfont, _ := (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&font)).Data)), cgoAllocsUnknown
	cpOut, _ := unpackArgSVec2(pOut)
	csize, _ := (C.float)(size), cgoAllocsUnknown
	cmaxWidth, _ := (C.float)(maxWidth), cgoAllocsUnknown
	cwrapWidth, _ := (C.float)(wrapWidth), cgoAllocsUnknown
	textBegin = safeString(textBegin)
	ctextBegin, _ := unpackPCharString(textBegin)
	textEnd = safeString(textEnd)
	ctextEnd, _ := unpackPCharString(textEnd)
	cremaining, _ := unpackArgSString(remaining)
	C.ImFont_CalcTextSizeA(cfont, cpOut, csize, cmaxWidth, cwrapWidth, ctextBegin, ctextEnd, cremaining)
	packSString(remaining, cremaining)
	runtime.KeepAlive(textEnd)
	runtime.KeepAlive(textBegin)
	packSVec2(pOut, cpOut)
}

// FontCalcWordWrapPositionA function as declared in cimgui/cimgui.h:1124
func FontCalcWordWrapPositionA(font []Font, scale float32, text string, textEnd string, wrapWidth float32) string {
	cfont, _ := (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&font)).Data)), cgoAllocsUnknown
	cscale, _ := (C.float)(scale), cgoAllocsUnknown
	text = safeString(text)
	ctext, _ := unpackPCharString(text)
	textEnd = safeString(textEnd)
	ctextEnd, _ := unpackPCharString(textEnd)
	cwrapWidth, _ := (C.float)(wrapWidth), cgoAllocsUnknown
	__ret := C.ImFont_CalcWordWrapPositionA(cfont, cscale, ctext, ctextEnd, cwrapWidth)
	runtime.KeepAlive(textEnd)
	runtime.KeepAlive(text)
	__v := packPCharString(__ret)
	return __v
}

// FontRenderChar function as declared in cimgui/cimgui.h:1125
func FontRenderChar(font []Font, drawList []DrawList, size float32, pos Vec2, col U32, c uint16) {
	cfont, _ := (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&font)).Data)), cgoAllocsUnknown
	cdrawList, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&drawList)).Data)), cgoAllocsUnknown
	csize, _ := (C.float)(size), cgoAllocsUnknown
	cpos, _ := pos.PassValue()
	ccol, _ := (C.ImU32)(col), cgoAllocsUnknown
	cc, _ := (C.ushort)(c), cgoAllocsUnknown
	C.ImFont_RenderChar(cfont, cdrawList, csize, cpos, ccol, cc)
}

// FontRenderText function as declared in cimgui/cimgui.h:1126
func FontRenderText(font []Font, drawList []DrawList, size float32, pos Vec2, col U32, clipRect []Vec4, textBegin string, textEnd string, wrapWidth float32, cpuFineClip bool) {
	cfont, _ := (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&font)).Data)), cgoAllocsUnknown
	cdrawList, _ := (*C.struct_ImDrawList)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&drawList)).Data)), cgoAllocsUnknown
	csize, _ := (C.float)(size), cgoAllocsUnknown
	cpos, _ := pos.PassValue()
	ccol, _ := (C.ImU32)(col), cgoAllocsUnknown
	cclipRect, _ := unpackArgSVec4(clipRect)
	textBegin = safeString(textBegin)
	ctextBegin, _ := unpackPCharString(textBegin)
	textEnd = safeString(textEnd)
	ctextEnd, _ := unpackPCharString(textEnd)
	cwrapWidth, _ := (C.float)(wrapWidth), cgoAllocsUnknown
	ccpuFineClip, _ := (C._Bool)(cpuFineClip), cgoAllocsUnknown
	C.ImFont_RenderText(cfont, cdrawList, csize, cpos, ccol, cclipRect, ctextBegin, ctextEnd, cwrapWidth, ccpuFineClip)
	runtime.KeepAlive(textEnd)
	runtime.KeepAlive(textBegin)
	packSVec4(clipRect, cclipRect)
}

// FontGlyphsSize function as declared in cimgui/cimgui.h:1128
func FontGlyphsSize(font []Font) int32 {
	cfont, _ := (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&font)).Data)), cgoAllocsUnknown
	__ret := C.ImFont_Glyphs_size(cfont)
	__v := (int32)(__ret)
	return __v
}

// FontGlyphsIndex function as declared in cimgui/cimgui.h:1129
func FontGlyphsIndex(font []Font, index int32) *Glyph {
	cfont, _ := (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&font)).Data)), cgoAllocsUnknown
	cindex, _ := (C.int)(index), cgoAllocsUnknown
	__ret := C.ImFont_Glyphs_index(cfont, cindex)
	__v := *(**Glyph)(unsafe.Pointer(&__ret))
	return __v
}

// FontIndexXAdvanceSize function as declared in cimgui/cimgui.h:1131
func FontIndexXAdvanceSize(font []Font) int32 {
	cfont, _ := (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&font)).Data)), cgoAllocsUnknown
	__ret := C.ImFont_IndexXAdvance_size(cfont)
	__v := (int32)(__ret)
	return __v
}

// FontIndexXAdvanceIndex function as declared in cimgui/cimgui.h:1132
func FontIndexXAdvanceIndex(font []Font, index int32) float32 {
	cfont, _ := (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&font)).Data)), cgoAllocsUnknown
	cindex, _ := (C.int)(index), cgoAllocsUnknown
	__ret := C.ImFont_IndexXAdvance_index(cfont, cindex)
	__v := (float32)(__ret)
	return __v
}

// FontIndexLookupSize function as declared in cimgui/cimgui.h:1134
func FontIndexLookupSize(font []Font) int32 {
	cfont, _ := (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&font)).Data)), cgoAllocsUnknown
	__ret := C.ImFont_IndexLookup_size(cfont)
	__v := (int32)(__ret)
	return __v
}

// FontIndexLookupIndex function as declared in cimgui/cimgui.h:1135
func FontIndexLookupIndex(font []Font, index int32) uint16 {
	cfont, _ := (*C.struct_ImFont)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&font)).Data)), cgoAllocsUnknown
	cindex, _ := (C.int)(index), cgoAllocsUnknown
	__ret := C.ImFont_IndexLookup_index(cfont, cindex)
	__v := (uint16)(__ret)
	return __v
}
