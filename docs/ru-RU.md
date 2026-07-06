#Использование

## Локальный запуск

Скачайте предварительно скомпилированный бинарный файл из https://github.com/pybuild-org/pybuild/releases

Запуск `pybuild` начнёт компиляцию `target.xml`

Укажите файл конфигурации `pybuild custom.xml`

## Запуск в GitHub Actions

```yaml
- name: setup pybuild
  uses: pybuild-org/pybuild@main
  with: # optional
      version: 'latest' # default
      goos: 'linux' # default
      goarch: 'amd64' # default
```

Затем команда `pybuild` будет доступна в последующих задачах