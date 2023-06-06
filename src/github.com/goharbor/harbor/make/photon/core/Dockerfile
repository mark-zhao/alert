#ARG harbor_base_image_version
#FROM goharbor/harbor-core-base:${harbor_base_image_version}
FROM docker-registry.ztgame.com.cn/library/goharbor/harbor-core-base:ztgame20200526
#FROM goharbor/harbor-core-base:v1.10.1

HEALTHCHECK CMD curl --fail -s http://127.0.0.1:8081/api/ping || exit 1
COPY harbor_core /harbor/
COPY views /harbor/views
COPY migrations /harbor/migrations

RUN chmod u+x /harbor/harbor_core
WORKDIR /harbor/
USER harbor
ENTRYPOINT ["/harbor/harbor_core"]
