###############################################################################
# Build Static Website                                                        #
###############################################################################
FROM docker.io/hugomods/hugo:node-0.143.0 AS node_builder

# ability to run makefiles
RUN apk add --no-cache make

# avoid running as root
USER hugo
WORKDIR /usr/share/hugo

# copy makefile for build instructions
COPY --chown=hugo:hugo ./Makefile ./Makefile

# copy and install npm packages
COPY --chown=hugo:hugo package.json ./package.json
COPY --chown=hugo:hugo package-lock.json ./package-lock.json
RUN npm ci --omit=dev

# copy hugo site configuration
COPY --chown=hugo:hugo ./hugo.yaml ./hugo.yaml

# copy hugo css configuration
COPY --chown=hugo:hugo ./postcss.config.js ./postcss.config.js
COPY --chown=hugo:hugo ./tailwind.config.js ./tailwind.config.js

# copy source files, ordered from most to least often changed to take
# advantage of Docker's built-in caching.
COPY --chown=hugo:hugo ./static ./static
COPY --chown=hugo:hugo ./assets ./assets
COPY --chown=hugo:hugo ./layouts ./layouts
# COPY    --chown=hugo:hugo   ./data      ./data    # currently unused
# COPY    --chown=hugo:hugo   ./i18n      ./i18n    # currently unused
COPY --chown=hugo:hugo ./content ./content

# build website
RUN make ui

###############################################################################
# Compile Server                                                              #
#                                                                             #
# a debian or ubuntu based distribution is required for building and running  #
# this binary due to its dependency of CGO_ENABLED=1.                         #
###############################################################################
FROM docker.io/golang:1.24-bookworm AS golang_builder

# compile with c modules
ENV CGO_ENABLED=1

WORKDIR /usr/share/golang

# copy makefile for build instructions
COPY ./Makefile ./Makefile

# copy go project metadata
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

# copy go server source
COPY ./cmd ./cmd
COPY ./internal ./internal

RUN make server

###############################################################################
# Run Server                                                                  #
#                                                                             #
# a debian or ubuntu based distribution is required for building and running  #
# this binary due to its dependency of CGO_ENABLED=1.                         #
###############################################################################
FROM docker.io/debian:bookworm-20250203-slim

WORKDIR /app

COPY --from=node_builder /usr/share/hugo/public ./public
COPY --from=golang_builder /usr/share/golang/main ./main

CMD ["./main"]