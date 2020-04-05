package model

import "math/rand"

type Config struct {
	Mysql []*Mysql `json:"mysql",yaml:"mysql"`
	Redis []*Redis `json:"redis",yaml:"redis"`
	Node  int64    `json:"node",yaml:"node"`
	Log *Log `json:"log",yaml:"node"`
}

type Mysql struct {
	Ip       string `json:"ip",yaml:"ip"`
	Port     string `json:"port",yaml:"port"`
	UserName string `json:"user_name",yaml:"username"`
	Password string `json:"password",yaml:"password"`
	DbName   string `json:"db_name",yaml:"db_name"`
	Network  string `json:"network",yaml:"network"`
}

type Redis struct {
	Ip      string `json:"ip",yaml:"ip"`
	Port    string `json:"port",yaml:"port"`
	Network string `json:"network",yaml:"network"`
	Password string `json:"password,yaml:"password"`
}

type Log struct {
	Stdout bool   `json:"stdout",yaml:"stdout"`
	Path   string `json:"path",yaml:"path"`
}

func (c *Config) GetMysql() *Mysql {
	index := rand.Int63n(int64(len(c.Mysql)))
	return c.Mysql[index]
}

func (c *Config) GetRedis() *Redis {
	index := rand.Int63n(int64(len(c.Mysql)))
	return c.Redis[index]
}
