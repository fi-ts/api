import type { Interceptor, UnaryResponse } from "@connectrpc/connect";
import type { DescMessage, MessageShape } from "@bufbuild/protobuf";
export type ClientCall<Req extends DescMessage = DescMessage, Res extends DescMessage = DescMessage> = {
    wantRequest: MessageShape<Req>;
    wantRequestSchema: Req;
    wantResponse?: () => UnaryResponse<Req, Res>;
    wantError?: Error;
};
export declare function newTestInterceptor(calls: ClientCall[]): Interceptor;
