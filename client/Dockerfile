# Use the Node.js 20 image (non-alpine) for better compatibility
FROM node:20 as development

# Set the working directory inside the container
WORKDIR /app

# Install Bun globally using npm
RUN npm install -g bun

# Verify Bun installation
RUN bun --version || echo "Bun not found"

# Copy package.json and bun.lockb
COPY package.json bun.lockb ./

# Install dependencies using Bun
RUN bun install

# Copy the rest of the application code
COPY . .

# Expose port 3000 (default for Nuxt.js development)
EXPOSE 3000

# Start the Nuxt.js application in development mode with hot reloading
CMD ["bun", "run", "dev"]
