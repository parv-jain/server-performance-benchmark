# Use an official Node.js runtime as a parent image
FROM node:18

# Set the working directory inside the container
WORKDIR /app

# Copy your Node.js server code and package.json/package-lock.json
COPY app.js package*.json ./

# Install Node.js dependencies
RUN npm install

# Expose the port your Node.js server is listening on
EXPOSE 3000

# Start your Node.js server
CMD ["node", "app.js"]
