import { Solution } from '@src/solution';

class DaySolution extends Solution {
  private lineNumbers: string[][];

  private static readonly nums = [
    'one',
    'two',
    'three',
    'four',
    'five',
    'six',
    'seven',
    'eight',
    'nine',
  ];

  private static readonly numsRegex = new RegExp(
    this.nums
      .map((num) => `(?=(${num}))`)
      .join('|')
      .concat('|\\d'),
    'g'
  );

  parsePart1() {
    this.lineNumbers = this.lines.map((line) => {
      return line.replace(/\D/g, '').split('');
    });

    return 'DONE';
  }

  part1(): string {
    return this.lineNumbers
      .reduce((acc, curr) => {
        const first = curr.at(0)!;
        const last = curr.at(-1)!;

        const number = parseInt(first + last);
        return acc + number;
      }, 0)
      .toString();
  }

  parsePart2() {
    this.lineNumbers = this.lines.map((line) => {
      return Array.from(line.matchAll(DaySolution.numsRegex))
        .map((x) => x.find(Boolean)!)
        .map(DaySolution.parseNum);
    });

    return 'DONE';
  }

  part2(): string {
    return this.part1();
  }

  private static parseNum(num: string): string {
    const parsedNumber = parseInt(num);
    if (!isNaN(parsedNumber)) {
      return num;
    }

    return (DaySolution.nums.indexOf(num) + 1).toString();
  }
}

export const solution = new DaySolution();
