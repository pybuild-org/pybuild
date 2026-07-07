# docker 構建器配置

docker 構建器可以構建出符合 docker image 格式的 tar 檔案，而無需系統上有 docker cli，支援多架構，建置目標僅支援 linux 平台

docker 構建器依賴 builder 配置和本地 python 配置

```xml
<config name="docker image meta">
    <prop name="tag" value="" />
</config>
```

`tag` 是 docker image 的標籤

```xml
<config type="group" name="docker image targets">
    <config name="image">
        <prop name="base" value="" />
        <prop name="arch" value="" />
    </config>
    <config name="python">
        <prop name="arch" value="" />
        <prop name="os" value="" />
    </config>
    <config name="pip">
        <prop name="platform" value="" />
        <prop name="download" value="" />
    </config>
    <config name="launcher">
        <prop name="run" value="" />
    </config>
</config>

<run job="docker build" />
```

如果需要構建多個目標，則 `docker image targets` 需要編寫多次

`image.base` 是 docker image 的底層映像

`image.arch` 是 docker image 的執行架構

`python.arch` 是目標 python 直譯器執行架構

`python.os` 是目標 python 直譯器執行系統

`python.arch` 和 `python.os` 參照 https://github.com/astral-sh/python-build-standalone/releases 編寫

`pip.platform` 參考 https://packaging.python.org/en/latest/specifications/platform-compatibility-tags/#platform-tag 編寫

`pip.download` 用於聲明專案依賴，如果有多個依賴，則需編寫多個 `pip.download`

`launcher.run` 是專案入口點，可以是一個具體的 python 腳本檔案或一個模組（如 `-m http.server`）