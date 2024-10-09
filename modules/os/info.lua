-- os name
function name()
    local handle = io.popen("uname -s && uname -r")  -- For Linux and macOS
    local result = handle:read("*a")
    handle:close()
    print("OS Name and Version: " .. result)
end

-- Get current username
function user()
    local handle = io.popen("whoami")
    local result = handle:read("*a")
    handle:close()
    print("Current User: " .. result)
end

-- Get operating system architecture
function architecture()
    local handle = io.popen("uname -m")
    local result = handle:read("*a")
    handle:close()
    print( result)
end

-- Get total memory
function total_memory()
    local osName = "{{os.Name}}" -- Check the OS name

    if osName == "darwin" then
        local handle = io.popen("sysctl -n hw.memsize")
        local result = handle:read("*a")
        handle:close()
        print(  tonumber(result) / (1024 * 1024)  )
    elseif osName == "linux" then
        local handle = io.popen("grep MemTotal /proc/meminfo | awk '{print $2}'")
        local result = handle:read("*a")
        handle:close()
        print(  tonumber(result) / 1024  )
    else
        print("Unsupported OS")
    end
end

-- Get free memory
function free_memory()
    local osName = "{{os.Name}}" -- Check the OS name

    if osName == "darwin" then
        local handle = io.popen("vm_stat | grep 'Pages free:'")
        local result = handle:read("*a")
        handle:close()
        local pages = tonumber(result:match("%d+")) or 0
        print( pages * 4)  -- 4 KB per page
    elseif osName == "linux" then
        local handle = io.popen("free -m | grep 'Mem:' | awk '{print $4}'")
        local result = handle:read("*a")
        handle:close()
        print( result)
    else
        print("Unsupported OS")
    end
end

-- Get system uptime
function uptime()
    local osName = "{{os.Name}}" -- Check the OS name

    if osName == "darwin" then
        local handle = io.popen("uptime")
        local result = handle:read("*a")
        handle:close()
        print("Uptime: " .. result)
    elseif osName == "linux" then
        local handle = io.popen("uptime -p")
        local result = handle:read("*a")
        handle:close()
        print("Uptime: " .. result)
    else
        print("Uptime: Unsupported OS")
    end
end

-- Return functions as a table
return {
    name = name,
    user = user,
    architecture = architecture,
    total_memory = total_memory,
    free_memory = free_memory,
    uptime = uptime,
}
