FROM python:3.9.0a3-alpine3.10

COPY  requirement.txt /mnt

WORKDIR /mnt


RUN /usr/local/bin/python -m pip install --upgrade pip \
    && pip3 install -r requirement.txt

CMD ["python3","-x","app.py"]


