import grpc
import pb.example_pb2
import pb.example_pb2_grpc

if __name__ == "__main__":
    with grpc.insecure_channel("127.0.0.1:8080") as channel:
        stub = pb.example_pb2_grpc.ExampleStub(channel)

        # This is pretty infuriating -- Python protobuf code isn't generated
        # directly, it generates a *Python* generator to generate the code at
        # runtime (in *_pb2.py). As such, you don't get completion hints in
        # editors because there's no data types to autofill etc. So, look at
        # your proto file and determine how you should be setting up your
        # requests!
        echo_request = pb.example_pb2.Echoable(msg = "hello grpc")
        echo_response = stub.Echo(echo_request)
        print(echo_response)

        getrecord_request = pb.example_pb2.GetRecordRequest(name = "Thomas Anderson")
        getrecord_response = stub.GetRecord(getrecord_request)
        print(getrecord_response)
