FROM csighub.tencentyun.com/admin/tlinux2.2-bridge-base-tcloud:latest

LABEL maintainer="yifenglu@tencent.com"

USER root

STOPSIGNAL 37

ENV MICRO_DEBUG_PROFILE "http"
ENV MICRO_SERVER_NAME "bossapi"
ENV MICRO_SERVER_ADDRESS ":8070"


ENV JAEGER_SAMPLER_TYPE "probabilistic"
ENV JAEGER_SAMPLER_PARAM 0.5
ENV JAEGER_SAMPLER_MANAGER_HOST_PORT=127.0.0.1:6832

RUN mkdir -p /data/app
RUN rm -rf /tmp/* /var/tmp/*

# http perf
EXPOSE 6060
# http server
EXPOSE 8080
# prometheus
EXPOSE 8088

ADD pkg_tools_v2/ /data/pkg_tools_v2/
ADD pkg_install_batch.sh /usr/local/bin/pkg_install_batch.sh
ADD cmd/cmd /data/app/cmd
ADD start.sh /etc/kickStart.d/start.sh

