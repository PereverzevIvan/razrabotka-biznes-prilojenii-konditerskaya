package config_loader_utils

import (
	"flag"
)

// Функция для считывания пути до конфига, из значения флага командной строки
func FetchCmdParamValue(param_name string) string {
	var res string

	flag.StringVar(&res, param_name, "", "") // Сначала пробуем считать путь из консоли

	// if res == "" { // Если путь не был передан через кмнд. строку, то пробуем считать из переменной среды
	// 	res = os.Getenv("CONFIG_PATH")
	// }

	return res
}
