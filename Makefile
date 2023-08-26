# styles
tw_src = './static/css/src.css'
tw_dst = './static/css/out.css'
# build
ifeq ($(OS),Windows_NT)
	target = './bin/bookly.exe'
else
	target = './bin/bookly'
endif

bookly:
	go build -o $(target) ./cmd/main.go

run: bookly
	./bin/bookly --no-migration

# compile styles with standalone tailwindcss
tw:
	tailwindcss -i $(tw_src) -o $(tw_dst)