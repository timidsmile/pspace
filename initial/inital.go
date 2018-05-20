package initial

import (
	"github.com/timidsmile/pspace/components"
	"github.com/timidsmile/pspace/setting"
)

func InitDb() error {
	return components.InitDb(setting.Cfg)
}

func InitRedis() error {
	return components.InitRedis(setting.Cfg)
}
