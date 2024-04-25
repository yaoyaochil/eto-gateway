package initialize

import (
	"gateway/config"
	"gateway/global"
	"gateway/initialize/internal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GormMysql initializes multiple Mysql databases and returns a map of them
func GormMysql() map[string]*gorm.DB {
	dbs := make(map[string]*gorm.DB)

	// Initialize GormMysqlDTaiwan
	m := global.GatewayConf.GormMysqlDTaiwan
	if m.Dbname != "" {
		db := initializeDB(m)
		if db != nil {
			dbs["d_taiwan"] = db
		}
	}

	// Initialize GormMysqlSystem
	m = global.GatewayConf.GormMysqlSystem
	if m.Dbname != "" {
		db := initializeDB(m)
		if db != nil {
			dbs["system"] = db
		}
	}

	return dbs
}

// initializeDB initializes a Mysql database with the given config
func initializeDB(m config.GormMysql) *gorm.DB {
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config(m.Dbname, m.Prefix, m.Singular)); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}
