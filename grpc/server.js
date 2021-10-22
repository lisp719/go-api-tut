const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");

const PROTO_PATH = __dirname + "/hello.proto";

const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
});

const hello_proto = grpc.loadPackageDefinition(packageDefinition).hello;

function sayHello(call, callback) {
  callback(null, { message: "Hello " + call.request.name });
}

function main() {
  const server = new grpc.Server();

  server.addService(hello_proto.Greeter.service, { sayHello: sayHello });

  server.bindAsync(
    "0.0.0.0:50051",
    grpc.ServerCredentials.createInsecure(),
    () => {
      server.start();
      console.log("start server");
    }
  );
}

main();
