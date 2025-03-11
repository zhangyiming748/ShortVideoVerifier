package sql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)
var engine *xorm.Engine

func GetMysql()*xorm.Engine{
	return engine
}

func InitMySQLEngine() (*xorm.Engine, error) {
	// 不指定数据库名称进行连接
	dsn := fmt.Sprintf("root:123456@tcp(127.0.0.1:3306)/")
	var err error
	engine, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("连接MySQL失败: %v", err)
	}

	// 检查数据库是否存在
	exists, err := isDatabaseExists(engine, "test")
	if err != nil {
		return nil, fmt.Errorf("检查数据库存在失败: %v", err)
	}

	// 如果数据库不存在，创建数据库
	if !exists {
		if err := createDatabase(engine); err != nil {
			return nil, fmt.Errorf("创建数据库失败: %v", err)
		}
	}

	// 重新连接，这次指定数据库
	engine.Close()
	dsn = fmt.Sprintf("root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4")
	engine, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("重新连接MySQL失败: %v", err)
	}

	// 设置连接池
	engine.SetMaxIdleConns(5)
	engine.SetMaxOpenConns(10)

	return engine, nil
}

func isDatabaseExists(engine *xorm.Engine, dbName string) (bool, error) {
	sql := "SELECT SCHEMA_NAME FROM information_schema.SCHEMATA WHERE SCHEMA_NAME = ?"
	results, err := engine.Query(sql, dbName)
	if err != nil {
		return false, err
	}
	return len(results) > 0, nil
}

func createDatabase(engine *xorm.Engine) error {
	_, err := engine.Exec("CREATE DATABASE test CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci")
	return err
}