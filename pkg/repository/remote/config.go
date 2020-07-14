package remote

import (
	"ddd/pkg/domain"
)

type configStore struct {
}

func NewConfigStore() *configStore {
	return &configStore{}
}

func (c *configStore) Find() (domain.Config, error) {
	// 假設利用 遠端配置中心 讀取資料, 產生 config 結構
	return domain.Config{}, nil
}
