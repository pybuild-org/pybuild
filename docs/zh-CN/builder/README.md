# 构建器配置

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

`app` 是应用程序名称

`source` 是源码目录，这个目录会在构建产物中完整保留

`output` 是构建产物输出目录

`version` 是 python 解释器版本

`release` 是 python 解释器构建版本

`version` 和 `release` 参照 https://github.com/astral-sh/python-build-standalone/releases 编写

# 本地 python 配置

```xml
<config name="python">
    <prop name="arch" value="" />
    <prop name="os" value="" />
</config>

<run job="setup python" />
```

`arch` 是本地 python 解释器运行架构

`os` 是本地 python 解释器运行系统

`arch` 和 `os` 参照 https://github.com/astral-sh/python-build-standalone/releases 编写

[standalone 构建器配置](./standalone.md)

[docker 构建器配置](./docker.md)