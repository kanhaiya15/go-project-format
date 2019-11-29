package gopfmysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/jmoiron/sqlx"
	"github.com/kanhaiya15/gopf/constants"
	"github.com/kanhaiya15/gopf/utils"
)

// Pool sqlx.DB
var (
	Pool *sqlx.DB
)

// DBConfig Db Config struct
type DBConfig struct {
	Driver       string
	Protocol     string
	Host         string
	Port         string
	Username     string
	Password     string
	DatabaseName string
}

// Init initialize
func Init() {
	conn, err := getConnection()
	if err != nil {
		panic(err.Error())
	}
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conn.Username, conn.Password, conn.Host, conn.Port, conn.DatabaseName)
	Pool, err = sqlx.Connect("mysql", connection)
	if err != nil {
		panic(err.Error())
	}
	Pool.SetMaxIdleConns(constants.MysqlMaxIdleConnection)
	Pool.SetMaxOpenConns(constants.MysqlMaxOpenConnection)
	Pool.SetConnMaxLifetime(constants.MysqlMaxConnectionLifetime)

	err = Pool.Ping()
	if err != nil {
		panic(err.Error())
	}
}

//GetConn Get Pool
func GetConn() *sqlx.DB {
	return Pool
}

//DBstats Get DB Stats
func DBstats() interface{} {
	return Pool.DB.Stats()
}

//DBstatus Get Conn Ping
func DBstatus() error {
	return Pool.Ping()
}

func getConnection() (conn DBConfig, err error) {
	host, err := utils.GetConfValue("DB_HOST")
	if err != nil {
		panic(err.Error())
	}

	port, err := utils.GetConfValue("DB_PORT")
	if err != nil {
		panic(err.Error())
	}

	username, err := utils.GetConfValue("DB_USERNAME")
	if err != nil {
		panic(err.Error())
	}

	password, err := utils.GetConfValue("DB_PASSWORD")
	if err != nil {
		panic(err.Error())
	}

	dbName, err := utils.GetConfValue("DB_NAME")
	if err != nil {
		panic(err.Error())
	}
	conn = DBConfig{
		Host:         host,
		Port:         port,
		Username:     username,
		Password:     password,
		DatabaseName: dbName,
	}
	return conn, nil
}
