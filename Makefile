.PHONY: templ	run

templ:
	@templ generate

css: 
	@npx tailwindcss -i ./routes/app/public/input.css -o ./routes/app/public/app.css

clean:
	@find . -type f -name '*_templ.go'
	
run: templ 
	@go run cmd/main.go