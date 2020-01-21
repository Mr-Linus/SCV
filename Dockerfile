FROM nvidia/cuda:10.0-base

WORKDIR /

COPY scv /usr/local/bin

CMD ["scv"]