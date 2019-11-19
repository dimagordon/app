#!/usr/bin/env bash
PRIVACYDBURI="postgres://dev-user:testpassword@db/dev-user?sslmode=disable"
echo $PRIVACYDBURI


goose -dir ./migrations postgres "$PRIVACYDBURI" up-all-unapplied fix
sqlboiler -c ./migrations/sqlboiler.yml --no-tests --no-hooks --wipe --pkgname model --output ./db/model psql
echo "Done"
