gen-ent:
	go generate ./ent

gen-graph:
	go run github.com/99designs/gqlgen

start:
	go run github.com/cosmtrek/air
