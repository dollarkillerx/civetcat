/**
 * @Author: DollarKillerX
 * @Description: os.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:41 2019/12/28
 */
package utils

import "runtime"

func GetOs() string {
	return runtime.GOOS
}
