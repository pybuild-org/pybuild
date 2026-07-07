# Docker Builder Configuration

The Docker builder can produce tar files that conform to the Docker image format without requiring the Docker CLI on the system, supports multiple architectures, and the build target only supports the Linux platform.

The Docker builder depends on builder configuration and local Python configuration.

```xml
<config name="docker image meta">
    <prop name="tag" value="" />
</config>
```

`tag` is the Docker image tag

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

If you need to build multiple targets, the `docker image targets` configuration should be written multiple times.

`image.base` is the base image of the Docker image.

`image.arch` is the runtime architecture of the Docker image.

`python.arch` is the target Python interpreter's architecture.

`python.os` is the target Python interpreter's operating system.

`python.arch` and `python.os` should be written according to https://github.com/astral-sh/python-build-standalone/releases.

`pip.platform` should be written referencing https://packaging.python.org/en/latest/specifications/platform-compatibility-tags/#platform-tag.

`pip.download` is used to declare project dependencies; if there are multiple dependencies, multiple `pip.download` entries should be written.

`launcher.run` is the project entry point, which can be a specific Python script file or a module (e.g., `-m http.server`).