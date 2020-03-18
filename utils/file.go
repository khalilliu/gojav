package utils

import "os"

func IsExist(file string) bool {
	_, err := os.Stat(file)
	return err == nil || os.IsExist(err)
}

func EnsureNestDir(dir string){
	if ok := IsExist(dir); ok == false {
		os.MkdirAll(dir, os.ModePerm)
	}
}