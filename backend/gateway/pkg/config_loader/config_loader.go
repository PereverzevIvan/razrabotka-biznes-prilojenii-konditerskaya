package config_loader

import (
	"os"
	"reflect"

	"github.com/ilyakaznacheev/cleanenv"
)

// Функция для парсинга конфигурации из файла
func MustLoad(cfg_path string, cfg interface{}) {
	if reflect.ValueOf(cfg).Kind() != reflect.Ptr {
		panic("config must be a pointer")
	}

	// path := fetchConfigPath(cfg_path) // Считываем путь до конфига из командной строки
	if cfg_path == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(cfg_path); os.IsNotExist(err) { // Проверяем существование файла по данному пути
		panic(err)
	}

	cleanenv.ReadConfig(cfg_path, cfg) // Парсим информацию из файла в структуру через специальную библиотеку
}

// Функция-обёртка для парсинга конфигурации из командной строки
// Вызывает функцию FetchCmdParamValue() для получения пути до конфига
// Вызывает функцию MustLoad() для парсинга конфигурации
func MustLoadFromCmd(param_name string, cfg interface{}) {
	if param_name == "" {
		panic("param name is empty")
	}

	cfg_path := configloader_utils.FetchCmdParamValue(param_name)
	if cfg_path == "" {
		panic("config path is empty")
	}

	MustLoad(cfg_path, cfg)
}
