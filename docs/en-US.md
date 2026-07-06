# Usage

## Local Run

Download the precompiled binary from https://github.com/pybuild-org/pybuild/releases

Run `pybuild` to start compiling `target.xml`

Specify configuration file `pybuild custom.xml`

## Github Action Run

```yaml
- name: setup pybuild
  uses: pybuild-org/pybuild@main
  with: # optional
      version: 'latest' # default
      goos: 'linux' # default
      goarch: 'amd64' # default
```

Then the `pybuild` command is available in subsequent jobs

# Tag Usage

## xml tag

```xml
<xml></xml>
```

A placeholder with no effect, does not occupy the stack, can be written any number of times, any number of layers

## use tag

```xml
<use file="" />
```

Introduces a segment of external XML content and executes it immediately in the current context

The `file` parameter can be a file path or a link starting with http(s)://

If the `.xml` suffix is not added, the `use` tag will automatically add it

## config tag

```xml
<config [type="group"] name=""><config>
```

Tag used to declare configuration, `name` is used to declare the configuration name

`type="group"` is optional, used to declare a group configuration, in which case this `config` tag is declared together by several child `config` tags

## prop tag

```xml
<prop name="" value="" />
```

Tag used to declare the specific value of a configuration item, `name` is used to declare the field name, `value` is used to declare the field value

## run tag

```xml
<run job="" | command="" />
```

Tag used to execute a task, `job` and `command` are mutually exclusive

`job` is used to declare the task name, used to execute a pre-configured task

`command` is used to declare the command content, used to execute a system command

## log tag

```xml
<log></log>
```

Used to print a line of log