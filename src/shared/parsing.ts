export function isDigit(digit: string): boolean {
  return digit >= '0' && digit <= '9';
}

export function pureParseInt(n: string) {
  return parseInt(n);
}

export function isSpace(n: string) {
  return [' ', '\n', '\r'].includes(n);
}

export function splitOnSpaces(
  str: string,
  {
    ignoreParts,
  }: {
    ignoreParts: number;
  } = { ignoreParts: 0 }
) {
  const result: string[] = [];

  let temp = '';
  let count = 0;
  for (const char of str.split('')) {
    if (temp && isSpace(char)) {
      if (ignoreParts <= count) {
        result.push(temp);
      }
      temp = '';
      count++;
    } else if (!isSpace(char)) {
      temp = temp.concat(char);
    }
  }

  if (temp && ignoreParts <= count) {
    result.push(temp);
  }

  return result;
}
