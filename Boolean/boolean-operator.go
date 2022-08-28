package main

import "fmt"

func main() {
	var ujian = 80
	var absensi = 90

	var lulusUjian = ujian >= 80
	var lulusAbsensi = absensi >= 80
	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var lulus = lulusAbsensi && lulusUjian
	fmt.Println(lulus)

	// biar lebih cepat

	fmt.Println(ujian >= 76 && absensi >= 76)
}
