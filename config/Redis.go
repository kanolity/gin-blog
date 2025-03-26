package config

import "strconv"

type Redis struct {
	IP       string `json:"ip" yaml:"ip"`
	Port     int    `json:"port" yaml:"port"`
	Password string `json:"password" yaml:"password"`
	PoolSize int    `json:"pool_size" yaml:"pool_size"`
}

func (r *Redis) Addr() string {
	return r.IP + ":" + strconv.Itoa(r.Port)
}
