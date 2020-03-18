package config

import (
	"fmt"
	"reflect"
	"sync"
)

type Config struct {
	Parallel int
	Timeout  int64
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
	lock = sync.RWMutex{}
)

func (c *Config) Set (key string,  val interface{}) {
	lock.Lock()
	ps := reflect.ValueOf(c)
	s := ps.Elem()
	if s.Kind() == reflect.Struct{
		f := s.FieldByName(key)
		if f.IsValid() {
			if f.CanSet() {
				fmt.Println(f.Kind() == reflect.Int)
				switch f.Kind() {
				case reflect.Int:
					f.SetInt(int64(val.(int)))
				case reflect.String:
					f.SetString(val.(string))
				case reflect.Bool:
					f.SetBool(val.(bool))
				}
			}
		}
	}
	defer lock.Unlock()
}