import { isDigit } from '@src/shared/parsing';
import { Solution } from '@src/solution';

interface Position {
  x: number;
  y: number;
}

class SSymbol {
  constructor(
    public value: string,
    public position: Position
  ) {}
}

class SNumber {
  constructor(
    public value: string,
    public position: Position,
    public lastDigitX: number
  ) {}

  toNumber(): number {
    return parseInt(this.value);
  }

  isAdjacentToASymbol(symbols: SSymbol[]): boolean {
    return symbols.some(this.isAdjacentTo.bind(this));
  }

  isAdjacentTo(symbol: SSymbol): boolean {
    if (Math.abs(symbol.position.y - this.position.y) > 1) return false;

    return (
      symbol.position.x >= this.position.x - 1 && symbol.position.x <= this.lastDigitX + 1
    );
  }
}

class DaySolution extends Solution {
  // TODO: split in a map by Y, so we can reduce the amount of items we have to search for.
  private numbers: SNumber[] = [];
  private symbols: SSymbol[] = [];

  parsePart1() {
    let actualNumber: SNumber | undefined;

    this.lines.forEach((line, y) => {
      line.split('').forEach((char, x) => {
        if (isDigit(char)) {
          if (actualNumber) {
            actualNumber.value += char;
            actualNumber.lastDigitX = x;
          } else {
            actualNumber = new SNumber(char, { x, y }, x);
          }
          return;
        }

        if (actualNumber) {
          this.numbers.push(actualNumber);
          actualNumber = undefined;
        }

        if (char === '.') return;

        this.symbols.push(new SSymbol(char, { x, y }));
      });
    });

    return 'DONE';
  }

  part1(): string {
    return this.numbers
      .filter((num) => {
        const a = num.isAdjacentToASymbol(this.symbols);
        return a;
      })
      .reduce((acc, cur) => {
        return acc + cur.toNumber();
      }, 0)
      .toString();
  }

  parsePart2() {
    return 'Nothing to do';
  }

  part2(): string {
    return this.symbols
      .filter((s) => s.value === '*')
      .map((s) => {
        const value = {
          gearRatio: 1,
          adjacentNumbersCount: 0,
        };

        const adjacentNumbers = this.numbers.filter((num) => num.isAdjacentTo(s));

        value.adjacentNumbersCount = adjacentNumbers.length;

        if (value.adjacentNumbersCount === 2) {
          const [numA, numB] = adjacentNumbers;
          value.gearRatio = numA.toNumber() * numB.toNumber();
        }

        return value;
      })
      .filter((v) => v.adjacentNumbersCount === 2)
      .reduce((acc, actual) => {
        return acc + actual.gearRatio;
      }, 0)
      .toString();
  }
}

export const solution = new DaySolution();
