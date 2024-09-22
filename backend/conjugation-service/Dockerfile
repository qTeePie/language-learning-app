# Start with a lightweight Python base image
FROM python:3.11-slim

# Set environment variables for Miniconda installation
ENV CONDA_DIR /opt/conda
ENV PATH $CONDA_DIR/bin:$PATH

# Install Miniconda
RUN apt-get update && apt-get install -y \
    wget \
    && wget --quiet https://repo.anaconda.com/miniconda/Miniconda3-latest-Linux-x86_64.sh -O /tmp/miniconda.sh \
    && /bin/bash /tmp/miniconda.sh -b -p $CONDA_DIR \
    && rm /tmp/miniconda.sh \
    && apt-get clean

# Copy the environment.yml file that contains the conda environment details
COPY environment.yml /tmp/environment.yml

# Create the Conda environment from the environment.yml file
RUN conda env create -f /tmp/environment.yml

# Set the default environment to the one created
ENV PATH /opt/conda/envs/myenv/bin:$PATH
RUN echo "source activate myenv" > ~/.bashrc

# Copy your application code into the container (adjust the path as needed)
COPY . /app

# Set the working directory
WORKDIR /app

# Specify the command to run your Python app
CMD ["python", "main.py"]
