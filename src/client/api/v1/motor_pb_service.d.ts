// package: car
// file: motor.proto

import * as motor_pb from "./motor_pb";
import {grpc} from "@improbable-eng/grpc-web";

type CarMove = {
  readonly methodName: string;
  readonly service: typeof Car;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof motor_pb.MoveVector;
  readonly responseType: typeof motor_pb.CurrentData;
};

type CarGetSensorData = {
  readonly methodName: string;
  readonly service: typeof Car;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof motor_pb.SensorParameters;
  readonly responseType: typeof motor_pb.CurrentData;
};

export class Car {
  static readonly serviceName: string;
  static readonly Move: CarMove;
  static readonly GetSensorData: CarGetSensorData;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class CarClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  move(
    requestMessage: motor_pb.MoveVector,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: motor_pb.CurrentData|null) => void
  ): UnaryResponse;
  move(
    requestMessage: motor_pb.MoveVector,
    callback: (error: ServiceError|null, responseMessage: motor_pb.CurrentData|null) => void
  ): UnaryResponse;
  getSensorData(
    requestMessage: motor_pb.SensorParameters,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: motor_pb.CurrentData|null) => void
  ): UnaryResponse;
  getSensorData(
    requestMessage: motor_pb.SensorParameters,
    callback: (error: ServiceError|null, responseMessage: motor_pb.CurrentData|null) => void
  ): UnaryResponse;
}

