-- This file is ran during startup of the server
require("app/prelude")

DATABASE:exec [[
    CREATE TABLE IF NOT EXISTS todos (id INTEGER PRIMARY KEY, done INTEGER, content);

    INSERT INTO todos VALUES (NULL, 0, 'Add some todos');
    INSERT INTO todos VALUES (NULL, 0, 'Add a form to add todos');
]]
