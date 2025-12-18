import { describe, it, expect } from "vitest";
import { randomString } from "../src/node";

describe("randomString", () => {
  it("should generate random string with specified length", () => {
    const length = 16;
    const result = randomString(length);
    expect(result.length).toBe(length);
  });
});
