import { Runner } from './runner';

interface Day {
  runner: Runner;
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
    const module = (await import(
      `@src/years/${year}/days/day-${day.toString().padStart(2, '0')}.ts`
    )) as Day;

    const { runner } = module;

    const start = performance.now();

    await timeit(async () => {
      await runner.getInput(year, day);
      return 'Loading input...\n';
    });

    await timeit(() => {
      runner.parsePart1();
      return 'Parsing input 1...';
    });

    await timeit(() => {
      const answer = runner.part1();
      return `Part 1: ${answer}\n`;
    });

    await timeit(() => {
      runner.parsePart2();
      return 'Parsing input 2...';
    });

    await timeit(() => {
      const answer = runner.part2();
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
