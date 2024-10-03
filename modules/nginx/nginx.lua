-- restart nginx server
function restart()

    local command = ""
    if "{{os.Name}}" == "windows" then
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
