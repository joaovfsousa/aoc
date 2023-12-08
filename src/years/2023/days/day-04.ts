import { Solution } from '@src/solution';

class Card {
  public cardNumbers: Set<number>;
  public copies = 1;
  private matches: number[];
  constructor(
    public winningNumbers: number[],
    cardNumbers: number[]
  ) {
    this.cardNumbers = new Set(cardNumbers);
  }

  getMatches() {
    if (!this.matches) {
      this.matches = this.winningNumbers.filter((wn) => this.cardNumbers.has(wn));
    }

    return this.matches;
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
        const cardPoints = card.getMatches().length;

        if (!cardPoints) return acc;

        return acc + 2 ** (cardPoints - 1);
      }, 0)
      .toString();
  }

  parsePart2() {
    return 'DONE';
  }

  part2(): string {
    return this.cards
      .reduce((acc, card, index) => {
        const matches = card.getMatches().length;

        this.addCopiesToNextCards(matches, card.copies, index);

        return acc + card.copies;
      }, 0)
      .toString();
  }

  private addCopiesToNextCards(
    copiesToAdd: number,
    originalCardCopies: number,
    originalCardIndex: number
  ) {
    for (let i = 1; i <= copiesToAdd; i++) {
      const index = originalCardIndex + i;
      if (index < this.cards.length) {
        this.cards[index].copies += originalCardCopies;
      }
    }
  }
}

export const solution = new DaySolution();
