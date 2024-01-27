# syntax=docker/dockerfile:experimental

FROM scratch
COPY bin/zookeepercli /bin/zookeepercli
ENTRYPOINT ["/bin/zookeepercli"]
