var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
import { equals } from "@bufbuild/protobuf";
export function newTestInterceptor(calls) {
    let count = 0;
    return (next) => (req) => __awaiter(this, void 0, void 0, function* () {
        var _a, _b;
        var _c;
        const reqAny = req;
        const expected = calls[count];
        if (!expected) {
            throw new Error(`received an unexpected client call of type ${(_c = (_b = (_a = reqAny.method) === null || _a === void 0 ? void 0 : _a.input) === null || _b === void 0 ? void 0 : _b.typeName) !== null && _c !== void 0 ? _c : "unknown"}`);
        }
        count++;
        if (!equals(expected.wantRequestSchema, reqAny.message, expected.wantRequest)) {
            throw new Error(`request mismatch for call #${count - 1}: got ${JSON.stringify(reqAny.message)}`);
        }
        if (expected.wantError) {
            throw expected.wantError;
        }
        if (expected.wantResponse) {
            return expected.wantResponse();
        }
        throw new Error(`no wantResponse or wantError configured for call #${count - 1}`);
    });
}
