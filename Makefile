# Makefile

run:
	go run main.go

template:
	go run libs/templates/main.go $(name)
