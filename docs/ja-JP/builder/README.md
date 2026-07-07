# builder 設定

```xml
<config name="builder">
    <prop name="app" value="" />
    <prop name="source" value="" />
    <prop name="output" value="" />
    <prop name="version" value="" />
    <prop name="release" value="" />
</config>

<run job="setup builder" />
```

`app` はアプリケーション名です  

`source` はソースコードディレクトリで、このディレクトリはビルド成果物に完全に保持されます  

`output` はビルド成果物の出力ディレクトリです  

`version` は Python インタプリタのバージョンです  

`release` は Python インタプリタのビルドバージョンです  

`version` と `release` は https://github.com/astral-sh/python-build-standalone/releases を参照して記述してください  

# ローカル Python 設定

```xml
<config name="python">
    <prop name="arch" value="" />
    <prop name="os" value="" />
</config>

<run job="setup python" />
```

`arch` はローカル Python インタプリタの実行アーキテクチャです  

`os` はローカル Python インタプリタの実行システムです  

`arch` と `os` は https://github.com/astral-sh/python-build-standalone/releases を参照して記述してください  

[standalone ビルダー設定](./standalone.md)

[docker ビルダー設定](./docker.md)