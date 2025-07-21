FROM python:3.9
WORKDIR /app
COPY ./requirements-docker.txt /app/
RUN pip install --upgrade pip
RUN pip install --only-binary=gevent gevent==21.12.0
RUN pip install -r /app/requirements-docker.txt
COPY run.sh /app/
ENV PYTHONPATH /app
ENV PATH $PATH:/app
RUN chmod 755 /app/run.sh
CMD /app/run.sh
