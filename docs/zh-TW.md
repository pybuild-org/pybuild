#使用方法

## 本機執行

在 https://github.com/pybuild-org/pybuild/releases 中下載預編譯的二進位檔案

執行 `pybuild` 即可開始編譯 `target.xml`

指定配置檔案 `pybuild custom.xml`

## GitHub Action 執行

```yaml
- name: setup pybuild
  uses: pybuild-org/pybuild@main
  with: # optional
      version: 'latest' # default
      goos: 'linux' # default
      goarch: 'amd64' # default
```

然後 `pybuild` 命令在之後的任務中可用