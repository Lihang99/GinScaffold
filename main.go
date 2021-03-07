package main

import (
	"GinScaffold/dao/mysql"
	"GinScaffold/dao/redis"
	"GinScaffold/logger"
	"GinScaffold/router"
	"GinScaffold/settings"
	"GinScaffold/utils/snowflake"
	"fmt"
)

func main() {
	//修改所需要的配置文件
	if err := settings.Init("./conf/dev.yaml"); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	if settings.Conf.MySQLConfig.Enable {
		if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
			fmt.Printf("init mysql failed, err:%v\n", err)
			return
		}
	}
	defer mysql.Close() // 程序退出关闭数据库连接
	if settings.Conf.MySQLConfig.Enable {
		if err := redis.Init(settings.Conf.RedisConfig); err != nil {
			fmt.Printf("init redis failed, err:%v\n", err)
			return
		}
	}
	defer redis.Close()
	if settings.Conf.SnowFlakeConfig.Enable {
		if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
			fmt.Printf("Init id gennerator failed,err:%v\n", err)
			return
		}
	}
	r := router.SetupRouter(settings.Conf.Mode)
	err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		fmt.Printf("Run server failed , err : %v\n", err)
		return
	}
}
