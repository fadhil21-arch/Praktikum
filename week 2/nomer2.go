package main

import "fmt"

func main(){
	var nama, prodi, kelas string
	var nim int

	fmt.Println("Tuliskan nama mahasiswa")
	fmt.Scanln(&nama)
	fmt.Println("Tuliskan prodi mahasiswa")
	fmt.Scanln(&prodi)
	fmt.Println("Tuliskan kelas mahasiswa")
	fmt.Scanln(&kelas)
	fmt.Println("Tuliskan nim mahasiswa")
	fmt.Scanln(&nim)
	fmt.Print("Perkenalkan Saya adalah ", nama, " salah satu mahssiswa prodi ", prodi , " dari kelas ", kelas, "dengan nim ", nim)
}