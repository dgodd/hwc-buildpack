package main

import (
    "fmt"
    "syscall"
    "unsafe"
)

func abort(funcname string, err error) {
    panic(fmt.Sprintf("%s failed: %v", funcname, err))
}

var (
    hwebcore, err1 = syscall.LoadLibrary("hwebcore.dll")
    webCoreActivate, err2 = syscall.GetProcAddress(hwebcore, "WebCoreActivate")
    // webCoreShutDown, _ = syscall.GetProcAddress(hwebcore, "WebCoreShutDown")
)

func WebCoreActivate(appHostConfig, rootWebConfig, instanceName string) {
    var nargs uintptr = 3
    _, _, callErr := syscall.Syscall(uintptr(webCoreActivate),
        nargs,
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(appHostConfig))),
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(rootWebConfig))),
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(instanceName))))

    if callErr != 0 {
        abort("Call WebCoreActivate", callErr)
    }
    return
}

func main() {
    defer syscall.FreeLibrary(hwebcore)

    fmt.Print(err1)
    fmt.Print(err2)
    fmt.Print("Hello\n")

    WebCoreActivate(
	    "App.config",
	    "c:\\Users\\IEUser\\Desktop\\NoraPublished\\Web.config",
	    "daveinstance10")
}
