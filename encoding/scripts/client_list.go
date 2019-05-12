package scripts

/*
local list = ARGV[1]

local result = {}

for line in string.gmatch(list, '[^\r\n]+') do
    local kvp = {}
    for client in string.gmatch(line, '[^ ]+') do
        for k, v in string.gmatch(client, '(.*)=(.*)') do
            kvp[k] = v
        end
    end
    table.insert(result, kvp)
end

return cjson.encode(result)
*/
const ClientList = `local a=ARGV[1]local b={}for c in string.gmatch(a,'[^\r\n]+')do local d={}for e in string.gmatch(c,'[^ ]+')do for f,g in string.gmatch(e,'(.*)=(.*)')do d[f]=g end end;table.insert(b,d)end;return cjson.encode(b)`
