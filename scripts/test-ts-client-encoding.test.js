// CI guard for the RFC 3986 paramsSerializer in ts-client/utils.ts and the
// import injected into ts-client/*/rest.ts by scripts/patch-ts-client-encoding.sh.
// Run with: node --test (Node >= 20).

const { test } = require("node:test");
const assert = require("node:assert/strict");
const fs = require("node:fs");
const path = require("node:path");

const tsClientDir = path.resolve(__dirname, "..", "ts-client");
const utilsPath = path.join(tsClientDir, "utils.ts");

// Extract and execute the actual function from utils.ts so the algorithm tests
// run against the shipped code, not a copy. Strips the few TS annotations
// inside the function so node can compile it as JS.
function loadParamsSerializer() {
  const src = fs.readFileSync(utilsPath, "utf8");
  const match = src.match(/export function paramsSerializer\(([^)]*)\):\s*string\s*\{([\s\S]*?)\n\}/);
  if (!match) throw new Error("paramsSerializer not found in ts-client/utils.ts");
  const params = match[1].replace(/:\s*Record<string,\s*any>/g, "");
  const body = match[2].replace(/:\s*string\[\]/g, "");
  return new Function(params, body);
}

const paramsSerializer = loadParamsSerializer();

test("base64 pagination key with '+' is percent-encoded to %2B", () => {
  assert.equal(
    paramsSerializer({ "pagination.key": "BPyFnvTHtPIy2pm3+Ih2Apuyiok=" }),
    "pagination.key=BPyFnvTHtPIy2pm3%2BIh2Apuyiok%3D",
  );
});

test("all base64 specials (+ / =) are encoded", () => {
  assert.equal(
    paramsSerializer({ "pagination.key": "ab+cd/ef==" }),
    "pagination.key=ab%2Bcd%2Fef%3D%3D",
  );
});

test("dots in keys remain literal (gRPC-gateway routes on 'pagination.key')", () => {
  const out = paramsSerializer({ "pagination.key": "abc" });
  assert.ok(out.startsWith("pagination.key="), `dot was encoded: ${out}`);
});

test("undefined and null values are skipped, not stringified", () => {
  assert.equal(
    paramsSerializer({ a: "x", b: undefined, c: null, d: "y" }),
    "a=x&d=y",
  );
});

test("arrays produce repeated key=value pairs (events= form)", () => {
  assert.equal(
    paramsSerializer({ events: ["tx.height=10", "message.action=/cosmos.bank.MsgSend"] }),
    "events=tx.height%3D10&events=message.action%3D%2Fcosmos.bank.MsgSend",
  );
});

test("booleans and numbers stringify", () => {
  assert.equal(
    paramsSerializer({ "pagination.reverse": true, "pagination.offset": 0 }),
    "pagination.reverse=true&pagination.offset=0",
  );
});

test("ts-client/utils.ts exists and exports paramsSerializer", () => {
  assert.ok(fs.existsSync(utilsPath), "ts-client/utils.ts is missing -- run scripts/patch-ts-client-encoding.sh");
  const src = fs.readFileSync(utilsPath, "utf8");
  assert.match(src, /export function paramsSerializer/);
});

test("every ts-client/*/rest.ts imports paramsSerializer and uses it on axios.create", () => {
  const restFiles = [];
  for (const name of fs.readdirSync(tsClientDir)) {
    const candidate = path.join(tsClientDir, name, "rest.ts");
    if (fs.existsSync(candidate)) restFiles.push(candidate);
  }
  assert.ok(restFiles.length >= 16, `expected >=16 rest.ts files, found ${restFiles.length}`);
  for (const f of restFiles) {
    const src = fs.readFileSync(f, "utf8");
    const rel = path.relative(tsClientDir, f);
    assert.match(src, /from "\.\.\/utils"/, `${rel} missing utils import -- run scripts/patch-ts-client-encoding.sh`);
    assert.match(src, /paramsSerializer,\s*\n\s*\}\);/, `${rel} not using paramsSerializer in axios.create`);
  }
});
