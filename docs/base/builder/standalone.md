# standalone 构建器配置

standalone 构建器可以构建出独立的可执行文件，支持多架构、跨平台，以 zip 压缩包格式生成产物

standalone 构建器依赖 builder 配置和本地 python 配置

```xml
<config type="group" name="standalone targets">
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

<run job="build standalone" />
```

如果需要构建多个目标，则 `standalone targets` 需要编写多次

`arch` 是目标 python 解释器运行架构

`os` 是目标 python 解释器运行系统

`arch` 和 `os` 参照 https://github.com/astral-sh/python-build-standalone/releases 编写

`platform` 参考 https://packaging.python.org/en/latest/specifications/platform-compatibility-tags/#platform-tag 编写

`download` 用于声明项目依赖，如果由多个依赖，则需编写多个 `download`

`run` 是项目入口点，可以是一个具体的 python 脚本文件或一个模块 (如 `-m http.server`)
