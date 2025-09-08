package main

// TODO: other imports
import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// TODO: definisi Model/Struktur Tabel
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

}
