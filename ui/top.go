package ui

// #include <windows.h>
// #include <winuser.h>
import "C"

import (
	_ "embed"
	"unsafe"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver"
)

// AlwaysOnTop sets a window to be always on top regardless of if its focused. Should be ran after [fyne.Window.Show], but before [fyne.App.Run].
func AlwaysOnTop(win fyne.Window) {
	// Fyne still has no way of making a window always-on-top, so this is the best alternative...
	win.(driver.NativeWindow).RunNative(func(context any) {
		if handle, ok := context.(driver.WindowsWindowContext); ok {
			const swpNomove = 0x0002
			var hwndTopmost = C.HWND(unsafe.Pointer(^uintptr(0))) // -1 as uintptr
			hwnd := C.HWND(unsafe.Pointer(handle.HWND))

			C.SetWindowPos(
				hwnd,
				hwndTopmost, // make the window "always on top"
				0, 0, 0, 0,  // coordinates and stuff (we dont care about these, see below)
				swpNomove, // don't move window
			)
		}
	})
}
