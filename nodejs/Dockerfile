FROM node:22

EXPOSE 3000

# switch working directory
WORKDIR /app

# install packages
COPY nodejs/package.json ./

RUN npm install

# copy app
COPY nodejs/ ./

CMD node index.js
