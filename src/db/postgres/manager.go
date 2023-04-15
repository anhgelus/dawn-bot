package postgres

import (
	"dawn-bot/src/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
)

type PostgresOptions struct {
	Host     string `toml:"host"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Db       string `toml:"db"`
	Port     uint16 `toml:"port"`
	Timezone string `toml:"timezone"`
}

// Connect to the database
func (o *PostgresOptions) Connect() {
	GenerateDns(*o)
	Connect()
}

var options string
var Db *gorm.DB

// GenerateDns generate the dns to connect to the postgres database
func GenerateDns(o PostgresOptions) {
	options = "host=" + o.Host + " user=" + o.User + " password=" + o.Password + " dbname=" + o.Db + " port=" + strconv.Itoa(int(o.Port))
	options = options + " sslmod=disable TimeZone=" + o.Timezone
}

// Connect the database
func Connect() {
	var err error
	Db, err = gorm.Open(postgres.Open(options), &gorm.Config{})
	utils.PanicError(err)
}
