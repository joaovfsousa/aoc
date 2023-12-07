import { getSolutionFileName } from './file-helpers';
import { Solution } from './solution';

interface Day {
  solution: Solution;
}

function getEllapsedTime(start: number, end: number) {
  return `${(end - start).toFixed(3)}ms`;
}

async function timeit(func: () => Promise<string> | string) {
  const start = performance.now();
  const logMessage = await func();
  const end = performance.now();

  console.log(`${getEllapsedTime(start, end)}: ${logMessage}`);
}

(async () => {
  const [, , yearAsString, dayAsString] = process.argv;

  const year = parseInt(yearAsString);
  const day = parseInt(dayAsString);

  console.log(`---- ${year}, Day ${day} ----\n`);

  try {
    const module = (await import(getSolutionFileName(year, day))) as Day;

    const { solution } = module;

    const start = performance.now();

    await timeit(async () => {
      await solution.getInput(year, day);
      return 'Loading input...\n';
    });

    await timeit(() => {
      solution.parsePart1();
      return 'Parsing input 1...';
    });

    await timeit(() => {
      const answer = solution.part1();
      return `Part 1: ${answer}\n`;
    });

    await timeit(() => {
      solution.parsePart2();
      return 'Parsing input 2...';
    });

    await timeit(() => {
      const answer = solution.part2();
      return `Part 2: ${answer}`;
    });

    const end = performance.now();

    console.log(`\nTotal time ellapsed: ${getEllapsedTime(start, end)}`);
  } catch (error) {
    console.error('---- Error ----');

    if ((error as Error).message.includes('Cannot find module')) {
      console.error('No runner found for that day');
      return;
    }

    console.error(error);
  }
})();
