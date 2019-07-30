// package: car
// file: motor.proto

import * as jspb from "google-protobuf";

export class Vector extends jspb.Message {
  getX(): number;
  setX(value: number): void;

  getY(): number;
  setY(value: number): void;

  getZ(): number;
  setZ(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Vector.AsObject;
  static toObject(includeInstance: boolean, msg: Vector): Vector.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Vector, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Vector;
  static deserializeBinaryFromReader(message: Vector, reader: jspb.BinaryReader): Vector;
}

export namespace Vector {
  export type AsObject = {
    x: number,
    y: number,
    z: number,
  }
}

export class MoveVector extends jspb.Message {
  hasDirection(): boolean;
  clearDirection(): void;
  getDirection(): Vector | undefined;
  setDirection(value?: Vector): void;

  hasRotation(): boolean;
  clearRotation(): void;
  getRotation(): Vector | undefined;
  setRotation(value?: Vector): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MoveVector.AsObject;
  static toObject(includeInstance: boolean, msg: MoveVector): MoveVector.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: MoveVector, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MoveVector;
  static deserializeBinaryFromReader(message: MoveVector, reader: jspb.BinaryReader): MoveVector;
}

export namespace MoveVector {
  export type AsObject = {
    direction?: Vector.AsObject,
    rotation?: Vector.AsObject,
  }
}

export class SensorParameters extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SensorParameters.AsObject;
  static toObject(includeInstance: boolean, msg: SensorParameters): SensorParameters.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SensorParameters, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SensorParameters;
  static deserializeBinaryFromReader(message: SensorParameters, reader: jspb.BinaryReader): SensorParameters;
}

export namespace SensorParameters {
  export type AsObject = {
  }
}

export class CurrentData extends jspb.Message {
  getErrormessage(): string;
  setErrormessage(value: string): void;

  hasLocation(): boolean;
  clearLocation(): void;
  getLocation(): Vector | undefined;
  setLocation(value?: Vector): void;

  hasRotation(): boolean;
  clearRotation(): void;
  getRotation(): Vector | undefined;
  setRotation(value?: Vector): void;

  hasVelocity(): boolean;
  clearVelocity(): void;
  getVelocity(): Vector | undefined;
  setVelocity(value?: Vector): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CurrentData.AsObject;
  static toObject(includeInstance: boolean, msg: CurrentData): CurrentData.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CurrentData, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CurrentData;
  static deserializeBinaryFromReader(message: CurrentData, reader: jspb.BinaryReader): CurrentData;
}

export namespace CurrentData {
  export type AsObject = {
    errormessage: string,
    location?: Vector.AsObject,
    rotation?: Vector.AsObject,
    velocity?: Vector.AsObject,
  }
}

