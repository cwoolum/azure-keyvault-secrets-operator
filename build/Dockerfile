FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

ENV OPERATOR=/usr/local/bin/az-keyvault-secrets-operator \
    USER_UID=1001 \
    USER_NAME=az-keyvault-secrets-operator

# install operator binary
COPY build/_output/bin/src ${OPERATOR}

COPY build/bin /usr/local/bin
RUN chmod +x /usr/local/bin/user_setup
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
