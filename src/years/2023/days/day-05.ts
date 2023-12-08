import { pureParseInt } from '@src/shared/parsing';
import { Solution } from '@src/solution';

class Range {
  constructor(
    public destinationStart: number,
    public sourceStart: number,
    public rangeLength: number
  ) {}

  sourceToDestination(sourceIndex: number): number | null {
    if (
      sourceIndex >= this.sourceStart &&
      sourceIndex < this.sourceStart + this.rangeLength
    ) {
      return this.destinationStart + (sourceIndex - this.sourceStart);
    }
    return null;
  }
}

class SMap {
  constructor(public ranges: Range[] = []) {}

  sort() {
    this.ranges.sort((rA, rB) => rA.sourceStart - rB.sourceStart);
  }

  addRange(range: Range) {
    this.ranges.push(range);
  }
}

class DaySolution extends Solution {
  private maps: SMap[] = [];
  private seeds: number[];

  parsePart1() {
    this.seeds = this.lines.at(0)!.split(': ').at(1)!.split(' ').map(pureParseInt);

    this.lines.slice(2, this.lines.length).forEach((line) => {
      if (line === '') {
        return;
      }

      if (line.endsWith(':')) {
        this.maps.push(new SMap());
        return;
      }

      this.maps
        .at(-1)!
        .addRange(
          new Range(...(line.split(' ').map(pureParseInt) as [number, number, number]))
        );
    });

    this.maps.forEach((map) => map.sort());

    return 'DONE';
  }

  part1(): string {
    return this.seeds
      .map((seed) => {
        let lastValue = seed;
        this.maps.forEach((map) => {
          const correspondingRange = map.ranges.find(
            (r) => !!r.sourceToDestination(lastValue)
          );

          if (correspondingRange) {
            lastValue = correspondingRange.sourceToDestination(lastValue)!;
          }
        });
        return lastValue;
      })
      .reduce((prev, curr) => {
        if (prev < curr) {
          return prev;
        }
        return curr;
      }, Infinity)
      .toString();
  }

  parsePart2() {
    this.lines.map((line) => {
      void line;
    });

    return 'TODO';
  }

  part2(): string {
    return 'TODO';
  }
}

export const solution = new DaySolution();
