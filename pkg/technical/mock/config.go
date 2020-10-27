package mock

import "ddd/pkg/technical/configs"

var Config = configs.NewLocalRepo("dev").Find()
