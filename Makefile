build: image
	kustomize build --enable-alpha-plugins .testdata/

image:
	docker build -t bluebrown/test-validator .
