package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

// 配置文件路径 (包名 + 配置文件名 )
const USCardsConfigFile = "configs/account_cards_us.yaml"
const UkCardsConfigFile = "configs/account_cards_uk.yaml"
const QiNiuConfigFile = "configs/qiniu.yaml"

func NewSetting(configs ...string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	for _, config := range configs {
		if config != "" {
			vp.AddConfigPath(config)
		}
	}
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	s := &Setting{vp}
	s.WatchSettingChange()
	return s, nil
}
func (s *Setting) WatchSettingChange() {
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			_ = s.ReloadAllSection()
		})
	}()
}

func NewUSAccountCardsSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigFile(USCardsConfigFile)
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	s := &Setting{vp}
	s.WatchSettingChange()
	return s, nil
}

func NewUKAccountCardsSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigFile(UkCardsConfigFile)
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	s := &Setting{vp}
	s.WatchSettingChange()
	return s, nil
}

func NewQiNiuSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigFile(QiNiuConfigFile)
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	s := &Setting{vp}
	s.WatchSettingChange()
	return s, nil
}
