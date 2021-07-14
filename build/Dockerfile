FROM registry.cn-hangzhou.aliyuncs.com/riet/ubi-minimal:latest

ENV OPERATOR=/usr/local/bin/kubernetes-pod-health-check \
    USER_UID=1001 \
    USER_NAME=kubernetes-pod-health-check

# install operator binary
COPY build/_output/bin/kubernetes-pod-health-check ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
