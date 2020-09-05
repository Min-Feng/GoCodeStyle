package mock

import "ddd/pkg/configs"

var Config = configs.NewLocalRepo("dev").Find()
