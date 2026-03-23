.PHONY: build clean

build: clean
	@mkdir -p dist
	dagger call build --source=. export --path=./dist/srcom

clean:
	rm -rf dist/
	rm -rf build/
