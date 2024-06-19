run: frontend backend
	build/gm

frontend-vue:
	bash ./scripts/build_frontend_vue.sh

frontend-htmx:
	bash ./scripts/build_frontend_htmx.sh

backend:
	mkdir -p build
	go build -o build/gm ./cmd/gm

clean:
	mkdir -p build
	rm -rf build/*
	rmdir build
