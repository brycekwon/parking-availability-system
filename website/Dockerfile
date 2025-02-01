###############################################################################
# Build Web Assets                                                            #
###############################################################################
FROM docker.io/node:22.13.1-alpine as node_builder

USER node
WORKDIR /usr/share/node

COPY --chown=node:node package.json ./package.json
COPY --chown=node:node package-lock.json ./package-lock.json
RUN npm ci --omit=dev

COPY --chown=node:node ./postcss.config.js ./postcss.config.js
COPY --chown=node:node ./tailwind.config.js ./tailwind.config.js

COPY --chown=node:node ./templates ./templates
COPY --chown=node:node ./assets ./assets

RUN npm run build

###############################################################################
# Compile Executable                                                          #
###############################################################################
FROM docker.io/golang:1.23.6-alpine as go_builder

WORKDIR /usr/share/go

COPY ./go.sum ./go.sum
COPY ./go.mod ./go.mod

COPY ./main.go ./main.go
COPY ./internal ./internal

COPY --from=node_builder /usr/share/node/dist ./dist
COPY ./templates ./templates

RUN go build -o main .

###############################################################################
# Run Web Server                                                              #
###############################################################################
FROM scratch

WORKDIR /app

COPY --from=go_builder /usr/share/go/main ./main

ENTRYPOINT ["./main"]