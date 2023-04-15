package postgres

import (
	"dawn-bot/src/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	DiscordID  string
	Name       string `gorm:"size:256"`
	XP         uint32
	DistrictID int
	District   District `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type District struct {
	gorm.Model
	Name string `gorm:"size:256"`
	Max  uint8
}

type Sanction struct {
	gorm.Model
	Name        string `gorm:"size:256"`
	Description string `gorm:"size:1024"`
}

type Mod struct {
	gorm.Model
	UserID     int
	User       User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SanctionID int
	Sanction   Sanction `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ModID      int
	Mod        User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Migrate do the migration of database
func Migrate() {
	err := Db.AutoMigrate(&District{})
	utils.PanicError(err)
	err = Db.AutoMigrate(&User{})
	utils.PanicError(err)
	err = Db.AutoMigrate(&Sanction{})
	utils.PanicError(err)
	err = Db.AutoMigrate(&Mod{})
	utils.PanicError(err)
}
