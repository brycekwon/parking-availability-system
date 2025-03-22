import-lorawan-device-profiles:
	docker-compose run --rm --entrypoint sh --user root chirpstack -c '\
		apk add --no-cache git && \
		git clone --depth 1 https://github.com/brocaar/lorawan-devices /tmp/lorawan-devices && \
		chirpstack -c /etc/chirpstack import-legacy-lorawan-devices-repository -d /tmp/lorawan-devices'
