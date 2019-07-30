// package: car
// file: motor.proto

var motor_pb = require("./motor_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var Car = (function () {
  function Car() {}
  Car.serviceName = "car.Car";
  return Car;
}());

Car.Move = {
  methodName: "Move",
  service: Car,
  requestStream: false,
  responseStream: false,
  requestType: motor_pb.MoveVector,
  responseType: motor_pb.CurrentData
};

Car.GetSensorData = {
  methodName: "GetSensorData",
  service: Car,
  requestStream: false,
  responseStream: false,
  requestType: motor_pb.SensorParameters,
  responseType: motor_pb.CurrentData
};

exports.Car = Car;

function CarClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

CarClient.prototype.move = function move(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Car.Move, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

CarClient.prototype.getSensorData = function getSensorData(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Car.GetSensorData, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.CarClient = CarClient;

