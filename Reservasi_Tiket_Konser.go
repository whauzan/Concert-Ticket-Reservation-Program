//Program Reservasi Tiket Konser

package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

const (
	N = 4
	M = 4
	K = 4
	T = 3
)

type tiket struct {
	nama      string
	identitas string
	reservasi int
	umur      int
	harga     int
}
type vip [N][M]tiket
type regkiri [N][K]tiket
type regkanan [N][K]tiket

var tabVIP, vipbaru vip
var tabRegKiri, regkiribaru regkiri
var tabRegKanan, regkananbaru regkanan
var anak, dewasa, reservasi, harga, umur, x, y int
var clear map[string]func()
var nama, identitas, jawab, jenistiket string
var t tiket

func generateTicket() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(100000)
}

func kosong(x int, y int, jenistiket string) bool {
	kosong := true
	if jenistiket == "VIP" {
		if x < 0 || y < 0 || x > N || y > M || tabVIP[x][y].identitas != "" {
			kosong = false
		}
	} else if jenistiket == "REGULERKIRI" {
		if x < 0 || y < 0 || x > N || y > M || tabRegKiri[x][y].identitas != "" {
			kosong = false
		}
	} else if jenistiket == "REGULERKANAN" {
		if x < 0 || y < 0 || x > N || y > M || tabRegKanan[x][y].identitas != "" {
			kosong = false
		}
	}
	return kosong
}

func input() {
	fmt.Print("Nama :")
	fmt.Scanln(&nama)
	fmt.Print("Nomor Identitas :")
	fmt.Scanln(&identitas)
	fmt.Print("Umur :")
	fmt.Scanln(&umur)
}

func data() {

	if umur >= 0 && umur <= 5 {
		fmt.Println("Umur Belum Cukup")
	}

	if jenistiket == "VIP" {
		harga = 100000
	} else if jenistiket == "REGULERKIRI" || jenistiket == "REGULERKANAN" {
		harga = 50000
	}

	if umur >= 6 && umur <= 18 {
		harga = harga / 10 * 6
		anak = anak + 1
	} else if umur > 18 {
		dewasa = dewasa + 1
	}

	if x >= 0 {
		j := 0
		for j < x {
			harga = harga / 100 * 97
			j++
		}
	}
	if umur >= 6 {
		fmt.Println("=====================================================================")
		fmt.Println("=                                                                   =")
		fmt.Println("=                         KONSER MUSIK KLASIK                       =")
		fmt.Println("=                                                                   =")
		fmt.Println("=====================================================================")
		fmt.Println("Nama             : ", nama)
		fmt.Println("Nomor Identitas  : ", identitas)
		fmt.Println("Umur             : ", umur)
		fmt.Println("Harga            : ", harga)
		fmt.Println("Tempat Duduk     : ", x, y)
		fmt.Println("Kode Reservasi   : ", reservasi)
		fmt.Print("Ingin Cetak?(iya/tidak)")
		fmt.Scanln(&jawab)
		if jawab == "iya" {
			ubahkuy(nama, jenistiket, reservasi, harga, x, y)
		}
		if jenistiket == "VIP" {
			tabVIP[x][y].reservasi = reservasi
			tabVIP[x][y].nama = nama
			tabVIP[x][y].identitas = identitas
			tabVIP[x][y].umur = umur
			tabVIP[x][y].harga = harga
		} else if jenistiket == "REGULERKIRI" {
			tabRegKiri[x][y].reservasi = reservasi
			tabRegKiri[x][y].nama = nama
			tabRegKiri[x][y].identitas = identitas
			tabRegKiri[x][y].umur = umur
			tabRegKiri[x][y].harga = harga
		} else if jenistiket == "REGULERKANAN" {
			tabRegKanan[x][y].reservasi = reservasi
			tabRegKanan[x][y].nama = nama
			tabRegKanan[x][y].identitas = identitas
			tabRegKanan[x][y].umur = umur
			tabRegKanan[x][y].harga = harga
		}
		fmt.Println()
	}
}

