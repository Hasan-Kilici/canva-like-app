package utils

import "github.com/go-ini/ini"

func GetConfig(section, key string) string {
    cfg, err := ini.Load("config.ini")
    if err != nil {
        panic(err)
    }

    return cfg.Section(section).Key(key).String()
}
