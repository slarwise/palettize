FROM grafana/tempo:latest

ADD tempo.yaml /etc/tempo.yaml
ADD tempo-data /tmp/tempo

ENTRYPOINT ["/tempo"]