func datavip() {
	fmt.Print("Kolom : ")
	fmt.Scanln(&x)
	fmt.Print("Baris : ")
	fmt.Scanln(&y)
	if x < N && y < M {
		if kosong(x, y, jenistiket) {
			input()
			data()
		} else {
			fmt.Println("Kursi Sudah Terisi")
		}
	} else {
		fmt.Println("Kursi Tidak Ada")
	}
}

func datareg() {
	fmt.Print("Baris : ")
	fmt.Scanln(&y)
	if y < K {
		if kosong(x, y, jenistiket) {
			input()
			data()
		} else {
			fmt.Println("Kursi Sudah Terisi")
		}
	} else {
		fmt.Println("Kursi Tidak Ada")
	}
}

func outvip() {
	fmt.Println("=====================================================================")
	fmt.Println("=                                                                   =")
	fmt.Println("=                         KONSER MUSIK KLASIK                       =")
	fmt.Println("=                                                                   =")
	fmt.Println("=====================================================================")
	fmt.Println("Jenis TIKET      : VIP")
	fmt.Println("Nama             : ", tabVIP[x][y].nama)
	fmt.Println("Nomor Identitas  : ", tabVIP[x][y].identitas)
	fmt.Println("Umur             : ", tabVIP[x][y].umur)
	fmt.Println("Harga            : ", tabVIP[x][y].harga)
	fmt.Println("Tempat Duduk     : ", x, y)
	fmt.Println("Kode Reservasi   : ", tabVIP[x][y].reservasi)
	fmt.Println("\n\n")
}

func outregkiri() {
	fmt.Println("=====================================================================")
	fmt.Println("=                                                                   =")
	fmt.Println("=                         KONSER MUSIK KLASIK                       =")
	fmt.Println("=                                                                   =")
	fmt.Println("=====================================================================")
	fmt.Println("Jenis TIKET     : REGULER KIRI")
	fmt.Println("Nama            : ", tabRegKiri[x][y].nama)
	fmt.Println("Nomor Identitas : ", tabRegKiri[x][y].identitas)
	fmt.Println("Umur            : ", tabRegKiri[x][y].umur)
	fmt.Println("Harga           : ", tabRegKiri[x][y].harga)
	fmt.Println("Tempat Duduk    : ", x, y)
	fmt.Println("Kode Reservasi  : ", tabRegKiri[x][y].reservasi)
	fmt.Println("\n\n")
}

func outregkanan() {
	fmt.Println("=====================================================================")
	fmt.Println("=                                                                   =")
	fmt.Println("=                         KONSER MUSIK KLASIK                       =")
	fmt.Println("=                                                                   =")
	fmt.Println("=====================================================================")
	fmt.Println("Jenis TIKET       : REGULER KANAN")
	fmt.Println("Nama              : ", tabRegKanan[x][y].nama)
	fmt.Println("Nomor Identitas   : ", tabRegKanan[x][y].identitas)
	fmt.Println("Umur              : ", tabRegKanan[x][y].umur)
	fmt.Println("Harga             : ", tabRegKanan[x][y].harga)
	fmt.Println("Tempat Duduk      : ", x, y)
	fmt.Println("Kode Reservasi    : ", tabRegKanan[x][y].reservasi)
	fmt.Println("\n\n")
}

func ubahkuy(nama, jenistiket string, reservasi, harga, x, y int) {
	f, err := os.OpenFile("Data Tiket.txt", os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf("Nama: %s", nama))
	_, err = f.WriteString(fmt.Sprintf("\nKelas: %02s", jenistiket))
	_, err = f.WriteString(fmt.Sprintf("\nKode Reservasi: %02d", reservasi))
	_, err = f.WriteString(fmt.Sprintf("\nTempat Duduk: %02d %02d", x, y))
	_, err = f.WriteString(fmt.Sprintf("\nHarga: %d\n\n", harga))
	if err != nil {
		fmt.Printf("error writing string: %v", err)
	}
}

func init() {
	clear = make(map[string]func())
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Platform anda tidak support! Aku tidak bisa membersihkan layar terminal :(")
	}
}

