FROM ubuntu:22.04
RUN mkdir /build
WORKDIR /build
COPY Dedeket /build/backend
EXPOSE 8080
ENTRYPOINT ["./backend"]