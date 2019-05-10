package main

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	var res byte = 0
	for i := 0; i <= 7; i++ {
		res += pc[byte(x>>(uint64(i)*8))]
	}
	return int(res)
}

func main() {

}
