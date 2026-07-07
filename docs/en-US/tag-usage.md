# xml tag

```xml
<xml></xml>
```

A placeholder with no effect, does not occupy the stack, and can be written any number of times and at any depth.

# use tag

```xml
<use file="" />
```

Imports an external XML snippet and executes it immediately in the current context.

The `file` attribute can be a file path or a URL that starts with `http(s)://`.

If the `.xml` suffix is omitted, the `use` tag will automatically add it.

# config tag

```xml
<config [type="group"] name=""><config>
```

Declares a configuration tag; `name` specifies the configuration name.

`type="group"` is optional and indicates a group configuration, in which case this `config` tag is defined together with several child `config` tags.

# prop tag

```xml
<prop name="" value="" />
```

Declares a specific configuration item; `name` specifies the field name, and `value` specifies the field value.

# run tag

```xml
<run job="" | command="" />
```

Executes a task; choose either `job` or `command`.

`job` specifies the name of a predefined task to run.

`command` specifies the command content to execute as a system command.

# log tag

```xml
<log></log>
```

Prints a single line of log output.

Next, see: [Builder Configuration](./builder)