# time.Time 筆記.

## 寫入 到 資料庫
在 mysql 中, 若欄位型別為 datetime 且 允許 NULL  

- 當 go 的輸入值為 time.Time{} 空值  
    會返回錯誤  
    ```
    *mysql.MySQLError("Error 1292: Incorrect datetime value: '0000-00-00' for column 'created_date' at row 1")
    ```

- 當 go 的輸入值為有效的 time.Time   
    若用 DBeaver 工具查看  
    會看到資料庫存成 毫秒格式, 實際應該沒有毫秒  
    `2099-10-17 00:31:20.0`

## 從 資料庫 讀取
在 mysql 中, 若欄位型別為 datetime 且 允許 NULL  
使用 spew.Sdump 輸出資訊  

- 當 go 的欄位 為 time.Time 可得  
    ` CreatedDate: (time.Time) 2771-01-02 12:23:01 +0800 CST,`

- 當 go 的欄位 為 string 可得  
    ` CreatedDate: (string) (len=25) "2771-01-02T12:23:01+08:00",`
