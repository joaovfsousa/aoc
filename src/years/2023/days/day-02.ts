import { Solution } from '@src/solution';

interface GameSet {
  red: number;
  green: number;
  blue: number;
}

type Game = GameSet[];

class DaySolution extends Solution {
  private games: Game[];

  parsePart1() {
    this.games = this.lines.map((line) =>
      line
        .split(': ') // game: sets
        .at(1)!
        .split('; ') // set1; set 2; set 3
        .map((set) => {
          return set
            .split(', ') // 1 color1, 2 color2, 3 color3
            .reduce(
              (acc, countAndColor) => {
                const [count, color] = countAndColor.split(' '); // 1 color1 => ['1', 'color1']
                return { ...acc, [color]: parseInt(count) };
              },
              { red: 0, green: 0, blue: 0 }
            );
        })
    );

    return 'DONE';
  }

  part1(): string {
    const maxCubesPerSet: GameSet = {
      red: 12,
      green: 13,
      blue: 14,
    };

    return this.games
      .reduce((acc, game, index) => {
        const isPossible = game.every(
          (set) =>
            set.red <= maxCubesPerSet.red &&
            set.green <= maxCubesPerSet.green &&
            set.blue <= maxCubesPerSet.blue
        );

        if (isPossible) {
          return acc + index + 1;
        }

        return acc;
      }, 0)
      .toString();
  }

  parsePart2() {
    return 'Nothing to do';
  }

  part2(): string {
    return this.games
      .map((game) => {
        return game.reduce(
          (acc, set) => {
            (Object.keys(set) as (keyof GameSet)[]).forEach((color) => {
              if (set[color] > acc[color]) {
                acc[color] = set[color];
              }
            });

            return acc;
          },
          { red: 0, green: 0, blue: 0 }
        );
      })
      .reduce((acc, set) => {
        return acc + set.red * set.green * set.blue;
      }, 0)
      .toString();
  }
}

export const solution = new DaySolution();
