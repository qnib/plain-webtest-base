FROM qnib/go-webtest:v0.1.3 AS src

FROM scratch
ENV WEBTEST_HTTP_PORT="9999"
COPY --from=src /webtest /webtest
CMD ["/webtest"]
