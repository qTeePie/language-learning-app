# Use a base image that includes Apertium
FROM apertium/apertium-apy:latest

# Install additional packages if needed (ensure you have the correct POS tagger)
RUN apt-get update && \
    apt-get install -y apertium-en-it

# Expose the port for Apertium if needed
EXPOSE 8080

# Run Apertium server for POS tagging
CMD ["apertium", "en-it", "-p", "8081"]
