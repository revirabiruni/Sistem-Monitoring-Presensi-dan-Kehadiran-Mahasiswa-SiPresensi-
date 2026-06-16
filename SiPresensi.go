package main
import "fmt"

const NMAX int = 100
const NKELAS int = 10

type mahasiswa struct {
	nim   string
	nama  string
	hadir int
	izin  int
	sakit int
	alpa  int
}
type jadwal struct {
	kodeKelas  string
	mataKuliah string
	hari       string
	jam        string
	ruangan    string
}
type logKehadiran struct {
	nim        string
	namaMhs    string
	kodeKelas  string
	mataKuliah string
	pertemuan  int
	status     string
}
type dataMahasiswa [NMAX]mahasiswa
type dataJadwal [NKELAS]jadwal
type dataLog [NMAX * NKELAS]logKehadiran

func inisialisasiJadwal(J *dataJadwal) {
	J[0] = jadwal{"001", "Bahasa Inggris", "Senin", "07.30-09.30", "KTT2.26"}
	J[1] = jadwal{"002", "Organisasi dan Arsitektur Komputer", "Senin", "12.30-15.30", "KTT B-1.07"}
	J[2] = jadwal{"003", "Matriks dan Ruang Vektor", "Selasa", "06.30-09.30", "KTT2.26"}
	J[3] = jadwal{"004", "Etika dalam AI", "Selasa", "10.30-12.30", "KTT2.35"}
	J[4] = jadwal{"005", "Algoritma dan Pemrograman 2 (Praktikum)", "Rabu", "06.30-09.30", "CPROG.LAB"}
	J[5] = jadwal{"006", "Algoritma dan Pemrograman 2 (Teori)", "Rabu", "14.30-17.30", "KTT2.35"}
	J[6] = jadwal{"007", "Kalkulus Lanjut", "Kamis", "09.30-12.30", "KTT B-1.08"}
	J[7] = jadwal{"008", "Pemodelan Basis Data", "Kamis", "13.30-16.30", "KTT2.26"}
}

