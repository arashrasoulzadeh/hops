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

    -- Fetch and print CPU info based on the detected OS
    if os_name == "windows" then
        os.execute("wmic cpu get caption,deviceid,numberofcores,maxclockspeed")
    elseif os_name == "mac" then
        os.execute("sysctl -n machdep.cpu.brand_string")
    else
        os.execute("lscpu")
    end
end

-- Function to get system usage
function usage()
    -- Fetch and print system usage information based on the OS
    local os_name
    if package.config:sub(1,1) == '\\' then
        -- Windows system usage command. Add appropriate Windows command here.
        print("System usage information not available for Windows.")
    else
        os_name = io.popen("uname"):read("*l")
        if os_name == "Darwin" then
            os.execute("top -l 1 | grep -E '^CPU|Phys'")
        else
            os.execute("top -bn1 | grep -E '^%Cpu|Mem'")
        end
    end
end

-- Return functions as a table
return {
    cpu = cpu,
    ram_usage = ram_usage,
    usage = usage
}