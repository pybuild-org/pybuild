# 使用方法

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