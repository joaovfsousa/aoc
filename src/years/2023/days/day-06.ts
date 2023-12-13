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
  race: Race;

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
        const total = this.countWaysToWin(race);

        if (total) {
          return total * prev;
        }
        return prev;
      }, 1)
      .toString();
  }

  parsePart2() {
    const [timeLine, distanceLine] = this.lines;

    const time = splitOnSpaces(timeLine, { ignoreParts: 1 }).join('');
    const distance = splitOnSpaces(distanceLine, { ignoreParts: 1 }).join('');

    this.race = new Race(parseInt(time), parseInt(distance));

    return 'DONE';
  }

  part2(): string {
    return this.countWaysToWin(this.race).toString();
  }

  countWaysToWin(race: Race) {
    let total = 0;
    for (let time = Math.floor(race.time / 2); time >= 1; time--) {
      const speed = race.time - time;

      if (speed * time > race.distance) {
        total = total += time === speed ? 1 : 2;
      }
    }

    return total;
  }
}

export const solution = new DaySolution();
