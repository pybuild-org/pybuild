# Конфигурация Docker‑строителя

Docker‑строитель может создавать tar‑файлы в формате Docker‑image без необходимости наличия Docker CLI в системе, поддерживает мультиархитектуру, цель сборки поддерживает только платформу Linux.

Docker‑строитель зависит от конфигурации builder и локальной конфигурации Python.

```xml
<config name="docker image meta">
    <prop name="tag" value="" />
</config>
```

`tag` — тег Docker‑image.

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

Если необходимо собрать несколько целей, то `docker image targets` следует указать несколько раз.

`image.base` — базовый образ Docker‑image.

`image.arch` — архитектура, на которой будет работать Docker‑image.

`python.arch` — архитектура целевого интерпретатора Python.

`python.os` — операционная система целевого интерпретатора Python.

`python.arch` и `python.os` следует указывать согласно <https://github.com/astral-sh/python-build-standalone/releases>.

`pip.platform` следует указывать согласно <https://packaging.python.org/en/latest/specifications/platform-compatibility-tags/#platform-tag>.

`pip.download` используется для указания зависимостей проекта; если зависимостей несколько, необходимо добавить несколько `pip.download`.

`launcher.run` — точка входа проекта; может быть конкретным файлом скрипта Python или модулем (например, `-m http.server`).