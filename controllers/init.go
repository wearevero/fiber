package controllers

import (
	"gorm.io/gorm"

	"github.com/wearevero/fiber/controllers/Laporan/absenhariancontroller"
	"github.com/wearevero/fiber/controllers/Laporan/absenjamcontroller"
	"github.com/wearevero/fiber/controllers/Laporan/absenlemburcontroller"
	"github.com/wearevero/fiber/controllers/MasterData/bagiancontroller"
	"github.com/wearevero/fiber/controllers/MasterData/jabatancontroller"
	"github.com/wearevero/fiber/controllers/MasterData/karyawancontroller"
	"github.com/wearevero/fiber/controllers/MasterData/usercontroller"
)

func InitControllers(db *gorm.DB) {
	bagiancontroller.SetDB(db)
	jabatancontroller.SetDB(db)
	karyawancontroller.SetDB(db)
	usercontroller.SetDB(db)
	absenjamcontroller.SetDB(db)
	absenhariancontroller.SetDB(db)
	absenlemburcontroller.SetDB(db)
}
