FROM nickblah/luajit:2-luarocks-ubuntu

RUN apt-get update
RUN apt-get install -y tree build-essential sqlite3 libsqlite3-dev
RUN luarocks install --server=https://luarocks.org lsqlite3

# Create a directory inside the container
WORKDIR /server

# Copy the app and bin directories into the container
COPY ./bin ./bin

RUN chmod +x ./bin/server

# The command to run the application
ENTRYPOINT [ "./bin/server" ]
