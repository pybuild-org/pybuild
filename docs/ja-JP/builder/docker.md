# docker ビルダー設定

docker ビルダーは docker image 形式に合致した tar ファイルを構築でき、システムに docker CLI がなくても動作し、マルチアーキテクチャをサポートし、ビルド対象は linux プラットフォームのみです。

docker ビルダーは builder 設定とローカル python 設定に依存します。

```xml
<config name="docker image meta">
    <prop name="tag" value="" />
</config>
```

`tag` は docker image のタグです。

```xml
<config type="group" name="docker image targets">
    <config name="image">
        <prop name="base" value="" />
        <prop name="arch" value="" />
    </config>
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

<run job="docker build" />
```

複数のターゲットを構築する必要がある場合は、`docker image targets` を複数回記述する必要があります。

`image.base` は docker image のベースイメージです。

`image.arch` は docker image の実行アーキテクチャです。

`python.arch` は対象の python インタプリタの実行アーキテクチャです。

`python.os` は対象の python インタプリタの実行 OS です。

`python.arch` と `python.os` は https://github.com/astral-sh/python-build-standalone/releases を参照して記述します。

`pip.platform` は https://packaging.python.org/en/latest/specifications/platform-compatibility-tags/#platform-tag を参考に記述します。

`pip.download` はプロジェクトの依存関係を宣言するために使用します。複数の依存がある場合は、複数の `pip.download` を記述する必要があります。

`launcher.run` はプロジェクトのエントリーポイントで、具体的な python スクリプトファイルやモジュール（例: `-m http.server`）にすることができます。