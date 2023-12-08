
# Use a base image with JRE
FROM openjdk:8-jdk-alpine

# Set the working directory in the container
WORKDIR /app

# Copy the WAR file into the container at /app
COPY ./webapp.war /app/webapp.war

# Expose the port that your application will run on
EXPOSE 80

# Specify the command to run on container start
CMD ["java", "-jar", "webapp.war"]
