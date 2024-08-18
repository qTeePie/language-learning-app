# Use an Ubuntu base image
FROM ubuntu:latest

RUN apt-get -qq update && apt-get -qq install python3-full python3-pip pipenv locales \
    && locale-gen en_US.UTF-8 \
    && update-locale LANG=en_US.UTF-8 LC_ALL=en_US.UTF-8

# Install necessary packages and Apertium-APY
RUN apt-get update && \
    apt-get install -y wget gnupg && \
    wget -qO- https://apertium.projectjj.com/apt/install-nightly.sh | bash && \
    apt-get install -y apertium-apy apertium-eng apertium-ita apertium-eng-ita

# Set the environment variables for locale
ENV LANG=en_US.UTF-8
ENV LC_ALL=en_US.UTF-8

# Expose the port used by Apertium-APY
EXPOSE 2737

# Run Apertium-APY
CMD ["apertium-apy", "-p", "2737", "/usr/share/apertium"]