build:
	docker rmi go-fetcher -f
	docker build -t go-fetcher .

run:
	docker run -d -p 5000:1234 -e PORT="1234" go-fetcher