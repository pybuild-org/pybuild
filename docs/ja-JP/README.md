# ローカルでの使用

https://github.com/pybuild-org/pybuild/releases で事前コンパイルされたバイナリをダウンロード

`pybuild` を実行するとデフォルトで `target.xml` をビルドスクリプトとして使用

`pybuild custom.xml` でカスタムビルドスクリプトを指定

カスタムビルドスクリプトを使用する場合、`.xml` 拡張子は省略可能

# Github Action での使用

```yaml
- name: setup pybuild
  uses: pybuild-org/pybuild@main
  with: # optional
      version: 'latest' # default
      goos: 'linux' # default
      goarch: 'amd64' # default
```

その後、`pybuild`（Windows では `pybuild.exe`）コマンドが以降のジョブで使用可能

次は：[タグの使用](./tag-usage.md)