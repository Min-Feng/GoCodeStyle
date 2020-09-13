package mock

import "ddd/pkg/assistant/configs"

var Config = configs.NewLocalRepo("dev").Find()
