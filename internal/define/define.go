package define

var ConfigName = "redis-client.conf"

type Connection struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	Addr     string `json:"addr"`
	Port     string `json:"port"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type Config struct {
	Connections []*Connection `json:"connections"`
}
