package conf

const DriverName = "mysql"
const MasterDataSourceName = "root:186liuyuJQK@tcp(127.0.0.1:3306)/superstar?charset=utf8"

type DbConf struct {
	Host   string
	Port   int
	User   string
	Pwd    string
	DbName string
}

var MasterDbConfig DbConf = DbConf{
	Host:   "127.0.0.1:3306",
	Port:   3306,
	User:   "root",
	Pwd:    "root",
	DbName: "superstar",
}

var SlaveDbConfig DbConf = DbConf{
	Host:   "127.0.0.1",
	Port:   3306,
	User:   "root",
	Pwd:    "186liuyuJQK",
	DbName: "superstar",
}
