# Локальное использование

Скачайте предварительно скомпилированные бинарные файлы в https://github.com/pybuild-org/pybuild/releases

Запуск `pybuild` по умолчанию использует `target.xml` в качестве скрипта сборки

С помощью `pybuild custom.xml` задайте пользовательский скрипт сборки

При использовании пользовательского скрипта суффикс `.xml` может быть опущен

# Использование в Github Action

```yaml
- name: setup pybuild
  uses: pybuild-org/pybuild@main
  with: # optional
      version: 'latest' # default
      goos: 'linux' # default
      goarch: 'amd64' # default
```

Затем команда `pybuild` (в Windows это `pybuild.exe`) будет доступна в последующих задачах

Далее см. [использование тегов](./tag-usage.md)