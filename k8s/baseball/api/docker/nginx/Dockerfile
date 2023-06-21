FROM nginx:1.16.1

# Nginxの設定ファイルをコンテナにコピー
COPY nginx.conf /etc/nginx/nginx.conf

# ホストシステム上のWebサイトのファイルをコンテナにコピー
COPY website /usr/share/nginx/html

# Nginxが動作する80番ポートを公開
EXPOSE 80

# Nginxの起動
CMD ["nginx", "-g", "daemon off;"]