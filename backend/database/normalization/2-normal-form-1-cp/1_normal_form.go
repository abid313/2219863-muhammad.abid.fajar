package main

// pada tahap normalisai 1 (1NF), kita akan menyederhanakan bentuk unormal sesuai dengan kaidah bentuk normalisasi 1
// dengan memisahkan data rekap dengan nomor bon yang sama dari 1 row ke beberapa row

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

//the struct are irrelevant for the code, but hint for column
type Rekap struct {
	NoBon      string
	NamaBarang string
	Harga      int
	Jumlah     int
	Biaya      int
	SubTotal   int
	Discount   int
	Total      int
	Bayar      int
	Kembalian  int
	Kasir      string
	Tanggal    string
	Waktu      string
}

// Migrate digunakan untuk melakukan migrasi database dengan data yang dibutuhkan
// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi Migrate
// Buatlah tabel dengan nama rekap dan insert data seperti pada contoh di bagian bawah file ini
func Migrate() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `CREATE TABLE IF NOT EXISTS rekap (
		no_bon VARCHAR(10),
		nama_produk VARCHAR(30),
		harga_produk INTEGER,
		jumlah_produk INTEGER,
		harga_total INTEGER,
		sub_total INTEGER,
		discount INTEGER,
		total INTEGER,
		bayar INTEGER,
		kembalian INTEGER,
		nama_kasir VARCHAR(30),
		tanggal VARCHAR(30),
		waktu VARCHAR(30)
	) ;` // TODO: replace this

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
	INSERT OR REPLACE INTO 
	rekap (no_bon, nama_produk, harga_produk, jumlah_produk, harga_total, sub_total, discount, total, bayar, kembalian, nama_kasir, tanggal, waktu)
	VALUES 
	("00001", "Disket", 4500, 3, 13500, 13500, 0, 13500, 100000, 23000, "Rosi", "04-05-2022", "12:00:00"),
	("00001", "Refil Tinta", 22500, 1, 22500, 36000, 0, 36000, 100000, 23000, "Rosi", "04-05-2022", "12:00:00"),
	("00001", "CD Blank", 1500, 4, 6000, 42000, 0, 42000, 100000, 23000, "Rosi", "04-05-2022", "12:00:00"),
	("00001", "CD Mouse", 17500, 2, 35000, 77000, 0, 77000, 100000, 23000, "Rosi", "04-05-2022", "12:00:00"),
	("00002", "Disket", 4500, 1, 4500, 4500, 0, 4500, 17500, 0, "Dewi", "04-05-2022", "12:00:00"),
	("00002", "Mouse", 17400, 1, 17500, 22000, 0, 22000, 117500, 0, "Dewi", "04-05-2022", "12:00:00"),
	("00002", "Flash Disk", 100000, 1, 100000, 117500, 0, 117500, 117500, 0, "Dewi", "04-05-2022", "12:00:00");`) // TODO: replace this

	if err != nil {
		panic(err)
	}

	return db, nil
}

// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi countByNoBon:
// countByNoBon digunakan untuk menghitung jumlah data yang ada berdasarkan no_bon
func countByNoBon(noBon string) (int, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `SELECT no_bon FROM rekap WHERE no_bon = ?;` // TODO: replace this

	row, err := db.Query(sqlStmt, noBon)
	if err != nil {
		panic(err)
	}
	var countBon int
	defer row.Close()
	for row.Next() {
		countBon++
	}
	return countBon, nil
}

// insert value hint
// ("00001", "Disket", 4500, 3, 13500, 13500, 0, 13500, 100000, 23000, "Rosi", "04-05-2022", "12:00:00"),
// ("00001", "Refil Tinta", 22500, 1, 22500, 36000, 0, 36000, 100000, 23000, "Rosi", "04-05-2022", "12:00:00"),
// ("00001", "CD Blank", 1500, 4, 6000, 42000, 0, 42000, 100000, 23000, "Rosi", "04-05-2022", "12:00:00"),
// ("00001", "CD Mouse", 17500, 2, 35000, 77000, 0, 77000, 100000, 23000, "Rosi", "04-05-2022", "12:00:00"),
// ("00002", "Disket", 4500, 1, 4500, 4500, 0, 4500, 17500, 0, "Dewi", "04-05-2022", "12:00:00"),
// ("00002", "Mouse", 17400, 1, 17500, 22000, 0, 22000, 117500, 0, "Dewi", "04-05-2022", "12:00:00"),
// ("00002", "Flash Disk", 100000, 1, 100000, 117500, 0, 117500, 117500, 0, "Dewi", "04-05-2022", "12:00:00")
