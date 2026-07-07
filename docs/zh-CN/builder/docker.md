# docker 构建器配置

docker 构建器可以构建出符合 docker image 格式的 tar 文件，而无需系统上有 docker cli，支持多架构，构建目标仅支持 linux 平台

docker 构建器依赖 builder 配置和本地 python 配置

```xml
<config name="docker image meta">
    <prop name="tag" value="" />
</config>
```

`tag` 是 docker image 的标签

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

如果需要构建多个目标，则 `docker image targets` 需要编写多次

`image.base` 是 docker image 的底层镜像

`image.arch` 是 docker image 的运行架构

`python.arch` 是目标 python 解释器运行架构

`python.os` 是目标 python 解释器运行系统

`python.arch` 和 `python.os` 参照 https://github.com/astral-sh/python-build-standalone/releases 编写

`pip.platform` 参考 https://packaging.python.org/en/latest/specifications/platform-compatibility-tags/#platform-tag 编写

`pip.download` 用于声明项目依赖，如果由多个依赖，则需编写多个 `pip.download`

`launcher.run` 是项目入口点，可以是一个具体的 python 脚本文件或一个模块 (如 `-m http.server`)