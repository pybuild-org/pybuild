# xml 標籤

```xml
<xml></xml>
```

沒有任何用處的佔位符，不佔用棧，可以寫任意多次，任意多層

# define 標籤

```xml
<define name="" value="" />
```

用於定義值別名的標籤，`name` 用於聲明值別名，`value` 用於聲明值

`value` 參數可以透過 `{name}` 使用之前的 define 標籤中定義過的值別名

# use 標籤

```xml
<use src="" />
```

引入一段外部的 xml 內容，並立即在當前上下文中執行

`src` 參數可以是一個檔案路徑，或者一個以 `http(s)://` 開頭的連結

`src` 參數可以透過 `{name}` 使用 define 標籤中定義過的值別名

如果不加 `.xml` 副檔名，`use` 標籤會自動加上

# config 標籤

```xml
<config [type="group"] name=""><config>
```

用於聲明配置的標籤，`name` 用於聲明配置名稱

`type="group"` 是可選的，用於聲明一個組配置，此時該 `config` 標籤由若干子 `config` 標籤共同聲明

[通用配置項](./common-config.md)

# prop 標籤

```xml
<prop name="" value="" />
```

用於聲明配置項具體值的標籤，`name` 用於聲明欄位名稱，`value` 用於聲明欄位值

# run 標籤

```xml
<run job="" | command="" />
```

用於執行任務的標籤，`job` 和 `command` 二選一

`job` 用於聲明任務名稱，用於執行一個預配置的任務

`command` 用於聲明命令內容，用於執行一個系統命令

# log 標籤

```xml
<log></log>
```

用於列印一行日誌

接下來看：[構建器配置](./builder)