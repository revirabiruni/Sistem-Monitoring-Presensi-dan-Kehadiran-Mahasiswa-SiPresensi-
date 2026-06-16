# Sistem-Monitoring-Presensi-dan-Kehadiran-Mahasiswa-SiPresensi-
SiPresensi — Sistem Monitoring Presensi Mahasiswa
Aplikasi berbasis terminal untuk mencatat dan memantau kehadiran mahasiswa dalam perkuliahan.
Dibuat menggunakan bahasa pemrograman Go (Golang) sebagai Tugas Besar mata kuliah Algoritma dan Pemrograman 2.
Anggota Kelompok
Anggota 1   : Thania Brova Sastaviana / 108072500122
Anggota 2   : Biruni Revira Ramadhani / 108072500182


Deskripsi Program
SiPresensi adalah aplikasi untuk mencatat dan memantau tingkat kehadiran mahasiswa dalam perkuliahan. Data utama yang digunakan adalah:
•	Data Mahasiswa — menyimpan identitas dan rekap kehadiran tiap mahasiswa
•	Data Jadwal Kelas — menyimpan 8 kelas dengan info hari, jam, dan ruangan
•	Data Log Kehadiran — mencatat status absensi (Hadir/Izin/Sakit/Alpa) per mahasiswa per pertemuan
Pengguna aplikasi adalah dosen atau staf administrasi akademik.

Langkah Menjalankan
1. Clone repository ini
2. Jalankan program
go run main.go
3. Tampilan awal program
  +++ Selamat Datang di SiPresensi +++
  Sistem Monitoring Presensi Mahasiswa
+++ SiPresensi +++
--- Absensi ---
1.  Mulai Sesi Absensi
2.  Tampil Log Kehadiran
--- Data Mahasiswa ---
3.  Tambah Mahasiswa
...
Pilihan :
Masukkan angka sesuai menu yang ingin diakses, lalu tekan Enter.

 
Struktur Data
Program menggunakan tiga struct utama yang didefinisikan sebagai tipe bentukan:
// Menyimpan data dan rekap kehadiran satu mahasiswa
type mahasiswa struct {
    nim   string
    nama  string
    hadir int   // jumlah hadir
    izin  int   // jumlah izin
    sakit int   // jumlah sakit
    alpa  int   // jumlah alpa
}

// Menyimpan informasi satu kelas/mata kuliah
type jadwal struct {
    kodeKelas  string
    mataKuliah string
    hari       string
    jam        string
    ruangan    string
}

// Menyimpan satu catatan absensi
type logKehadiran struct {
    nim        string
    namaMhs    string
    kodeKelas  string
    mataKuliah string
    pertemuan  int
    status     string  // "H" / "I" / "S" / "A"
}
Array statis menggunakan type alias:
const NMAX   int = 100   // maks mahasiswa
const NKELAS int = 8     // jumlah kelas

type dataMahasiswa [NMAX]mahasiswa
type dataJadwal    [NKELAS]jadwal
type dataLog       [NMAX * NKELAS]logKehadiran

 
Penjelasan Fitur
1. Sesi Absensi (Menu 1)
Dosen memilih kode kelas, lalu memasukkan NIM mahasiswa satu per satu.
Setiap mahasiswa diinput status kehadirannya: H (Hadir), I (Izin), S (Sakit), A (Alpa).
Program otomatis:
•	Mencegah absen ganda di pertemuan yang sama
•	Menambah counter H/I/S/A pada data mahasiswa
•	Menyimpan log ke array dataLog
•	Menaikkan nomor pertemuan setelah sesi selesai
Contoh penggunaan:
NIM (atau SELESAI) : 001
Nama     : George Russell
Status (H/I/S/A) : H
George Russell - Bahasa Inggris - Pertemuan 3 - Status: H tercatat.

2. Tampil Log Kehadiran (Menu 2)
Menampilkan seluruh riwayat absensi. Bisa difilter per kode kelas atau tampil semua.

3. Tambah Mahasiswa (Menu 3)
Menambah mahasiswa baru ke sistem. Program memvalidasi:
•	Array tidak melebihi batas NMAX (100 mahasiswa)
•	NIM tidak boleh duplikat (dicek dengan Sequential Search)

4. Ubah & Hapus Mahasiswa (Menu 4 & 5)
•	Ubah: Mengubah nama mahasiswa berdasarkan NIM
•	Hapus: Menghapus data mahasiswa dengan menggeser elemen array ke kiri

5. Tampil Data Mahasiswa (Menu 6)
Menampilkan seluruh data mahasiswa beserta rekap H/I/S/A dan total kehadiran.
No   NIM     Nama                  H    I    S    A    Total
1    001     George Russell        8    0    0    0    8
2    002     Kimi Antonelli        7    0    1    0    8
...

6. Sequential Search by NIM (Menu 8)
Mencari data mahasiswa berdasarkan NIM menggunakan algoritma Sequential Search.
Program menelusuri array satu per satu dari indeks 0 hingga data ditemukan atau habis.
func sequentialSearch(A dataMahasiswa, n int, nim string) int {
    var found bool = false
    var i int = 0
    for i < n && !found {
        found = A[i].nim == nim
        i = i + 1
    }
    if found {
        return i   // mengembalikan posisi (1-based)
    }
    return -1      // tidak ditemukan
}
Kompleksitas: O(n) — linear terhadap jumlah data.

