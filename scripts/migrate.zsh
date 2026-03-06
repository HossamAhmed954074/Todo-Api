#!/bin/zsh

# Load environment variables from .env
if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
fi

COMMAND=$1
NAME=$2

case $COMMAND in
  "up")
    migrate -path migrations -database "$DATABASE_URL" up
    ;;
  "down")
    COUNT=${NAME:-1}
    echo "Rolling back $COUNT migration(s). Continue? [y/N]"
    read -r CONFIRM
    if [[ "$CONFIRM" == "y" ]]; then
        migrate -path migrations -database "$DATABASE_URL" down "$COUNT"
    fi
    ;;
  "create")
    migrate create -ext sql -dir migrations -seq "$NAME"
    ;;
  "force")
    migrate -path migrations -database "$DATABASE_URL" force "$NAME"
    ;;
  *)
    echo "Usage: $0 {up|down|create|force}"
    exit 1
    ;;
esac