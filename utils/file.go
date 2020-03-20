package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
)

func IsExist(file string) bool {
	_, err := os.Stat(file)
	return err == nil || os.IsExist(err)
}

func EnsureNestDir(dir string){
	if ok := IsExist(dir); ok == false {
		os.MkdirAll(dir, os.ModePerm)
	}
}

func DeleteFile(dir string) {
	if ok := IsExist(dir); ok == true {
		os.RemoveAll(dir)
	}
}

func RootPath() string {
	pwd, _ :=  os.Getwd()
	return pwd
}

func SaveFileToJson(item interface{}, dest string) error {
	jsonf, err := json.MarshalIndent(item, "", "\t")
	if err != nil {
		log.Fatal("saveJsonToFile", err)
		return err
	}

	file, err := os.Create(dest)
	defer file.Close()

	if err != nil {
		log.Fatal("saveJsonToFile", err)
		return err
	}

	buf := bytes.NewBuffer(jsonf)
	_, err = io.Copy(file, buf)
	return err
}