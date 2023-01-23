FROM python:3.7

RUN mkdir /app
ADD requirements.txt /app

WORKDIR /app
RUN pip3 install -r requirements.txt

EXPOSE 5000

CMD ["python", "main.py"]
