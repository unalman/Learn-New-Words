import { describe, expect, it } from "vitest";
import { textValidationAndTrim } from "../../libs/common";

describe("textValidation tests", () => {
  it("when send empty string, function should be return false", () => {
    expect(textValidationAndTrim("")).toBe(false)
  });
  it("when send white-space string, function should be return false", () => {
    expect(textValidationAndTrim(" ")).toBe(false)
  });
  it("when send any character string, function should be return true", () => {
    expect(textValidationAndTrim("test")).toBe(true)
  });
});
