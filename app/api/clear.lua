require "lib/prelude"

DATABASE:exec [[
    DELETE FROM todos WHERE done = 1;
]]

-- This is a small hack to make sure the todos are reloaded after they are cleared
print("<div hx-get=\"/api/todos\" hx-swap=\"outerHTML\" hx-trigger=\"load\"></div>")
