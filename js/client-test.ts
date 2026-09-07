import { describe, it, expect } from "bun:test";
import { create } from "@bufbuild/protobuf";
import type { DescMessage, MessageShape } from "@bufbuild/protobuf";
import {
  VersionSchema,
  VersionServiceGetRequestSchema,
  VersionServiceGetResponseSchema,
} from "./fits/api/v1/version_pb";
import type { UnaryResponse } from "@connectrpc/connect";
import { newClient } from "./client";
import { newTestInterceptor } from "./test-interceptor";

function unaryResponse<Req extends DescMessage, Res extends DescMessage>(
  message: MessageShape<Res>,
): UnaryResponse<Req, Res> {
  return {
    stream: false as const,
    message: message as any,
    header: new Headers(),
    trailer: new Headers(),
    service: undefined as any,
    method: undefined as any,
  };
}

describe("client", () => {
  it("works with test interceptor", async () => {
    const c = newClient({
      baseUrl: "http://this-is-just-for-testing",
      interceptors: [
        newTestInterceptor([
          {
            wantRequest: create(VersionServiceGetRequestSchema, {}),
            wantRequestSchema: VersionServiceGetRequestSchema,
            wantResponse: () =>
              unaryResponse(
                create(VersionServiceGetResponseSchema, {
                  version: create(VersionSchema, { version: "1.0" }),
                }),
              ),
          },
        ]),
      ],
    });

    const v = await c.apiv1().version().get({});
    expect(v.version?.version).toBe("1.0");
  });
});
