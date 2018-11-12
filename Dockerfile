FROM iron/go

WORKDIR /app

ADD ./dist /app

ENTRYPOINT ["./user-service"]
