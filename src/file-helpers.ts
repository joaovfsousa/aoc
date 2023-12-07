import { access } from 'node:fs/promises';
import path from 'node:path';

export async function fileExists(filename: string) {
  try {
    await access(filename);
    return true;
  } catch (err) {
    if ((err as { code: string }).code === 'ENOENT') {
      return false;
    } else {
      throw err;
    }
  }
}

const getFileName = (year: number, day: number, dir: 'inputs' | 'days', ext: string) => {
  return path.join(
    __dirname,
    `/years/${year}/${dir}/day-${day.toString().padStart(2, '0')}.${ext}`
  );
};

export const getInputFileName = (year: number, day: number) => {
  return getFileName(year, day, 'inputs', 'txt');
};

export const getSolutionFileName = (year: number, day: number) => {
  return getFileName(year, day, 'days', 'ts');
};
