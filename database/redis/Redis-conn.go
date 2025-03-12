package redis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"localhost/business_code_test/config"
	log "localhost/business_code_test/logger"
)

type Redis struct {
	rdb    *redis.Client
	logger *zap.SugaredLogger
}

var rdbMaster Redis
var rdbSlave1 Redis
var rdbSlave2 Redis

func init() {
	var c = config.GetConfig()
	masterHost := c.Database.Redis.MasterHost
	slaveHost1 := c.Database.Redis.SlaveHost1
	slaveHost2 := c.Database.Redis.SlaveHost2
	port := c.Database.Redis.Port
	passwd := c.Database.Redis.Password

	rdbLog := log.GetLog("Redis")
	masterURL := fmt.Sprintf("redis://root:%s@%s:%d/0", passwd, masterHost, port)
	master, err := redis.ParseURL(masterURL)
	if err != nil {
		rdbLog.Errorf("Redis Master connect faild: %s\n", err)
	} else {
		rdbLog.Infof("Redis Master connect success: %s\n", masterURL)
	}
	rdbMaster.rdb = redis.NewClient(master)

	slaveURL1 := fmt.Sprintf("redis://root:%s@%s:%d/0", passwd, slaveHost1, port)
	slave1, err := redis.ParseURL(slaveURL1)
	if err != nil {
		rdbLog.Errorf("Redis Slave1 connect faild: %s\n", err)
	} else {
		rdbLog.Infof("Redis Slave1 connect success: %s\n", slaveURL1)
	}
	rdbSlave1.rdb = redis.NewClient(slave1)

	slaveURL2 := fmt.Sprintf("redis://root:%s@%s:%d/0", passwd, slaveHost2, port)
	slave2, err := redis.ParseURL(slaveURL2)
	if err != nil {
		rdbLog.Errorf("Redis Slave2 connect faild: %s\n", err)
	} else {
		rdbLog.Infof("Redis Slave2 connect success: %s\n", slaveURL2)
	}
	rdbSlave2.rdb = redis.NewClient(slave2)
}

func (rdb Redis) GetLogger() *zap.SugaredLogger {
	rdb.logger = log.GetLog("Redis")
	return rdb.logger
}

func GetRDBMaster() *redis.Client {
	return rdbMaster.rdb
}

func GetRDBSlave1() *redis.Client {
	return rdbSlave1.rdb
}

func GetRDBSlave2() *redis.Client {
	return rdbSlave2.rdb
}
