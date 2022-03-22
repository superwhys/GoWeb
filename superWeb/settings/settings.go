package settings

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/superwhys/superLog"
)

func InitSetting() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		superLog.Info("config has change ! ")
	})
	return nil
}
