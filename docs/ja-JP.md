#使い方

## ローカルで実行

https://github.com/pybuild-org/pybuild/releases からプリコンパイル済みのバイナリをダウンロード

`pybuild` を実行すると `target.xml` のコンパイルが開始されます

設定ファイルを指定するには `pybuild custom.xml` を実行

## GitHub Actions での実行

```yaml
- name: setup pybuild
  uses: pybuild-org/pybuild@main
  with: # optional
      version: 'latest' # default
      goos: 'linux' # default
      goarch: 'amd64' # default
```

その後、後のジョブで `pybuild` コマンドが利用可能になります