package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

type Config struct {
	App      AppConfig
	Database Database
	Log      LogConf
}

type AppConfig struct {
	Name    string
	Env     string
	Port    uint64
	Version string
}

type Database struct {
	MySQL MySQLConf `mapstructure:"mysql"`
	Redis RedisConf `mapstructure:"redis"`
}

type MySQLConf struct {
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	MasterHost   string `mapstructure:"master_host"`
	SlaveHost1   string `mapstructure:"slave_host_1"`
	SlaveHost2   string `mapstructure:"slave_host_2"`
	Port         int    `mapstructure:"port"`
	DBName       string `mapstructure:"db_name"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	Timeout      string `mapstructure:"timeout"`
}

type RedisConf struct {
	MasterHost string `mapstructure:"master_host"`
	SlaveHost1 string `mapstructure:"slave_host_1"`
	SlaveHost2 string `mapstructure:"slave_host_2"`
	Port       uint64 `mapstructure:"port"`
	Password   string `mapstructure:"password"`
}

type LogConf struct {
	Version    string
	Level      string
	Path       string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

var c Config

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	logFile, err := os.OpenFile("E:/var/log/Star-PPT/conf.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	// 读取配置文件
	err = viper.ReadInConfig()
	if err != nil {
		log.Printf("config load Error: %v \n", err)
	} else {
		log.Println("configuration file was read successfully")
	}

	// 将 viper 读到的数据序列化写入 config
	if err := viper.Unmarshal(&c); err != nil {
		now := time.Now()
		log.Printf("%v: viper Unmarshal err:%s \n", now.Format("2006-01-02 15:04:05"), err)
	}
}

func GetConfig() *Config {
	return &c
}
