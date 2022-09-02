package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type Config struct {
	LyricsPath string `env:"LYRCS_PATH"`
}

func BuildFromEnv() *Config {
	c := Config{}
	x := reflect.TypeOf(c)
	fields := reflect.VisibleFields(x)
	values := reflect.ValueOf(&c)
	for _, field := range fields {
		envKey := field.Tag.Get("env")
		envValue, ok := os.LookupEnv(envKey)
		if ok {
			rValue := reflect.Indirect(values).FieldByName(field.Name)
			fmt.Println(rValue)
			if rValue.Kind() == reflect.String {

				rValue.SetString(envValue)
			}
			if rValue.Kind() == reflect.Int {
				envInt, err := strconv.Atoi(envValue)
				if err != nil {
					panic(err)
				}
				rValue.SetInt(int64(envInt))
			}
		}
	}
	return &c
}
