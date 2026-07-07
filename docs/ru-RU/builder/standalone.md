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

Если необходимо собрать несколько целей, то `standalone targets` следует писать несколько раз.

`python.arch` — архитектура целевого интерпретатора python.

`python.os` — операционная система целевого интерпретатора python.

`python.arch` и `python.os` следует указывать, ориентируясь на https://github.com/astral-sh/python-build-standalone/releases.

`pip.platform` следует указывать согласно https://packaging.python.org/en/latest/specifications/platform-compatibility-tags/#platform-tag.

`pip.download` используется для указания зависимостей проекта; если зависимостей несколько, необходимо добавить несколько `pip.download`.

`launcher.run` — точка входа проекта, может быть конкретным python‑скриптом или модулем (например, `-m http.server`).