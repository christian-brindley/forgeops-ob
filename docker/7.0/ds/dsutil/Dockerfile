# Utility ds image for testing, benchmarking, etc.
# Also useful as a general purpose debugging tool
FROM gcr.io/forgerock-io/ds/pit1:7.1.0-7c0e8c80719b88e015fe3cf3dd2a2aeb54ff2dfe

USER root
RUN apt-get update && apt-get install -y bash vim curl dnsutils procps netcat


COPY --chown=forgerock:root bin/*  bin/
COPY --chown=forgerock:root ds-idrepo.template config/MakeLDIF

# reset entrypoint of ds container
ENTRYPOINT []
CMD ["bash"]
