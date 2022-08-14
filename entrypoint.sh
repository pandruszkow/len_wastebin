#!/bin/sh
set -e

RUN_CMD="lenpaste"


# WASTEBIN_ADDRESS
if [[ -n "$WASTEBIN_ADDRESS" ]]; then
	RUN_CMD="$RUN_CMD -address $WASTEBIN_ADDRESS"
fi


# WASTEBIN_DB_DRIVER
if [[ -n "$WASTEBIN_DB_DRIVER" ]]; then
	RUN_CMD="$RUN_CMD -db-driver '$WASTEBIN_DB_DRIVER'"
fi


# WASTEBIN_DB_SOURCE
if [[ -z "$WASTEBIN_DB_DRIVER" || "$WASTEBIN_DB_DRIVER" == "sqlite3" ]]; then
	RUN_CMD="$RUN_CMD -db-source /data/lenpaste.db"

else
	RUN_CMD="$RUN_CMD -db-source '$WASTEBIN_DB_SOURCE'"
fi


# WASTEBIN_DB_CLEANUP_PERIOD
if [[ -n "$WASTEBIN_DB_CLEANUP_PERIOD" ]]; then
	RUN_CMD="$RUN_CMD -db-cleanup-period '$WASTEBIN_DB_CLEANUP_PERIOD'"
fi

# WASTEBIN_ROBOTS_DISALLOW
if [[ "$WASTEBIN_ROBOTS_DISALLOW" == "true" ]]; then
	RUN_CMD="$RUN_CMD -robots-disallow"
	
else
	if [[ "$WASTEBIN_ROBOTS_DISALLOW" != "" && "$WASTEBIN_ROBOTS_DISALLOW" != "false" ]]; then
		echo "[ENTRYPOINT] Error: unknown: WASTEBIN_ROBOTS_DISALLOW = $WASTEBIN_ROBOTS_DISALLOW"
		exit 2
	fi
fi


# WASTEBIN_TITLE_MAX_LENGTH
if [[ -n "$WASTEBIN_TITLE_MAX_LENGTH" ]]; then
	RUN_CMD="$RUN_CMD -title-max-length '$WASTEBIN_TITLE_MAX_LENGTH'"
fi


# WASTEBIN_BODY_MAX_LENGTH
if [[ -n "$WASTEBIN_BODY_MAX_LENGTH" ]]; then
	RUN_CMD="$RUN_CMD -body-max-length '$WASTEBIN_BODY_MAX_LENGTH'"
fi


# WASTEBIN_MAX_PASTE_LIFETIME
if [[ -n "$WASTEBIN_MAX_PASTE_LIFETIME" ]]; then
	RUN_CMD="$RUN_CMD -max-paste-lifetime '$WASTEBIN_MAX_PASTE_LIFETIME'"
fi


# Server about
if [[ -f "/data/about.html" ]]; then
	RUN_CMD="$RUN_CMD -server-about /data/about.html"
fi


# Server rules
if [[ -f "/data/rules.html" ]]; then
	RUN_CMD="$RUN_CMD -server-rules /data/rules.html"
fi


# WASTEBIN_ADMIN_NAME
if [[ -n "$WASTEBIN_ADMIN_NAME" ]]; then
	RUN_CMD="$RUN_CMD -admin-name '$WASTEBIN_ADMIN_NAME'"
fi


# WASTEBIN_ADMIN_MAIL
if [[ -n "$WASTEBIN_ADMIN_MAIL" ]]; then
	RUN_CMD="$RUN_CMD -admin-mail '$WASTEBIN_ADMIN_MAIL'"
fi


# Run Lenpaste
echo "[ENTRYPOINT] $RUN_CMD"
sh -c "$RUN_CMD"
