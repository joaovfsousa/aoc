import { Runner } from './runner';

interface Day {
  runner: Runner;
}

function getEllapsedTime(start: number, end: number) {
  return `${(end - start).toFixed(3)}ms`;
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

    await runner.getInput(year, day);
    const afterInput = performance.now();
    console.log(`${getEllapsedTime(start, afterInput)}: Loading input...\n`);

    runner.parsePart1();
    const afterParse1 = performance.now();
    console.log(`${getEllapsedTime(afterInput, afterParse1)}: Parsing input 1...`);

    const part1Answer = runner.part1();
    const afterPart1 = performance.now();
    console.log(`${getEllapsedTime(afterParse1, afterPart1)}: Part 1: ${part1Answer}`);

    runner.parsePart2();
    const afterParse2 = performance.now();
    console.log(`\n${getEllapsedTime(afterPart1, afterParse2)}: Parsing input 2...`);

    const part2Answer = runner.part2();
    const end = performance.now();
    console.log(`${getEllapsedTime(afterParse2, end)}: Part 2: ${part2Answer}`);

    console.log(`\nTotal time ellapsed: ${end - start}ms`);
  } catch (error) {
    console.error('---- Error ----');

    if ((error as Error).message.includes('Cannot find module')) {
      console.error('No runner found for that day');
      return process.exit(1);
    }

    console.error(error);
  }
})();
