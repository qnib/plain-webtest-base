FROM qnib/go-webtest:v0.1.1 AS src

FROM scratch
ENV WEBTEST_HTTP_PORT="9999"
COPY --from=src /go/src/github.com/qnib/go-webtest/webtest /webtest
CMD ["/webtest"]
