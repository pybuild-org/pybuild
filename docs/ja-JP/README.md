# ローカルでの使用

https://github.com/pybuild-org/pybuild/releases から事前にコンパイルされたバイナリファイルをダウンロードします

`pybuild` を実行すると、デフォルトで `target.xml` がビルドスクリプトとして使用されます

`pybuild custom.xml` でビルドスクリプトをカスタマイズします

カスタムビルドスクリプトを使用する場合、`.xml` 拡張子は省略可能です

# Github Action での使用

```yaml
- name: setup pybuild
  uses: pybuild-org/pybuild@main
  with: # optional
      version: 'latest' # default
      goos: 'linux' # default
      goarch: 'amd64' # default
```

その後、`pybuild`（Windows では `pybuild.exe`）コマンドが以降のジョブで使用可能になります

次は、[タグの使用](./tag-usage.md) をご覧ください