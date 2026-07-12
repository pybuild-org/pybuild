# xml tag

```xml
<xml></xml>
```

A placeholder with no effect, does not occupy the stack, can be written any number of times, any number of layers.

# define tag

```xml
<define name="" value="" />
```

Tag used to define value aliases, `name` declares the alias name, `value` declares the value.  
The `value` parameter can use previously defined aliases from a define tag via `{name}`.

# use tag

```xml
<use src="" />
```

Introduces an external XML content and executes it immediately in the current context.  
The `src` parameter can be a file path or a link starting with `http(s)://`.  
The `src` parameter can use value aliases defined by a define tag via `{name}`.  
If the `.xml` suffix is omitted, the `use` tag will automatically add it.

# config tag

```xml
<config [type="group"] name=""><config>
```

Tag used to declare a configuration, `name` declares the configuration name.  
`type="group"` is optional, used to declare a group configuration, in which case this `config` tag is composed of several child `config` tags.  

[General configuration items](./common-config.md)

# prop tag

```xml
<prop name="" value="" />
```

Tag used to declare a specific value for a configuration item, `name` declares the field name, `value` declares the field value.

# run tag

```xml
<run job="" | command="" />
```

Tag used to execute a task, choose either `job` or `command`.  
`job` declares the task name to run a predefined task.  
`command` declares the command content to execute a system command.

# log tag

```xml
<log></log>
```

Tag used to print a line of log.

Next, see: [Builder configuration](./builder)