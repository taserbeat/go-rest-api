# 概要

golamg と React で作る RestAPI のサンプルアプリ

# 開発環境

TODO: 追記

# 実行方法

## 1. Docker コンテナの起動

```bash
# コンテナの起動
docker-compose up -d
```

## 2. DB のマイグレーション

```bash
GO_ENV=dev go run migrate/migrate.go

>> Connected!
>> Successfully Migrated!
```

\* pgweb でマイグレーションを確認しておく

pgweb: http://localhost:8081

## 3. バックエンドの起動

```bash
GO_ENV=dev go run main.go
```

# 終了方法

## 1. Docker コンテナの終了

```bash
# コンテナの終了
docker-compose down --volumes --remove-orphans
```
