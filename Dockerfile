FROM python:2
COPY ./mybin ./
RUN chmod 0755 ./mybin
ENTRYPOINT [ "./mybin" ]
