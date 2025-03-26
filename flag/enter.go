package flag

import sysflag "flag"

type Option struct {
	DB   bool
	User string //-u admin,-u user
}

// Parse 解析命令行参数
func Parse() Option {
	db := sysflag.Bool("db", false, "初始化数据库")
	user := sysflag.String("u", "", "创建用户")
	sysflag.Parse()
	return Option{
		DB:   *db,
		User: *user,
	}
}

// IsWbeStop 是否停止web项目
func IsWbeStop(option Option) bool {
	if option.DB {
		return true
	}
	if option.User != "" {
		return true
	}
	return false
}

// SwitchOption 根据命令执行不同函数
func SwitchOption(option Option) {
	if option.DB {
		MakeMigration()
		return
	}
	if option.User == "admin" || option.User == "user" {
		CreateUser(option.User)
		return
	}
}