func main() {
	var a int
	var run bool

	run = true

	for run {
		fmt.Println("================================================================================")
		fmt.Println("=                       SELAMAT DATANG DI E-TICKET BOOKING                     =")
		fmt.Println("=    PASTIKAN ANDA BERUSIA 6 TAHUN ATAU LEBIH UNTUK DAPAT MEMESAN TIKET INI    =")
		fmt.Println("================================================================================")
		fmt.Println("\n1. RESERVASI TIKET(VIP, REGULER KIRI, REGULER KANAN)                           =")
		fmt.Println("2. CARI DENGAN KODE RESERVASI                                                  =")
		fmt.Println("3. CARI DENGAN NOMOR IDENTITAS                                                 =")
		fmt.Println("4. DATA PENONTON VIP                                                           =")
		fmt.Println("5. OKUPANSI                                                                    =")
		fmt.Println("6. KOMPOSISI PENONTON DEWASA DAN ANAK PADA AREA VIP DAN REGULER                =")
		fmt.Println("7. DATA PENONTON BERDASAR USIA                                                 =")
		fmt.Println("8. JUMLAH KURSI KOSONG                                                         =")
		fmt.Println("9. KELUAR                                                                      =")
		fmt.Println("================================================================================")
		fmt.Print("\nMASUKKAN PILIHAN : ")
		fmt.Scanln(&a)

		switch a {
		case 1:
			Reservasi_Tiket()
		case 2:
			Kode_Reservasi()
		case 3:
			Nomor_Identitas()
		case 4:
			Data_VIP()
		case 5:
			Okupansi()
		case 6:
			Komposisi_Penonton()
		case 7:
			Data_Usia()
		case 8:
			Kursi_Kosong()
		case 9:
			run = false
		default:
			main()
		}
		time.Sleep(3 * time.Second)
		CallClear()
	}
	fmt.Println("================================================================================")
	fmt.Println("=                  TERIMA KASIH TELAH MENGGUNAKAN JASA KAMI                    =")
	fmt.Println("================================================================================")
}

func Reservasi_Tiket() {
	var jumtiket int

	fmt.Print("\n\nJUMLAH TIKET YANG DIPESAN (MAKS 3) : ")
	fmt.Scanln(&jumtiket)
	for jumtiket > T {
		fmt.Println("Melebihi Batas Pemesanan")
		fmt.Print("\n\nJUMLAH TIKET YANG DIPESAN (MAKS 3) : ")
		fmt.Scanln(&jumtiket)
	}
	reservasi = generateTicket()

	i := 0

	for i < jumtiket {
		fmt.Print("\nJenis Tiket(VIP/REGULERKIRI/REGULERKANAN) : ")
		fmt.Scanln(&jenistiket)

		if jenistiket == "VIP" {
			datavip()
		} else if jenistiket == "REGULERKIRI" {
			datareg()
		} else if jenistiket == "REGULERKANAN" {
			datareg()
		}
		i++
	}
}

func Kode_Reservasi() {
	var reservasi int
	var x, y int
	var found bool

	fmt.Print("\nMasukkan Kode Reservasi : ")
	fmt.Scanln(&reservasi)

	found = false
	x = 0
	for x < N {
		y = 0

		for y < M {
			if tabVIP[x][y].reservasi == reservasi {
				outvip()
				found = true
			}
			y++
		}
		x++
	}

	x = 0
	for x < N {
		y = 0

		for y < K {
			if tabRegKiri[x][y].reservasi == reservasi {
				outregkiri()
				found = true
			}
			y++
		}
		x++
	}

	x = 0
	for x < N {
		y = 0

		for y < K {
			if tabRegKanan[x][y].reservasi == reservasi {
				outregkanan()
				found = true
			}
			y++
		}
		x++
	}

	if found == false {
		fmt.Println("Kursi Kosong")
	}
}

