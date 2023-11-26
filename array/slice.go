package main

import "fmt"

func main() {
	names := [...]string{"Ucup", "Asep", "Bambang", "Otong", "Udin", "Yanto"}

	slice1 := names[4:6]
	fmt.Println(slice1) // [Udin Yanto]

	slice2 := names[:3]
	fmt.Println(slice2) // [Ucup Asep Bambang]

	var slice3 []string = names[3:]
	fmt.Println(slice3) // [Otong Udin Yanto]

	slice4 := names[:]
	fmt.Println(slice4) // [Ucup Asep Bambang Otong Udin Yanto]

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daySlice1 := days[5:] // [Sabtu Minggu]
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru") // daysBaru := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu", "Libur Baru"}
	daySlice2[0] = "Sabtu Lama"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Ucup"
	newSlice[1] = "Otong"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Udin")

	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Budi"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
