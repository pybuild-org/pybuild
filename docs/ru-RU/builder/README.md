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

`app` — имя приложения  

`source` — каталог исходного кода, этот каталог будет полностью сохранён в артефактах сборки  

`output` — каталог вывода артефактов сборки  

`version` — версия интерпретатора Python  

`release` — версия сборки интерпретатора Python  

`version` и `release` указываются согласно https://github.com/astral-sh/python-build-standalone/releases

# Локальная конфигурация Python

```xml
<config name="python">
    <prop name="arch" value="" />
    <prop name="os" value="" />
</config>

<run job="setup python" />
```

`arch` — архитектура локального интерпретатора Python  

`os` — операционная система локального интерпретатора Python  

`arch` и `os` указываются согласно https://github.com/astral-sh/python-build-standalone/releases

[Конфигурация standalone‑сборщика](./standalone.md)

[docker‑конфигурация сборщика](./docker.md)