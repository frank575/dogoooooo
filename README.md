# 狗語言練習範例
```command
# 生成 README.md 指令
go run file/readme-generate/main.go
```
## 項目結構

- [file](./file) 
  - [concurrency-file](./file/concurrency-file) 
    - [write](./file/concurrency-file/write) 
      - [chan](./file/concurrency-file/write/chan) 併發(channel 版本)寫入檔案
      - [wait-group](./file/concurrency-file/write/wait-group) 併發(waitGroup 版本)寫入檔案
    - [write-and-read](./file/concurrency-file/write-and-read) 併發寫入隨機檔案數的檔案值(+1)
  - [readme-generate](./file/readme-generate) 生成 README 目錄腳本
  - [simple](./file/simple) 
    - [read-file](./file/simple/read-file) 基本讀取檔案
    - [write-file](./file/simple/write-file) 基本寫入檔案
- [reg-exp](./reg-exp) 
  - [1](./reg-exp/1) 沒什麼內容的正則範例
