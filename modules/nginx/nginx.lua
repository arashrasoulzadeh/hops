-- restart nginx server: restart
function restart()
    print("restart nginx!")
end
-- Return functions as a table
return {
    restart = restart,
}
