import { readFile } from 'node:fs/promises';

import { getInputFileName } from './file-helpers';

export abstract class Solution {
  lines: string[];

  async getInput(year: number, day: number) {
    const input = await readFile(getInputFileName(year, day));

    this.lines = input.toString().split('\n');
    this.lines.pop();
    console.log(`Input size: ${this.lines.length} lines\n`);
  }

  abstract parsePart1(): void;
  abstract part1(): string;
  abstract parsePart2(): void;
  abstract part2(): string;
}
