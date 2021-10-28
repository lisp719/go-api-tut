const grpc = require("@grpc/grpc-js");

const messages = require("./hello_pb");
const services = require("./hello_grpc_pb");

function sayHello(call, callback) {
  const reply = new messages.HelloReply();
  reply.setMessage("Hello " + call.request.getName());
  callback(null, reply);
}

function main() {
  const server = new grpc.Server();

  server.addService(services.GreeterService, { sayHello: sayHello });

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
