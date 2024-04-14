# Start with a base image containing Node.js runtime
FROM node:14

# Set the working directory in the container
WORKDIR /app

# Copy package.json and package-lock.json
COPY ./package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the code
COPY ./ .

# Expose port 3000 to the outside world
EXPOSE 3000

# Start the app
CMD ["npm", "start"]