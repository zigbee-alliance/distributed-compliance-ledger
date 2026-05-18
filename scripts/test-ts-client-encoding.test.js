// CI guard for the RFC 3986 paramsSerializer injected into ts-client/*/rest.ts
// by scripts/patch-ts-client-encoding.sh. Run with: node --test (Node >= 20).

const { test } = require("node:test");
const assert = require("node:assert/strict");
const fs = require("node:fs");
const path = require("node:path");

// Keep in sync with scripts/patch-ts-client-encoding.sh
const paramsSerializer = (params) => {
  const parts = [];
  for (const key of Object.keys(params)) {
    const value = params[key];
    if (value === undefined || value === null) continue;
    const ek = encodeURIComponent(key);
    if (Array.isArray(value)) {
      for (const item of value) parts.push(`${ek}=${encodeURIComponent(String(item))}`);
    } else {
      parts.push(`${ek}=${encodeURIComponent(String(value))}`);
    }
  }
  return parts.join("&");
};

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

test("every ts-client/*/rest.ts has the paramsSerializer patch applied", () => {
  const tsClientDir = path.resolve(__dirname, "..", "ts-client");
  const restFiles = [];
  for (const name of fs.readdirSync(tsClientDir)) {
    const candidate = path.join(tsClientDir, name, "rest.ts");
    if (fs.existsSync(candidate)) restFiles.push(candidate);
  }
  assert.ok(restFiles.length >= 16, `expected >=16 rest.ts files, found ${restFiles.length}`);
  for (const f of restFiles) {
    const src = fs.readFileSync(f, "utf8");
    assert.ok(
      src.includes("paramsSerializer") && src.includes("encodeURIComponent"),
      `${path.relative(tsClientDir, f)} is missing the patch -- run scripts/patch-ts-client-encoding.sh`,
    );
  }
});