-- Contains all the libraries that are required by default in the application
local sqlite3 = require("lsqlite3")

DATABASE = sqlite3.open("/server/data/db.sqlite")

REQUEST_METHOD = os.getenv("REQUEST_METHOD") or "GET"
REQUEST_BODY = os.getenv("REQUEST_BODY") or ""

function UriDecode(s)
    s = string.gsub(s, '%%(%x%x)', function(hex)
        return string.char(tonumber(hex, 16))
    end)
    return s
end

function ParseQueryString(query)
    local params = {}
    for pair in string.gmatch(query, "[^&]+") do
        local key, value = string.match(pair, "([^=]+)=([^=]+)")
        if key and value then
            params[key] = UriDecode(value)
        end
    end
    return params
end

function Cond(cond, T, F)
    if cond then
        return T
    else
        return F
    end
end

function HTML(type, attrs, childContent)
    local tag = "<" .. type
    for p, v in pairs(attrs) do
        if p == "_" then
            tag = tag .. " " .. v
        else
            tag = tag .. " " .. p .. "=\'" .. v .. "\'"
        end
    end
    tag = tag .. ">"
    tag = tag .. childContent
    tag = tag .. "</" .. type .. ">"
    return tag
end

print("") -- Always print an empty line to prevent the server from marking the script as failed
