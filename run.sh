docker run -it -p 8080:8080 \
    -v ./app:/server/app \
    -v ./bin:/server/bin \
    -v ./data:/server/data \
    hulas-server