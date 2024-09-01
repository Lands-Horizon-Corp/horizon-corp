FROM node:20.17-alpine as development

# Set the working directory inside the container
WORKDIR /app

# Install Bun
RUN curl -fsSL https://bun.sh/install | bash

# Add Bun to the PATH
ENV PATH="/root/.bun/bin:${PATH}"

# Copy package.json and bun.lockb
COPY package.json bun.lockb ./

# Install dependencies using Bun
RUN bun install

# Copy the rest of the application code
COPY . .

# Expose port 3000 (default for Nuxt.js development)
EXPOSE 3000

# Start the Nuxt.js application in development mode
CMD ["bun", "run", "dev"]
