# Usage

## Local Run

Download the precompiled binary from https://github.com/pybuild-org/pybuild/releases

Run `pybuild` to start compiling `target.xml`

Specify the configuration file `pybuild custom.xml`

## GitHub Action Run

```yaml
- name: setup pybuild
  uses: pybuild-org/pybuild@main
  with: # optional
      version: 'latest' # default
      goos: 'linux' # default
      goarch: 'amd64' # default
```

Then the `pybuild` command is available in subsequent jobs.