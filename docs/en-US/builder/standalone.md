# Standalone Builder Configuration

The standalone builder can produce independent executable files, supporting multiple architectures and cross‑platform builds, and generates artifacts in a zip archive format.

The standalone builder depends on the builder configuration and the local Python configuration.

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

If you need to build multiple targets, the `standalone targets` section must be written multiple times.

`python.arch` is the architecture on which the target Python interpreter runs.

`python.os` is the operating system on which the target Python interpreter runs.

Write `python.arch` and `python.os` according to the information on https://github.com/astral-sh/python-build-standalone/releases.

`pip.platform` should be written based on the specifications at https://packaging.python.org/en/latest/specifications/platform-compatibility-tags/#platform-tag.

`pip.download` is used to declare project dependencies; if there are multiple dependencies, write multiple `pip.download` entries.

`launcher.run` is the entry point of the project and can be a specific Python script file or a module (e.g., `-m http.server`).