import { execSync } from 'node:child_process';
import { copyFile, writeFile } from 'node:fs/promises';
import path from 'node:path';

import { config } from 'dotenv';

import { fileExists, getInputFileName, getSolutionFileName } from './file-helpers';

(async () => {
  config();
  const [, , yearAsString, dayAsString] = process.argv;

  const year = parseInt(yearAsString);
  const day = parseInt(dayAsString);

  const solutionFileName = getSolutionFileName(year, day);
  const inputFileName = getInputFileName(year, day);

  const solutionFileExists = await fileExists(solutionFileName);
  const inputFileExists = await fileExists(inputFileName);

  if (solutionFileExists) {
    console.warn('Solution already exists');
  } else {
    await copyFile(path.join('templates/day'), solutionFileName);
  }

  if (inputFileExists) {
    console.warn('Input already exists');
  } else {
    const aocSession = process.env.AOC_SESSION;

    let fileContent = '';
    if (aocSession) {
      const response = await fetch(`https://adventofcode.com/${year}/day/${day}/input`, {
        headers: {
          cookie: `session=${aocSession}`,
        },
      });

      fileContent = await response.text();
    }

    await writeFile(inputFileName, fileContent);
  }

  console.log('Files created successfully');

  execSync(
    `tmux splitw -v -c '#{pane_current_path}' zsh -c 'nvim ${solutionFileName} ${inputFileName}; zsh' && tmux select-pane -U \\; swap-pane -s '!' \\; select-pane -U`
  );

  execSync(`pnpm run dev ${year} ${day}`, { stdio: 'inherit' });
})();
