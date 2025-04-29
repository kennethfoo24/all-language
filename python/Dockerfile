# start by pulling the python image
FROM python:3.8-slim

EXPOSE 5000

# switch working directory
WORKDIR /app

# install the dependencies and packages in the requirements file
COPY python/requirements.txt ./
RUN python -m pip install --upgrade pip \
 && pip install --no-cache-dir -r requirements.txt

# Then copy the rest of your code
COPY python/. .

# Tell Flask where your app lives
ENV FLASK_APP=app.py

CMD ["flask", "run", "--host=0.0.0.0", "--port=5000"]
