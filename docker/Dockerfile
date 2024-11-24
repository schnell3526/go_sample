FROM golang:1.23-bookworm AS devcontainer
SHELL ["/bin/bash", "-eo", "pipefail", "-c"]

ARG USER_NAME=devuser
ARG USER_UID=1000
ARG USER_GID=$USER_UID

RUN apt-get update && apt-get install -y --no-install-recommends \
    sudo \
    vim \
    tree \
    && rm -rf /var/lib/apt/lists/*
# ユーザーとグループを作成し、デフォルトシェルを bash に設定
RUN getent group ${USER_GID} || groupadd --gid $USER_GID $USER_NAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USER_NAME -s /bin/bash \
    && usermod -aG sudo $USER_NAME \
    && echo "$USER_NAME ALL=(ALL) NOPASSWD:ALL" > /etc/sudoers.d/$USER_NAME \
    && chmod 0440 /etc/sudoers.d/$USER_NAME
USER ${USER_NAME}
CMD [ "bash" ]

FROM golang:1.23.3-bookworm AS build
SHELL ["/bin/bash", "-eo", "pipefail", "-c"]
ARG MAIN_PATH=cmd/main.go
ARG APP_NAME=main
WORKDIR /app
COPY ../go.mod* ../go.sum* ./
COPY ../cmd ./cmd
COPY ../internal ./internal
RUN CGO_ENABLED=0 go build -o ${APP_NAME} -ldflags="-w -s" ${MAIN_PATH}

FROM gcr.io/distroless/static-debian12:nonroot AS prod
ARG APP_NAME=main
COPY --from=build --chmod=755 /app/${APP_NAME} /
ENTRYPOINT [ "${APP_NAME}" ]
