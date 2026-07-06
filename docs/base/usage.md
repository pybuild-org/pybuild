# 本地使用

在 https://github.com/pybuild-org/pybuild/releases 中下载预编译的二进制文件

运行 `pybuild` 默认使用 `target.xml` 作为构建脚本

通过 `pybuild custom.xml` 自定义构建脚本

# 在 Github Action 中使用

```yaml
- name: setup pybuild
  uses: pybuild-org/pybuild@main
  with: # optional
      version: 'latest' # default
      goos: 'linux' # default
      goarch: 'amd64' # default
```

然后 `pybuild` (Windows 上是 `pybuild.exe`) 命令在之后的任务中可用
