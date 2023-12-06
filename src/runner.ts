import { readFile } from 'node:fs/promises';
import path from 'node:path';

export abstract class Runner {
  lines: string[];

  async getInput(year: number, day: number) {
    const input = await readFile(
      path.join(
        __dirname,
        `years/${year}/inputs/day-${day.toString().padStart(2, '0')}.txt`
      )
    );

    this.lines = input.toString().split('\n');
    this.lines.pop();
    console.log(`Input size: ${this.lines.length} lines\n`);
  }

  abstract parsePart1(): void;
  abstract part1(): string;
  abstract parsePart2(): void;
  abstract part2(): string;
}
