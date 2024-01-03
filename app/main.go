package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_init "acs/core"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	_ "github.com/denisenkom/go-mssqldb"
)

func init() {
	viper.SetConfigFile(`config.json`)
	// viper.SetConfigFile(`/home/regate/r_backend/app/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}

	
}

func main(){
	loc, err := time.LoadLocation("America/Chicago")
	if err != nil {
		log.Println(loc)
	}
	host := viper.GetString(`database.host`)
	port := viper.GetInt(`database.port`)
	// user := viper.GetString(`database.user`)
	// password := viper.GetString(`database.pass`)
	dbname := viper.GetString(`database.name`)
	time.Local = loc
	// psqlInfo := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
	// 	 user, password,host,port, dbname)


	// stringConn := fmt.Sprintf("mssql://%s:%s@%s:%s/instance?database=%s",
	// 	 user, password,host,"1434", dbname)
	connString := fmt.Sprintf("server=%s;port=%d;database=%s;trusted_connection=yes", host, port,dbname)

	
	db, err := sql.Open("sqlserver", connString)	 

	// db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
	}

	_init.InitServer(db)
	defer db.Close()
}