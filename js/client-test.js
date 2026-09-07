var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
import { describe, it, expect } from "bun:test";
import { create } from "@bufbuild/protobuf";
import { VersionSchema, VersionServiceGetRequestSchema, VersionServiceGetResponseSchema, } from "./fits/api/v1/version_pb";
import { newClient } from "./client";
import { newTestInterceptor } from "./test-interceptor";
function unaryResponse(message) {
    return {
        stream: false,
        message: message,
        header: new Headers(),
        trailer: new Headers(),
        service: undefined,
        method: undefined,
    };
}
describe("client", () => {
    it("works with test interceptor", () => __awaiter(void 0, void 0, void 0, function* () {
        var _a;
        const c = newClient({
            baseUrl: "http://this-is-just-for-testing",
            interceptors: [
                newTestInterceptor([
                    {
                        wantRequest: create(VersionServiceGetRequestSchema, {}),
                        wantRequestSchema: VersionServiceGetRequestSchema,
                        wantResponse: () => unaryResponse(create(VersionServiceGetResponseSchema, {
                            version: create(VersionSchema, { version: "1.0" }),
                        })),
                    },
                ]),
            ],
        });
        const v = yield c.apiv1().version().get({});
        expect((_a = v.version) === null || _a === void 0 ? void 0 : _a.version).toBe("1.0");
    }));
});
