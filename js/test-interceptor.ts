import type { Interceptor, UnaryRequest, UnaryResponse } from "@connectrpc/connect";
import { equals } from "@bufbuild/protobuf";
import type { DescMessage, MessageShape } from "@bufbuild/protobuf";

export type ClientCall<Req extends DescMessage = DescMessage, Res extends DescMessage = DescMessage> = {
  wantRequest: MessageShape<Req>;
  wantRequestSchema: Req;
  wantResponse?: () => UnaryResponse<Req, Res>;
  wantError?: Error;
};

export function newTestInterceptor(calls: ClientCall[]): Interceptor {
  let count = 0;

  return (next) => async (req) => {
    const reqAny = req as UnaryRequest;

    const expected = calls[count];
    if (!expected) {
      throw new Error(
        `received an unexpected client call of type ${reqAny.method?.input?.typeName ?? "unknown"}`,
      );
    }

    count++;

    if (
      !equals(
        expected.wantRequestSchema,
        reqAny.message,
        expected.wantRequest as MessageShape<typeof expected.wantRequestSchema>,
      )
    ) {
      throw new Error(
        `request mismatch for call #${count - 1}: got ${JSON.stringify(reqAny.message)}`,
      );
    }

    if (expected.wantError) {
      throw expected.wantError;
    }

    if (expected.wantResponse) {
      return expected.wantResponse();
    }

    throw new Error(`no wantResponse or wantError configured for call #${count - 1}`);
  };
}
