# 專案開發心得

## 事前準備

我還沒學怎麼撰寫 make 指令  
所以拿到專案第一步  
先執行以下命令, 確保程式應該沒問題

```bash
golangci-lint run ./... && go test ./... -count=1 | grep -v "no test files"
```

靜態分析工具 golangci-lint  
參考以下連結  
https://golangci-lint.run/  

程式碼中 出現 //noinspection  
這是 ide goland 的檢查功能  
跟靜態分析的用途差不多  

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

