package fragment

import "fmt"

func Print_fragments(fragments [][]byte) {
	fmt.Printf("len(fragments): %d\n", len(fragments))
	for i := 0; i < len(fragments); i++ {
		fmt.Printf("len(fragments[%.2d]): %.2d , content:", i, len(fragments[i]))
		for j := 0; j < len(fragments[i]); j++ {
			fmt.Printf("%c", fragments[i][j])
		}
		fmt.Print("\n")
	}
}

func Breakdown_bytes(rawdata []byte, n int) [][]byte {
	var L int = len(rawdata)
	if n > L {
		n = L
	}
	var normal int = L / n
	var remainder int = L % n
	var output = make([][]byte, n)
	var piece_length = normal
	var index = 0

	for i := 0; i < n; i++ {
		if i == n-remainder {
			piece_length += 1
		}
		output[i] = make([]byte, piece_length)
		for j := 0; j < piece_length; j++ {
			output[i][j] = rawdata[index]
			index += 1
		}
	}
	return output
}
