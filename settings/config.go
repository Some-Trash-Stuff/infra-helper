package config

import (
	"encoding/json"
	"log"
	"os"
	"reflect"
	"strconv"
)

// Load carrega as configurações do JSON e sobrescreve com variáveis de ambiente
// T é o tipo de struct que será carregado (ex: AppSettings)
func Load[T any]() T {
	cfg := loadFromJSON[T]()
	overrideWithEnv(&cfg)

	return cfg
}

func loadFromJSON[T any]() T {
	file, err := os.Open("configs/AppSettings.json")
	if err != nil {
		log.Fatalf("[config] failed to open AppSettings.json: %v", err)
	}
	defer file.Close()

	var cfg T
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		log.Fatalf("[config] failed to parse AppSettings.json: %v", err)
	}

	return cfg
}

func overrideWithEnv(cfg interface{}) {
	v := reflect.ValueOf(cfg).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		// Se o campo é uma struct, processar recursivamente
		if field.Kind() == reflect.Struct {
			overrideWithEnv(field.Addr().Interface())
			continue
		}

		// Buscar a tag 'env'
		envTag := fieldType.Tag.Get("env")
		if envTag == "" {
			continue
		}

		// Buscar o valor da variável de ambiente
		envValue := os.Getenv(envTag)
		if envValue == "" {
			continue
		}

		// Setar o valor baseado no tipo do campo
		if !field.CanSet() {
			continue
		}

		switch field.Kind() {
		case reflect.String:
			field.SetString(envValue)
			log.Printf("[config] override %s from ENV: %s", fieldType.Name, envTag)

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if intVal, err := strconv.ParseInt(envValue, 10, 64); err == nil {
				field.SetInt(intVal)
				log.Printf("[config] override %s from ENV: %s", fieldType.Name, envTag)
			} else {
				log.Printf("[config] warning: failed to parse %s as int: %v", envTag, err)
			}

		case reflect.Bool:
			if boolVal, err := strconv.ParseBool(envValue); err == nil {
				field.SetBool(boolVal)
				log.Printf("[config] override %s from ENV: %s", fieldType.Name, envTag)
			} else {
				log.Printf("[config] warning: failed to parse %s as bool: %v", envTag, err)
			}

		case reflect.Float32, reflect.Float64:
			if floatVal, err := strconv.ParseFloat(envValue, 64); err == nil {
				field.SetFloat(floatVal)
				log.Printf("[config] override %s from ENV: %s", fieldType.Name, envTag)
			} else {
				log.Printf("[config] warning: failed to parse %s as float: %v", envTag, err)
			}

		default:
			log.Printf("[config] warning: unsupported type for env override: %s (%s)", fieldType.Name, field.Kind())
		}
	}
}
