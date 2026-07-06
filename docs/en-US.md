#Usage

## Local Execution

Download the precompiled binary from https://github.com/pybuild-org/pybuild/releases

Run `pybuild` to start compiling `target.xml`

Specify configuration file `pybuild custom.xml`

## GitHub Action Execution

```yaml
- name: setup pybuild
  uses: pybuild-org/pybuild@main
  with: # optional
      version: 'latest' # default
      goos: 'linux' # default
      goarch: 'amd64' # default
```

Then the `pybuild` command is available in subsequent tasks.

# Tag Usage

## xml tag

```xml
<xml></xml>
```

A useless placeholder that does not occupy the stack; it can be written any number of times, at any depth.

## use tag

```xml
<use file="" />
```

Imports an external XML snippet and immediately executes it in the current context.

The `file` parameter can be a file path or a URL starting with http(s)://.

If the `.xml` suffix is omitted, the `use` tag will automatically add it.

## config tag

```xml
<config [type="group"] name=""><config>
```

Used to declare a configuration; `name` declares the configuration name.

`type="group"` is optional, used to declare a group configuration; in this case, the `config` tag is declared collectively by several child `config` tags.

## prop tag

```xml
<prop name="" value="" />
```

Used to declare a specific configuration value; `name` declares the field name, `value` declares the field value.

## run tag

```xml
<run job="" | command="" />
```

Used to execute a task; `job` and `command` are mutually exclusive.

`job` declares the task name, used to execute a pre-configured task.

`command` declares the command content, used to execute a system command.

## log tag

```xml
<log></log>
```

Used to print a line of log.