FROM golang:1.19

WORKDIR /app

RUN set -ex; \
   ln -fs /usr/share/zoneinfo/Asia/Tokyo /etc/localtime; \
   go install github.com/cortesi/modd/cmd/modd@latest; \
   go install github.com/go-delve/delve/cmd/dlv@latest; \
   go install gotest.tools/gotestsum@latest; \
   :

CMD ["modd", "-f", "tools/api/modd.conf"]
