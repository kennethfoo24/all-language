FROM ruby:3.1.2

RUN apt-get update && \
    apt-get install -y build-essential && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY ruby/ .

RUN gem install bundler && \
    bundle install

EXPOSE 4567

CMD ["ruby", "app.rb", "-o", "0.0.0.0", "-p", "4567"]