func Nomor_Identitas() {

	fmt.Print("\nMasukkan Nomer Identitas : ")
	fmt.Scanln(&identitas)

	x = 0
	for x < N {
		y = 0

		for y < M {
			if tabVIP[x][y].identitas == identitas {
				outvip()
			} else {
				fmt.Println("Kursi Kosong")
			}
			y++
		}
		x++
	}

	x = 0
	for x < N {
		y = 0

		for y < K {
			if tabRegKiri[x][y].identitas == identitas {
				outregkiri()
			} else {
				fmt.Println("Kursi Kosong")
			}
			y++
		}
		x++
	}
	x = 0
	for x < N {
		y = 0

		for y < K {
			if tabRegKanan[x][y].identitas == identitas {
				outregkanan()
			} else {
				fmt.Println("Kursi Kosong")
			}
			y++
		}
		x++
	}
}

func Data_VIP() {

	fmt.Print("\nMasukkan Kolom Kursi : ")
	fmt.Scanln(&x)
	fmt.Print("Masukkan Baris Kursi : ")
	fmt.Scanln(&y)

	if tabVIP[x][y].umur != 0 {
		outvip()
	} else {
		fmt.Println("Kursi Kosong")
	}
}

func Okupansi() {
	var x, y int

	vip := 0
	regkiri := 0
	regkanan := 0

	x = 0
	for x < N {
		y = 0
		for y < M {
			if tabVIP[x][y].reservasi != 0 {
				vip++
			}
			y++
		}
		x++
	}

	x = 0
	for x < N {
		y = 0
		for y < K {
			if tabRegKiri[x][y].reservasi != 0 {
				regkiri++
			}
			y++
		}
		x++
	}

	x = 0
	for x < N {
		y = 0
		for y < K {
			if tabRegKanan[x][y].reservasi != 0 {
				regkanan++
			}
			y++
		}
		x++
	}
	fmt.Println("\nKursi VIP terisi:", vip, "kursi")
	fmt.Println("Kursi Reguler Kiri terisi:", regkiri, "kursi")
	fmt.Println("Kursi Reguler Kanan terisi:", regkanan, "kursi")
	fmt.Println("\n\n")
}

func Komposisi_Penonton() {
	fmt.Println("\nJumlah Penonton Anak-Anak sebanyak", anak, "orang")
	fmt.Println("Jumlah Penonton Dewasa sebanyak", dewasa, "orang")
	fmt.Println("\n\n")
}

