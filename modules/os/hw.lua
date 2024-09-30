-- Function to get CPU information based on the OS
function cpu()
    -- Detect OS based on package.config
    local os_name
    if package.config:sub(1,1) == '\\' then
        os_name = "windows"
    else
        os_name = io.popen("uname"):read("*l")
        if os_name == "Darwin" then
            os_name = "mac"
        else
            os_name = "unix"
        end
    end

    -- Fetch CPU info based on the detected OS
    local cpu_info
    if os_name == "windows" then
        -- Command to fetch CPU info in Windows
        cpu_info = io.popen("wmic cpu get caption,deviceid,numberofcores,maxclockspeed"):read("*a")
    elseif os_name == "mac" then
        -- Command to fetch CPU info in macOS
        cpu_info = io.popen("sysctl -n machdep.cpu.brand_string"):read("*a")
    else
        -- Command to fetch CPU info in Linux/Unix
        cpu_info = io.popen("lscpu"):read("*a")
    end

    print( cpu_info)
end

-- Return functions as a table
return {
    cpu = cpu,
}
