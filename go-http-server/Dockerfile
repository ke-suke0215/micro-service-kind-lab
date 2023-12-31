# 基本イメージとしてgolangの最新版を使用
FROM golang:1.21.5

# 作業ディレクトリを設定
WORKDIR /app

# 依存関係ファイルをコピーし、依存関係をダウンロード
COPY go.mod ./
# COPY go.sum ./  パッケージの外部依存がないため不要
RUN go mod download

# ソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o main .

# 実行ポートを指定
EXPOSE 8080

# 実行コマンド
CMD ["./main"]