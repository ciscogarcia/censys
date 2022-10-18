docker:
	docker build --tag censys .
	docker run -itd -p 10000:10000 censys:latest