7. Sequential Search by Status (Menu 9)
Mencari dan menampilkan semua mahasiswa yang memiliki minimal 1 catatan status tertentu (H/I/S/A).

8. Binary Search by NIM (Menu 10)
Mencari NIM menggunakan algoritma Binary Search yang lebih efisien.
Array diurutkan terlebih dahulu berdasarkan NIM, lalu pencarian dilakukan dengan membagi rentang pencarian menjadi dua di setiap langkah.
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
Kompleksitas: O(log n) — jauh lebih cepat dari Sequential Search untuk data besar.

9. Selection Sort by Total Absensi (Menu 11)
Mengurutkan data mahasiswa secara ascending berdasarkan total kehadiran menggunakan Selection Sort.
Di setiap iterasi, dicari elemen dengan total terkecil dari sisa array, lalu ditukar ke posisi terdepan.
for pass = 0; pass < n-1; pass++ {
    idx = pass                         // anggap posisi pass adalah minimum
    for i = pass + 1; i < n; i++ {
        if totalAbsensi(A[i]) < totalAbsensi(A[idx]) {
            idx = i                    // update jika ada yang lebih kecil
        }
    }
    // tukar A[pass] dengan A[idx]
    temp = A[pass]; A[pass] = A[idx]; A[idx] = temp
}

10. Insertion Sort by Nama (Menu 12)
Mengurutkan data mahasiswa secara ascending berdasarkan nama menggunakan Insertion Sort.
Setiap elemen "disisipkan" ke posisi yang tepat di bagian array yang sudah terurut.
for pass = 1; pass < n; pass++ {
    temp = A[pass]
    i = pass
    for i > 0 && temp.nama < A[i-1].nama {
        A[i] = A[i-1]   // geser elemen ke kanan
        i--
    }
    A[i] = temp          // sisipkan di posisi yang tepat
}

11. Statistik Kehadiran (Menu 13)
Menampilkan persentase kehadiran setiap mahasiswa dan menampilkan nama mahasiswa dengan jumlah alpa terbanyak.
STATISTIK KEHADIRAN
No   NIM     Nama                  % Hadir   H    I    S    A
1    001     George Russell          100.00%  8    0    0    0
2    002     Kimi Antonelli           87.50%  7    0    1    0
...
Mahasiswa alpa terbanyak : Carlos Sainz (1 kali alpa)
Rumus persentase:
% Hadir = (jumlah Hadir / Total Pertemuan) × 100

Daftar Seluruh Fungsi/Prosedur
No	Nama Fungsi/Prosedur	Jenis 	Deskripsi Singkat
1	inisialisasiJadwal	Prosedur	Mengisi 8 data jadwal kelas
2	inisialisasiDummy	Prosedur	Mengisi data dummy mahasiswa & log awal
3	cariJadwal	Fungsi	Mencari indeks jadwal berdasarkan kode kelas
4	tampilJadwal	Prosedur	Menampilkan semua jadwal kelas
5	tambahMahasiswa	Prosedur	Menambah mahasiswa baru (dengan validasi)
6	ubahMahasiswa	Prosedur	Mengubah nama mahasiswa
7	hapusMahasiswa	Prosedur	Menghapus mahasiswa + geser array
8	tampilData	Prosedur	Menampilkan tabel data mahasiswa
9	sesiAbsensi	Prosedur	Menjalankan sesi absensi kelas
10	tampilLog	Prosedur	Menampilkan log kehadiran (dengan filter)
11	sequentialSearch	Fungsi	Pencarian NIM secara sekuensial
12	sequentialSearchStatus	Prosedur	Menampilkan mahasiswa berdasarkan status
13	insertionSortByNIM	Prosedur	Mengurutkan array by NIM (untuk Binary Search)
14	binarySearch	Fungsi	Pencarian NIM dengan Binary Search
15	selectionSort	Prosedur	Mengurutkan by total absensi (Selection Sort)
16	insertionSort	Prosedur	Mengurutkan by nama (Insertion Sort)
17	totalAbsensi	Fungsi	Menghitung total H+I+S+A seorang mahasiswa
18	persentase	Fungsi	Menghitung % kehadiran seorang mahasiswa
19	statistik	Prosedur	Menampilkan statistik & mahasiswa alpa terbanyak
20	tampilMenu	Prosedur	Menampilkan menu utama
21	main	Fungsi	Program utama: inisialisasi + loop menu
22	inisialisasiDummy	Prosedur	Mengisi 8 data jadwal kelas

Referensi Modul
Modul	Topik	Digunakan Pada
Modul 3	Fungsi	totalAbsensi, persentase, sequentialSearch, binarySearch, cariJadwal
Modul 4	Prosedur & Pointer	tambahMahasiswa, hapusMahasiswa, sesiAbsensi, semua sort
Modul 7	Struct & Tipe Bentukan	mahasiswa, jadwal, logKehadiran, type alias array
Modul 9	Array Statis	dataMahasiswa, dataJadwal, dataLog, semua operasi array
Modul 12	Searching	sequentialSearch, sequentialSearchStatus, binarySearch
Modul 14	Sorting	selectionSort, insertionSort, insertionSortByNIM