func inisialisasiDummy(A *dataMahasiswa, nA *int, L *dataLog, nL *int, pertemuanKe *[NKELAS]int, J dataJadwal) {
	A[0] = mahasiswa{"001", "George Russell", 0, 0, 0, 0}
	A[1] = mahasiswa{"002", "Kimi Antonelli", 0, 0, 0, 0}
	A[2] = mahasiswa{"003", "Charles Leclerc", 0, 0, 0, 0}
	A[3] = mahasiswa{"004", "Lewis Hamilton", 0, 0, 0, 0}
	A[4] = mahasiswa{"005", "Carlos Sainz", 0, 0, 0, 0}
	A[5] = mahasiswa{"006", "Alex Albon", 0, 0, 0, 0}
	A[6] = mahasiswa{"007", "Max Verstappen", 0, 0, 0, 0}
	*nA = 7
	catat := func(idxMhs int, kodeKelas string, mataKuliah string, pertemuan int, status string) {
		L[*nL] = logKehadiran{A[idxMhs].nim, A[idxMhs].nama, kodeKelas, mataKuliah, pertemuan, status}
		*nL++
		if status == "H" {
			A[idxMhs].hadir++
		} else if status == "I" {
			A[idxMhs].izin++
		} else if status == "S" {
			A[idxMhs].sakit++
		} else if status == "A" {
			A[idxMhs].alpa++
		}
	}
	catat(0, "001", J[0].mataKuliah, 1, "H")
	catat(0, "001", J[0].mataKuliah, 2, "H")
	catat(1, "001", J[0].mataKuliah, 1, "H")
	catat(1, "001", J[0].mataKuliah, 2, "S")
	catat(2, "001", J[0].mataKuliah, 1, "H")
	catat(2, "001", J[0].mataKuliah, 2, "H")
	catat(3, "001", J[0].mataKuliah, 1, "I")
	catat(3, "001", J[0].mataKuliah, 2, "H")
	catat(4, "001", J[0].mataKuliah, 1, "H")
	catat(4, "001", J[0].mataKuliah, 2, "H")
	catat(5, "001", J[0].mataKuliah, 1, "H")
	catat(5, "001", J[0].mataKuliah, 2, "I")
	catat(6, "001", J[0].mataKuliah, 1, "H")
	catat(6, "001", J[0].mataKuliah, 2, "H")
	catat(0, "002", J[1].mataKuliah, 1, "H")
	catat(0, "002", J[1].mataKuliah, 2, "H")
	catat(1, "002", J[1].mataKuliah, 1, "H")
	catat(1, "002", J[1].mataKuliah, 2, "H")
	catat(2, "002", J[1].mataKuliah, 1, "S")
	catat(2, "002", J[1].mataKuliah, 2, "H")
	catat(3, "002", J[1].mataKuliah, 1, "H")
	catat(3, "002", J[1].mataKuliah, 2, "H")
	catat(4, "002", J[1].mataKuliah, 1, "H")
	catat(4, "002", J[1].mataKuliah, 2, "A")
	catat(5, "002", J[1].mataKuliah, 1, "H")
	catat(5, "002", J[1].mataKuliah, 2, "H")
	catat(6, "002", J[1].mataKuliah, 1, "H")
	catat(6, "002", J[1].mataKuliah, 2, "H")
	catat(0, "003", J[2].mataKuliah, 1, "H")
	catat(0, "003", J[2].mataKuliah, 2, "H")
	catat(1, "003", J[2].mataKuliah, 1, "H")
	catat(1, "003", J[2].mataKuliah, 2, "H")
	catat(2, "003", J[2].mataKuliah, 1, "H")
	catat(2, "003", J[2].mataKuliah, 2, "I")
	catat(3, "003", J[2].mataKuliah, 1, "H")
	catat(3, "003", J[2].mataKuliah, 2, "H")
	catat(4, "003", J[2].mataKuliah, 1, "A")
	catat(4, "003", J[2].mataKuliah, 2, "H")
	catat(5, "003", J[2].mataKuliah, 1, "H")
	catat(5, "003", J[2].mataKuliah, 2, "H")
	catat(6, "003", J[2].mataKuliah, 1, "H")
	catat(6, "003", J[2].mataKuliah, 2, "H")
	catat(0, "004", J[3].mataKuliah, 1, "H")
	catat(0, "004", J[3].mataKuliah, 2, "H")
	catat(1, "004", J[3].mataKuliah, 1, "A")
	catat(1, "004", J[3].mataKuliah, 2, "H")
	catat(2, "004", J[3].mataKuliah, 1, "H")
	catat(2, "004", J[3].mataKuliah, 2, "H")
	catat(3, "004", J[3].mataKuliah, 1, "H")
	catat(3, "004", J[3].mataKuliah, 2, "H")
	catat(4, "004", J[3].mataKuliah, 1, "H")
	catat(4, "004", J[3].mataKuliah, 2, "H")
	catat(5, "004", J[3].mataKuliah, 1, "S")
	catat(5, "004", J[3].mataKuliah, 2, "H")
	catat(6, "004", J[3].mataKuliah, 1, "H")
	catat(6, "004", J[3].mataKuliah, 2, "H")
	var k int
	for k = 0; k < NKELAS; k++ {
		pertemuanKe[k] = 2
	}
}
func cariJadwal(J dataJadwal, kode string) int {
	var i int
	for i = 0; i < NKELAS; i++ {
		if J[i].kodeKelas == kode {
			return i
		}
	}
	return -1
}
func tampilJadwal(J dataJadwal, pertemuanKe [NKELAS]int) {
	fmt.Printf("\n%-5s  %-40s  %-8s  %-11s  %-10s  %s\n",
		"Kode", "Mata Kuliah", "Hari", "Jam", "Ruangan", "Pertemuan")
	fmt.Println("----  ----------------------------------------  --------  -----------  ----------  ---------")
	var i int
	for i = 0; i < NKELAS; i++ {
		fmt.Printf("%-5s  %-40s  %-8s  %-11s  %-10s  %d\n",
			J[i].kodeKelas,
			J[i].mataKuliah,
			J[i].hari,
			J[i].jam,
			J[i].ruangan,
			pertemuanKe[i],
		)
	}
}
func tambahMahasiswa(A *dataMahasiswa, n *int) {
	if *n >= NMAX {
		fmt.Println("Data mahasiswa sudah penuh!")
		return
	}
	var nimBaru string
	fmt.Print("NIM  : ")
	fmt.Scan(&nimBaru)

	if sequentialSearch(*A, *n, nimBaru) != -1 {
		fmt.Println("NIM sudah terdaftar!")
		return
	}
	A[*n].nim = nimBaru
	fmt.Print("Nama : ")
	fmt.Scan(&A[*n].nama)
	*n++
	fmt.Println("Mahasiswa berhasil ditambahkan.")
}
func ubahMahasiswa(A *dataMahasiswa, n int) {
	var nim string
	fmt.Print("NIM yang dicari : ")
	fmt.Scan(&nim)
	idx := sequentialSearch(*A, n, nim)
	if idx != -1 {
		idx = idx - 1
		fmt.Print("Nama baru : ")
		fmt.Scan(&A[idx].nama)
		fmt.Println("Data berhasil diubah.")
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}
func hapusMahasiswa(A *dataMahasiswa, n *int) {
	var nim string
	var i, idx int
	fmt.Print("NIM yang dihapus : ")
	fmt.Scan(&nim)
	idx = sequentialSearch(*A, *n, nim)
	if idx != -1 {
		idx = idx - 1
		for i = idx; i < *n-1; i++ {
			A[i] = A[i+1]
		}
		*n--
		fmt.Println("Data berhasil dihapus.")
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}
func tampilData(A dataMahasiswa, n int) {
	if n == 0 {
		fmt.Println("Belum ada data mahasiswa.")
		return
	}
	fmt.Printf("\n%-3s  %-6s  %-20s  %-3s  %-3s  %-3s  %-3s  %s\n",
		"No", "NIM", "Nama", "H", "I", "S", "A", "Total")
	fmt.Println("---  ------  --------------------  ---  ---  ---  ---  -----")
	var i int
	for i = 0; i < n; i++ {
		fmt.Printf("%-3d  %-6s  %-20s  %-3d  %-3d  %-3d  %-3d  %d\n",
			i+1,
			A[i].nim,
			A[i].nama,
			A[i].hadir,
			A[i].izin,
			A[i].sakit,
			A[i].alpa,
			totalAbsensi(A[i]),
		)
	}
}
func sesiAbsensi(A *dataMahasiswa, nA int, J dataJadwal, L *dataLog, nL *int, pertemuanKe *[NKELAS]int) {
	var kode string
	tampilJadwal(J, *pertemuanKe)
	fmt.Print("\nMasukkan Kode Kelas : ")
	fmt.Scan(&kode)
	idxJ := cariJadwal(J, kode)
	if idxJ == -1 {
		fmt.Println("Kode kelas tidak ditemukan!")
		return
	}
	fmt.Println("\n================================================")
	fmt.Printf("  Kelas     : %s - %s\n", J[idxJ].kodeKelas, J[idxJ].mataKuliah)
	fmt.Printf("  Hari/Jam  : %s, %s\n", J[idxJ].hari, J[idxJ].jam)
	fmt.Printf("  Ruangan   : %s\n", J[idxJ].ruangan)
	fmt.Printf("  Pertemuan : ke-%d\n", pertemuanKe[idxJ]+1)
	fmt.Println("================================================")
	fmt.Println("Ketik NIM untuk absen, ketik SELESAI untuk mengakhiri sesi.")
	for {
		var input string
		fmt.Print("\nNIM (atau SELESAI) : ")
		fmt.Scan(&input)
		if input == "SELESAI" {
			pertemuanKe[idxJ]++
			fmt.Printf("\nSesi pertemuan ke-%d untuk %s selesai.\n",
				pertemuanKe[idxJ], J[idxJ].mataKuliah)
			fmt.Printf("Pertemuan berikutnya otomatis menjadi ke-%d.\n", pertemuanKe[idxJ]+1)
			return
		}
		idxA := sequentialSearch(*A, nA, input)
		if idxA == -1 {
			fmt.Println("NIM tidak ditemukan! Coba lagi.")
			continue
		}
		idxA = idxA - 1
		fmt.Printf("Nama     : %s\n", A[idxA].nama)
		fmt.Print("Status (H/I/S/A) : ")
		var status string
		fmt.Scan(&status)
		if status != "H" && status != "I" && status != "S" && status != "A" {
			fmt.Println("Status tidak valid! Gunakan H / I / S / A. Absen dibatalkan.")
			continue
		}
		var sudahAbsen bool = false
		var k int
		for k = 0; k < *nL; k++ {
			if L[k].nim == input && L[k].kodeKelas == kode && L[k].pertemuan == pertemuanKe[idxJ]+1 {
				sudahAbsen = true
			}
		}
		if sudahAbsen {
			fmt.Printf("%s sudah diabsen di pertemuan ini!\n", A[idxA].nama)
			continue
		}
		L[*nL].nim = input
		L[*nL].namaMhs = A[idxA].nama
		L[*nL].kodeKelas = J[idxJ].kodeKelas
		L[*nL].mataKuliah = J[idxJ].mataKuliah
		L[*nL].pertemuan = pertemuanKe[idxJ] + 1
		L[*nL].status = status
		*nL++
		if status == "H" {
			A[idxA].hadir++
		} else if status == "I" {
			A[idxA].izin++
		} else if status == "S" {
			A[idxA].sakit++
		} else if status == "A" {
			A[idxA].alpa++
		}
		fmt.Printf("✓ %s - %s - Pertemuan %d - Status: %s tercatat.\n",
			A[idxA].nama, J[idxJ].mataKuliah, pertemuanKe[idxJ]+1, status)
	}
}
func tampilLog(L dataLog, nL int, J dataJadwal, pertemuanKe [NKELAS]int) {
	if nL == 0 {
		fmt.Println("Belum ada data log kehadiran.")
		return
	}
	var kode string
	tampilJadwal(J, pertemuanKe)
	fmt.Print("\nFilter by Kode Kelas (atau ketik 'semua') : ")
	fmt.Scan(&kode)
	fmt.Printf("\n%-6s  %-20s  %-5s  %-40s  %-9s  %s\n",
		"NIM", "Nama", "Kelas", "Mata Kuliah", "Pertemuan", "Status")
	fmt.Println("------  --------------------  -----  ----------------------------------------  ---------  ------")
	var i int
	var ada bool = false
	for i = 0; i < nL; i++ {
		if kode == "semua" || L[i].kodeKelas == kode {
			fmt.Printf("%-6s  %-20s  %-5s  %-40s  %-9d  %s\n",
				L[i].nim,
				L[i].namaMhs,
				L[i].kodeKelas,
				L[i].mataKuliah,
				L[i].pertemuan,
				L[i].status,
			)
			ada = true
		}
	}
	if !ada {
		fmt.Println("Tidak ada data untuk kelas tersebut.")
	}
}
func sequentialSearch(A dataMahasiswa, n int, nim string) int {
	var found bool = false
	var i int = 0
	for i < n && !found {
		found = A[i].nim == nim
		i = i + 1
	}
	if found {
		return i
	}
	return -1
}
func sequentialSearchStatus(A dataMahasiswa, n int, status string) {
	var found bool = false
	var i int = 0
	fmt.Println("\nHasil pencarian status", status, ":")
	for i < n {
		var cocok bool = false
		if status == "H" {
			cocok = A[i].hadir > 0
		} else if status == "I" {
			cocok = A[i].izin > 0
		} else if status == "S" {
			cocok = A[i].sakit > 0
		} else if status == "A" {
			cocok = A[i].alpa > 0
		}
		if cocok {
			fmt.Println(A[i].nim, A[i].nama,
				"H:", A[i].hadir, "I:", A[i].izin,
				"S:", A[i].sakit, "A:", A[i].alpa)
			found = true
		}
		i = i + 1
	}
	if !found {
		fmt.Println("Tidak ada mahasiswa dengan status tersebut.")
	}
}
func insertionSortByNIM(A *dataMahasiswa, n int) {
	var pass, i int
	var temp mahasiswa
	for pass = 1; pass < n; pass++ {
		temp = A[pass]
		i = pass
		for i > 0 && temp.nim < A[i-1].nim {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
	}
}
func binarySearch(A *dataMahasiswa, n int, nim string) int {
	insertionSortByNIM(A, n)
	var left, right, mid int
	var found bool = false
	left = 0
	right = n - 1
	for left <= right && !found {
		mid = (left + right) / 2
		if A[mid].nim == nim {
			found = true
		} else if nim < A[mid].nim {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if found {
		return mid + 1
	}
	return -1
}
func selectionSort(A *dataMahasiswa, n int) {
	var pass, idx, i int
	var temp mahasiswa
	for pass = 0; pass < n-1; pass++ {
		idx = pass
		for i = pass + 1; i < n; i++ {
			if totalAbsensi(A[i]) < totalAbsensi(A[idx]) {
				idx = i
			}
		}
		temp = A[pass]
		A[pass] = A[idx]
		A[idx] = temp
	}
}
func insertionSort(A *dataMahasiswa, n int) {
	var pass, i int
	var temp mahasiswa
	for pass = 1; pass < n; pass++ {
		temp = A[pass]
		i = pass
		for i > 0 && temp.nama < A[i-1].nama {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
	}
}
func totalAbsensi(m mahasiswa) int {
	return m.hadir + m.izin + m.sakit + m.alpa
}
func persentase(m mahasiswa) float64 {
	total := totalAbsensi(m)
	if total == 0 {
		return 0
	}
	return float64(m.hadir) / float64(total) * 100
}
func statistik(A dataMahasiswa, n int) {
	var i int
	var idxMax int = 0
	if n == 0 {
		fmt.Println("Belum ada data mahasiswa.")
		return
	}
	fmt.Println("\nSTATISTIK KEHADIRAN")
	fmt.Printf("%-3s  %-6s  %-20s  %-8s  %-3s  %-3s  %-3s  %s\n",
		"No", "NIM", "Nama", "% Hadir", "H", "I", "S", "A")
	fmt.Println("---  ------  --------------------  --------  ---  ---  ---  ---")
	for i = 0; i < n; i++ {
		fmt.Printf("%-3d  %-6s  %-20s  %7.2f%%  %-3d  %-3d  %-3d  %d\n",
			i+1,
			A[i].nim,
			A[i].nama,
			persentase(A[i]),
			A[i].hadir, A[i].izin, A[i].sakit, A[i].alpa,
		)
		if A[i].alpa > A[idxMax].alpa {
			idxMax = i
		}
	}
	fmt.Printf("\nMahasiswa alpa terbanyak : %s (%d kali alpa)\n",
		A[idxMax].nama, A[idxMax].alpa)
}
func tampilMenu() {
	fmt.Println("\n+++ SiPresensi +++")
	fmt.Println("--- Absensi ---")
	fmt.Println("1.  Mulai Sesi Absensi")
	fmt.Println("2.  Tampil Log Kehadiran")
	fmt.Println("--- Data Mahasiswa ---")
	fmt.Println("3.  Tambah Mahasiswa")
	fmt.Println("4.  Ubah Mahasiswa")
	fmt.Println("5.  Hapus Mahasiswa")
	fmt.Println("6.  Tampil Data Mahasiswa")
	fmt.Println("--- Jadwal ---")
	fmt.Println("7.  Tampil Jadwal Kelas")
	fmt.Println("--- Pencarian ---")
	fmt.Println("8.  Sequential Search by NIM")
	fmt.Println("9.  Sequential Search by Status")
	fmt.Println("10. Binary Search by NIM")
	fmt.Println("--- Pengurutan ---")
	fmt.Println("11. Selection Sort (by Total Absen)")
	fmt.Println("12. Insertion Sort (by Nama)")
	fmt.Println("--- Lainnya ---")
	fmt.Println("13. Statistik Kehadiran")
	fmt.Println("0.  Keluar")
}
func main() {
	var A dataMahasiswa
	var nA int
	var J dataJadwal
	var L dataLog
	var nL int
	var pilih int
	var nim string
	var idx int
	var pertemuanKe [NKELAS]int
	inisialisasiJadwal(&J)
	inisialisasiDummy(&A, &nA, &L, &nL, &pertemuanKe, J)
	fmt.Println("================================================")
	fmt.Println("  +++ Selamat Datang di SiPresensi +++")
	fmt.Println("  Sistem Monitoring Presensi Mahasiswa")
	fmt.Println("================================================")
	for {
		tampilMenu()
		fmt.Print("Pilihan : ")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			sesiAbsensi(&A, nA, J, &L, &nL, &pertemuanKe)
		case 2:
			tampilLog(L, nL, J, pertemuanKe)
		case 3:
			tambahMahasiswa(&A, &nA)
		case 4:
			ubahMahasiswa(&A, nA)
		case 5:
			hapusMahasiswa(&A, &nA)
		case 6:
			tampilData(A, nA)
		case 7:
			tampilJadwal(J, pertemuanKe)
		case 8:
			fmt.Print("Cari NIM : ")
			fmt.Scan(&nim)
			idx = sequentialSearch(A, nA, nim)
			if idx != -1 {
				idx = idx - 1
				fmt.Println("\nHasil Sequential Search:")
				fmt.Printf("%-6s  %-20s  %-3s  %-3s  %-3s  %-3s  %s\n",
					"NIM", "Nama", "H", "I", "S", "A", "Total")
				fmt.Println("------  --------------------  ---  ---  ---  ---  -----")
				fmt.Printf("%-6s  %-20s  %-3d  %-3d  %-3d  %-3d  %d\n",
					A[idx].nim, A[idx].nama,
					A[idx].hadir, A[idx].izin, A[idx].sakit, A[idx].alpa,
					totalAbsensi(A[idx]),
				)
			} else {
				fmt.Println("NIM tidak ditemukan.")
			}
		case 9:
			var status string
			fmt.Print("Cari Status (H/I/S/A) : ")
			fmt.Scan(&status)
			sequentialSearchStatus(A, nA, status)
		case 10:
			fmt.Print("Cari NIM : ")
			fmt.Scan(&nim)
			idx = binarySearch(&A, nA, nim)
			if idx != -1 {
				idx = idx - 1
				fmt.Println("\nHasil Binary Search:")
				fmt.Printf("%-6s  %-20s  %-3s  %-3s  %-3s  %-3s  %s\n",
					"NIM", "Nama", "H", "I", "S", "A", "Total")
				fmt.Println("------  --------------------  ---  ---  ---  ---  -----")
				fmt.Printf("%-6s  %-20s  %-3d  %-3d  %-3d  %-3d  %d\n",
					A[idx].nim, A[idx].nama,
					A[idx].hadir, A[idx].izin, A[idx].sakit, A[idx].alpa,
					totalAbsensi(A[idx]),
				)
			} else {
				fmt.Println("NIM tidak ditemukan.")
			}
		case 11:
			selectionSort(&A, nA)
			fmt.Println("Data diurutkan by total absen:")
			tampilData(A, nA)
		case 12:
			insertionSort(&A, nA)
			fmt.Println("Data diurutkan by nama:")
			tampilData(A, nA)
		case 13:
			statistik(A, nA)
		case 0:
			fmt.Println("Terima kasih telah menggunakan SiPresensi!")
			return
		}
	}
}