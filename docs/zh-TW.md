#使用方法

## 本地運行

在 https://github.com/pybuild-org/pybuild/releases 中下載預編譯的二進位檔案

執行 `pybuild` 即可開始編譯 `target.xml`

指定配置檔案 `pybuild custom.xml`

## Github Action 執行

```yaml
- name: setup pybuild
  uses: pybuild-org/pybuild@main
  with: # optional
      version: 'latest' # default
      goos: 'linux' # default
      goarch: 'amd64' # default
```

然後 `pybuild` 指令在之後的任務中可用

# 標籤用法

## xml 標籤

```xml
<xml></xml>
```

沒有任何用處的佔位符，不佔用棧，可以寫任意多次，任意多層

## use 標籤

```xml
<use file="" />
```

引入一段外部的 xml 內容，並立即在當前上下文中執行

`file` 參數可以是一個檔案路徑，或者一個以 `http(s)://` 開頭的連結

如果不加 `.xml` 副檔名，`use` 標籤會自動加上

## config 標籤

```xml
<config [type="group"] name=""><config>
```

用於聲明配置的標籤，`name` 用於聲明配置名稱

`type="group"` 是可選的，用於聲明一個組配置，此時該 `config` 標籤由若干子 `config` 標籤共同聲明

## prop 標籤

```xml
<prop name="" value="" />
```

用於聲明配置項具體值的標籤，`name` 用於聲明欄位名稱，`value` 用於聲明欄位值

## run 標籤

```xml
<run job="" | command="" />
```

用於執行任務的標籤，`job` 和 `command` 二選一

`job` 用於聲明任務名稱，用於執行一個預配置的任務

`command` 用於聲明指令內容，用於執行一個系統指令

## log 標籤

```xml
<log></log>
```

用於列印一行日誌