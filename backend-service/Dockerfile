FROM alpine:latest
MAINTAINER anonymous<pgl888999@icloud.com>
ENV VERSION 1.0

# 设置工作目录
WORKDIR /app
# 拷贝可执行文件
COPY bin/public-system-app /app/public-system-app
# 拷贝配置文件
COPY config.yaml /app
# 设置时区
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime &&  \
    echo 'Asia/Shanghai' > /etc/timezone
# 设置编码
ENV LANG C.UTF-8
# 暴露端口
EXPOSE 8085
ENTRYPOINT ["/app/public-system-app"]