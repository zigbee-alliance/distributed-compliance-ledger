FROM ubuntu:latest

RUN apt-get update && apt-get install -y \
        curl \
        git \
        unzip \
        gcc \
        make \
     && rm -rf /var/lib/apt/lists/*

# GoLang
ARG GO_VERSION
ENV GO_VERSION=${GO_VERSION:-1.17.2}
ENV BASH_ENV=/etc/bashrc
ENV PATH="${PATH}:/usr/local/go/bin"
RUN curl -L https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz -o /tmp/go.tar.gz \
    && rm -rf /usr/local/go \
    && tar -C /usr/local -xzf /tmp/go.tar.gz \ 
    && go version \
    && rm -f /tmp/go.tar.gz

# test user
ARG TEST_USER
ENV TEST_USER=${TEST_USER:-test}

ARG TEST_USER_GROUP
ENV TEST_USER_GROUP=${TEST_USER_GROUP:-test}

ARG uid=1000
ARG gid=1000

RUN (groupadd -g ${gid} ${TEST_USER_GROUP} || true) \
    && useradd -c "test user" -d /home/${TEST_USER} -u ${uid} -g ${gid} -m ${TEST_USER}

# Starport
ARG STARPORT_VERSION
#ENV STARPORT_VERSION=${STARPORT_VERSION:-0.18.5}
ENV STARPORT_VERSION=${STARPORT_VERSION:-dev}
#ENV PATH="${PATH}:/usr/local/bin"
#RUN curl https://get.starport.network/starport@v${STARPORT_VERSION}! -o /tmp/startport \
#    && bash /tmp/startport \
#    && rm /tmp/startport
RUN curl -L https://github.com/tendermint/starport/archive/refs/heads/develop.zip -o /tmp/starport.zip
RUN cd /tmp \
    && unzip starport.zip \
    && cd starport-develop \
    && make build \
    && cp ./dist/starport /usr/local/bin
ENV PATH="${PATH}:/usr/local/bin"
