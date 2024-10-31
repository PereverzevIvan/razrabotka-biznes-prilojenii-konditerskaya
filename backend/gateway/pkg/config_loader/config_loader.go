package config_loader

import (
	"errors"
	"os"
	"reflect"

	config_loader_utils "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/pkg/config_loader/utils"
	"github.com/ilyakaznacheev/cleanenv"
)

// Функция для парсинга конфигурации из файла
func Load(path_to_cfg string, cfg interface{}) error {
	if reflect.ValueOf(cfg).Kind() != reflect.Ptr {
		return errors.New("config must be a pointer")
	}

	// path := fetchConfigPath(cfg_path) // Считываем путь до конфига из командной строки
	if path_to_cfg == "" {
		return errors.New("config path is empty")
	}

	if _, err := os.Stat(path_to_cfg); os.IsNotExist(err) { // Проверяем существование файла по данному пути
		return err
	}

	cleanenv.ReadConfig(path_to_cfg, cfg) // Парсим информацию из файла в структуру через специальную библиотеку
	return nil
}

// Функция для парсинга конфигурации из файла
func MustLoad(path_to_cfg string, cfg interface{}) {
	err := Load(path_to_cfg, cfg)
	if err != nil {
		panic(err)
	}
}

// Функция-обёртка для парсинга конфигурации из командной строки
// Вызывает функцию FetchCmdParamValue() для получения пути до конфига
// Вызывает функцию MustLoad() для парсинга конфигурации
func MustLoadFromCmd(param_name string, cfg interface{}) {
	if param_name == "" {
		panic("param name is empty")
	}

	cfg_path := config_loader_utils.FetchCmdParamValue(param_name)
	if cfg_path == "" {
		panic("config path is empty")
	}

	MustLoad(cfg_path, cfg)
}
