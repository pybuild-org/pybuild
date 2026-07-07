# 本地使用

在 https://github.com/pybuild-org/pybuild/releases 中下載預編譯的二進位檔案

執行 `pybuild` 預設使用 `target.xml` 作為建置腳本

透過 `pybuild custom.xml` 自訂建置腳本

使用自訂建置腳本時，`.xml` 副檔名允許省略

# 在 Github Action 中使用

```yaml
- name: setup pybuild
  uses: pybuild-org/pybuild@main
  with: # optional
      version: 'latest' # default
      goos: 'linux' # default
      goarch: 'amd64' # default
```

然後 `pybuild`（Windows 上為 `pybuild.exe`）指令在之後的任務中可用

接下來看：[標籤使用](./tag-usage.md)