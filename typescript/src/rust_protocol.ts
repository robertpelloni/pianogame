/**
 * Typed client protocol for the bobmani Rust hub (src/server.rs).
 *
 * Wire format mirrors serde's internally-tagged enums exactly:
 *   request:  {"action":"ping"}  |  {"action":"parse_sm","payload":{...}}
 *   response: {"type":"pong","data":{...}}  |  {"type":"error","data":{...}}
 *
 * Why this file exists: docs claimed a formatRustRequest/parseRustResponse
 * bridge since v0.2.6, but ws_bridge.ts is only a raw socket upgrade
 * handler. This is the real, tested protocol layer for all 8 actions.
 */

export type RustAction =
  | "ping"
  | "status"
  | "parse_sm"
  | "analyze_density"
  | "extract_audio_features"
  | "detect_onsets"
  | "calculate_ex_score"
  | "predict_placement";

export interface ParseSmPayload { content: string }
export interface AnalyzeDensityPayload { chart_notes: Array<[number, string]> }
export interface ExtractAudioPayload { samples: number[] }
export interface DetectOnsetsPayload { onset_salience: number[]; threshold: number }
export interface CalculateExScorePayload { epg: number; lpg: number; egr: number; lgr: number; notes: number }
export interface PredictPlacementPayload { onset_salience: number[]; threshold: number }

export type RustRequestPayload =
  | ParseSmPayload
  | AnalyzeDensityPayload
  | ExtractAudioPayload
  | DetectOnsetsPayload
  | CalculateExScorePayload
  | PredictPlacementPayload;

/** Serializes a request exactly as the Rust WsRequest enum expects it. */
export function formatRustRequest(action: "ping" | "status"): string;
export function formatRustRequest(action: Exclude<RustAction, "ping" | "status">, payload: RustRequestPayload): string;
export function formatRustRequest(action: RustAction, payload?: RustRequestPayload): string {
  if (action === "ping" || action === "status") {
    return JSON.stringify({ action });
  }
  if (payload === undefined) {
    throw new Error(`action ${action} requires a payload`);
  }
  return JSON.stringify({ action, payload });
}

export type RustResponse =
  | { type: "pong"; data: { timestamp: number } }
  | { type: "status"; data: { version: string; healthy: boolean; submodules: string[] } }
  | { type: "sm_metadata"; data: { title: string; artist: string; charts_count: number } }
  | { type: "density_analysis"; data: { nps: number; length: number; stream_percentage: number; jack_percentage: number } }
  | { type: "audio_features"; data: { frames_count: number; bands: number; channels: number } }
  | { type: "onset_detections"; data: { peaks: number[]; filtered_peaks: number[]; total_frames: number } }
  | { type: "score_analysis"; data: { ex_score: number; max_ex_score: number; accuracy_rate: number; is_valid: boolean } }
  | { type: "placement_prediction"; data: { placements: number[]; total_frames: number; model_loaded: boolean; source: string } }
  | { type: "error"; data: { message: string } };

const KNOWN_TYPES = new Set([
  "pong", "status", "sm_metadata", "density_analysis", "audio_features",
  "onset_detections", "score_analysis", "placement_prediction", "error",
]);

/** Parses a hub response; throws on malformed JSON or an unknown type tag. */
export function parseRustResponse(text: string): RustResponse {
  const parsed = JSON.parse(text) as RustResponse;
  if (!parsed || typeof parsed !== "object" || !KNOWN_TYPES.has((parsed as { type: string }).type)) {
    throw new Error(`unknown rust response: ${text.slice(0, 120)}`);
  }
  return parsed;
}

/** All 8 actions, for completeness checks against GET /health. */
export const RUST_ACTIONS: RustAction[] = [
  "ping", "status", "parse_sm", "analyze_density",
  "extract_audio_features", "detect_onsets", "calculate_ex_score", "predict_placement",
];
