import assert from "assert";
import { formatRustRequest, parseRustResponse, RUST_ACTIONS } from "./rust_protocol.js";

// All 8 Rust hub actions are covered (mirrors GET /health "actions").
assert.strictEqual(RUST_ACTIONS.length, 8);
assert.ok(RUST_ACTIONS.includes("predict_placement"));

// Unit-variant requests carry no payload key (serde internally-tagged).
assert.strictEqual(formatRustRequest("ping"), '{"action":"ping"}');
assert.strictEqual(formatRustRequest("status"), '{"action":"status"}');

// Payload requests match src/server.rs WsRequest field names exactly.
assert.strictEqual(
  formatRustRequest("parse_sm", { content: "#TITLE:X;" }),
  '{"action":"parse_sm","payload":{"content":"#TITLE:X;"}}'
);
assert.strictEqual(
  formatRustRequest("predict_placement", { onset_salience: [0.9, 0.1], threshold: 0.5 }),
  '{"action":"predict_placement","payload":{"onset_salience":[0.9,0.1],"threshold":0.5}}'
);
assert.strictEqual(
  formatRustRequest("calculate_ex_score", { epg: 50, lpg: 50, egr: 10, lgr: 10, notes: 120 }),
  '{"action":"calculate_ex_score","payload":{"epg":50,"lpg":50,"egr":10,"lgr":10,"notes":120}}'
);
assert.throws(() => formatRustRequest("parse_sm" as never, undefined as never));

// Responses: every WsResponse tag parses; unknown tags and junk throw.
const pong = parseRustResponse('{"type":"pong","data":{"timestamp":42}}');
assert.strictEqual(pong.type, "pong");
const placement = parseRustResponse(
  '{"type":"placement_prediction","data":{"placements":[2,5],"total_frames":8,"model_loaded":false,"source":"onset_heuristic_no_model"}}'
);
assert.strictEqual(placement.type, "placement_prediction");
if (placement.type === "placement_prediction") {
  assert.deepStrictEqual(placement.data.placements, [2, 5]);
  assert.strictEqual(placement.data.model_loaded, false);
}
assert.throws(() => parseRustResponse('{"type":"nope","data":{}}'));
assert.throws(() => parseRustResponse("not json"));

console.log("TypeScript RustProtocol tests passed cleanly!");
