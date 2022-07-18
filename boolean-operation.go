package main

import "fmt"

func main() {

	var (
		ujian = 90
		absensi = 80
	)

	var lulusUjian = ujian > 80
	var lulusAbsensi = absensi > 80

	var lulus = lulusUjian || lulusAbsensi
	fmt.Println(lulus)

}