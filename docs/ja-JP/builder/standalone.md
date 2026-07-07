# standalone ビルダー設定

standalone ビルダーは独立した実行ファイルを構築でき、マルチアーキテクチャ、クロスプラットフォームをサポートし、zip 圧縮パッケージ形式で成果物を生成します。

standalone ビルダーは builder 設定とローカル python 設定に依存します。

```xml
<config type="group" name="standalone targets">
    <config name="python">
        <prop name="arch" value="" />
        <prop name="os" value="" />
    </config>
    <config name="pip">
        <prop name="platform" value="" />
        <prop name="download" value="" />
    </config>
    <config name="launcher">
        <prop name="run" value="" />
    </config>
</config>

<run job="build standalone" />
```

複数のターゲットを構築する必要がある場合は、`standalone targets` を複数回記述する必要があります。

`python.arch` は対象 python インタプリタの実行アーキテクチャです。

`python.os` は対象 python インタプリタの実行システムです。

`python.arch` と `python.os` は https://github.com/astral-sh/python-build-standalone/releases を参照して記述します。

`pip.platform` は https://packaging.python.org/en/latest/specifications/platform-compatibility-tags/#platform-tag を参照して記述します。

`pip.download` はプロジェクトの依存関係を宣言するために使用し、複数の依存がある場合は複数の `pip.download` を記述する必要があります。

`launcher.run` はプロジェクトのエントリーポイントで、具体的な python スクリプトファイルやモジュール（例: `-m http.server`）で指定できます。