FROM node:22

EXPOSE 3000

# switch working directory
WORKDIR /app

# install packages
COPY package.json ./

RUN npm install

# copy app
COPY ./ ./

CMD node index.js