# 狗語言學習歷程

- [我的筆記 - go 語法篇](https://hackmd.io/IrpAln1QQ4GsVW-_fW6nNA?view)

# 項目目錄結構

```command
# 使用以下指令生成項目目錄結構
go run file/concurrency-file/readme-generate2/main.go
```

<!--TOC-->
- **1. file**
  - [1-1. (舊)生成 README 目錄簡述](./file/readme-generate/main.go)
  - [1-2. 併發寫入自訂檔案數的隨機檔案值(+1)](./file/concurrency-file/write-and-read/main.go)
  - [1-3. 基本寫入檔案](./file/simple/write-file/main.go)
  - [1-4. (新)併發生成 README 目錄簡述](./file/concurrency-file/readme-generate2/main.go)
  - [1-5. 併發(channel 版本)寫入檔案](./file/concurrency-file/write/chan/main.go)
  - [1-6. 併發(waitGroup 版本)寫入檔案](./file/concurrency-file/write/wait-group/main.go)
  - [1-7. 基本讀取檔案](./file/simple/read-file/main.go)
- **2. reflect**
  - [2-1. 基本反射](./reflect/main.go)
- **3. regexp**
  - [3-1. 基本正則](./regexp/main.go)
<!--TOC-->

# 大項目介紹
