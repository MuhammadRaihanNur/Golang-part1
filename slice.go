package main

import "fmt"

func main() {
	names := [...]string{"Muhammad", "Raihan", "Nur", "Rizqi", "Amin", "Ganteng"}

	slice1 := names[2:5]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(days)

	dayslice2 := append(daySlice1, "Libur Baru", "Hah Libur")
	dayslice2[0] = "Ups"
	fmt.Println(dayslice2)
	fmt.Println(days)
}
