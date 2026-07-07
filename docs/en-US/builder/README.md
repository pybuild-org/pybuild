# builder configuration

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

`app` is the application name

`source` is the source code directory, which will be fully retained in the build artifacts

`output` is the output directory for the build artifacts

`version` is the Python interpreter version

`release` is the Python interpreter build version

`version` and `release` should be written according to https://github.com/astral-sh/python-build-standalone/releases

# Local python configuration

```xml
<config name="python">
    <prop name="arch" value="" />
    <prop name="os" value="" />
</config>

<run job="setup python" />
```

`arch` is the architecture of the local Python interpreter

`os` is the operating system of the local Python interpreter

`arch` and `os` should be written according to https://github.com/astral-sh/python-build-standalone/releases

[standalone builder configuration](./standalone.md)

[docker builder configuration](./docker.md)