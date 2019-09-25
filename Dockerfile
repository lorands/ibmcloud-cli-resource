FROM golang:alpine as builder
#RUN apk add --no-cache curl jq
RUN mkdir -p /assets
WORKDIR /assets
COPY . /go/src/github.com/lorands/ibmcloud-cli-resource
ENV CGO_ENABLED 0
RUN go build -o /assets/in github.com/lorands/ibmcloud-cli-resource/in/cmd/in
RUN go build -o /assets/out github.com/lorands/ibmcloud-cli-resource/out/cmd/out
RUN go build -o /assets/check github.com/lorands/ibmcloud-cli-resource/check/cmd/check
WORKDIR /go/src/github.com/lorands/ibmcloud-cli-resource
RUN set -e; for pkg in $(go list ./... | grep -v "acceptance"); do \
		go test -o "/tests/$(basename $pkg).test" -c $pkg; \
	done

FROM frolvlad/alpine-glibc AS base_resource
RUN apk add --no-cache bash tzdata ca-certificates

RUN apk add --no-cache bash ca-certificates curl jq python3 git build-base
RUN update-ca-certificates

RUN apk add --no-cache python3 && \
    python3 -m ensurepip && \
    rm -r /usr/lib/python*/ensurepip && \
    pip3 install --upgrade pip setuptools && \
    if [ ! -e /usr/bin/pip ]; then ln -s pip3 /usr/bin/pip ; fi && \
    if [[ ! -e /usr/bin/python ]]; then ln -sf /usr/bin/python3 /usr/bin/python; fi && \
    rm -r /root/.cache

#RUN \
#pip3 install git+https://github.com/lorands/concourse-python-tooling.git

# Install IBN Cloud cli (bx)
WORKDIR /tmp
ADD https://clis.ng.bluemix.net/download/bluemix-cli/0.18.2/linux64 /tmp/bx-cli.tgz
RUN \
  tar zxvf bx-cli.tgz && \
  Bluemix_CLI/install_bluemix_cli && \
  rm -f /tmp/bx-cli.tgz && \
  rm -rf /tmp/Bluemix_CLI
WORKDIR /

RUN ibmcloud plugin install cloud-databases -f
#RUN bx plugin install activity-tracker -r Bluemix -f
#RUN bx plugin install analytics-engine -r Bluemix -f
#RUN bx plugin install auto-scaling -r Bluemix -f
#RUN bx plugin install cloud-functions -r Bluemix -f
#RUN bx plugin install cloud-internet-services -r Bluemix -f
#RUN bx plugin install container-registry -r Bluemix -f
#RUN bx plugin install container-service -r Bluemix -f
#RUN bx plugin install dev -r Bluemix -f
#RUN bx plugin install infrastructure-service -r Bluemix -f
#RUN bx plugin install logging-cli -r Bluemix -f
#RUN bx plugin install machine-learning -r Bluemix -f
#RUN bx plugin install vpn -r Bluemix -f

RUN bx config --check-version=false

# FROM resource AS tests
# COPY --from=builder /tests /go-tests
# COPY out/assets /go-tests/assets
# WORKDIR /go-tests
# RUN set -e; for test in /go-tests/*.test; do \
# 		$test; \
# 	done

FROM base_resource AS resource
COPY --from=builder assets/ /opt/resource/
RUN chmod +x /opt/resource/*

FROM resource
