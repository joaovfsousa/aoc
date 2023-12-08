export function isDigit(digit: string): boolean {
  return digit >= '0' && digit <= '9';
}

export function pureParseInt(n: string) {
  return parseInt(n);
}
