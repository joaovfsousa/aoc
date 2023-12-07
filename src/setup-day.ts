import { copyFile, writeFile } from 'node:fs/promises';
import path from 'node:path';

import { fileExists, getInputFileName, getSolutionFileName } from './file-helpers';

(async () => {
  const [, , yearAsString, dayAsString] = process.argv;

  const year = parseInt(yearAsString);
  const day = parseInt(dayAsString);

  const solutionFileExists = await fileExists(getSolutionFileName(year, day));
  const inputFileExists = await fileExists(getInputFileName(year, day));

  if (solutionFileExists || inputFileExists) {
    console.error('Solution or input file already exists');
    process.exit(1);
  }

  await copyFile(path.join('templates/day'), getSolutionFileName(year, day));
  await writeFile(getInputFileName(year, day), '');

  console.log('Files created successfully');
  process.exit(0);
})();
