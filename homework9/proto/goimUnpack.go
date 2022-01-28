package proto

import "fmt"

func Unpack(buffer []byte, readerChannel chan []byte) []byte {
	length := len(buffer)
	var i int
	for i = 0; i < length; i++ {
		if length < i+PackageLength {
			break
		}
		packageLengthInt := BytesToInt32(buffer[:PackageLength])
		if length < i+packageLengthInt {
			break
		}
		fmt.Printf("length: %d buffer %d, packageLength %d, i:%d\n", length, len(buffer), packageLengthInt, i)

		data := buffer[i:packageLengthInt]
		readerChannel <- data
		i += packageLengthInt - 1
		if i >= packageLengthInt {
			break
		}
	}

	if i == length {
		return make([]byte, 0)
	}
	return buffer[i:]
}
