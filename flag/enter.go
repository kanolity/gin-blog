package flag

import sys_flag "flag"

type Option struct {
	DB bool
}

// Parse 解析命令行参数
func Parse() Option {
	db := sys_flag.Bool("db", false, "初始化数据库")
	sys_flag.Parse()
	return Option{*db}
}

// IsWbeStop 是否停止web项目
func IsWbeStop(option Option) bool {
	if option.DB {
		return true
	}
	return false
}

// SwitchOption 根据命令执行不同函数
func SwitchOption(option Option) {
	if option.DB {
		MakeMigration()
	}
}
