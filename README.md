# Connect 2 container via IP static
docker run --name simplebank bank-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:aA@123123@172.17.0.2:5432/sample_bank_db?sslmode=disable" simplebank:latest

# Connect 2 container via network
docker run --name simplebank --network bank-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:aA@123123@postgres14.4:5432/sample_bank_db?sslmode=disable" simplebank:latest
