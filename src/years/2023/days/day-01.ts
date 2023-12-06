import { Runner } from '@src/runner';

class Day01Runner extends Runner {
  private lineNumbers: string[];

  parse(): void {
    this.lineNumbers = this.lines.map((line) => {
      return line.replace(/\D/g, '');
    });
  }

  part1(): string {
    return this.lineNumbers
      .reduce((acc, curr) => {
        const first = curr.at(0);
        const last = curr.at(-1);

        if (!first || !last) return acc;
        const number = parseInt(first + last);
        return acc + number;
      }, 0)
      .toString();
  }
  part2(): string {
    throw new Error('Method not implemented.');
  }
}

export const runner = new Day01Runner();
