#使用方法

## 本地运行

在 https://github.com/pybuild-org/pybuild/releases 中下载预编译的二进制文件

运行 `pybuild` 即可开始编译 `target.xml`

指定配置文件 `pybuild custom.xml`

## Github Action 运行

```yaml
- name: setup pybuild
  uses: pybuild-org/pybuild@main
  with: # optional
      version: 'latest' # default
      goos: 'linux' # default
      goarch: 'amd64' # default
```

然后 `pybuild` 命令在之后的任务中可用

# 标签用法

## xml 标签

```xml
<xml></xml>
```

没有任何用处的占位符，不占用栈，可以写任意多次，任意多层

## use 标签

```xml
<use file="" />
```

引入一段外部的 xml 内容，并立即在当前上下文中执行

`file` 参数可以是一个文件路径，或者一个以 http(s):// 开头的链接

如果不加 `.xml` 后缀，`use` 标签会自动加上

## config 标签

```xml
<config [type="group"] name=""><config>
```

用于声明配置的标签，`name` 用于声明配置名称

`type="group"` 是可选的，用于声明一个组配置，此时该 `config` 标签由若干子 `config` 标签共同声明

## prop 标签

```xml
<prop name="" value="" />
```

用于声明配置项具体值的标签，`name` 用于声明字段名称，`value` 用于声明字段值

## run 标签

```xml
<run job="" | command="" />
```

用于执行任务的标签，`job` 和 `command` 二选一

`job` 用于声明任务名称，用于执行一个预配置的任务

`command` 用于声明命令内容，用于执行一个系统命令

## log 标签

```xml
<log></log>
```

用于打印一行日志