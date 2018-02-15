// WARNING: This file has automatically been generated on Fri, 16 Feb 2018 00:12:48 EET.
// By https://git.io/c-for-go. DO NOT EDIT.

package imgui

/*
#include "cimgui/cimgui/cimgui.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// DrawIdx type as declared in cimgui/cimgui.h:52
type DrawIdx uint16

// U32 type as declared in cimgui/cimgui.h:53
type U32 uint32

// Wchar type as declared in cimgui/cimgui.h:54
type Wchar uint16

// TextureID type as declared in cimgui/cimgui.h:55
type TextureID unsafe.Pointer

// ID type as declared in cimgui/cimgui.h:56
type ID uint32

// Col type as declared in cimgui/cimgui.h:57
type Col int32

// StyleVar type as declared in cimgui/cimgui.h:58
type StyleVar int32

// Key type as declared in cimgui/cimgui.h:59
type Key int32

// ColorEditFlags type as declared in cimgui/cimgui.h:60
type ColorEditFlags int32

// MouseCursor type as declared in cimgui/cimgui.h:61
type MouseCursor int32

// WindowFlags type as declared in cimgui/cimgui.h:62
type WindowFlags int32

// Cond type as declared in cimgui/cimgui.h:63
type Cond int32

// ColumnsFlags type as declared in cimgui/cimgui.h:64
type ColumnsFlags int32

// InputTextFlags type as declared in cimgui/cimgui.h:65
type InputTextFlags int32

// SelectableFlags type as declared in cimgui/cimgui.h:66
type SelectableFlags int32

// TreeNodeFlags type as declared in cimgui/cimgui.h:67
type TreeNodeFlags int32

// TextEditCallback type as declared in cimgui/cimgui.h:74
type TextEditCallback func(data *TextEditCallbackData) int32

// SizeConstraintCallback type as declared in cimgui/cimgui.h:75
type SizeConstraintCallback func(data *SizeConstraintCallbackData)

// DrawCallback type as declared in cimgui/cimgui.h:76
type DrawCallback func(parentList *DrawList, cmd *DrawCmd)

// U64 type as declared in cimgui/cimgui.h:80
type U64 uint64

// DrawCmd as declared in cimgui/cimgui.h:40
type DrawCmd struct {
	ElemCount        uint32
	ClipRect         Vec4
	TextureId        *TextureID
	UserCallback     DrawCallback
	UserCallbackData unsafe.Pointer
	ref111e6f2b      *C.struct_ImDrawCmd
	allocs111e6f2b   interface{}
}

// DrawData as declared in cimgui/cimgui.h:30
type DrawData struct {
	Valid          bool
	CmdLists       [][]DrawList
	CmdListsCount  int32
	TotalVtxCount  int32
	TotalIdxCount  int32
	ref9a158ae0    *C.struct_ImDrawData
	allocs9a158ae0 interface{}
}

// DrawList as declared in cimgui/cimgui.h:35
type DrawList C.struct_ImDrawList

// DrawVert as declared in cimgui/cimgui.h:499
type DrawVert struct {
	Pos            Vec2
	Uv             Vec2
	Col            U32
	ref5c8bfe45    *C.struct_ImDrawVert
	allocs5c8bfe45 interface{}
}

// Font as declared in cimgui/cimgui.h:37
type Font C.struct_ImFont

// FontAtlas as declared in cimgui/cimgui.h:39
type FontAtlas C.struct_ImFontAtlas

// FontConfig as declared in cimgui/cimgui.h:38
type FontConfig struct {
	FontData             unsafe.Pointer
	FontDataSize         int32
	FontDataOwnedByAtlas bool
	FontNo               int32
	SizePixels           float32
	OversampleH          int32
	OversampleV          int32
	PixelSnapH           bool
	GlyphExtraSpacing    Vec2
	GlyphOffset          Vec2
	GlyphRanges          []Wchar
	MergeMode            bool
	RasterizerFlags      uint32
	RasterizerMultiply   float32
	Name                 [32]byte
	DstFont              []Font
	ref5c4f67a0          *C.struct_ImFontConfig
	allocs5c4f67a0       interface{}
}

// IO as declared in cimgui/cimgui.h:28
type IO struct {
	DisplaySize             Vec2
	DeltaTime               float32
	IniSavingRate           float32
	IniFilename             string
	LogFilename             string
	MouseDoubleClickTime    float32
	MouseDoubleClickMaxDist float32
	MouseDragThreshold      float32
	KeyMap                  [19]int32
	KeyRepeatDelay          float32
	KeyRepeatRate           float32
	UserData                unsafe.Pointer
	Fonts                   []FontAtlas
	FontGlobalScale         float32
	FontAllowUserScaling    bool
	FontDefault             []Font
	DisplayFramebufferScale Vec2
	DisplayVisibleMin       Vec2
	DisplayVisibleMax       Vec2
	OptMacOSXBehaviors      bool
	OptCursorBlink          bool
	RenderDrawListsFn       *func(data []DrawData)
	GetClipboardTextFn      *func(userData unsafe.Pointer) string
	SetClipboardTextFn      *func(userData unsafe.Pointer, text string)
	ClipboardUserData       unsafe.Pointer
	MemAllocFn              *func(sz uint) unsafe.Pointer
	MemFreeFn               *func(ptr unsafe.Pointer)
	ESetInputScreenPosFn    *func(x int32, y int32)
	EWindowHandle           unsafe.Pointer
	MousePos                Vec2
	MouseDown               [5]bool
	MouseWheel              float32
	MouseDrawCursor         bool
	KeyCtrl                 bool
	KeyShift                bool
	KeyAlt                  bool
	KeySuper                bool
	KeysDown                [512]bool
	InputCharacters         [17]Wchar
	WantCaptureMouse        bool
	WantCaptureKeyboard     bool
	WantTextInput           bool
	Framerate               float32
	MetricsAllocs           int32
	MetricsRenderVertices   int32
	MetricsRenderIndices    int32
	MetricsActiveWindows    int32
	MouseDelta              Vec2
	MousePosPrev            Vec2
	MouseClicked            [5]bool
	MouseClickedPos         [5]Vec2
	MouseClickedTime        [5]float32
	MouseDoubleClicked      [5]bool
	MouseReleased           [5]bool
	MouseDownOwned          [5]bool
	MouseDownDuration       [5]float32
	MouseDownDurationPrev   [5]float32
	MouseDragMaxDistanceAbs [5]Vec2
	MouseDragMaxDistanceSqr [5]float32
	KeysDownDuration        [512]float32
	KeysDownDurationPrev    [512]float32
	ref4700f756             *C.struct_ImGuiIO
	allocs4700f756          interface{}
}

// ListClipper as declared in cimgui/cimgui.h:41
type ListClipper struct {
	StartPosY      float32
	ItemsHeight    float32
	ItemsCount     int32
	StepNo         int32
	DisplayStart   int32
	DisplayEnd     int32
	refd52a46bd    *C.struct_ImGuiListClipper
	allocsd52a46bd interface{}
}

// Payload as declared in cimgui/cimgui.h:43
type Payload struct {
	Data           unsafe.Pointer
	DataSize       int32
	SourceId       ID
	SourceParentId ID
	DataFrameCount int32
	DataType       [9]byte
	Preview        bool
	Delivery       bool
	refea983e4e    *C.struct_ImGuiPayload
	allocsea983e4e interface{}
}

// SizeConstraintCallbackData as declared in cimgui/cimgui.h:34
type SizeConstraintCallbackData struct {
	UserData       unsafe.Pointer
	Pos            Vec2
	CurrentSize    Vec2
	DesiredSize    Vec2
	ref24c5ad70    *C.struct_ImGuiSizeConstraintCallbackData
	allocs24c5ad70 interface{}
}

// Storage as declared in cimgui/cimgui.h:36
type Storage C.struct_ImGuiStorage

// Style as declared in cimgui/cimgui.h:29
type Style struct {
	Alpha                  float32
	WindowPadding          Vec2
	WindowRounding         float32
	WindowBorderSize       float32
	WindowMinSize          Vec2
	WindowTitleAlign       Vec2
	ChildRounding          float32
	ChildBorderSize        float32
	PopupRounding          float32
	PopupBorderSize        float32
	FramePadding           Vec2
	FrameRounding          float32
	FrameBorderSize        float32
	ItemSpacing            Vec2
	ItemInnerSpacing       Vec2
	TouchExtraPadding      Vec2
	IndentSpacing          float32
	ColumnsMinSpacing      float32
	ScrollbarSize          float32
	ScrollbarRounding      float32
	GrabMinSize            float32
	GrabRounding           float32
	ButtonTextAlign        Vec2
	DisplayWindowPadding   Vec2
	DisplaySafeAreaPadding Vec2
	AntiAliasedLines       bool
	AntiAliasedFill        bool
	CurveTessellationTol   float32
	Colors                 [43]Vec4
	ref6e47ee0d            *C.struct_ImGuiStyle
	allocs6e47ee0d         interface{}
}

// TextEditCallbackData as declared in cimgui/cimgui.h:33
type TextEditCallbackData struct {
	EventFlag      InputTextFlags
	Flags          InputTextFlags
	UserData       unsafe.Pointer
	ReadOnly       bool
	EventChar      Wchar
	EventKey       Key
	Buf            []byte
	BufTextLen     int32
	BufSize        int32
	BufDirty       bool
	CursorPos      int32
	SelectionStart int32
	SelectionEnd   int32
	ref61acec8     *C.struct_ImGuiTextEditCallbackData
	allocs61acec8  interface{}
}

// TextFilter as declared in cimgui/cimgui.h:42
type TextFilter C.struct_ImGuiTextFilter

// Vec2 as declared in cimgui/cimgui.h:31
type Vec2 struct {
	X              float32
	Y              float32
	ref74e98a33    *C.struct_ImVec2
	allocs74e98a33 interface{}
}

// Vec4 as declared in cimgui/cimgui.h:32
type Vec4 struct {
	X              float32
	Y              float32
	Z              float32
	W              float32
	ref9d8a2f06    *C.struct_ImVec4
	allocs9d8a2f06 interface{}
}
