# xml Tag

```xml
<xml></xml>
```

A placeholder with no purpose, does not occupy the stack, can be written any number of times and any number of layers.

# define Tag

```xml
<define name="" value="" />
```

Tag used to define value aliases; `name` declares the alias name, `value` declares the value.

# use Tag

```xml
<use src="" />
```

Imports an external XML content and executes it immediately in the current context.

`src` parameter can be a file path or a link starting with `http(s)://`.

`src` parameter can use a value alias defined by the define tag via `{name}`.

If the `.xml` suffix is omitted, the `use` tag will automatically add it.

# config Tag

```xml
<config [type="group"] name=""><config>
```

Tag used to declare a configuration; `name` declares the configuration name.

`type="group"` is optional and declares a group configuration, in which case this `config` tag is composed of several child `config` tags.

# prop Tag

```xml
<prop name="" value="" />
```

Tag used to declare a specific configuration item value; `name` declares the field name, `value` declares the field value.

# run Tag

```xml
<run job="" | command="" />
```

Tag used to execute a task; choose either `job` or `command`.

`job` declares the task name to execute a predefined task.

`command` declares the command content to execute a system command.

# log Tag

```xml
<log></log>
```

Tag used to print a line of log.

Next, see: [Builder Configuration](./builder)