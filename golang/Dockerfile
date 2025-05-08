FROM golang:latest
WORKDIR /app
COPY run.sh /app/
COPY go.mod /app/
COPY go.sum /app/
ENV PATH $PATH:/app
RUN chmod 755 /app/run.sh
CMD /app/run.sh

