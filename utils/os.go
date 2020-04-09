package utils

import "runtime"

func GetOs() string {
	return runtime.GOOS
}
