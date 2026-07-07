# xml tag

```xml
<xml></xml>
```

A placeholder with no purpose, does not occupy the stack, and can be written any number of times and any number of layers.

# use tag

```xml
<use file="" />
```

Imports an external XML content segment and executes it immediately in the current context.

The `file` parameter can be a file path or a link that starts with `http(s)://`.

If the `.xml` suffix is omitted, the `use` tag will automatically add it.

# config tag

```xml
<config [type="group"] name=""><config>
```

A tag for declaring configurations; `name` specifies the configuration name.

`type="group"` is optional and declares a group configuration, in which case this `config` tag is defined together with several child `config` tags.

# prop tag

```xml
<prop name="" value="" />
```

A tag for declaring the specific value of a configuration item; `name` specifies the field name, and `value` specifies the field value.

# run tag

```xml
<run job="" | command="" />
```

A tag for executing tasks; choose either `job` or `command`.

`job` declares the task name to run a predefined task.

`command` declares the command content to execute a system command.

# log tag

```xml
<log></log>
```

Prints a single line of log.

Next, see: [Builder Configuration](./builder)