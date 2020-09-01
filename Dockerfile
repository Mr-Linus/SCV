FROM nvidia/cuda:10.0-base

WORKDIR /

COPY bin/scv /usr/local/bin

CMD ["scv"]
