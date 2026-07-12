FROM ruby:3.4-slim

RUN apt update && apt install -y build-essential bundler

WORKDIR /app
RUN bundle config path vendor

ENTRYPOINT ["bundle", "exec", "jekyll", "serve", "--host", "0.0.0.0"]
