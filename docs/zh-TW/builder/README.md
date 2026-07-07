# builder 配置

```xml
<config name="builder">
    <prop name="app" value="" />
    <prop name="source" value="" />
    <prop name="output" value="" />
    <prop name="version" value="" />
    <prop name="release" value="" />
</config>

<run job="setup builder" />
```

`app` 是應用程式名稱

`source` 是原始碼目錄，這個目錄會在建置產物中完整保留

`output` 是建置產物輸出目錄

`version` 是 python 直譯器版本

`release` 是 python 直譯器建置版本

`version` 和 `release` 參照 https://github.com/astral-sh/python-build-standalone/releases 編寫

# 本地 python 配置

```xml
<config name="python">
    <prop name="arch" value="" />
    <prop name="os" value="" />
</config>

<run job="setup python" />
```

`arch` 是本地 python 直譯器執行架構

`os` 是本地 python 直譯器執行系統

`arch` 和 `os` 參照 https://github.com/astral-sh/python-build-standalone/releases 編寫

[standalone 構建器配置](./standalone.md)

[docker 構建器配置](./docker.md)