func Data_Usia() {
	var x, y, now1, now2, a, b int

	vipbaru = tabVIP
	regkiribaru = tabRegKiri
	regkananbaru = tabRegKanan

	x = 0
	for x < N {
		y = 0
		for y < M {
			now1 = x
			now2 = y
			a = 0
			for a < N {
				b = 0
				for b < M {
					if vipbaru[now1][now2].umur < vipbaru[a][b].umur {
						now1 = a
						now2 = b
					}
					temp := vipbaru[x][y]
					vipbaru[x][y] = vipbaru[now1][now2]
					vipbaru[now1][now2] = temp
					b++
				}
				a++
			}
			y++
		}
		x++
	}

	x = 0
	for x < N {
		y = 0
		for y < K {
			now1 = x
			now2 = y
			a = 0
			for a < N {
				b = 0
				for b < K {
					if regkiribaru[now1][now2].umur < regkiribaru[a][b].umur {
						now1 = a
						now2 = b
					}
					temp := regkiribaru[x][y]
					regkiribaru[x][y] = regkiribaru[now1][now2]
					regkiribaru[now1][now2] = temp
					b++
				}
				a++
			}
			y++
		}
		x++
	}

	x = 0
	for x < N {
		y = 0
		for y < K {
			now1 = x
			now2 = y
			a = 0
			for a < N {
				b = 0
				for b < K {
					if regkananbaru[now1][now2].umur < regkananbaru[a][b].umur {
						now1 = a
						now2 = b
					}
					temp := regkananbaru[x][y]
					regkananbaru[x][y] = regkananbaru[now1][now2]
					regkananbaru[now1][now2] = temp
					b++
				}
				a++
			}
			y++
		}
		x++
	}
	fmt.Println("\nArea VIP")

	x = 0
	for x < N {
		y = 0
		for y < M {
			if vipbaru[x][y].reservasi != 0 {
				fmt.Println("=====================================================================")
				fmt.Println("=                                                                   =")
				fmt.Println("=                         KONSER MUSIK KLASIK                       =")
				fmt.Println("=                                                                   =")
				fmt.Println("=====================================================================")
				fmt.Println("Jenis TIKET      : VIP")
				fmt.Println("Nama             : ", vipbaru[x][y].nama)
				fmt.Println("Nomor Identitas  : ", vipbaru[x][y].identitas)
				fmt.Println("Umur             : ", vipbaru[x][y].umur)
				fmt.Println("Harga            : ", vipbaru[x][y].harga)
				fmt.Println("Tempat Duduk     : ", x, y)
				fmt.Println("Kode Reservasi   : ", vipbaru[x][y].reservasi)
				fmt.Println("\n")
			}
			y++
		}
		x++
	}

	fmt.Println("\nArea Reguler Kiri")

	x = 0
	for x < N {
		y = 0
		for y < K {
			if regkiribaru[x][y].reservasi != 0 {
				fmt.Println("=====================================================================")
				fmt.Println("=                                                                   =")
				fmt.Println("=                         KONSER MUSIK KLASIK                       =")
				fmt.Println("=                                                                   =")
				fmt.Println("=====================================================================")
				fmt.Println("Jenis TIKET      : REGULER KIRI")
				fmt.Println("Nama             : ", regkiribaru[x][y].nama)
				fmt.Println("Nomor Identitas  : ", regkiribaru[x][y].identitas)
				fmt.Println("Umur             : ", regkiribaru[x][y].umur)
				fmt.Println("Harga            : ", regkiribaru[x][y].harga)
				fmt.Println("Tempat Duduk     : ", x, y)
				fmt.Println("Kode Reservasi   : ", regkiribaru[x][y].reservasi)
				fmt.Println("\n")
			}
			y++
		}
		x++
	}

	fmt.Println("\nArea Reguler Kanan")

	x = 0
	for x < N {
		y = 0
		for y < K {
			if regkananbaru[x][y].reservasi != 0 {
				fmt.Println("=====================================================================")
				fmt.Println("=                                                                   =")
				fmt.Println("=                         KONSER MUSIK KLASIK                       =")
				fmt.Println("=                                                                   =")
				fmt.Println("=====================================================================")
				fmt.Println("Jenis TIKET      : REGULER KANAN")
				fmt.Println("Nama             : ", regkananbaru[x][y].nama)
				fmt.Println("Nomor Identitas  : ", regkananbaru[x][y].identitas)
				fmt.Println("Umur             : ", regkananbaru[x][y].umur)
				fmt.Println("Harga            : ", regkananbaru[x][y].harga)
				fmt.Println("Tempat Duduk     : ", x, y)
				fmt.Println("Kode Reservasi   : ", regkananbaru[x][y].reservasi)
				fmt.Println("\n")
			}
			y++
		}
		x++
	}
}

func Kursi_Kosong() {
	var x, y int

	vip := 0
	regkiri := 0
	regkanan := 0

	x = 0
	for x < N {
		y = 0
		for y < M {
			if tabVIP[x][y].reservasi == 0 {
				vip++
			}
			y++
		}
		x++
	}

	x = 0
	for x < N {
		y = 0
		for y < K {
			if tabRegKiri[x][y].reservasi == 0 {
				regkiri++
			}
			y++
		}
		x++
	}

	x = 0
	for x < N {
		y = 0
		for y < K {
			if tabRegKanan[x][y].reservasi == 0 {
				regkanan++
			}
			y++
		}
		x++
	}
	fmt.Println("\nJumlah Kursi VIP Yang Kosong:", vip, "Kursi")
	fmt.Println("Jumlah Kursi Reguler Kiri Yang Kosong:", regkiri, "Kursi")
	fmt.Println("Jumlah Kursi Reguler Kanan Yang Kosong:", regkanan, "Kursi")
	fmt.Println("\n\n")
}
