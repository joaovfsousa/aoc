import { Solution } from '@src/solution';

class Card {
  public cardNumbers: Set<number>;
  constructor(
    public winningNumbers: number[],
    cardNumbers: number[]
  ) {
    this.cardNumbers = new Set(cardNumbers);
  }
}

class DaySolution extends Solution {
  private cards: Card[];

  parsePart1() {
    this.cards = this.lines.map((line) => {
      const [winningNumbersInput, cardNumbersInput] = line
        .split(': ')
        .at(1)!
        .split(' | ');

      const winningNumbers: number[] = winningNumbersInput
        .split(' ')
        .filter((n) => n !== '')
        .map((n) => parseInt(n.trim()));

      const cardNumbers: number[] = cardNumbersInput
        .split(' ')
        .filter((n) => n !== '')
        .map((n) => parseInt(n.trim()));

      return new Card(winningNumbers, cardNumbers);
    });

    return 'DONE';
  }

  part1(): string {
    return this.cards
      .reduce((acc, card) => {
        const cardPoints = card.winningNumbers.filter((wn) =>
          card.cardNumbers.has(wn)
        ).length;

        if (!cardPoints) return acc;

        return acc + 2 ** (cardPoints - 1);
      }, 0)
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
