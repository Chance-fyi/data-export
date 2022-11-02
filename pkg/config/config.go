package config

import (
	"data-export/pkg/console"
	"data-export/pkg/file"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
)

var cfg *viper.Viper

func Init() {
	dirname := "config"
	cfg = viper.New()

	viper.AddConfigPath(dirname)

	for _, cf := range readConfigFile(dirname) {
		viper.SetConfigType(cf["ext"])
		viper.SetConfigName(cf["name"])
		err := viper.ReadInConfig()
		console.ExitIf(err)
		s := viper.AllSettings()
		if len(s) > 0 {
			cfg.Set(cf["name"], s)
		}
		configChange(viper.GetViper())
	}

	setEnvVariables(cfg)
}

//读取目录下的配置文件
func readConfigFile(dirname string) (configFile []map[string]string) {
	dir, err := ioutil.ReadDir(dirname)
	console.ExitIf(err)

	for _, fileInfo := range dir {
		fileName := fileInfo.Name()
		ext := path.Ext(fileName)

		configFile = append(configFile, map[string]string{
			"name": fileName[0 : len(fileName)-len(ext)],
			"ext":  strings.Trim(ext, "."),
		})
	}
	return
}

func configChange(v *viper.Viper) {
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fileName := filepath.Base(in.Name)
		ext := path.Ext(fileName)
		name := fileName[0 : len(fileName)-len(ext)]
		v.SetConfigType(strings.Trim(ext, "."))
		v.SetConfigName(name)
		s := v.AllSettings()
		if len(s) > 0 {
			cfg.Set(name, s)
		}
	})
}

func setEnvVariables(cfg *viper.Viper) {
	if !file.Exists(".env") {
		return
	}
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	console.ExitIf(err)

	for _, key := range viper.AllKeys() {
		cfg.Set(key, viper.Get(key))
	}
}
