package main

import (
	"work_diary/config"
)

func main() {
	conf, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	_, err = config.LoadDatabase(conf)
	if err != nil {
		panic(err)
	}

}
