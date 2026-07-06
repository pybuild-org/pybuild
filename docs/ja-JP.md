# 使用方法

## ローカルでの実行

https://github.com/pybuild-org/pybuild/releases からプリコンパイル済みのバイナリをダウンロード

`pybuild` を実行すると `target.xml` のコンパイルを開始できます

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

その後、`pybuild` コマンドは後のジョブで利用可能になります

## タグの使い方

## xml タグ

```xml
<xml></xml>
```

何の役にも立たないプレースホルダーで、スタックを消費せず、任意の回数、任意のネストレベルで記述できます

## use タグ

```xml
<use file="" />
```

外部のXMLコンテンツを取り込み、現在のコンテキストで即座に実行します

`file` パラメータはファイルパスまたは `http(s)://` で始まるリンクのいずれかです

`.xml` 拡張子が付いていない場合、`use` タグは自動的に付加します

## config タグ

```xml
<config [type="group"] name=""><config>
```

設定を宣言するためのタグで、`name` は設定名を宣言します

`type="group"` はオプションで、グループ設定を宣言するために使用されます。この場合、その `config` タグは複数の子 `config` タグによって共に宣言されます

## prop タグ

```xml
<prop name="" value="" />
```

設定項目の具体的な値を宣言するためのタグで、`name`フィールド名を宣言し、`value`はフィールドの値を宣言します

## run タグ

```xml
<run job="" | command="" />
```

タスクを実行するためのタグで、`job` と `command` のいずれかを指定します

`job` はタスク名を宣言し、事前に設定されたタスクを実行するために使用されます

`command` はコマンド内容を宣言し、システムコマンドを実行するために使用されます

## log タグ

```xml
<log></log>
```

1行のログを出力するために使用されます