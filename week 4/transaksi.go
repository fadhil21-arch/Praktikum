package main

import "fmt"

type Saham struct {
	hargaBeli          float64
	hargaJual          float64
	jumlahSaham        float64
	totalInvestasiAwal float64
	totalPenjualan     float64
	keuntunganKotor    float64
	biayaTransaksi     float64
	pajakKeuntungan    float64
	keuntunganBersih   float64
}

func main() {
	var s Saham

	fmt.Println("Informasi Investasi Saham:")
	fmt.Print("Harga Beli: Rp ")
	fmt.Scan(&s.hargaBeli)
	fmt.Print("Harga Jual: Rp ")
	fmt.Scan(&s.hargaJual)
	fmt.Print("Jumlah Saham: ")
	fmt.Scan(&s.jumlahSaham)

	s.totalInvestasiAwal = s.hargaBeli * s.jumlahSaham
	s.totalPenjualan = s.hargaJual * s.jumlahSaham
	s.keuntunganKotor = s.totalPenjualan - s.totalInvestasiAwal
	s.biayaTransaksi = 0.002 * s.totalPenjualan

	if s.keuntunganKotor > 0 {
		s.pajakKeuntungan = 0.10 * s.keuntunganKotor
	} else {
		s.pajakKeuntungan = 0.0
	}

	s.keuntunganBersih = s.keuntunganKotor - s.biayaTransaksi - s.pajakKeuntungan

	fmt.Printf("Total Investasi Awal: Rp %.2f\n", s.totalInvestasiAwal)
	fmt.Printf("Total Penjualan: Rp %.2f\n", s.totalPenjualan)
	fmt.Printf("Keuntungan Kotor: Rp %.2f\n", s.keuntunganKotor)
	fmt.Printf("Biaya Transaksi: Rp %.2f\n", s.biayaTransaksi)
	fmt.Printf("Pajak Keuntungan: Rp %.2f\n", s.pajakKeuntungan)
	fmt.Printf("Keuntungan Bersih: Rp %.2f\n", s.keuntunganBersih)
}