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

  *toIter() {
    for (let i = 0; i < this.rangeLength; i++) {
      yield this.sourceStart + i;
    }
  }
}

class SMap {
  constructor(public ranges: Range[] = []) {}

  sortBy(key: 'sourceStart' | 'destinationStart' = 'sourceStart') {
    this.ranges.sort((rA, rB) => rA[key] - rB[key]);
  }

  addRange(range: Range) {
    this.ranges.push(range);
  }
}

class DaySolution extends Solution {
  private maps: SMap[] = [];
  private seeds: number[];
  private seedsAsRanges: SMap = new SMap();
  private minLoc = Infinity;

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

    this.maps.forEach((map) => map.sortBy('sourceStart'));

    return 'DONE';
  }

  part1(): string {
    this.seeds.forEach((seed) => {
      const loc = this.getLocationForSeed(seed);

      if (loc < this.minLoc) {
        this.minLoc = loc;
      }
    });

    return this.minLoc.toString();
  }

  parsePart2() {
    this.minLoc = Infinity;

    let lastValue: number;

    this.lines
      .at(0)!
      .split(': ')
      .at(1)!
      .split(' ')
      .forEach((numAsString, index) => {
        if (index % 2 === 0) {
          lastValue = pureParseInt(numAsString);
          return;
        }

        this.seedsAsRanges.addRange(new Range(0, lastValue, pureParseInt(numAsString)));
      });

    return 'DONE';
  }

  part2(): string {
    console.log('Total Ranges:', this.seedsAsRanges.ranges.length);

    this.seedsAsRanges.ranges.forEach((range, index) => {
      console.log(`Processing range: ${index + 1}/${this.seedsAsRanges.ranges.length}`);
      for (const seed of range.toIter()) {
        const loc = this.getLocationForSeed(seed);

        if (loc < this.minLoc) {
          this.minLoc = loc;
        }
      }
    });

    return this.minLoc.toString();
  }

  private getLocationForSeed(seed: number) {
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
  }
}

export const solution = new DaySolution();
