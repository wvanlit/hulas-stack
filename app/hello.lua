local sqlite3 = require("lsqlite3")

local db = sqlite3.open("/server/data/db.sqlite")

db:exec [[
    DROP TABLE IF EXISTS test;

    CREATE TABLE test (id INTEGER PRIMARY KEY, content);

    INSERT INTO test VALUES (NULL, 'Hi world!');
    INSERT INTO test VALUES (NULL, 'Hi lua!');
    INSERT INTO test VALUES (NULL, 'Hi sqlite3!');
    INSERT INTO test VALUES (NULL, 'Hi HULAS!');
]]

print("<ul>")
for row in db:nrows("SELECT * FROM test") do
    print("<li>", row.content, "</li>")
end
print("</ul>")
