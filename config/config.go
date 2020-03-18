package config

type Config struct {
	Parallel int
	Timeout  int
	Limit    int
	Proxy    string
	Search   string
	Base     string
	Output   string
	Nomag    bool
	Allmag   bool
	Nopic    bool
	Caption  bool
}

var (
	Cfg  = Config{}

	BaseUrl  = "https://www.javbus.com"
	SearchRoute = "/search"
)