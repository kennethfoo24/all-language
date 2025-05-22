FROM ruby:3.4.0

RUN apt-get update && apt-get install -y build-essential && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY ruby/ .

RUN gem install bundler && bundle install

EXPOSE 4567

ENV DD_TRACE_DEBUG=true
ENV DD_TRACE_STARTUP_LOGS=true

CMD ["bundle", "exec", "ruby", "app.rb", "-o", "0.0.0.0", "-p", "4567", "-s", "puma"]
