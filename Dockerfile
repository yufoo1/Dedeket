FROM ubuntu:22.04
RUN mkdir /build
WORKDIR /build
COPY E-TexSub-backend /build/backend
EXPOSE 8080
ENTRYPOINT ["./backend"]