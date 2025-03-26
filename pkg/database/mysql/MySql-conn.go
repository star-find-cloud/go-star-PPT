package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"localhost/business_code_test/config"
	log "localhost/business_code_test/logger"
)

type MySQL struct {
	_db    *sqlx.DB
	logger *zap.SugaredLogger
}

var master = &MySQL{}
var slave1 = &MySQL{}
var slave2 = &MySQL{}

func init() {
	c := config.GetConfig()
	dbUser := c.Database.MySQL.User
	dbPass := c.Database.MySQL.Password
	dbMasterHost := c.Database.MySQL.MasterHost
	//dbSlaveHost1 := c.Database.MySQL.SlaveHost1
	//dbSlaveHost2 := c.Database.MySQL.SlaveHost2
	dbPort := c.Database.MySQL.Port
	dbName := c.Database.MySQL.DBName
	timeout := c.Database.MySQL.Timeout

	dbLog := log.GetLog("MySQL")
	masterDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&timeout=%s", dbUser, dbPass, dbMasterHost, dbPort, dbName, timeout)
	m_db, err := sqlx.Connect("mysql", masterDSN)
	if err != nil {
		fmt.Println("Database error, please check the logs.")
		dbLog.Errorf("MySQL master connect faild: %s \n", err)
	} else {
		dbLog.Infof("MySQL master connection successful: %s\n", dbMasterHost)
	}

	//slaveDsn1 := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&timeout=%s", dbUser, dbPass, dbSlaveHost1, dbPort, dbName, timeout)
	//s1_db, err := sqlx.Connect("mysql", slaveDsn1)
	//if err != nil {
	//	dbLog = log.GetLog("MySQL")
	//	dbLog.Errorf("MySQL slave1 connect faild: %s \n", err)
	//} else {
	//	dbLog.Infof("MySQL slave1 connection successful: %s\n", dbSlaveHost1)
	//}
	//slaveDsn2 := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&timeout=%s", dbUser, dbPass, dbSlaveHost2, dbPort, dbName, timeout)
	//s2_db, err := sqlx.Connect("mysql", slaveDsn2)
	//if err != nil {
	//	dbLog = log.GetLog("MySQL")
	//	dbLog.Errorf("MySQL slave2 connect faild: %s \n", err)
	//} else {
	//	dbLog.Infof("MySQL slave2 connection successful: %s \n", dbSlaveHost2)
	//}

	fmt.Println(m_db)

	m_db.SetMaxOpenConns(c.Database.MySQL.MaxOpenConns)
	m_db.SetMaxIdleConns(c.Database.MySQL.MaxIdleConns)
	//s1_db.SetMaxOpenConns(c.Database.MySQL.MaxOpenConns)
	//s1_db.SetMaxIdleConns(c.Database.MySQL.MaxIdleConns)
	//s2_db.SetMaxOpenConns(c.Database.MySQL.MaxOpenConns)
	//s2_db.SetMaxIdleConns(c.Database.MySQL.MaxIdleConns)

	master._db = m_db
	//slave1._db = s1_db
	//slave2._db = s2_db
}

func (db *MySQL) GetLog() *zap.SugaredLogger {
	db.logger = log.GetLog("MySQL")
	return db.logger
}

func GetMasterDB() *sqlx.DB {
	return master._db
}

func GetSlaveDB1() *sqlx.DB {
	return slave1._db
}

func GetSlaveDB2() *sqlx.DB {
	return slave2._db
}
