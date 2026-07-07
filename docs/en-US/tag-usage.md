# xml tag

```xml
<xml></xml>
```

A placeholder with no purpose, does not occupy the stack, can be written any number of times, any number of layers.

# use tag

```xml
<use file="" />
```

Introduces an external XML content and executes it immediately in the current context.

The `file` parameter can be a file path or a link starting with `http(s)://`.

If the `.xml` suffix is omitted, the `use` tag will automatically add it.

# config tag

```xml
<config [type="group"] name=""><config>
```

Tag used to declare a configuration; `name` declares the configuration name.

`type="group"` is optional, used to declare a group configuration; in this case the `config` tag is declared together with several child `config` tags.

# prop tag

```xml
<prop name="" value="" />
```

Tag used to declare the specific value of a configuration item; `name` declares the field name, `value` declares the field value.

# run tag

```xml
<run job="" | command="" />
```

Tag used to execute a task; `job` and `command` are mutually exclusive.

`job` declares the task name to execute a predefined task.

`command` declares the command content to execute a system command.

# log tag

```xml
<log></log>
```

Tag used to print a line of log.

Next, see: [Builder Configuration](./builder)