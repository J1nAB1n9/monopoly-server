package utils

import "runtime"

func GetCurrentOS() string {
	return runtime.GOOS
}

func GetSystemLineEnding(sys string ) string {
	if sys == "Windows" {
		return "\r\n"
	}

	return "\n"
}