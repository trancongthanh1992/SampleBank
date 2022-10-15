# Lưu ý: Khi sử dụng mutil-stage 
# ta có thể sử dụng COPY –from từ bất kì một image 
# nào ko được tạo ra trên Dockerfile hiện tại mà có thể trên cả Docker hotst.


# Một trong những vấn đề lớn nhất khi sử dụng Docker là docker image có dung lượng lớn. 
# Nguyên nhân là sau mỗi layer có thêm data cho các layer đó được lưu vào image. 
# Vậy nên, giữ được kích thước Image ở mức càng thấp càng tốt là một bài toán người viết Dockerfile cần phải quan tâm.


# Golang and all packages that are required by our images project. Using build multiple-stage.


# Downloaded image `golang:1.19-alpine3.16`, and marked image name as `builder`. `builder` is a stage that use for command `COPY`.
FROM golang:1.19-alpine3.16 AS builder

# Created directory of working image in docker container.
WORKDIR /app 
# `.` mean copy everything from current folder to `/app`.
COPY . .
# Build single binary file. `-o main` name of output, `main.go` entry point.
RUN go build -o main main.go

### Run `stage`
# Downloaded image `alpine:3.16`.
FROM alpine:3.16
# At `/app` directory, copy the executable binary file from the stage `builder` into `app/main`.
# `app/main` is run stage image
WORKDIR /app
COPY --from=builder /app/main .

# Copy app.env
COPY app.env .

# Container listen on specified networking port at run time.
EXPOSE 8080

# Go into app directory and run file main binary.
# Golang includes all packages that required by our project. The size of docker images is very big.
# Problem can solve to build by `stage`.
CMD [ "/app/main" ]