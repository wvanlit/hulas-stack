require "lib/prelude"

if REQUEST_METHOD == "POST" then
    local content = ParseQueryString(REQUEST_BODY)["todo-item"]
    DATABASE:exec(
        "INSERT INTO todos VALUES (NULL, 0, '" .. content .. "')")

    REQUEST_METHOD = "GET" -- Also show the list of todos after adding a new one
end

if REQUEST_METHOD == "PATCH" then
    local params = ParseQueryString(REQUEST_BODY)

    DATABASE:exec("UPDATE todos SET done = " .. Cond(params.state == "1", 0, 1) .. " WHERE id = " .. params.id)

    REQUEST_METHOD = "GET" -- Also show the list of todos after adding a new one
end

if REQUEST_METHOD == "GET" then
    for row in DATABASE:nrows("SELECT * FROM todos") do
        local item = HTML("li",
            {
                class = Cond(row.done == 1, "line-through", "") ..
                    " flex flex-row items-center gap-4"
            },
            HTML(
                "input",
                {
                    type = "checkbox",
                    ["hx-vals"] = "{" ..
                        "\"id\": \"" .. row.id .. "\"," ..
                        " \"state\": \"" .. row.done .. "\"" ..
                        "}",
                    ["hx-patch"] = "/api/todos",
                    ["hx-target"] = "#todos",
                    ["hx-swap"] = "innerHTML",
                    _ = Cond(row.done == 1, "checked", ""),
                    ["class"] = "checkbox checkbox-primary"
                },
                ""
            )
            .. " " .. HTML("span", {}, row.content)
        )

        print(item)
    end
end
