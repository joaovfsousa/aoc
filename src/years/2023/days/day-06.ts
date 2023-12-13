import { splitOnSpaces } from '@src/shared/parsing';
import { Solution } from '@src/solution';

class Race {
  constructor(
    public time: number,
    public distance: number
  ) {}
}

class DaySolution extends Solution {
  races: Race[] = [];
  parsePart1() {
    const [timeLine, distanceLine] = this.lines;

    const times = splitOnSpaces(timeLine, { ignoreParts: 1 });
    const distances = splitOnSpaces(distanceLine, { ignoreParts: 1 });

    for (let i = 0; i < times.length; i++) {
      const time = times[i];
      const distance = distances[i];

      this.races.push(new Race(parseInt(time), parseInt(distance)));
    }

    return 'DONE';
  }

  part1(): string {
    return this.races
      .reduce((prev, race) => {
        let total = 0;
        for (let time = Math.floor(race.time / 2); time >= 1; time--) {
          const speed = race.time - time;

          if (speed * time > race.distance) {
            total = total += time === speed ? 1 : 2;
          }
        }

        if (total) {
          return total * prev;
        }

        return prev;
      }, 1)
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
