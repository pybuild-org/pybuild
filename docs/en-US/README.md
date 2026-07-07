# Local Usage

Download the precompiled binaries from https://github.com/pybuild-org/pybuild/releases

Running `pybuild` uses `target.xml` as the build script by default

Use `pybuild custom.xml` to specify a custom build script

# Using in GitHub Action

```yaml
- name: setup pybuild
  uses: pybuild-org/pybuild@main
  with: # optional
      version: 'latest' # default
      goos: 'linux' # default
      goarch: 'amd64' # default
```

Then the `pybuild` command (on Windows it's `pybuild.exe`) is available in subsequent steps

Next, see: [Tag Usage](./tag-usage.md)