-- os name
function name()
    print("{{os.Name}} {{os.Version}}")
    return "TEST"
end

-- Return functions as a table
return {
    name = name,
}
