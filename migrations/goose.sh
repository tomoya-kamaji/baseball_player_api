DIR="migrations"
DRIVER="mysql"
DBSTRING="root:password@tcp(127.0.0.1:3306)/mydb" 

# TODO
# DBSTRING_STAGING="user=liam dbname=tester_stg sslmode=disable"
# DBSTRING_PRODUCTION="user=liam dbname=tester_prd sslmode=disable"

# https://github.com/pressly/goose
function status() {
    goose -dir ${DIR} ${DRIVER} "${DBSTRING}" status
}

function up() {
    goose -dir ${DIR} ${DRIVER} "${DBSTRING}" up
}

function down() {
   goose -dir ${DIR} ${DRIVER} "${DBSTRING}" down
}

ENV=$(echo "$2" | tr '[:lower:]' '[:upper:]')
$1