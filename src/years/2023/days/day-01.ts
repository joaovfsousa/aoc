import { Runner } from '@src/runner';

const nums = ['one', 'two', 'three', 'four', 'five', 'six', 'seven', 'eight', 'nine'];

const numsRegex = new RegExp(
  nums
    .map((num) => `(?=(${num}))`)
    .join('|')
    .concat('|\\d'),
  'g'
);

const parseNum = (num: string): string => {
  const parsedNumber = parseInt(num);
  if (!isNaN(parsedNumber)) {
    return num;
  }
  return (nums.indexOf(num) + 1).toString();
};

class DayRunner extends Runner {
  private lineNumbers: string[][];

  parse(): void {
    this.lineNumbers = this.lines.map((line) => {
      return line.replace(/\D/g, '').split('');
    });
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
  part2(): string {
    this.lineNumbers = this.lines.map((line) => {
      return Array.from(line.matchAll(numsRegex))
        .map((x) => {
          return x.find(Boolean)!;
        })
        .map(parseNum);
    });

    return this.part1();
  }
}

export const runner = new DayRunner();
