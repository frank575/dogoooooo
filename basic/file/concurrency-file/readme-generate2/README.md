# 目錄生成器

1. 撈取所有檔案裡有 toc#.+#toc 規則的敘述
	- 可以帶入 ::: 來寫入自訂標籤，e.g. toc#file:::write#toc
		```txt
		// 生成的 markdown 代碼
		- **1. file** // 標籤名稱
			- **[1-1. write](檔案路徑)**
		```
	- 如果未帶 ::: 自動寫入 `無分類`
2. 若有 README.md 檔將生成目錄名為 `project` 的目錄，撈取規則為第一行 \# title 
    ```txt
    // 生成的 markdown 代碼
    - **⭐project** // 標籤名稱
        - **[1. 專案名稱](檔案路徑)**
    ```
3. 會強制過濾掉 .gitignore 下的所有檔案(如 node_modules, .idea 等等)
4. 可以客製化過濾檔案，只要在跟目錄下新建 `gtoc.txt` 寫入配置
    - ROOT_PATH 爬取的跟目錄
    - IGNORE= 要過濾的檔名
	```txt
	ROOT_PATH=basic
	IGNORE=readme-generate2/main.go
	```

# 如何使用

1. 使用 <\!--TOC--><\!--TOC--> 放在 README.md 你要生成目錄的地方
2. 點擊 `運行 main.go` 來生成項目目錄
