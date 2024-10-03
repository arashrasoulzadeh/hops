-- Function to get CPU information based on the OS
function cpu()

    -- Fetch and print CPU info based on the detected OS
    if "{{os.Name}}" == "windows" then
        os.execute("wmic cpu get caption,deviceid,numberofcores,maxclockspeed")
    elseif "{{os.Name}}" == "darwin" then
        os.execute("sysctl -n machdep.cpu.brand_string")
    else
        os.execute("lscpu")
    end
end

-- Function to get system usage
function usage()
    -- Fetch and print system usage information based on the OS
    if "{{os.Name}}" == "darwin" then
        os.execute("top -l 1 | grep -E '^CPU|Phys'")
    else
        os.execute("top -bn1 | grep -E '^%Cpu|Mem'")
    end
end

-- Return functions as a table
return {
    cpu = cpu,
    ram_usage = ram_usage,
    usage = usage
}