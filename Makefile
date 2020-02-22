export:
	@while read LINE; do export "$LINE"; done < go.env

run:
	@echo "---- Running Application ----"
	@go run .