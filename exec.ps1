docker build -t golang-webapp .
docker run --env-file .env -p 3000:3000 -it golan-webapp
