package main

import (
	"7task/pegawai"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setup(db *gorm.DB) {
	// TODO: Auto Migrate your schema (struct)
	db.AutoMigrate(&pegawai.Pegawai{})
	seed(db)
}

func seed(db *gorm.DB) {
	db.Exec("TRUNCATE TABLE pegawais")

	// TODO: Seed your database with initial data
	pegawais := []pegawai.Pegawai{
		{Nama: "Purnama", Posisi: "Software Engineer", GajiBulanan: 30000000},
		{Nama: "Farhani", Posisi: "System Analyst", GajiBulanan: 50000000},
		{Nama: "Faizah", Posisi: "UI Designer", GajiBulanan: 20000000},
		{Nama: "Ririn", Posisi: "UX Designer", GajiBulanan: 18000000},
		{Nama: "Akmal", Posisi: "Backend Developer", GajiBulanan: 20000000},
	}

	if err := db.Create(&pegawais).Error; err != nil {
		log.Fatal("Failed to create table:", err)
	}
}

func showAllPegawai(db *gorm.DB) {
	var pegawais []pegawai.Pegawai
	if err := db.Find(&pegawais).Error; err != nil {
		log.Fatal("Failed to show all pegawai:", err)
	}
	for _, p := range pegawais {
		p.TampilkanInformasi()
	}
}

func showPegawaiByID(db *gorm.DB, id uint) {
	var pegawai pegawai.Pegawai
	if err := db.First(&pegawai, id).Error; err != nil {
		log.Fatal("Failed to show pegawai:", err)
	}
	pegawai.TampilkanInformasi()
}

func insert(db *gorm.DB, pegawai pegawai.Pegawai) {
	if err := db.Create(&pegawai).Error; err != nil {
		log.Fatal("Failed to insert pegawai:", err)
	}
}

func read(db *gorm.DB, id uint) pegawai.Pegawai {
	var pegawai pegawai.Pegawai
	if err := db.First(&pegawai, id).Error; err != nil {
		log.Fatal("Failed to read pegawai:", err)
	}
	return pegawai
}

func update(db *gorm.DB, pegawai pegawai.Pegawai) {
	if err := db.Save(&pegawai).Error; err != nil {
		log.Fatal("Failed to update pegawai:", err)
	}
}

func delete(db *gorm.DB, id uint) {
	if err := db.Delete(&pegawai.Pegawai{}, id).Error; err != nil {
		log.Fatal("Failed to delete pegawai:", err)
	}
}

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/sandbox_golang?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	setup(db)

	fmt.Println("List Pegawai")
	showAllPegawai(db)

	pegawai := pegawai.Pegawai{Nama: "Johan Budi", Posisi: "Oracle Developer", GajiBulanan: 40000000}
	insert(db, pegawai)

	pegawai = read(db, 1)
	pegawai.GajiBulanan = 12000000
	update(db, pegawai)

	fmt.Println("Pegawai by ID")
	showPegawaiByID(db, 1)

	delete(db, 1)
	fmt.Println("List Pegawai After Delete")
	showAllPegawai(db)
}
