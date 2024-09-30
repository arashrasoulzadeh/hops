-- os name
function name()
    local os_name = package.config:sub(1,1) == '\\' and "windows" or "unix" or "mac"
    print(os_name)
end


-- Return functions as a table
return {
    name = name,
}
