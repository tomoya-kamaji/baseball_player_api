source .env

DIR="migrations"
DRIVER="mysql"
DBSTRING="${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DB}"

# TODO
# DBSTRING_STAGING="user=liam dbname=tester_stg sslmode=disable"
# DBSTRING_PRODUCTION="user=liam dbname=tester_prd sslmode=disable"

# https://github.com/pressly/goose
function create() {
    if [ -z "$MIGRATION_FILE" ]; then
        echo "引数がありません"
        exit 1
    fi
    goose -dir ${DIR} ${DRIVER} "${DBSTRING}" create $MIGRATION_FILE sql
}

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