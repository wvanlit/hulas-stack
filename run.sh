docker run -it -p 8080:8080 \
    -v ./app:/server/app \
    -v ./bin:/server/bin \
    -v ./data:/server/data \
    -v ./lib:/server/lib \
    hulas-server