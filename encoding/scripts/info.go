package scripts

/*
local info = redis.call('info', 'all')

local result = {
    cmdstats = {},
    keyspace = {},
    stats = {},
}

for line in string.gmatch(info, '[^\r\n]+') do
    if not string.find(line, "#")
    then
        for prop, value in string.gmatch(line, "(.*):(.*)") do
            if string.find(prop, "cmdstat") == 1
            then
                local kvp = {}
                for kv in string.gmatch(value, '[^,]+') do
                    for k, v in string.gmatch(kv, '(.*)=(.*)') do
                        kvp[k] = v
                    end
                end
                result.cmdstats[prop] = kvp
            elseif string.find(prop, "db") == 1
            then
                local kvp = {}
                for kv in string.gmatch(value, '[^,]+') do
                    for k, v in string.gmatch(kv, '(.*)=(.*)') do
                        kvp[k] = v
                    end
                end
                result.keyspace[prop] = kvp
            else
                result.stats[prop] = value
            end
        end
    end
end

return cjson.encode(result)
*/
const Info = `local a=redis.call('info','all')local b={cmdstats={},keyspace={},stats={}}for c in string.gmatch(a,'[^\r\n]+')do if not string.find(c,"#")then for d,e in string.gmatch(c,"(.*):(.*)")do if string.find(d,"cmdstat")==1 then local f={}for g in string.gmatch(e,'[^,]+')do for h,i in string.gmatch(g,'(.*)=(.*)')do f[h]=i end end;b.cmdstats[d]=f elseif string.find(d,"db")==1 then local f={}for g in string.gmatch(e,'[^,]+')do for h,i in string.gmatch(g,'(.*)=(.*)')do f[h]=i end end;b.keyspace[d]=f else b.stats[d]=e end end end end;return cjson.encode(b)`