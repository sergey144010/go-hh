
# Build
docker build -t go-hh-image .

# Run
docker run -d --rm --name=go-hh-container -p 8181:80 go-hh-image