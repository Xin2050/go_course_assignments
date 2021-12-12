package fmt

import (
	"encoding/hex"
	"fmt"
	"testing"
)

// 输出两位小数
func printNumWith2(float642 float64) string {
	return fmt.Sprintf("%.2f", float642)
}

func printBytes(data []byte) string {
	return hex.Dump(data)
	//return hex.EncodeToString(data)
}

func Test_fmt(t *testing.T) {
	fmt.Println(printNumWith2(1110.23421))
	fmt.Println(printNumWith2(23.23567))

	fmt.Println(printBytes([]byte("Go is an open source programming language.")))
	//output
	/*
		00000000  47 6f 20 69 73 20 61 6e  20 6f 70 65 6e 20 73 6f  |Go is an open so|
		00000010  75 72 63 65 20 70 72 6f  67 72 61 6d 6d 69 6e 67  |urce programming|
		00000020  20 6c 61 6e 67 75 61 67  65 2e                    | language.|
	*/
}
