import { readFile } from 'node:fs/promises';

import { getInputFileName } from './file-helpers';

export abstract class Solution {
  lines: string[];

  async getInput(year: number, day: number) {
    const input = await readFile(getInputFileName(year, day));

    this.lines = input.toString().split('\n');
    this.lines.pop();
    return `Input size: ${this.lines.length} lines\n`;
  }

  abstract parsePart1(): string;
  abstract part1(): string;
  abstract parsePart2(): string;
  abstract part2(): string;
}
