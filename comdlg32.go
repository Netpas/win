// This file was automatically generated by https://github.com/kbinani/win/blob/generator/internal/cmd/gen/gen.go
// go run internal/cmd/gen/gen.go

// +build windows

package win

import (
	"unsafe"
)

var (
	// Library
	libcomdlg32 uintptr

	// Functions
	chooseColor          uintptr
	chooseFont           uintptr
	commDlgExtendedError uintptr
	findText             uintptr
	getFileTitle         uintptr
	getOpenFileName      uintptr
	getSaveFileName      uintptr
	pageSetupDlg         uintptr
	printDlgEx           uintptr
	printDlg             uintptr
	replaceText          uintptr
)

func init() {
	// Library
	libcomdlg32 = doLoadLibrary("comdlg32.dll")

	// Functions
	chooseColor = doGetProcAddress(libcomdlg32, "ChooseColorW")
	chooseFont = doGetProcAddress(libcomdlg32, "ChooseFontW")
	commDlgExtendedError = doGetProcAddress(libcomdlg32, "CommDlgExtendedError")
	findText = doGetProcAddress(libcomdlg32, "FindTextW")
	getFileTitle = doGetProcAddress(libcomdlg32, "GetFileTitleW")
	getOpenFileName = doGetProcAddress(libcomdlg32, "GetOpenFileNameW")
	getSaveFileName = doGetProcAddress(libcomdlg32, "GetSaveFileNameW")
	pageSetupDlg = doGetProcAddress(libcomdlg32, "PageSetupDlgW")
	printDlgEx = doGetProcAddress(libcomdlg32, "PrintDlgExW")
	printDlg = doGetProcAddress(libcomdlg32, "PrintDlgW")
	replaceText = doGetProcAddress(libcomdlg32, "ReplaceTextW")
}

func ChooseColor(unnamed0 *CHOOSECOLOR) bool {
	ret1 := syscall3(chooseColor, 1,
		uintptr(unsafe.Pointer(unnamed0)),
		0,
		0)
	return ret1 != 0
}

func ChooseFont(unnamed0 LPCHOOSEFONT) bool {
	ret1 := syscall3(chooseFont, 1,
		uintptr(unsafe.Pointer(unnamed0)),
		0,
		0)
	return ret1 != 0
}

func CommDlgExtendedError() DWORD {
	ret1 := syscall3(commDlgExtendedError, 0,
		0,
		0,
		0)
	return DWORD(ret1)
}

func FindText(unnamed0 LPFINDREPLACE) HWND {
	ret1 := syscall3(findText, 1,
		uintptr(unsafe.Pointer(unnamed0)),
		0,
		0)
	return HWND(ret1)
}

func GetFileTitle(unnamed0 string, unnamed1 LPWSTR, unnamed2 WORD) int16 {
	unnamed0Str := unicode16FromString(unnamed0)
	ret1 := syscall3(getFileTitle, 3,
		uintptr(unsafe.Pointer(&unnamed0Str[0])),
		uintptr(unsafe.Pointer(unnamed1)),
		uintptr(unnamed2))
	return int16(ret1)
}

func GetOpenFileName(unnamed0 LPOPENFILENAME) bool {
	ret1 := syscall3(getOpenFileName, 1,
		uintptr(unsafe.Pointer(unnamed0)),
		0,
		0)
	return ret1 != 0
}

func GetSaveFileName(unnamed0 LPOPENFILENAME) bool {
	ret1 := syscall3(getSaveFileName, 1,
		uintptr(unsafe.Pointer(unnamed0)),
		0,
		0)
	return ret1 != 0
}

func PageSetupDlg(unnamed0 LPPAGESETUPDLG) bool {
	ret1 := syscall3(pageSetupDlg, 1,
		uintptr(unsafe.Pointer(unnamed0)),
		0,
		0)
	return ret1 != 0
}

func PrintDlgEx(unnamed0 LPPRINTDLGEX) HRESULT {
	ret1 := syscall3(printDlgEx, 1,
		uintptr(unsafe.Pointer(unnamed0)),
		0,
		0)
	return HRESULT(ret1)
}

func PrintDlg(unnamed0 LPPRINTDLG) bool {
	ret1 := syscall3(printDlg, 1,
		uintptr(unsafe.Pointer(unnamed0)),
		0,
		0)
	return ret1 != 0
}

func ReplaceText(unnamed0 LPFINDREPLACE) HWND {
	ret1 := syscall3(replaceText, 1,
		uintptr(unsafe.Pointer(unnamed0)),
		0,
		0)
	return HWND(ret1)
}
