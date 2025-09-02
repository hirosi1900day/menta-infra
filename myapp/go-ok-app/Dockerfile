# ベースイメージ
FROM golang:1.21-alpine

# 作業ディレクトリ作成
WORKDIR /app

# Goモジュールファイルを先にコピーして依存解決
COPY go.mod .
RUN go mod tidy

# 残りのファイルをコピー
COPY . .

# ビルド
RUN go build -o app main.go

# ポート開放
EXPOSE 8080

# 実行
CMD ["./app"]
