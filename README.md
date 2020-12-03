# 狗語言學習歷程

- [我的筆記 - go 語法篇](https://hackmd.io/IrpAln1QQ4GsVW-_fW6nNA?view)

# 項目目錄結構

```command
# 使用以下指令生成項目目錄結構
go run file/readme-generate/main.go
```

<!--TOC-->
- [file](./file) 
  - [concurrency-file](./file/concurrency-file) 
    - [write](./file/concurrency-file/write) 
      - [chan](./file/concurrency-file/write/chan) 併發(channel 版本)寫入檔案
      - [wait-group](./file/concurrency-file/write/wait-group) 併發(waitGroup 版本)寫入檔案
    - [write-and-read](./file/concurrency-file/write-and-read) 併發寫入自訂檔案數的隨機檔案值(+1)
  - [readme-generate](./file/readme-generate) 生成 README 的項目目錄結構腳本
  - [simple](./file/simple) 
    - [read-file](./file/simple/read-file) 基本讀取檔案
    - [write-file](./file/simple/write-file) 基本寫入檔案
- [regexp](./regexp) 
  - [1](./regexp/1) 基本正則
<!--TOC-->

# 大項目介紹
