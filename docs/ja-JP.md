# 使用方法

## ローカルでの実行

https://github.com/pybuild-org/pybuild/releases からプリコンパイルされたバイナリをダウンロード

`pybuild` を実行すると `target.xml` のコンパイルが開始されます

構成ファイルを指定するには `pybuild custom.xml`

## GitHub Action での実行

```yaml
- name: setup pybuild
  uses: pybuild-org/pybuild@main
  with: # optional
      version: 'latest' # default
      goos: 'linux' # default
      goarch: 'amd64' # default
```

その後、後のタスクで `pybuild` コマンドを使用できます

# タグの使い方

## xml タグ

```xml
<xml></xml>
```

何の役にも立たないプレースホルダーで、スタックを消費せず、何回でも何層でも記述できます

## use タグ

```xml
<use file="" />
```

外部の XML フрагメントをインポートし、現在のコンテキストで即座に実行します

`file` パラメータはファイルパスまたは http(s):// で始まるリンクのいずれかを指定できます

もし `.xml` 拡張子を付けない場合、`use` タグは自動で付加します

## config タグ

```xml
<config [type="group"] name=""><config>
```

構成を宣言するタグで、`name` は構成名を宣言します

`type="group"` はオプションで、グループ構成を宣言するために使用されます。この場合、その `config` タグは複数の子 `config` タグによって宣言されます

## prop タグ

```xml
<prop name="" value="" />
```

構成項目の具体的な値を宣言するタグで、`name` はフィールド名を、`value` はフィールド値を宣言します

## run タグ

```xml
<run job="" | command="" />
```

タスクを実行するタグで、`job` と `command` のいずれかを指定します

`job` はタスク名を宣言し、プリセットされたタスクを実行するために使用されます

`command` はコマンド内容を宣言し、システムコマンドを実行するために使用されます

## log タグ

```xml
<log></log>
```

1 行のログを出力するために使用されます