package main

import (
	"7task/pegawai"
	"fmt"

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
	pg := []pegawai.Pegawai{
		{Nama: "Purnama", Posisi: "Software Engineer", GajiBulanan: 30000000},
		{Nama: "Farhani", Posisi: "System Analyst", GajiBulanan: 40000000},
		{Nama: "Faizah", Posisi: "UI Designer", GajiBulanan: 15000000},
		{Nama: "Ririn", Posisi: "UX Designer", GajiBulanan: 18000000},
		{Nama: "Akmal", Posisi: "Backend Developer", GajiBulanan: 15000000},
	}

	if err := db.Create(&pg).Error; err != nil {
		fmt.Println("Failed to create table:", err)
	}
}

func showAllPegawai(db *gorm.DB) {
	var pg []pegawai.Pegawai
	if err := db.Find(&pg).Error; err != nil {
		fmt.Println("Failed to show all pegawai:", err)
	}
	for _, p := range pg {
		p.TampilkanInformasi()
	}
}

func showPegawaiByID(db *gorm.DB, id uint) {
	var pg pegawai.Pegawai
	if err := db.First(&pg, id).Error; err != nil {
		fmt.Println("Failed to show pegawai:", err)
		return
	}
	pg.TampilkanInformasi()
}

func insert(db *gorm.DB, pg pegawai.Pegawai) {
	if err := db.Create(&pg).Error; err != nil {
		fmt.Println("Failed to insert pegawai:", err)
	}
}

func update(db *gorm.DB, pg pegawai.Pegawai) {
	if err := db.Save(&pg).Error; err != nil {
		fmt.Println("Failed to update pegawai:", err)
	}
}

func delete(db *gorm.DB, id uint) {
	if err := db.Delete(&pegawai.Pegawai{}, id).Error; err != nil {
		fmt.Println("Failed to delete pegawai:", err)
	}
}

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/sandbox_golang?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect database:", err)
	}

	setup(db)

	p := pegawai.Pegawai{Nama: "Johan Budi", Posisi: "Oracle Developer", GajiBulanan: 40000000}
	insert(db, p)
	fmt.Println("\nList Pegawai")
	showAllPegawai(db)

	if err := db.First(&p, 1).Error; err != nil {
		fmt.Println("Failed to find pegawai:", err)
	}
	p.GajiBulanan = 20000000
	update(db, p)
	fmt.Println("\nPegawai by First Record")
	showPegawaiByID(db, p.ID)

	delete(db, p.ID)
	fmt.Println("\nList Pegawai After Delete")
	showAllPegawai(db)
}
