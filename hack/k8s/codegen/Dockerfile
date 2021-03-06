FROM golang:1.9
ARG CODEGEN_VERSION="1.9.1"

RUN apt-get update && \
    apt-get install -y \
    git

# Code generator stuff
# Check: https://github.com/kubernetes/kubernetes/pull/57656
RUN wget http://github.com/kubernetes/code-generator/archive/kubernetes-${CODEGEN_VERSION}.tar.gz && \
    mkdir -p /go/src/k8s.io/code-generator/ && \
    tar zxvf kubernetes-${CODEGEN_VERSION}.tar.gz --strip 1 -C /go/src/k8s.io/code-generator/ && \
    mkdir -p /go/src/k8s.io/kubernetes/hack/boilerplate/ && \
    touch /go/src/k8s.io/kubernetes/hack/boilerplate/boilerplate.go.txt && \
    go get  k8s.io/apimachinery/pkg/apimachinery/registered

# Create user
ARG uid=1000
ARG gid=1000
RUN addgroup --gid $gid codegen && \
    adduser --gecos "First Last,RoomNumber,WorkPhone,HomePhone" --disabled-password --uid $uid --ingroup codegen codegen && \
    chown codegen:codegen -R /go

COPY . /hack
RUN chown codegen:codegen -R /hack


USER codegen

WORKDIR /hack
CMD ["./update-codegen.sh"]
