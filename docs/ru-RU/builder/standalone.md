# standalone Конфигурация сборщика

standalone сборщик может создавать независимые исполняемые файлы, поддерживает несколько архитектур, кроссплатформенность и генерирует артефакты в виде zip‑архивов.

standalone сборщик зависит от конфигурации builder и локальной конфигурации python.

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

Если необходимо собрать несколько целей, то `standalone targets` нужно указать несколько раз.

`python.arch` — архитектура целевого интерпретатора Python.

`python.os` — операционная система целевого интерпретатора Python.

Значения `python.arch` и `python.os` следует задавать, ориентируясь на https://github.com/astral-sh/python-build-standalone/releases.

Для `pip.platform` используйте рекомендации из https://packaging.python.org/en/latest/specifications/platform-compatibility-tags/#platform-tag.

`pip.download` используется для указания зависимостей проекта; если зависимостей несколько, необходимо добавить несколько элементов `pip.download`.

`launcher.run` — точка входа проекта; это может быть конкретный python‑скрипт или модуль (например, `-m http.server`).