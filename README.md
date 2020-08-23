# 專案開發心得

## 啟動專案


```bash
CONF_SRC=local FILE_NAME=dev go run ./cmd/main.go
```

非運營環境才啟動 -race 檢測
```bash
CONF_SRC=local FILE_NAME=dev CGO_ENABLED=1 go run ./cmd/main.go -race
```

## 對全域變數的想法

狀態會改變的變數  
不應該放在全域值  
比如 sql.DB 內部狀態是一直在變化的  
若定為全域值, 可能因為其他函數的副作用, 導致未知的狀態變化  

使用依賴注入的方式, 從 struct 內容  
可以確實把握到底, 哪些變數有使用 sql.DB  
而不用深入 method or function 的細節

反過來說, 狀態不會變化的變數  
比如說 規則類 的變數  
就很適合放在全域值  
可以避免重複 malloc memory, 藉此增加效能  

## 確保專案的正確性

專案第一步  
先執行以下命令, 確保程式沒問題

```bash
golangci-lint run ./... && go test ./... -count=1 | grep -v "no test files"
```

Data Race Detector need to enable CGO
```bash
golangci-lint run ./... && CGO_ENABLED=1 go test ./... -race -count=1 | grep -v "no test files"
```

靜態分析工具 golangci-lint  
參考以下連結  
https://golangci-lint.run/  

程式碼中 出現 //noinspection  
這是 ide goland 的檢查功能  
跟靜態分析 linter 的用途差不多  

## 整合測試

進行真實連線的整合測試時  
要加上 tags flag  
如下方所示  

```bash
go test ./... -tags=integration -count=1 -v | grep -v "no test files"
```

## log level 的選擇
靠近程式入口的相關程式碼  
log 都用 Fatal level 以求程式剛運行時  
以求快速發現錯誤  
相關議題, 可參可以下連結  
https://blog.maddevs.io/how-to-start-with-logging-in-golang-projects-part-1-3e3a708b75be

## err 應該用什麼樣的形式返回

心裡有想法  
但目前這專案還沒使用  
目前會使用 https://github.com/morikuni/failure  
作為錯誤處理的套件  



## 為什麼目錄要這樣設定

以後找工作再寫...

