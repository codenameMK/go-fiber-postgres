

# Use the official Golang image as the base image
FROM golang:1.19

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .


# Expose the port that the application will run on
EXPOSE 4000

# Command to run the executable
CMD ["make" , "run"]