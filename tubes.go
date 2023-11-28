package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

var (
	facilitate       daftafasilitas
	pemakai          daftarU
	tour             daftarWisata
	currentUserLogin user
)

type destinasi struct { // tipe data bentukan untuk daftar wisata atau array yang menampilkan tempat wisata, biaya, jarak ke outket dan kategori wisata nya
	nama, kategori     string
	jarakoutlet, biaya float64
}

type daftarWisata struct { // array tujuan wisata, biaya, kategori dan jarak ke outlet
	wisata     [10]destinasi
	jumlahTour int
}

type fasilitas struct { // tipe data bentukan untuk daftar fasilitas atau array yang menampilkan detail fasilitas atau wahana yang tersedia di tujuan wisata yang tersedia
	tujuan, destinasi, penginapan, transpot string
}

type daftafasilitas struct { // array fasilitas / wahana
	detail           [10]fasilitas
	jumlahFacilitate int
}

type user struct { // tipe data bentukan untuk daftarU atau array pengguna
	nama, password, nomorT string
	role                   string
	usia                   int
}

type daftarU struct { // array pengguna
	pengguna [10]user
	jumlahP  int
}

func clearScreen() { // function menghapus layar output agar lebih rapi
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func registrasi() { // function registrasi pengguna
	fmt.Print("Username             : ")
	fmt.Scan(&pemakai.pengguna[pemakai.jumlahP].nama)
	fmt.Print("Usia                 : ")
	fmt.Scan(&pemakai.pengguna[pemakai.jumlahP].usia)
	fmt.Print("Nomor Telepon Aktif  : ")
	fmt.Scan(&pemakai.pengguna[pemakai.jumlahP].nomorT)
	fmt.Print("Password             : ")
	fmt.Scan(&pemakai.pengguna[pemakai.jumlahP].password)

	pemakai.pengguna[pemakai.jumlahP].role = "Pengguna"
	pemakai.jumlahP++
	fmt.Println("                   ========   REGISTRASI BERHASIL   ========")
	clearScreen()
}

func authentication(username, password string) int { //autentikasi
	var (
		idx int = -1
	)
	for i := 0; i < pemakai.jumlahP; i++ {
		if pemakai.pengguna[i].nama == username && pemakai.pengguna[i].password == password {
			idx = i
		}
	}
	return idx
}

func inputOption(option *int, maxOpt int) { // pembatasan opsi input
	fmt.Print("Pilih Opsi: ")
	fmt.Scan(&*option)

	for *option > maxOpt || *option < 0 {
		fmt.Println("------- PILIHAN TIDAK TERSEDIA, SILAHKAN PILIH ULANG ! -------")
		fmt.Print("Pilih Opsi : ")
		fmt.Scan(&*option)
	}
}

func login() user { // function untuk fitur login
	var (
		pilih              int
		username, password string
	)
	fmt.Print("Username :")
	fmt.Scan(&username)
	fmt.Print("Password :")
	fmt.Scan(&password)

	var masuk int = authentication(username, password)
	for masuk == -1 {
		fmt.Println("---------------------------------------------------------------------------------------")
		fmt.Println("                               USER TIDAK DITEMUKAN !   ")
		fmt.Println("---------------------------------------------------------------------------------------")
		fmt.Println("                                  SILAHKAN LAKUKAN :")
		fmt.Println("                      0.Login Ulang      ATAU     1.Registrasi")
		fmt.Println("---------------------------------------------------------------------------------------")
		inputOption(&pilih, 2)
		if pilih == 0 {
			fmt.Println("---------------------------------------------------------------------------------------")
			fmt.Println("                           ========   LOGIN   ========")

			var reInputUsername, reInputPassword string
			fmt.Print("Username :")
			fmt.Scan(&reInputUsername)
			fmt.Print("Password :")
			fmt.Scan(&reInputPassword)
			masuk = authentication(reInputUsername, reInputPassword)
		} else if pilih == 1 {
			fmt.Println("---------------------------------------------------------------------------------------")
			fmt.Println("                   ========   SILAHKAN REGISTRASI   ========")
			registrasi()
			var reInputUsername, reInputPassword string
			fmt.Println("---------------------------------------------------------------------------------------")
			fmt.Println("                   ========   SILAHKAN LOGIN ULANG   ========")
			fmt.Println("---------------------------------------------------------------------------------------")
			fmt.Print("Username :")
			fmt.Scan(&reInputUsername)
			fmt.Print("Password :")
			fmt.Scan(&reInputPassword)
			masuk = authentication(reInputUsername, reInputPassword)
		}
	}

	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("                     ========   LOGIN BERHASIL    ========")
	fmt.Println("---------------------------------------------------------------------------------------")

	return pemakai.pengguna[masuk]
}

func tampilanUser() { //tampilan / menu pengguna
	var (
		tampilanPOption int
	)

	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("           ===== Halo", currentUserLogin.nama, "! |", "Selamat Datang di TAMPILAN PENGGUNA    =====")
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("1. Profil")
	fmt.Println("2. Tujuan Wisata")
	fmt.Println("3. log out")
	fmt.Println("---------------------------------------------------------------------------------------")
	inputOption(&tampilanPOption, 3)
	if tampilanPOption == 1 {
		clearScreen()
		profil()
		fmt.Println()
	} else if tampilanPOption == 2 {
		clearScreen()
		tujuanWisata()
	} else if tampilanPOption == 3 {
		clearScreen()
		logout()
	}

}

func profil() { // fitur menampilkan profil pengguna
	var (
		profilOption int
	)
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("                                    PROFIL")
	fmt.Println("---------------------------------------------------------------------------------------")
	if currentUserLogin.role == "admin" {
		fmt.Println("Username       :", pemakai.pengguna[0].nama)
		fmt.Println("Usia           :", pemakai.pengguna[0].usia)
		fmt.Println("Nomor Telepon  :", pemakai.pengguna[0].nomorT)
		fmt.Println("Sebagai        :", pemakai.pengguna[0].role)
	} else {
		for i := 1; i < pemakai.jumlahP; i++ {
			fmt.Println("Username       :", currentUserLogin.nama)
			fmt.Println("Usia           :", currentUserLogin.usia)
			fmt.Println("Nomor Telepon  :", currentUserLogin.nomorT)
			fmt.Println("Sebagai        :", currentUserLogin.role)
		}
	}

	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("  ================== ! Masukkan 0 untuk kembali ke Tampilan Awal ! ==================")
	inputOption(&profilOption, 1)
	if profilOption == 0 {
		if currentUserLogin.role == "admin" {
			clearScreen()
			tampilanAdmin()
			fmt.Println("---------------------------------------------------------------------------------------")
		} else {
			clearScreen()
			tampilanUser()
		}
	}
}

func tampilDaftarWisata() { // function untuk menampilkan daftar wisata
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("                           Daftar Tujuan Wisata : ")
	fmt.Println("       Tujuan Wisata |   Kategori  | Jarak Ke Outlet (km) | Biaya (Rp)")
	fmt.Println("---------------------------------------------------------------------------------------")
	for i := 0; i < tour.jumlahTour; i++ {
		fmt.Printf("%-20s | %-20s |  %-20.2f | %-20.2f\n",
			tour.wisata[i].nama, tour.wisata[i].kategori, tour.wisata[i].jarakoutlet, tour.wisata[i].biaya)
	}
	fmt.Println("***Berminat Hubungi : 081726352415(Cecep)***")
	fmt.Println("---------------------------------------------------------------------------------------")
}

func tampilfasilitas() { // function untuk menampilkan daftar fasilitas
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("                       Daftar Fasilitas per Tujuan Wisata : ")
	fmt.Println("               Tujuan | Destinasi Wisata | Penginapan | Transportasi")
	fmt.Println("---------------------------------------------------------------------------------------")
	for i := 0; i < facilitate.jumlahFacilitate; i++ {
		fmt.Printf("%-20s | %-20s |  %-20s | %-20s\n",
			facilitate.detail[i].tujuan, facilitate.detail[i].destinasi, facilitate.detail[i].penginapan, facilitate.detail[i].transpot)
	}
	fmt.Println("***Berminat Hubungi : 081726352415(Cecep)***")
}

func tujuanWisata() { // menu tujuan wisata
	var (
		opsitujuan int
	)

	tampilDaftarWisata()
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("   Pilih menu : ")
	fmt.Println("1. Cari Tujuan Wisata")
	fmt.Println("2. Urutkan Tujuan wisata")
	fmt.Println("3. Lihat Daftar Fasilitas per Tujuan Wisata")
	fmt.Println("0. kembali ke tampilan pengguna ")
	fmt.Println("---------------------------------------------------------------------------------------")
	inputOption(&opsitujuan, 4)
	if opsitujuan == 1 {
		clearScreen()
		caridaftarW()
	} else if opsitujuan == 2 {
		clearScreen()
		urutTujuanWisata()
	} else if opsitujuan == 0 {
		clearScreen()
		tampilanUser()
	} else if opsitujuan == 3 {
		clearScreen()
		daftarfasilitas()
	}
}

func daftarfasilitas() { // menu daftar fasilitas
	var (
		opsidaftarF int
	)
	fmt.Println("---------------------------------------------------------------------------------------")
	tampilfasilitas()
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("   Pilih menu : ")
	fmt.Println("1. Urutkan Daftar Fasilitas")
	fmt.Println("2. Cari Daftar Fasilitas")
	fmt.Println("3. Kembali Lihat Daftar Wisata")
	fmt.Println("0. kembali ke tampilan pengguna ")
	fmt.Println("---------------------------------------------------------------------------------------")
	inputOption(&opsidaftarF, 4)
	fmt.Println("---------------------------------------------------------------------------------------")
	if opsidaftarF == 3 {
		clearScreen()
		tujuanWisata()
	} else if opsidaftarF == 0 {
		clearScreen()
		tampilanUser()
	} else if opsidaftarF == 1 {
		clearScreen()
		urutdaftarFasilitas()
	} else if opsidaftarF == 2 {
		clearScreen()
		caridaftarFasilitas()
	}
}

func urutdaftarFasilitas() { // mengurutkan daftar fasilitas berdasarkan tujuan/destinasi/penginapan/transportasi
	var opsiurutf int
	fmt.Println("Urutkan dari yang terkecil Berdasarkan : ")
	fmt.Println("1. Tujuan | 2. Destinasi  | 3. Penginapan	| 4. Transportasi")
	fmt.Println("---------------------------------------------------------------------------------------")
	inputOption(&opsiurutf, 4)
	if opsiurutf == 1 {
		clearScreen()
		sortfbytujuan()
		tampilfasilitas()
		fmt.Println("---------------------------------------------------------------------------------------")
	} else if opsiurutf == 2 {
		clearScreen()
		sortfbydestinasi()
		tampilfasilitas()
		fmt.Println("---------------------------------------------------------------------------------------")
	} else if opsiurutf == 3 {
		clearScreen()
		sortfbypenginapan()
		tampilfasilitas()
		fmt.Println("---------------------------------------------------------------------------------------")
	} else if opsiurutf == 4 {
		clearScreen()
		sortfbytranspot()
		tampilfasilitas()
		fmt.Println("---------------------------------------------------------------------------------------")
	}
	fmt.Println("   Pilih menu : ")
	fmt.Println("1. Cari Daftar fasilitas")
	fmt.Println("2. Lihat Tujuan Wisata")
	fmt.Println("3. Lihat Daftar Fasitilas per Tujuan Wisata")
	fmt.Println("0. kembali ke tampilan pengguna ")
	fmt.Println("---------------------------------------------------------------------------------------")
	inputOption(&opsiurutf, 4)
	if opsiurutf == 1 {
		clearScreen()
		caridaftarFasilitas()
	} else if opsiurutf == 2 {
		clearScreen()
		tujuanWisata()
	} else if opsiurutf == 0 {
		clearScreen()
		tampilanUser()
	} else if opsiurutf == 3 {
		clearScreen()
		daftarfasilitas()
	}
}

func caridaftarFasilitas() { // mencari daftar fasilitas berdasarkan tujuan/destinasi/penginapan/transportasi
	var (
		keyf      string
		opsicarif int
	)
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println(" ! Untuk Mencari, Masukkan Tujuan Wisata atau Destinasi atau Penginapan atau Transportasi !")
	fmt.Scan(&keyf)
	clearScreen()
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("                       Daftar Fasilitas per Tujuan Wisata : ")
	fmt.Println("               Tujuan | Destinasi Wisata | Penginapan | Transportasi")
	fmt.Println("---------------------------------------------------------------------------------------")
	keywordFasilitas(keyf)
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("   Pilih menu : ")
	fmt.Println("1. Urutkan Daftar Fasilitas")
	fmt.Println("2. Lihat Tujuan Wisata")
	fmt.Println("3. Lihat Daftar Fasitilas per Tujuan Wisata")
	fmt.Println("0. kembali ke tampilan pengguna ")
	fmt.Println("---------------------------------------------------------------------------------------")
	inputOption(&opsicarif, 4)
	if opsicarif == 1 {
		clearScreen()
		urutdaftarFasilitas()
	} else if opsicarif == 2 {
		clearScreen()
		tujuanWisata()
	} else if opsicarif == 0 {
		clearScreen()
		tampilanUser()
	} else if opsicarif == 3 {
		clearScreen()
		daftarfasilitas()
	}
}

func urutTujuanWisata() { // mengurutkan tujuan wisata berdasarkan nama tujuan/kategori/jarak/biaya
	var opsiurutW int
	fmt.Println("Urutkan dari yang terkecil Berdasarkan : ")
	fmt.Println("1. Nama Tujuan | 2. Kategori Wisata  | 3. Jarak Daftar Wisata Dari Outlet (km)| 4. Biaya (Rp)")
	fmt.Println("---------------------------------------------------------------------------------------")
	inputOption(&opsiurutW, 4)
	if opsiurutW == 1 {
		clearScreen()
		sortbynama()
		tampilDaftarWisata()
		fmt.Println("---------------------------------------------------------------------------------------")
	} else if opsiurutW == 2 {
		clearScreen()
		sortbykategori()
		tampilDaftarWisata()
		fmt.Println("---------------------------------------------------------------------------------------")
	} else if opsiurutW == 3 {
		clearScreen()
		sortbyjarak()
		tampilDaftarWisata()
		fmt.Println("---------------------------------------------------------------------------------------")

	} else if opsiurutW == 4 {
		clearScreen()
		sortbybiaya()
		tampilDaftarWisata()
		fmt.Println("---------------------------------------------------------------------------------------")
	}
	fmt.Println("   Pilih menu : ")
	fmt.Println("1. Cari Tujuan Wisata")
	fmt.Println("2. Lihat Tujuan Wisata")
	fmt.Println("3. Lihat Daftar Fasitilas per Tujuan Wisata")
	fmt.Println("0. kembali ke tampilan pengguna ")
	fmt.Println("---------------------------------------------------------------------------------------")
	inputOption(&opsiurutW, 4)
	if opsiurutW == 1 {
		clearScreen()
		caridaftarW()
	} else if opsiurutW == 2 {
		clearScreen()
		tujuanWisata()
	} else if opsiurutW == 0 {
		clearScreen()
		tampilanUser()
	} else if opsiurutW == 3 {
		clearScreen()
		daftarfasilitas()
	}
}

func caridaftarW() { // mencari tujuan wisata berdasarkan nama tujuan/kategori/jarak/biaya
	var (
		opsicaridaftarW int
		keyw            string
	)
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("     ! Untuk Mencari, Masukkan Tujuan Wisata atau Kategori atau Jarak atau Biaya !")
	fmt.Scan(&keyw)
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("                           Daftar Tujuan Wisata : ")
	fmt.Println("       Tujuan Wisata |   Kategori  | Jarak Ke Outlet (km) | Biaya (Rp)")
	fmt.Println("---------------------------------------------------------------------------------------")
	keywordTujuanWisata(keyw)
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("   Pilih menu : ")
	fmt.Println("1. Urutkan Tujuan Wisata")
	fmt.Println("2. Lihat Tujuan Wisata")
	fmt.Println("3. Lihat Daftar Fasitilas per Tujuan Wisata")
	fmt.Println("0. kembali ke tampilan pengguna ")
	fmt.Println("---------------------------------------------------------------------------------------")
	inputOption(&opsicaridaftarW, 4)
	if opsicaridaftarW == 1 {
		clearScreen()
		urutTujuanWisata()
	} else if opsicaridaftarW == 2 {
		clearScreen()
		tujuanWisata()
	} else if opsicaridaftarW == 0 {
		clearScreen()
		tampilanUser()
	} else if opsicaridaftarW == 3 {
		clearScreen()
		daftarfasilitas()
	}
}

func keywordTujuanWisata(keyw string) { // func mencari tujuan wisata berdasarkan nama tujuan/kategori/jarak/biaya
	angka, _ := strconv.ParseFloat(keyw, 64)
	for i := 0; i < tour.jumlahTour; i++ {
		if tour.wisata[i].nama == keyw || tour.wisata[i].kategori == keyw ||
			tour.wisata[i].jarakoutlet == angka || tour.wisata[i].biaya == angka {
			fmt.Printf("%-20s | %-20s |  %-20.2f | %-20.2f\n",
				tour.wisata[i].nama, tour.wisata[i].kategori, tour.wisata[i].jarakoutlet, tour.wisata[i].biaya)
		}
	}
}

func keywordFasilitas(keyf string) { // mencari daftar fasilitas berdasarkan tujuan/destinasi/penginapan/transportasi
	for i := 0; i < facilitate.jumlahFacilitate; i++ {
		if facilitate.detail[i].tujuan == keyf || facilitate.detail[i].destinasi == keyf ||
			facilitate.detail[i].penginapan == keyf || facilitate.detail[i].transpot == keyf {
			fmt.Printf("%-20s | %-20s |  %-20s | %-20s\n",
				facilitate.detail[i].tujuan, facilitate.detail[i].destinasi, facilitate.detail[i].penginapan, facilitate.detail[i].transpot)
		}
	}
}

// sorting array daftar fasilitas
func sortfbytujuan() {
	for i := 1; i < facilitate.jumlahFacilitate; i++ {
		key := facilitate.detail[i]
		j := i - 1
		for j >= 0 && facilitate.detail[j].tujuan > key.tujuan {
			facilitate.detail[j+1] = facilitate.detail[j]
			j--
		}
		facilitate.detail[j+1] = key
	}
}

func sortfbydestinasi() {
	for i := 1; i < facilitate.jumlahFacilitate; i++ {
		key := facilitate.detail[i]
		j := i - 1
		for j >= 0 && facilitate.detail[j].destinasi > key.destinasi {
			facilitate.detail[j+1] = facilitate.detail[j]
			j--
		}
		facilitate.detail[j+1] = key
	}
}

func sortfbypenginapan() {
	for i := 1; i < facilitate.jumlahFacilitate; i++ {
		key := facilitate.detail[i]
		j := i - 1
		for j >= 0 && facilitate.detail[j].penginapan > key.penginapan {
			facilitate.detail[j+1] = facilitate.detail[j]
			j--
		}
		facilitate.detail[j+1] = key
	}
}

func sortfbytranspot() {
	for i := 1; i < facilitate.jumlahFacilitate; i++ {
		key := facilitate.detail[i]
		j := i - 1
		for j >= 0 && facilitate.detail[j].transpot > key.transpot {
			facilitate.detail[j+1] = facilitate.detail[j]
			j--
		}
		facilitate.detail[j+1] = key
	}
}

// sorting array tujuan wisata
func sortbynama() {
	for i := 1; i < tour.jumlahTour; i++ {
		key := tour.wisata[i]
		j := i - 1
		for j >= 0 && tour.wisata[j].nama > key.nama {
			tour.wisata[j+1] = tour.wisata[j]
			j--
		}
		tour.wisata[j+1] = key
	}
}

func sortbykategori() {
	for i := 1; i < tour.jumlahTour; i++ {
		key := tour.wisata[i]
		j := i - 1
		for j >= 0 && tour.wisata[j].kategori > key.kategori {
			tour.wisata[j+1] = tour.wisata[j]
			j--
		}
		tour.wisata[j+1] = key
	}
}

func sortbyjarak() {
	for i := 1; i < tour.jumlahTour; i++ {
		key := tour.wisata[i]
		j := i - 1
		for j >= 0 && tour.wisata[j].jarakoutlet > key.jarakoutlet {
			tour.wisata[j+1] = tour.wisata[j]
			j--
		}
		tour.wisata[j+1] = key
	}
}

func sortbybiaya() {
	for i := 1; i < tour.jumlahTour; i++ {
		key := tour.wisata[i]
		j := i - 1
		for j >= 0 && tour.wisata[j].biaya > key.biaya {
			tour.wisata[j+1] = tour.wisata[j]
			j--
		}
		tour.wisata[j+1] = key
	}
}

func tampilanAdmin() { // menu admin
	var (
		adminoption int
	)
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("     ===== Halo admin", currentUserLogin.nama, "! |", "Selamat Datang di TAMPILAN ADMIN   =====")
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("1. Profil")
	fmt.Println("2. Menambah Daftar")
	fmt.Println("3. Menghapus Daftar")
	fmt.Println("4. Edit Daftar")
	fmt.Println("5. log out")
	fmt.Println("---------------------------------------------------------------------------------------")
	inputOption(&adminoption, 5)
	if adminoption == 1 {
		clearScreen()
		profil()
		fmt.Println()
	} else if adminoption == 2 {
		clearScreen()
		tambahAdmin()
	} else if adminoption == 3 {
		clearScreen()
		hapusAdmin()
	} else if adminoption == 4 {
		clearScreen()
		editAdmin()
	} else if adminoption == 5 {
		clearScreen()
		logout()
	}

}

func tambahAdmin() { // menu tambah daftar
	var (
		a          destinasi
		b          fasilitas
		opsitambah int
	)
	tampilDaftarWisata()
	tampilfasilitas()
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("  Tambah : 1. Daftar Wisata atau 2. Daftar Fasilitas atau 0. Kembali ke Tampilan Awal")
	inputOption(&opsitambah, 3)
	if opsitambah == 1 {
		clearScreen()
		tambahDaftarWisata(a)
	} else if opsitambah == 2 {
		clearScreen()
		tambahDaftarFasilitas(b)
	} else if opsitambah == 0 {
		clearScreen()
		tampilanAdmin()
	}
}

func tambahDaftarWisata(a destinasi) { // tambah daftar wisata
	var (
		opsitambahW int
	)
	tampilDaftarWisata()
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("               Silahkan Input Daftar Wisata Terbaru Secara Lengkap !")
	fmt.Println("           Urut mulai Nama Tujuan, Kategori, Jarak ke Outlet, hingga Biaya")
	fmt.Scan(&a.nama, &a.kategori, &a.jarakoutlet, &a.biaya)
	tour.wisata[tour.jumlahTour] = a
	tour.jumlahTour++
	clearScreen()
	tampilDaftarWisata()
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("                                   Pilih menu : ")
	fmt.Println("           1. Hapus Daftar | 2. Edit Daftar | 0. Kembali ke Tampilan Awal")
	fmt.Println("---------------------------------------------------------------------------------------")
	inputOption(&opsitambahW, 3)
	if opsitambahW == 1 {
		clearScreen()
		hapusAdmin()
	} else if opsitambahW == 2 {
		clearScreen()
		editAdmin()
	} else if opsitambahW == 0 {
		clearScreen()
		tampilanAdmin()
	}
}

func tambahDaftarFasilitas(b fasilitas) { // tambah daftar fasilitas
	var (
		opsitambahF int
	)
	tampilfasilitas()
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("               Silahkan Input Daftar Wisata Terbaru Secara Lengkap !")
	fmt.Println("           Urut mulai Nama Tujuan, Kategori, Jarak ke Outlet, hingga Biaya")
	fmt.Scan(&b.tujuan, &b.destinasi, &b.penginapan, &b.transpot)
	facilitate.detail[facilitate.jumlahFacilitate] = b
	facilitate.jumlahFacilitate++
	clearScreen()
	tampilfasilitas()
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("                                   Pilih menu : ")
	fmt.Println("           1. Hapus Daftar | 2. Edit Daftar | 0. Kembali ke Tampilan Awal")
	fmt.Println("---------------------------------------------------------------------------------------")
	inputOption(&opsitambahF, 3)
	if opsitambahF == 1 {
		clearScreen()
		hapusAdmin()
	} else if opsitambahF == 2 {
		clearScreen()
		editAdmin()
	} else if opsitambahF == 0 {
		clearScreen()
		tampilanAdmin()
	}
}

func hapusAdmin() { // menu hapus daftar
	var (
		opsihapus      int
		tujuan, daerah string
	)
	tampilDaftarWisata()
	tampilfasilitas()
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("  Hapus : 1. Daftar Wisata | 2. Daftar Fasilitas | 0. Kembali ke Tampilan Awal")
	inputOption(&opsihapus, 3)
	if opsihapus == 1 {
		clearScreen()
		tampilDaftarWisata()
		fmt.Println("---------------------------------------------------------------------------------------")
		fmt.Print("Masukkan nomor urut daftar yang akan di hapus ! (*nomor urut dimulai dari 0) :")
		fmt.Scan(&tujuan)
		hapusDaftarWisata(tujuan)
	} else if opsihapus == 2 {
		clearScreen()
		tampilfasilitas()
		fmt.Println("---------------------------------------------------------------------------------------")
		fmt.Print("Masukkan nomor urut daftar yang akan di hapus ! (*nomor urut dimulai dari 0) :")
		fmt.Scan(&tujuan)
		hapusDaftarF(daerah)
	} else if opsihapus == 0 {
		clearScreen()
		tampilanAdmin()
	}
}

func hapusDaftarWisata(tujuan string) { // hapus daftar wisata
	var (
		opsihapusW int
	)
	for j := 0; j < tour.jumlahTour; j++ {
		if tour.wisata[j].nama == tujuan {
			for i := j; i < tour.jumlahTour; i++ {
				tour.wisata[i] = tour.wisata[i+1]
			}
		}
	}
	tour.jumlahTour--
	clearScreen()
	tampilDaftarWisata()
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("                                   Pilih menu : ")
	fmt.Println("           1. Edit Daftar | 2. Tambah Daftar  | 0. Kembali ke Tampilan Awal")
	fmt.Println("---------------------------------------------------------------------------------------")
	inputOption(&opsihapusW, 3)
	if opsihapusW == 1 {
		clearScreen()
		editAdmin()
	} else if opsihapusW == 2 {
		clearScreen()
		tambahAdmin()
	} else if opsihapusW == 0 {
		clearScreen()
		tampilanAdmin()
	}
}

func hapusDaftarF(daerah string) { // hapus daftar fasilitas
	var (
		opsihapusF int
	)
	for j := 0; j < facilitate.jumlahFacilitate; j++ {
		if facilitate.detail[j].tujuan == daerah {
			for i := j; i < facilitate.jumlahFacilitate; i++ {
				facilitate.detail[i] = facilitate.detail[i+1]
			}
		}
	}
	facilitate.jumlahFacilitate--
	tampilfasilitas()
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("                                   Pilih menu : ")
	fmt.Println("           1. Edit Daftar | 2. Tambah Daftar  | 0. Kembali ke Tampilan Awal")
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Scan(&opsihapusF)
	if opsihapusF == 1 {
		clearScreen()
		editAdmin()
	} else if opsihapusF == 2 {
		clearScreen()
		tambahAdmin()
	} else {
		clearScreen()
		tampilanAdmin()
	}
}

func editAdmin() { // menu edit daftar
	var opsiedit int
	tampilDaftarWisata()
	tampilfasilitas()
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("  Edit : 1. Daftar Wisata atau 2. Daftar Fasilitas atau 0. Kembali ke Tampilan Awal")
	inputOption(&opsiedit, 3)
	if opsiedit == 1 {
		clearScreen()
		editTujuanWisata()
	} else if opsiedit == 2 {
		clearScreen()
		editfasilitas()
	} else if opsiedit == 0 {
		clearScreen()
		tampilanAdmin()
	}
}

func editTujuanWisata() { //  edit daftar wisata
	var (
		a              destinasi
		idx, opsieditW int
	)
	tampilDaftarWisata()
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Print("Masukkan nomor urut daftar yang akan di edit ! (*nomor urut dimulai dari 0) :")
	fmt.Scan(&idx)
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("               Silahkan Input Daftar Wisata Terbaru Secara Lengkap !")
	fmt.Println("           Urut mulai Nama Tujuan, Kategori, Jarak ke Outlet, hingga Biaya")
	fmt.Scan(&a.nama, &a.kategori, &a.jarakoutlet, &a.biaya)
	tour.wisata[idx] = a
	clearScreen()
	tampilDaftarWisata()
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("                                   Pilih menu : ")
	fmt.Println("           1. Hapus Daftar | 2. Tambah Daftar  | 0. Kembali ke Tampilan Awal")
	fmt.Println("---------------------------------------------------------------------------------------")
	inputOption(&opsieditW, 3)
	if opsieditW == 1 {
		clearScreen()
		hapusAdmin()
	} else if opsieditW == 2 {
		clearScreen()
		tambahAdmin()
	} else if opsieditW == 0 {
		clearScreen()
		tampilanAdmin()
	}
}

func editfasilitas() { // edit daftar fasilitas
	var (
		opsieditF, idx int
		b              fasilitas
	)
	tampilfasilitas()
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Print("Masukkan nomor urut daftar yang akan di edit ! (*nomor urut dimulai dari 0) :")
	fmt.Scan(&idx)
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("               Silahkan Input Daftar Wisata Terbaru Secara Lengkap !")
	fmt.Println("           Urut mulai Nama Tujuan, Kategori, Jarak ke Outlet, hingga Biaya")
	fmt.Scan(&b.tujuan, &b.destinasi, &b.penginapan, &b.transpot)
	facilitate.detail[idx] = b
	clearScreen()
	tampilfasilitas()
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("                                   Pilih menu : ")
	fmt.Println("           1. Hapus Daftar | 2. Tambah Daftar  | 0. Kembali ke Tampilan Awal")
	fmt.Println("---------------------------------------------------------------------------------------")
	inputOption(&opsieditF, 3)
	if opsieditF == 1 {
		clearScreen()
		hapusAdmin()
	} else if opsieditF == 2 {
		clearScreen()
		tambahAdmin()
	} else if opsieditF == 0 {
		clearScreen()
		tampilanAdmin()
	}
}

func keluarVersilogout() { // function untuk fitur keluar dan login kembali setelah batal untuk keluar dari aplikasi
	var (
		opsikeluarVlogout int
	)
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("       ============ Apakah anda yakin ingin keluar dari aplikasi ? ============")
	fmt.Println("                                   0.YES | 1.NO")
	inputOption(&opsikeluarVlogout, 2)
	if opsikeluarVlogout == 0 {
		clearScreen()
		fmt.Println("---------------------------------------------------------------------------------------")
		fmt.Println("	=================    Terima kasih   =================")
		fmt.Println("---------------------------------------------------------------------------------------")
		os.Exit(0)
	} else {
		clearScreen()
		fmt.Println("---------------------------------------------------------------------------------------")
		fmt.Println("                           ========   LOGIN   ========")
		fmt.Println("---------------------------------------------------------------------------------------")
		currentUserLogin = login()

		if currentUserLogin.role == "admin" {
			clearScreen()
			tampilanAdmin()
			fmt.Println("---------------------------------------------------------------------------------------")
		} else {
			clearScreen()
			tampilanUser()
		}
	}
}

func logout() { // function untuk fitur keluar dan login kembali setelah batal untuk keluar dari aplikasi
	var (
		opsikeluar int
	)
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("                       ========   logout BERHASIL !   ========")
	fmt.Println("                               0. Login | 1. keluar")
	inputOption(&opsikeluar, 2)
	fmt.Println("---------------------------------------------------------------------------------------")
	if opsikeluar == 0 {
		clearScreen()
		fmt.Println("---------------------------------------------------------------------------------------")
		fmt.Println("                           ========   LOGIN   ========")
		fmt.Println("---------------------------------------------------------------------------------------")
		currentUserLogin = login()
		if currentUserLogin.role == "admin" {
			clearScreen()
			tampilanAdmin()
			fmt.Println("---------------------------------------------------------------------------------------")
		} else {
			clearScreen()
			tampilanUser()
		}

	} else {
		clearScreen()
		keluarVersilogout()
	}

}

func main() {
	facilitate = daftafasilitas{
		detail: [10]fasilitas{
			{"RajaAmpat", "pantai", "hotel*3", "bus"},
			{"Bunaken", "selam", "hotel*4", "bus"},
			{"PulauKomodo", "TamanNasional", "hotel*2", "pesawat"},
			{"Baluran", "TamanNasional", "hotel*3", "pesawat"},
			{"GunungBromo", "gunung", "hotel*3", "kereta"},
		},
		jumlahFacilitate: 5,
	}

	tour = daftarWisata{
		wisata: [10]destinasi{
			{"RajaAmpat", "keluarga", 10.5, 100000},
			{"Bunaken", "keluarga", 8.0, 200000},
			{"PulauKomodo", "pasangan", 9.5, 300000},
			{"Baluran", "pasangan", 5.0, 50000},
			{"GunungBromo", "keluarga", 8.5, 200000},
		},
		jumlahTour: 5,
	}

	pemakai.pengguna[pemakai.jumlahP].nama = "YUSTINUS"
	pemakai.pengguna[pemakai.jumlahP].usia = 25
	pemakai.pengguna[pemakai.jumlahP].nomorT = "08xxxxxxxxx"
	pemakai.pengguna[pemakai.jumlahP].password = "1004"
	pemakai.pengguna[pemakai.jumlahP].role = "admin"
	pemakai.jumlahP++

	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("                           TUGAS BESAR ALGORITMA PEMROGRAMAN           ")
	fmt.Println("                                   Aplikasi Pariwisata             ")
	fmt.Println("                       Dibuat Oleh : YustinusDwiA_1301223129_IF-46-10  ")
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("---------------------------------------------------------------------------------------")
	currentUserLogin = login()
	clearScreen()
	if currentUserLogin.role == "admin" {
		clearScreen()
		tampilanAdmin()
		fmt.Println("---------------------------------------------------------------------------------------")
		fmt.Println("                               Halo! Admin", pemakai.pengguna[0].nama)
	} else {
		clearScreen()
		tampilanUser()
	}
}
