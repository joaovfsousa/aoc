import { execSync } from 'node:child_process';
import { copyFile, writeFile } from 'node:fs/promises';
import path from 'node:path';

import { fileExists, getInputFileName, getSolutionFileName } from './file-helpers';

(async () => {
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
    await writeFile(inputFileName, '');
  }

  console.log('Files created successfully');

  execSync(
    `tmux splitw -v -c '#{pane_current_path}' nvim ${solutionFileName} ${inputFileName} && tmux select-pane -U \\; swap-pane -s '!' \\; select-pane -U`
  );

  execSync(`pnpm run dev ${year} ${day}`, { stdio: 'inherit' });
})();
