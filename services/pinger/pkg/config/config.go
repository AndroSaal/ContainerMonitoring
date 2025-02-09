package config

import (
	"errors"
	"flag"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type ServiceConfig struct {
	SrvConf ServerConfig
	Env     string `yaml:"env" env-default:"local"`
}

// конфигурация сервера
type ServerConfig struct {
	Port    string        `yaml:"port"`
	Host    string        `yaml:"host"`
	Timeout time.Duration `yaml:"timeout"`
	Env     string        `yaml:"env" env-default:"local"`
}

func MustLoadConfig() ServiceConfig {
	fi := "config.MustLoadConfig"

	//загружаем переменные окружения
	err := LoadEnv()
	if err != nil {
		log.Fatal(fi + ": " + err.Error())
	}

	//путь до файла конфигурации
	pathToConfDir, nameOfConfFile, err := getConfigLocation()
	if err != nil {
		log.Fatal(fi + ": " + err.Error())
	}

	//проверяем существует ли такие директория и файл
	if _, err := os.Stat(pathToConfDir + "/" + nameOfConfFile); os.IsNotExist(err) {
		log.Fatal(fi + ": " + err.Error())
	}

	//загружаем конфигурацию
	UserConf, err := LoadConfig(pathToConfDir, nameOfConfFile)
	if err != nil {
		log.Fatal(fi + ": " + err.Error())
	}

	return *UserConf

}

// MustLoadEnv загружает переменные окружения из файла .env,
// возвращает установленное окружение (local/dev/prod)
func LoadEnv() error {
	fi := "config.MustLoadEnv"

	//путь до файла .env
	if err := godotenv.Load("config/.env"); err != nil {
		log.Printf(fi + ": " + err.Error())
		return err
	}
	return nil
}

func getConfigLocation() (string, string, error) {
	fi := "config.getConfigLocation"

	//загрузка пути к директории с файлами конфигурции и имени файла из argv
	pathToConfDir, nameOfConfFile := getConfLocationFromArgv()

	//если имя директории - пустая строка, пробуем взять его из переменных окружения
	if pathToConfDir == "" {

		pathToConfDir = os.Getenv("CONFIG_DIR")

		if pathToConfDir == "" {
			return "", "", errors.New(fi + ": " + "pathToConfDir is empty at argv and env")
		}
	}

	//если имя файла - пустая строка, пробуем взять его из переменных окружения
	if nameOfConfFile == "" {

		nameOfConfFile = os.Getenv("CONFIG_FILE")

		if nameOfConfFile == "" {
			return "", "", errors.New(fi + ": " + "nameOfConfFile is empty at argv and env")
		}
	}

	return pathToConfDir, nameOfConfFile, nil
}

func getConfLocationFromArgv() (string, string) {

	var (
		pathToConfDir  string
		nameOfConfFile string
	)

	flag.StringVar(&pathToConfDir, "config_path", "", "name of directory with configs")
	flag.StringVar(&nameOfConfFile, "config_file", "", "name of config file")
	flag.Parse()

	return pathToConfDir, nameOfConfFile
}

func LoadConfig(path string, name string) (*ServiceConfig, error) {
	fi := "config.LoadConfig"

	var (
		srvConf ServerConfig
	)

	//инициализируем имя, папку и тип конфига
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.New(fi + ": " + err.Error())
	}

	//заполняем структуру сервера
	if err := viper.UnmarshalKey("server", &srvConf); err != nil {
		return nil, err
	}
	srvConf.Env = os.Getenv("ENVIRONMENT")
	if srvConf.Env == "" {
		log.Default().Println("ENVIRONMENT is not set, make it defaul: local")
		srvConf.Env = "local"
	}

	return &ServiceConfig{
		SrvConf: srvConf,
		Env:     srvConf.Env,
	}, nil

}
