#####################################
# GO BUILDER
#####################################
FROM golang:1.13-buster AS builder

WORKDIR /build/scim

COPY . ./

WORKDIR /build/scim/server

RUN make build

#####################################
# FINAL IMAGE
#####################################
FROM debian:buster-slim

RUN apt-get -yq update \
  && apt-get -yq install ca-certificates curl openssl \
  && apt-get -yq autoremove \
  && apt-get -yq clean \
  && rm -rf /var/lib/apt/lists/* \
  && truncate -s 0 /var/log/*log

# copy binary
COPY --from=builder /build/scim/server/bin/linux_amd64/scim /usr/bin/scim

#copy schemas
COPY --from=builder /build/scim/server/schemas /usr/share/schemas

# run binary
CMD ["/usr/bin/scim"]
