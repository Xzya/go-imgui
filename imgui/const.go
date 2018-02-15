// WARNING: This file has automatically been generated on Fri, 16 Feb 2018 00:12:48 EET.
// By https://git.io/c-for-go. DO NOT EDIT.

package imgui

/*
#include "cimgui/cimgui/cimgui.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

const (
	// IMFONTGLYPH as defined in cimgui/cimgui.h:49
	IMFONTGLYPH = 0
)

const (
	// WindowFlagsNoTitleBar as declared in cimgui/cimgui.h:96
	WindowFlagsNoTitleBar = 1 << 0
	// WindowFlagsNoResize as declared in cimgui/cimgui.h:97
	WindowFlagsNoResize = 1 << 1
	// WindowFlagsNoMove as declared in cimgui/cimgui.h:98
	WindowFlagsNoMove = 1 << 2
	// WindowFlagsNoScrollbar as declared in cimgui/cimgui.h:99
	WindowFlagsNoScrollbar = 1 << 3
	// WindowFlagsNoScrollWithMouse as declared in cimgui/cimgui.h:100
	WindowFlagsNoScrollWithMouse = 1 << 4
	// WindowFlagsNoCollapse as declared in cimgui/cimgui.h:101
	WindowFlagsNoCollapse = 1 << 5
	// WindowFlagsAlwaysAutoResize as declared in cimgui/cimgui.h:102
	WindowFlagsAlwaysAutoResize = 1 << 6
	// WindowFlagsNoSavedSettings as declared in cimgui/cimgui.h:104
	WindowFlagsNoSavedSettings = 1 << 8
	// WindowFlagsNoInputs as declared in cimgui/cimgui.h:105
	WindowFlagsNoInputs = 1 << 9
	// WindowFlagsMenuBar as declared in cimgui/cimgui.h:106
	WindowFlagsMenuBar = 1 << 10
	// WindowFlagsHorizontalScrollbar as declared in cimgui/cimgui.h:107
	WindowFlagsHorizontalScrollbar = 1 << 11
	// WindowFlagsNoFocusOnAppearing as declared in cimgui/cimgui.h:108
	WindowFlagsNoFocusOnAppearing = 1 << 12
	// WindowFlagsNoBringToFrontOnFocus as declared in cimgui/cimgui.h:109
	WindowFlagsNoBringToFrontOnFocus = 1 << 13
	// WindowFlagsAlwaysVerticalScrollbar as declared in cimgui/cimgui.h:110
	WindowFlagsAlwaysVerticalScrollbar = 1 << 14
	// WindowFlagsAlwaysHorizontalScrollbar as declared in cimgui/cimgui.h:111
	WindowFlagsAlwaysHorizontalScrollbar = 1 << 15
	// WindowFlagsAlwaysUseWindowPadding as declared in cimgui/cimgui.h:112
	WindowFlagsAlwaysUseWindowPadding = 1 << 16
	// WindowFlagsResizeFromAnySide as declared in cimgui/cimgui.h:113
	WindowFlagsResizeFromAnySide = 1 << 17
)

const (
	// InputTextFlagsCharsDecimal as declared in cimgui/cimgui.h:118
	InputTextFlagsCharsDecimal = 1 << 0
	// InputTextFlagsCharsHexadecimal as declared in cimgui/cimgui.h:119
	InputTextFlagsCharsHexadecimal = 1 << 1
	// InputTextFlagsCharsUppercase as declared in cimgui/cimgui.h:120
	InputTextFlagsCharsUppercase = 1 << 2
	// InputTextFlagsCharsNoBlank as declared in cimgui/cimgui.h:121
	InputTextFlagsCharsNoBlank = 1 << 3
	// InputTextFlagsAutoSelectAll as declared in cimgui/cimgui.h:122
	InputTextFlagsAutoSelectAll = 1 << 4
	// InputTextFlagsEnterReturnsTrue as declared in cimgui/cimgui.h:123
	InputTextFlagsEnterReturnsTrue = 1 << 5
	// InputTextFlagsCallbackCompletion as declared in cimgui/cimgui.h:124
	InputTextFlagsCallbackCompletion = 1 << 6
	// InputTextFlagsCallbackHistory as declared in cimgui/cimgui.h:125
	InputTextFlagsCallbackHistory = 1 << 7
	// InputTextFlagsCallbackAlways as declared in cimgui/cimgui.h:126
	InputTextFlagsCallbackAlways = 1 << 8
	// InputTextFlagsCallbackCharFilter as declared in cimgui/cimgui.h:127
	InputTextFlagsCallbackCharFilter = 1 << 9
	// InputTextFlagsAllowTabInput as declared in cimgui/cimgui.h:128
	InputTextFlagsAllowTabInput = 1 << 10
	// InputTextFlagsCtrlEnterForNewLine as declared in cimgui/cimgui.h:129
	InputTextFlagsCtrlEnterForNewLine = 1 << 11
	// InputTextFlagsNoHorizontalScroll as declared in cimgui/cimgui.h:130
	InputTextFlagsNoHorizontalScroll = 1 << 12
	// InputTextFlagsAlwaysInsertMode as declared in cimgui/cimgui.h:131
	InputTextFlagsAlwaysInsertMode = 1 << 13
	// InputTextFlagsReadOnly as declared in cimgui/cimgui.h:132
	InputTextFlagsReadOnly = 1 << 14
	// InputTextFlagsPassword as declared in cimgui/cimgui.h:133
	InputTextFlagsPassword = 1 << 15
	// InputTextFlagsNoUndoRedo as declared in cimgui/cimgui.h:134
	InputTextFlagsNoUndoRedo = 1 << 16
)

const (
	// TreeNodeFlagsSelected as declared in cimgui/cimgui.h:139
	TreeNodeFlagsSelected = 1 << 0
	// TreeNodeFlagsFramed as declared in cimgui/cimgui.h:140
	TreeNodeFlagsFramed = 1 << 1
	// TreeNodeFlagsAllowItemOverlap as declared in cimgui/cimgui.h:141
	TreeNodeFlagsAllowItemOverlap = 1 << 2
	// TreeNodeFlagsNoTreePushOnOpen as declared in cimgui/cimgui.h:142
	TreeNodeFlagsNoTreePushOnOpen = 1 << 3
	// TreeNodeFlagsNoAutoOpenOnLog as declared in cimgui/cimgui.h:143
	TreeNodeFlagsNoAutoOpenOnLog = 1 << 4
	// TreeNodeFlagsDefaultOpen as declared in cimgui/cimgui.h:144
	TreeNodeFlagsDefaultOpen = 1 << 5
	// TreeNodeFlagsOpenOnDoubleClick as declared in cimgui/cimgui.h:145
	TreeNodeFlagsOpenOnDoubleClick = 1 << 6
	// TreeNodeFlagsOpenOnArrow as declared in cimgui/cimgui.h:146
	TreeNodeFlagsOpenOnArrow = 1 << 7
	// TreeNodeFlagsLeaf as declared in cimgui/cimgui.h:147
	TreeNodeFlagsLeaf = 1 << 8
	// TreeNodeFlagsBullet as declared in cimgui/cimgui.h:148
	TreeNodeFlagsBullet = 1 << 9
	// TreeNodeFlagsFramePadding as declared in cimgui/cimgui.h:149
	TreeNodeFlagsFramePadding = 1 << 10
	// TreeNodeFlagsCollapsingHeader as declared in cimgui/cimgui.h:150
	TreeNodeFlagsCollapsingHeader = TreeNodeFlagsFramed | TreeNodeFlagsNoAutoOpenOnLog
)

const (
	// SelectableFlagsDontClosePopups as declared in cimgui/cimgui.h:155
	SelectableFlagsDontClosePopups = 1 << 0
	// SelectableFlagsSpanAllColumns as declared in cimgui/cimgui.h:156
	SelectableFlagsSpanAllColumns = 1 << 1
	// SelectableFlagsAllowDoubleClick as declared in cimgui/cimgui.h:157
	SelectableFlagsAllowDoubleClick = 1 << 2
)

// ComboFlags as declared in cimgui/cimgui.h:160
type ComboFlags int32

// ComboFlags enumeration from cimgui/cimgui.h:160
const (
	ComboFlagsPopupAlignLeft = 1 << 0
	ComboFlagsHeightSmall    = 1 << 1
	ComboFlagsHeightRegular  = 1 << 2
	ComboFlagsHeightLarge    = 1 << 3
	ComboFlagsHeightLargest  = 1 << 4
	ComboFlagsHeightMask     = ComboFlagsHeightSmall | ComboFlagsHeightRegular | ComboFlagsHeightLarge | ComboFlagsHeightLargest
)

// FocusedFlags as declared in cimgui/cimgui.h:170
type FocusedFlags int32

// FocusedFlags enumeration from cimgui/cimgui.h:170
const (
	FocusedFlagsChildWindows        = 1 << 0
	FocusedFlagsRootWindow          = 1 << 1
	FocusedFlagsRootAndChildWindows = FocusedFlagsRootWindow | FocusedFlagsChildWindows
)

// HoveredFlags as declared in cimgui/cimgui.h:177
type HoveredFlags int32

// HoveredFlags enumeration from cimgui/cimgui.h:177
const (
	HoveredFlagsChildWindows                 = 1 << 0
	HoveredFlagsRootWindow                   = 1 << 1
	HoveredFlagsAllowWhenBlockedByPopup      = 1 << 2
	HoveredFlagsAllowWhenBlockedByActiveItem = 1 << 4
	HoveredFlagsAllowWhenOverlapped          = 1 << 5
	HoveredFlagsRectOnly                     = HoveredFlagsAllowWhenBlockedByPopup | HoveredFlagsAllowWhenBlockedByActiveItem | HoveredFlagsAllowWhenOverlapped
	HoveredFlagsRootAndChildWindows          = HoveredFlagsRootWindow | HoveredFlagsChildWindows
)

// DragDropFlags as declared in cimgui/cimgui.h:189
type DragDropFlags int32

// DragDropFlags enumeration from cimgui/cimgui.h:189
const (
	DragDropFlagsSourceNoPreviewTooltip   = 1 << 0
	DragDropFlagsSourceNoDisableHover     = 1 << 1
	DragDropFlagsSourceNoHoldToOpenOthers = 1 << 2
	DragDropFlagsSourceAllowNullID        = 1 << 3
	DragDropFlagsSourceExtern             = 1 << 4
	DragDropFlagsAcceptBeforeDelivery     = 1 << 10
	DragDropFlagsAcceptNoDrawDefaultRect  = 1 << 11
	DragDropFlagsAcceptPeekOnly           = DragDropFlagsAcceptBeforeDelivery | DragDropFlagsAcceptNoDrawDefaultRect
)

const (
	// KeyTab as declared in cimgui/cimgui.h:203
	KeyTab = iota
	// KeyLeftArrow as declared in cimgui/cimgui.h:204
	KeyLeftArrow = 1
	// KeyRightArrow as declared in cimgui/cimgui.h:205
	KeyRightArrow = 2
	// KeyUpArrow as declared in cimgui/cimgui.h:206
	KeyUpArrow = 3
	// KeyDownArrow as declared in cimgui/cimgui.h:207
	KeyDownArrow = 4
	// KeyPageUp as declared in cimgui/cimgui.h:208
	KeyPageUp = 5
	// KeyPageDown as declared in cimgui/cimgui.h:209
	KeyPageDown = 6
	// KeyHome as declared in cimgui/cimgui.h:210
	KeyHome = 7
	// KeyEnd as declared in cimgui/cimgui.h:211
	KeyEnd = 8
	// KeyDelete as declared in cimgui/cimgui.h:212
	KeyDelete = 9
	// KeyBackspace as declared in cimgui/cimgui.h:213
	KeyBackspace = 10
	// KeyEnter as declared in cimgui/cimgui.h:214
	KeyEnter = 11
	// KeyEscape as declared in cimgui/cimgui.h:215
	KeyEscape = 12
	// KeyA as declared in cimgui/cimgui.h:216
	KeyA = 13
	// KeyC as declared in cimgui/cimgui.h:217
	KeyC = 14
	// KeyV as declared in cimgui/cimgui.h:218
	KeyV = 15
	// KeyX as declared in cimgui/cimgui.h:219
	KeyX = 16
	// KeyY as declared in cimgui/cimgui.h:220
	KeyY = 17
	// KeyZ as declared in cimgui/cimgui.h:221
	KeyZ = 18
	// KeyCOUNT as declared in cimgui/cimgui.h:222
	KeyCOUNT = 19
)

const (
	// ColText as declared in cimgui/cimgui.h:227
	ColText = iota
	// ColTextDisabled as declared in cimgui/cimgui.h:228
	ColTextDisabled = 1
	// ColWindowBg as declared in cimgui/cimgui.h:229
	ColWindowBg = 2
	// ColChildBg as declared in cimgui/cimgui.h:230
	ColChildBg = 3
	// ColPopupBg as declared in cimgui/cimgui.h:231
	ColPopupBg = 4
	// ColBorder as declared in cimgui/cimgui.h:232
	ColBorder = 5
	// ColBorderShadow as declared in cimgui/cimgui.h:233
	ColBorderShadow = 6
	// ColFrameBg as declared in cimgui/cimgui.h:234
	ColFrameBg = 7
	// ColFrameBgHovered as declared in cimgui/cimgui.h:235
	ColFrameBgHovered = 8
	// ColFrameBgActive as declared in cimgui/cimgui.h:236
	ColFrameBgActive = 9
	// ColTitleBg as declared in cimgui/cimgui.h:237
	ColTitleBg = 10
	// ColTitleBgActive as declared in cimgui/cimgui.h:238
	ColTitleBgActive = 11
	// ColTitleBgCollapsed as declared in cimgui/cimgui.h:239
	ColTitleBgCollapsed = 12
	// ColMenuBarBg as declared in cimgui/cimgui.h:240
	ColMenuBarBg = 13
	// ColScrollbarBg as declared in cimgui/cimgui.h:241
	ColScrollbarBg = 14
	// ColScrollbarGrab as declared in cimgui/cimgui.h:242
	ColScrollbarGrab = 15
	// ColScrollbarGrabHovered as declared in cimgui/cimgui.h:243
	ColScrollbarGrabHovered = 16
	// ColScrollbarGrabActive as declared in cimgui/cimgui.h:244
	ColScrollbarGrabActive = 17
	// ColCheckMark as declared in cimgui/cimgui.h:245
	ColCheckMark = 18
	// ColSliderGrab as declared in cimgui/cimgui.h:246
	ColSliderGrab = 19
	// ColSliderGrabActive as declared in cimgui/cimgui.h:247
	ColSliderGrabActive = 20
	// ColButton as declared in cimgui/cimgui.h:248
	ColButton = 21
	// ColButtonHovered as declared in cimgui/cimgui.h:249
	ColButtonHovered = 22
	// ColButtonActive as declared in cimgui/cimgui.h:250
	ColButtonActive = 23
	// ColHeader as declared in cimgui/cimgui.h:251
	ColHeader = 24
	// ColHeaderHovered as declared in cimgui/cimgui.h:252
	ColHeaderHovered = 25
	// ColHeaderActive as declared in cimgui/cimgui.h:253
	ColHeaderActive = 26
	// ColSeparator as declared in cimgui/cimgui.h:254
	ColSeparator = 27
	// ColSeparatorHovered as declared in cimgui/cimgui.h:255
	ColSeparatorHovered = 28
	// ColSeparatorActive as declared in cimgui/cimgui.h:256
	ColSeparatorActive = 29
	// ColResizeGrip as declared in cimgui/cimgui.h:257
	ColResizeGrip = 30
	// ColResizeGripHovered as declared in cimgui/cimgui.h:258
	ColResizeGripHovered = 31
	// ColResizeGripActive as declared in cimgui/cimgui.h:259
	ColResizeGripActive = 32
	// ColCloseButton as declared in cimgui/cimgui.h:260
	ColCloseButton = 33
	// ColCloseButtonHovered as declared in cimgui/cimgui.h:261
	ColCloseButtonHovered = 34
	// ColCloseButtonActive as declared in cimgui/cimgui.h:262
	ColCloseButtonActive = 35
	// ColPlotLines as declared in cimgui/cimgui.h:263
	ColPlotLines = 36
	// ColPlotLinesHovered as declared in cimgui/cimgui.h:264
	ColPlotLinesHovered = 37
	// ColPlotHistogram as declared in cimgui/cimgui.h:265
	ColPlotHistogram = 38
	// ColPlotHistogramHovered as declared in cimgui/cimgui.h:266
	ColPlotHistogramHovered = 39
	// ColTextSelectedBg as declared in cimgui/cimgui.h:267
	ColTextSelectedBg = 40
	// ColModalWindowDarkening as declared in cimgui/cimgui.h:268
	ColModalWindowDarkening = 41
	// ColDragDropTarget as declared in cimgui/cimgui.h:269
	ColDragDropTarget = 42
	// ColCOUNT as declared in cimgui/cimgui.h:270
	ColCOUNT = 43
)

const (
	// StyleVarAlpha as declared in cimgui/cimgui.h:275
	StyleVarAlpha = iota
	// StyleVarWindowPadding as declared in cimgui/cimgui.h:276
	StyleVarWindowPadding = 1
	// StyleVarWindowRounding as declared in cimgui/cimgui.h:277
	StyleVarWindowRounding = 2
	// StyleVarWindowBorderSize as declared in cimgui/cimgui.h:278
	StyleVarWindowBorderSize = 3
	// StyleVarWindowMinSize as declared in cimgui/cimgui.h:279
	StyleVarWindowMinSize = 4
	// StyleVarChildRounding as declared in cimgui/cimgui.h:280
	StyleVarChildRounding = 5
	// StyleVarChildBorderSize as declared in cimgui/cimgui.h:281
	StyleVarChildBorderSize = 6
	// StyleVarPopupRounding as declared in cimgui/cimgui.h:282
	StyleVarPopupRounding = 7
	// StyleVarPopupBorderSize as declared in cimgui/cimgui.h:283
	StyleVarPopupBorderSize = 8
	// StyleVarFramePadding as declared in cimgui/cimgui.h:284
	StyleVarFramePadding = 9
	// StyleVarFrameRounding as declared in cimgui/cimgui.h:285
	StyleVarFrameRounding = 10
	// StyleVarFrameBorderSize as declared in cimgui/cimgui.h:286
	StyleVarFrameBorderSize = 11
	// StyleVarItemSpacing as declared in cimgui/cimgui.h:287
	StyleVarItemSpacing = 12
	// StyleVarItemInnerSpacing as declared in cimgui/cimgui.h:288
	StyleVarItemInnerSpacing = 13
	// StyleVarIndentSpacing as declared in cimgui/cimgui.h:289
	StyleVarIndentSpacing = 14
	// StyleVarGrabMinSize as declared in cimgui/cimgui.h:290
	StyleVarGrabMinSize = 15
	// StyleVarButtonTextAlign as declared in cimgui/cimgui.h:291
	StyleVarButtonTextAlign = 16
	// StyleVarCount as declared in cimgui/cimgui.h:292
	StyleVarCount = 17
)

const (
	// ColorEditFlagsNoAlpha as declared in cimgui/cimgui.h:297
	ColorEditFlagsNoAlpha = 1 << 1
	// ColorEditFlagsNoPicker as declared in cimgui/cimgui.h:298
	ColorEditFlagsNoPicker = 1 << 2
	// ColorEditFlagsNoOptions as declared in cimgui/cimgui.h:299
	ColorEditFlagsNoOptions = 1 << 3
	// ColorEditFlagsNoSmallPreview as declared in cimgui/cimgui.h:300
	ColorEditFlagsNoSmallPreview = 1 << 4
	// ColorEditFlagsNoInputs as declared in cimgui/cimgui.h:301
	ColorEditFlagsNoInputs = 1 << 5
	// ColorEditFlagsNoTooltip as declared in cimgui/cimgui.h:302
	ColorEditFlagsNoTooltip = 1 << 6
	// ColorEditFlagsNoLabel as declared in cimgui/cimgui.h:303
	ColorEditFlagsNoLabel = 1 << 7
	// ColorEditFlagsNoSidePreview as declared in cimgui/cimgui.h:304
	ColorEditFlagsNoSidePreview = 1 << 8
	// ColorEditFlagsAlphaBar as declared in cimgui/cimgui.h:305
	ColorEditFlagsAlphaBar = 1 << 9
	// ColorEditFlagsAlphaPreview as declared in cimgui/cimgui.h:306
	ColorEditFlagsAlphaPreview = 1 << 10
	// ColorEditFlagsAlphaPreviewHalf as declared in cimgui/cimgui.h:307
	ColorEditFlagsAlphaPreviewHalf = 1 << 11
	// ColorEditFlagsHDR as declared in cimgui/cimgui.h:308
	ColorEditFlagsHDR = 1 << 12
	// ColorEditFlagsRGB as declared in cimgui/cimgui.h:309
	ColorEditFlagsRGB = 1 << 13
	// ColorEditFlagsHSV as declared in cimgui/cimgui.h:310
	ColorEditFlagsHSV = 1 << 14
	// ColorEditFlagsHEX as declared in cimgui/cimgui.h:311
	ColorEditFlagsHEX = 1 << 15
	// ColorEditFlagsUint8 as declared in cimgui/cimgui.h:312
	ColorEditFlagsUint8 = 1 << 16
	// ColorEditFlagsFloat as declared in cimgui/cimgui.h:313
	ColorEditFlagsFloat = 1 << 17
	// ColorEditFlagsPickerHueBar as declared in cimgui/cimgui.h:314
	ColorEditFlagsPickerHueBar = 1 << 18
	// ColorEditFlagsPickerHueWheel as declared in cimgui/cimgui.h:315
	ColorEditFlagsPickerHueWheel = 1 << 19
)

const (
	// MouseCursorNone as declared in cimgui/cimgui.h:320
	MouseCursorNone = -1
	// MouseCursorArrow as declared in cimgui/cimgui.h:321
	MouseCursorArrow = 0
	// MouseCursorTextInput as declared in cimgui/cimgui.h:322
	MouseCursorTextInput = 1
	// MouseCursorMove as declared in cimgui/cimgui.h:323
	MouseCursorMove = 2
	// MouseCursorResizeNS as declared in cimgui/cimgui.h:324
	MouseCursorResizeNS = 3
	// MouseCursorResizeEW as declared in cimgui/cimgui.h:325
	MouseCursorResizeEW = 4
	// MouseCursorResizeNESW as declared in cimgui/cimgui.h:326
	MouseCursorResizeNESW = 5
	// MouseCursorResizeNWSE as declared in cimgui/cimgui.h:327
	MouseCursorResizeNWSE = 6
	// MouseCursorCount as declared in cimgui/cimgui.h:328
	MouseCursorCount = 7
)

const (
	// CondAlways as declared in cimgui/cimgui.h:333
	CondAlways = 1 << 0
	// CondOnce as declared in cimgui/cimgui.h:334
	CondOnce = 1 << 1
	// CondFirstUseEver as declared in cimgui/cimgui.h:335
	CondFirstUseEver = 1 << 2
	// CondAppearing as declared in cimgui/cimgui.h:336
	CondAppearing = 1 << 3
)

// DrawCornerFlags as declared in cimgui/cimgui.h:339
type DrawCornerFlags int32

// DrawCornerFlags enumeration from cimgui/cimgui.h:339
const (
	DrawCornerFlagsTopLeft  = 1 << 0
	DrawCornerFlagsTopRight = 1 << 1
	DrawCornerFlagsBotLeft  = 1 << 2
	DrawCornerFlagsBotRight = 1 << 3
	DrawCornerFlagsTop      = DrawCornerFlagsTopLeft | DrawCornerFlagsTopRight
	DrawCornerFlagsBot      = DrawCornerFlagsBotLeft | DrawCornerFlagsBotRight
	DrawCornerFlagsLeft     = DrawCornerFlagsTopLeft | DrawCornerFlagsBotLeft
	DrawCornerFlagsRight    = DrawCornerFlagsTopRight | DrawCornerFlagsBotRight
	DrawCornerFlagsAll      = 0xF
)

// DrawListFlags as declared in cimgui/cimgui.h:352
type DrawListFlags int32

// DrawListFlags enumeration from cimgui/cimgui.h:352
const (
	DrawListFlagsAntiAliasedLines = 1 << 0
	DrawListFlagsAntiAliasedFill  = 1 << 1
)
