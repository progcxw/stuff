FROM docker.io/library/alpine
COPY ./mybin ./
RUN chmod 0755 ./mybin
ENTRYPOINT [ "./mybin" ]
