// @generated by protoc-gen-connect-es v1.5.0 with parameter "target=ts"
// @generated from file falin.proto (package within.website.x.falin.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GenerateImageRequest, GenerateImageResponse } from "./falin_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service within.website.x.falin.v1alpha1.ImageService
 */
export const ImageService = {
  typeName: "within.website.x.falin.v1alpha1.ImageService",
  methods: {
    /**
     * @generated from rpc within.website.x.falin.v1alpha1.ImageService.GenerateImage
     */
    generateImage: {
      name: "GenerateImage",
      I: GenerateImageRequest,
      O: GenerateImageResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

