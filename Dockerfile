FROM scratch
COPY ./main /
# 声明服务端口
EXPOSE 7001
ENV ENVIRONMENT=PROD
# 启动容器时运行的命令
CMD [ "/main"]
