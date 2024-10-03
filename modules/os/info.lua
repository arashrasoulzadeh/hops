-- os name
function name()
    print("{{os.Name}} {{os.Version}}")
    return "TEST"
end

function user()
    print("{{user.Username}}")
end

-- Return functions as a table
return {
    name = name,
    user = user,
}
