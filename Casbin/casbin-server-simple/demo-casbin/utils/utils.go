package utils

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	ENV     = "env"
	DEFAULT = "default"
)

// ReadConfig read config struct from environment
func ReadConfig(config interface{}) error {
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	// get value of interface
	v := reflect.ValueOf(config)
	// check if config is a pointer
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		field, typeField := v.Field(i), v.Type().Field(i)
		if field.CanSet() {
			// get os env and default value
			envName, defValue := typeField.Tag.Get(ENV), typeField.Tag.Get(DEFAULT)
			if envName == "" {
				envName = typeField.Name
			}

			switch field.Kind() {
			case reflect.String:
				if str, ok := os.LookupEnv(envName); ok {
					field.SetString(str)
				} else {
					field.SetString(defValue)
				}
			case reflect.Int, reflect.Int32, reflect.Int64:
				if str, ok := os.LookupEnv(envName); ok {
					if val, err := strconv.ParseInt(str, 10, 0); err == nil {
						field.SetInt(val)
					} else {
						return err
					}
				} else {
					if val, err := strconv.ParseInt(defValue, 10, 0); err == nil {
						field.SetInt(val)
					} else {
						return err
					}
				}
			case reflect.Uint, reflect.Uint32, reflect.Uint64:
				if str, ok := os.LookupEnv(envName); ok {
					if val, err := strconv.ParseUint(str, 10, 0); err == nil {
						field.SetUint(val)
					} else {
						return err
					}
				} else {
					if val, err := strconv.ParseUint(defValue, 10, 0); err == nil {
						field.SetUint(val)
					} else {
						return err
					}
				}
			case reflect.Float32, reflect.Float64:
				if str, ok := os.LookupEnv(envName); ok {
					if val, err := strconv.ParseFloat(str, 64); err == nil {
						field.SetFloat(val)
					} else {
						return err
					}
				} else {
					if val, err := strconv.ParseFloat(defValue, 64); err == nil {
						field.SetFloat(val)
					} else {
						return err
					}
				}
			case reflect.Bool:
				if str, ok := os.LookupEnv(envName); ok {
					if val, err := strconv.ParseBool(str); err == nil {
						field.SetBool(val)
					} else {
						return err
					}
				} else {
					if val, err := strconv.ParseBool(defValue); err == nil {
						field.SetBool(val)
					} else {
						return err
					}
				}
			}
		}
	}
	return nil
}
