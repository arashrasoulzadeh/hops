-- Create wordpress docker-compose.yml file
function wordpress(db_engine, db_user, db_password, db_name)
    -- Docker Compose template for WordPress
    local compose_template = [[
version: '3.1'

services:
  wordpress:
    image: wordpress:latest
    ports:
      - "8080:80"
    environment:
      WORDPRESS_DB_HOST: %s
      WORDPRESS_DB_USER: %s
      WORDPRESS_DB_PASSWORD: %s
      WORDPRESS_DB_NAME: %s
    volumes:
      - ./wordpress_data:/var/www/html

  %s:
    image: %s
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: %s
      MYSQL_DATABASE: %s
      MYSQL_USER: %s
      MYSQL_PASSWORD: %s
    volumes:
      - ./db_data:/var/lib/mysql
]]

    -- Database service name based on engine
    local db_service_name
    local db_image

    if db_engine == "mysql" then
        db_service_name = "db"
        db_image = "mysql:5.7"
    elseif db_engine == "mariadb" then
        db_service_name = "mariadb"
        db_image = "mariadb:latest"
    else
        error("Unsupported database engine: " .. db_engine)
    end

    -- Format the template with the provided inputs
    local docker_compose = string.format(compose_template,
            db_service_name,      -- WORDPRESS_DB_HOST
            db_user,              -- WORDPRESS_DB_USER
            db_password,          -- WORDPRESS_DB_PASSWORD
            db_name,              -- WORDPRESS_DB_NAME
            db_service_name,      -- Database service name (db or mariadb)
            db_image,             -- Database image (mysql or mariadb)
            db_password,          -- MYSQL_ROOT_PASSWORD
            db_name,              -- MYSQL_DATABASE
            db_user,              -- MYSQL_USER
            db_password           -- MYSQL_PASSWORD
    )

    -- Write the docker-compose.yml file
    local file = io.open("docker-compose.yml", "w")
    if file then
        file:write(docker_compose)
        file:close()
        print("docker-compose.yml created successfully!")
    else
        print("Error creating docker-compose.yml file")
    end
end

-- Return functions as a table
return {
    wordpress = wordpress,
}
