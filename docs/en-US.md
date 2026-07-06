# Usage

## Local Run

Download the precompiled binary from https://github.com/pybuild-org/pybuild/releases

Run `pybuild` to start compiling `target.xml`

Specify config file `pybuild custom.xml`

## GitHub Action Run

```yaml
- name: setup pybuild
  uses: pybuild-org/pybuild@main
  with: # optional
      version: 'latest' # default
      goos: 'linux' # default
      goarch: 'amd64' # default
```

Then the `pybuild` command is available in subsequent tasks

# Tag Usage

## xml Tag

```xml
<xml></xml>
```

Placeholder with no use, does not occupy the stack, can be written any number of times, any nesting level

## use Tag

```xml
<use file="" />
```

Import a piece of external XML content and execute immediately in the current context

The `file` parameter can be a file path, or a link starting with `http(s)://`

If the `.xml` suffix is omitted, the `use`file`` parameter can be a file path, or a link starting with `http(s)://`

If the `.xml` suffix is omitted, the `use` tag will automatically add it

## config Tag

```xml
<config [type="group"] name=""><config>
```

Tag used to declare configuration, `name` used to declare configuration name

`type="group"` is optional, used to declare a group configuration, in which case this `config` tag is declared together by several child `config` tags

## prop Tag

```xml
<prop name="" value="" />
```

Tag used to declare the specific value of a configuration item, `name` used to declare field name, `value` used to declare field value

## run Tag

```xml
<run job="" | command="" />
```

Tag used to execute a task, `job` and `command` are mutually exclusive

`job` used to declare task name, used to execute a pre‑configured task

`command` used to declare command content, used to execute a system command

## log Tag

```xml
<log></log>
```

Tag used to print a line of log