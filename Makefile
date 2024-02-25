.PHONY: templ	run

templ:
	@templ generate

css: 
	@npx tailwindcss -i ./routes/app/public/input.css -o ./routes/app/public/app.css --watch

clean:
	@find . -type f -name '*_templ.go'
	
run: templ 
	@go run cmd/main.go

drop:
	@go run cmd/db/main.go -drop

seed:
	@go run cmd/db/main.go -seed

migrate:
	@go run cmd/db/main.go