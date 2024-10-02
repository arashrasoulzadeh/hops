-- restart nginx server
function restart()
    -- Determine the operating system
    local os_name = package.config:sub(1,1) == '\\' and "windows" or "unix"

    local command = ""
    if os_name == "windows" then
        -- Windows command (assuming you're using a Windows service manager)
        command = "net stop nginx && net start nginx"
    else
        -- Linux/macOS (Unix-based) command
        command = "sudo systemctl restart nginx"
    end

    -- Execute the command
    local result = os.execute(command)
    if result == 0 then
        print("Nginx restarted successfully on " .. os_name .. ".")
    else
        print("Failed to restart Nginx on " .. os_name .. ".")
    end
end

-- Return functions as a table
return {
    restart = restart,
}
