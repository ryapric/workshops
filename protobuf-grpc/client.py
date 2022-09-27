import grpc
import pb.example_pb2_grpc

if __name__ == "__main__":
    with grpc.insecure_channel("127.0.0.1:8080") as channel:
        stub = pb.example_pb2_grpc.ExampleStub(channel)
        echo_response = stub.Echo("hello grpc")
        getrecord_response = stub.GetRecord()
