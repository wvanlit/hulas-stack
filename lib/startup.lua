-- This file is ran during startup of the server
require("app/lib/prelude")

DATABASE:exec [[
    CREATE TABLE IF NOT EXISTS todos (id INTEGER PRIMARY KEY, done INTEGER, content);
]]

print("Startup done!")
