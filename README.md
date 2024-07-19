# measurement-go
A data ingestion service that compares JSON payloads and protocol buffer measurements from mock producers

## Usage

1. Start the REST consumer and gRPC consumer.

```
go run ./cmd/rest-consumer
go run ./cmd/grpc-consumer
```

2. Run the REST producer.

```
go run ./cmd/rest-producer
```

3. Run the gRPC producer.

```
go run ./cmd/grpc-producer
```

## Results

Processing 1 million records with 1 thousand parallel workers.

### REST / JSON

```
Elapsed time: 2m20.224821625s
Error count: 938844
```

### gRPC / Protocol Buffers

```
Elapsed time: 13.796378958s
Error count: 0
```
