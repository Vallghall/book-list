# front
tw_src = ./static/css/src.css
tw_dst = ./static/css/out.css
htmx_url = https://unpkg.com/htmx.org@1.9.5/dist/htmx.min.js
htmx_path = ./static/js/htmx.min.js

# build
ifeq ($(OS),Windows_NT)
	target = ./bin/bookly.exe
else
	target = ./bin/bookly
endif

bookly: htmx
	go build -o $(target) ./cmd/main.go

run: bookly
	$(target) --no-migration

htmx:
ifeq (,$(wildcard ./static/js/htmx.min.js))
	curl -o $(htmx_path) $(htmx_url)
endif
# compile styles with standalone tailwindcss
tw:
	tailwindcss -i $(tw_src) -o $(tw_dst)