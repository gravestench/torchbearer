run: frontend backend
	build/gm

frontend:
	bash ./scripts/build_frontend.sh

backend:
	mkdir -p build
	go build -o build/gm ./cmd/gm

clean:
	mkdir -p build
	rm -rf build/*
	rmdir build
