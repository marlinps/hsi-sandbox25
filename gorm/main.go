package main

// TODO: other imports
import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// TODO: definisi Model/Struktur Tabel menjadi tabel ukur_suhus(default penamaan)
type UkurSuhu struct {
	ID             uint `gorm:"primaryKey"` // TODO: sesuaikan tipe data dan tag GORM
	SuhuCelcius    float64
	SuhuFahrenheit float64
	SuhuReamur     float64
}

func main() {
	// TODO: Migrasi dan inisialisasi database
	// TODO: 1. Mendefinisikan Koneksi Database
	dsn := "root:@tcp(127.0.0.1:3306)/sandbox_golang?charset=utf8mb4&parseTime=True&loc=Local"

	// TODO: 2. Membuka Koneksi Database
	koneksiDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err.Error())
	}

	// TODO: 3. Automigrasi (Migrasi struktur tabel dari struct)
	koneksiDb.AutoMigrate(&UkurSuhu{})

	// TODO: 4. Insert Data
	// fmt.Println("Insert Data Suhu")
	// newData := UkurSuhu{
	// 	SuhuCelcius:    100,
	// 	SuhuFahrenheit: 212,
	// 	SuhuReamur:     80,
	// }

	// operasiTambahData := koneksiDb.Create(&newData)
	// if operasiTambahData.Error != nil {
	// 	log.Fatal("Gagal menambahkan data:", operasiTambahData.Error.Error())
	// }
	// fmt.Println("Data berhasil ditambahkan dengan ID:", newData.ID)

	// TODO: 5.Read Data
	fmt.Println("\nRead Data Suhu")
	// Mengggunakakn fungsi First untuk mengambil data pertama
	var dataSuhu UkurSuhu
	operasiBacaData := koneksiDb.First(&dataSuhu)
	if operasiBacaData.Error != nil {
		log.Fatal("Gagal membaca data:", operasiBacaData.Error.Error())
	}
	fmt.Println("Data berhasil dibaca:", dataSuhu)

}
