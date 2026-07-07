# Конфигурация builder

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

`app` — название приложения  

`source` — каталог исходного кода, этот каталог будет полностью сохранён в артефактах сборки  

`output` — каталог вывода артефактов сборки  

`version` — версия интерпретатора Python  

`release` — версия сборки интерпретатора Python  

`version` и `release` оформляются согласно https://github.com/astral-sh/python-build-standalone/releases  

# Локальная конфигурация python

```xml
<config name="python">
    <prop name="arch" value="" />
    <prop name="os" value="" />
</config>

<run job="setup python" />
```

`arch` — архитектура, на которой работает локальный интерпретатор Python  

`os` — операционная система, на которой работает локальный интерпретатор Python  

`arch` и `os` оформляются согласно https://github.com/astral-sh/python-build-standalone/releases  

[Конфигурация standalone‑строителя](./standalone.md)

[Конфигурация docker‑строителя](./docker.md)