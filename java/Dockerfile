# FROM openjdk:17
# WORKDIR /app
# COPY java/app/ ./
# ARG VERSION
# RUN curl -L -H "Accept: application/zip" https://github.com/DataDog/dd-trace-java/releases/download/v1.49/dd-java-agent.jar -o dd-java-agent.jar
# RUN ./mvnw package -DskipTests 
# CMD java -javaagent:dd-java-agent.jar -jar target/spring-boot-0.0.1-SNAPSHOT.jar

FROM openjdk:17
WORKDIR /workspace
COPY java/app/ ./
RUN mvn -B clean package

FROM openjdk:17
WORKDIR /app
COPY --from=build /workspace/target/spring-boot-0.0.1-SNAPSHOT.jar app.jar
EXPOSE 8080
ENTRYPOINT ["java","-jar","/app/app.jar"]


