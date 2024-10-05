// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file falin.proto (package within.website.x.falin.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";

/**
 * @generated from message within.website.x.falin.v1alpha1.GenerateImageRequest
 */
export class GenerateImageRequest extends Message<GenerateImageRequest> {
  /**
   * @generated from field: string prompt = 1;
   */
  prompt = "";

  /**
   * @generated from field: string model = 2;
   */
  model = "";

  /**
   * @generated from field: int64 seed = 3;
   */
  seed = protoInt64.zero;

  /**
   * @generated from field: int32 num_images = 4;
   */
  numImages = 0;

  /**
   * @generated from field: bool enable_safety_checker = 5;
   */
  enableSafetyChecker = false;

  constructor(data?: PartialMessage<GenerateImageRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "within.website.x.falin.v1alpha1.GenerateImageRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "prompt", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "model", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "seed", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 4, name: "num_images", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "enable_safety_checker", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GenerateImageRequest {
    return new GenerateImageRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GenerateImageRequest {
    return new GenerateImageRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GenerateImageRequest {
    return new GenerateImageRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GenerateImageRequest | PlainMessage<GenerateImageRequest> | undefined, b: GenerateImageRequest | PlainMessage<GenerateImageRequest> | undefined): boolean {
    return proto3.util.equals(GenerateImageRequest, a, b);
  }
}

/**
 * @generated from message within.website.x.falin.v1alpha1.ImageOutput
 */
export class ImageOutput extends Message<ImageOutput> {
  /**
   * @generated from field: string url = 1;
   */
  url = "";

  /**
   * @generated from field: int32 width = 2;
   */
  width = 0;

  /**
   * @generated from field: int32 height = 3;
   */
  height = 0;

  /**
   * @generated from field: string content_type = 4;
   */
  contentType = "";

  constructor(data?: PartialMessage<ImageOutput>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "within.website.x.falin.v1alpha1.ImageOutput";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "width", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "height", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "content_type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ImageOutput {
    return new ImageOutput().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ImageOutput {
    return new ImageOutput().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ImageOutput {
    return new ImageOutput().fromJsonString(jsonString, options);
  }

  static equals(a: ImageOutput | PlainMessage<ImageOutput> | undefined, b: ImageOutput | PlainMessage<ImageOutput> | undefined): boolean {
    return proto3.util.equals(ImageOutput, a, b);
  }
}

/**
 * @generated from message within.website.x.falin.v1alpha1.GenerateImageResponse
 */
export class GenerateImageResponse extends Message<GenerateImageResponse> {
  /**
   * @generated from field: repeated within.website.x.falin.v1alpha1.ImageOutput images = 1;
   */
  images: ImageOutput[] = [];

  /**
   * @generated from field: string prompt = 2;
   */
  prompt = "";

  constructor(data?: PartialMessage<GenerateImageResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "within.website.x.falin.v1alpha1.GenerateImageResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "images", kind: "message", T: ImageOutput, repeated: true },
    { no: 2, name: "prompt", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GenerateImageResponse {
    return new GenerateImageResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GenerateImageResponse {
    return new GenerateImageResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GenerateImageResponse {
    return new GenerateImageResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GenerateImageResponse | PlainMessage<GenerateImageResponse> | undefined, b: GenerateImageResponse | PlainMessage<GenerateImageResponse> | undefined): boolean {
    return proto3.util.equals(GenerateImageResponse, a, b);
  }
}

