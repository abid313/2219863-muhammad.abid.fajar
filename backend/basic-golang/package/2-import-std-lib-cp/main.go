package main

import (
<<<<<<< HEAD
	"fmt"
	"time"
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
	// TODO: answer here
)

// Dari contoh yang telah diberikan dan eksplorasi yang dilakukan dari standard library golang, kamu dapat mencoba untuk mengimport salah satu package pada golang.
// Cobalah untuk mengimport package time dan lakukan perhitungan perbedaan hari antara dua waktu.

func CountDays(start, end time.Time) int {
	// TODO: answer here
<<<<<<< HEAD
	days := end.Sub(start).Hours()
	var y int = int(days)

	return y / 24
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}

func main() {
	start := time.Date(2022, time.March, 21, 0, 0, 0, 0, time.UTC)
	end := time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC)
	dayDifference := CountDays(start, end)
	fmt.Println(dayDifference)
}
