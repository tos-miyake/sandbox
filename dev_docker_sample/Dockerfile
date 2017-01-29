FROM ubuntu:16.04

RUN locale-gen en_US.UTF-8
ENV LANG en_US.utf8
ENV LC_ALL en_US.UTF-8

ARG DEBIAN_FRONTEND=noninteractive

RUN \
  apt-get update && \
  apt-get install -y curl && \
  # Add Node.js repository
  curl -sL https://deb.nodesource.com/setup_7.x | bash - && \
  apt-get install -y \
    build-essential \
    libmysqlclient-dev \
    libreadline-dev \
    libssl-dev \
    libxml2-dev \
    libxslt-dev \
    libyaml-dev \
    nodejs \
    ruby2.3 \
    ruby2.3-dev \
    zlib1g-dev

WORKDIR /app

# Install RubyGems
ARG BUNDLE_JOBS=2
RUN gem update --system && gem install bundler
COPY Gemfile /app/Gemfile
COPY Gemfile.lock /app/Gemfile.lock
RUN bundle install --path vendor/bundle

# Install npm modules
# RUN npm install -g yarn
# COPY package.json /app/package.json
# COPY yarn.lock /app/yarn.lock
# RUN yarn

EXPOSE 3000
CMD ["bundle", "exec", "rails", "server", "-b", "0.0.0.0"]
