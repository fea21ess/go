FROM golang:1.16.5

ENV GO111MODULE "on"
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /go/src/app
COPY . .

#RUN go get github.com/pilu/fresh
#RUN go get -u github.com/beego/beego/v2
#RUN go get -u github.com/beego/bee/v2
#RUN go get -u github.com/go-sql-driver/mysql
#RUN go get -u github.com/jackc/pgx/v4
#RUN go get -u go.mongodb.org/mongo-driver
#RUN go get -u github.com/go-ldap/ldap/v3
#RUN go get -u gopkg.in/xmlpath.v2
RUN go mod download
RUN go mod tidy

RUN go build -mod=mod -o main .
RUN chmod a+x ./main

RUN cp ./main /bin/main

#  fresh 将会自动运行项目的 main.go  
#CMD [ "fresh" ]
EXPOSE 8080

CMD ["/bin/main"]