package postgres

import "strconv"

type PostgresOptions struct {
	Host     string `toml:"host"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Db       string `toml:"db"`
	Port     uint16 `toml:"port"`
	Timezone string `toml:"timezone"`
}

var options string

// GenerateDns Generate the dns to connect to the postgres database
func GenerateDns(o PostgresOptions) {
	options = "host=" + o.Host + " user=" + o.User + " password=" + o.Password + " dbname=" + o.Db + " port=" + strconv.Itoa(int(o.Port))
	options = options + " sslmod=disable TimeZone=" + o.Timezone
}
