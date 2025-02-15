###############################################################################
# Hugo Website Build                                                          #
###############################################################################
FROM docker.io/hugomods/hugo:node-0.143.0 AS ui_builder

# ability to run makefiles
RUN apk add --no-cache make

# avoid running as root
USER hugo
WORKDIR /usr/share/hugo

# set build environment to production
ENV NODE_ENV=production
ENV HUGO_ENVIRONMENT=production

# copy makefile for build instructions
COPY --chown=hugo:hugo ./Makefile ./Makefile

# copy and install npm packages
COPY --chown=hugo:hugo package.json ./package.json
COPY --chown=hugo:hugo package-lock.json ./package-lock.json
RUN npm ci --omit=dev

# copy css configurations
COPY --chown=hugo:hugo ./tailwind.config.js ./tailwind.config.js
COPY --chown=hugo:hugo ./postcss.config.js ./postcss.config.js

# copy hugo configuration
COPY --chown=hugo:hugo ./config ./config

# copy source files, ordered from most to least often changed to take
# advantage of Docker's built-in caching.
COPY --chown=hugo:hugo ./static ./static
COPY --chown=hugo:hugo ./assets ./assets
COPY --chown=hugo:hugo ./layouts ./layouts
# COPY    --chown=hugo:hugo   ./data      ./data    # currently unused
# COPY    --chown=hugo:hugo   ./i18n      ./i18n    # currently unused
COPY --chown=hugo:hugo ./content ./content

# build website
RUN make ui-build

###############################################################################
# Go App Build                                                                #
###############################################################################
FROM docker.io/golang:1.24.0-alpine AS app_builder

# ability to run makefiles
RUN apk add --no-cache make

WORKDIR /usr/share/golang

# copy makefile for build instructions
COPY ./Makefile ./Makefile

# copy and install go packages
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
RUN go mod download

# copy application source
COPY ./cmd ./cmd
COPY ./internal ./internal

RUN make app-build

###############################################################################
# Running Application                                                         #
###############################################################################
FROM scratch

WORKDIR /usr/share/app

COPY --from=ui_builder /usr/share/hugo/public ./public
COPY --from=app_builder /usr/share/golang/website ./website

USER 10001:10001

CMD ["./website"]