FROM python:3-bookworm
ARG PUBSUB_PROJECT_ID
ARG PUBSUB_TOPIC_ID
ARG PUBSUB_PORT
COPY create.sh /usr/local/bin
WORKDIR /usr/local/app
RUN apt install git curl
RUN git clone --recursive https://github.com/googleapis/python-pubsub.git
WORKDIR /usr/local/app/python-pubsub/samples/snippets
ENV PUBSUB_EMULATOR_HOST=pubsub:${PUBSUB_PORT}
RUN --mount=type=cache,target=/root/.cache pip install -r requirements.txt
CMD ["sh", "-c", "/usr/local/bin/create.sh ${PUBSUB_PROJECT_ID} ${PUBSUB_TOPIC_ID}"]