# Use the official Node.js 16 image as base
FROM node:20.9-alpine3.17

# Set the working directory in the container
WORKDIR /app

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application files
COPY . .

# Build the Next.js application
RUN npm run build

# Expose the port Next.js runs on (default is 3000)
EXPOSE 3000

# Define the command to start the application in production
CMD ["npm", "run", "start"]
