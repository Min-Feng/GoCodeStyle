package local

import (
	"os"

	"ddd/pkg/domain"
)

type configStore struct {
	os.File
}

func NewConfigStore() *configStore {
	return &configStore{}
}

func (c *configStore) Find() (domain.Config, error) {
	// 假設利用 os.File 讀取本地檔案, 產生 config 結構
	return domain.Config{}, nil
}
