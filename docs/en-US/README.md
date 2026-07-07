# Local Use

Download the precompiled binaries from https://github.com/pybuild-org/pybuild/releases

Running `pybuild` uses `target.xml` as the build script by default

Use `pybuild custom.xml` to specify a custom build script

When using a custom build script, the `.xml` suffix may be omitted

# Using in Github Action

```yaml
- name: setup pybuild
  uses: pybuild-org/pybuild@main
  with: # optional
      version: 'latest' # default
      goos: 'linux' # default
      goarch: 'amd64' # default
```

Then the `pybuild` (on Windows it’s `pybuild.exe`) command is available in subsequent steps

Next, see: [Tag Usage](./tag-usage.md)