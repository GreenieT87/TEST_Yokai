FROM ubuntu
RUN apt update && apt install -y wget

RUN wget https://go.dev/dl/go1.22.2.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.22.2.linux-amd64.tar.gz && export PATH=$PATH:/usr/local/go/bin

WORKDIR /app

CMD [ "sleep", "365d" ]