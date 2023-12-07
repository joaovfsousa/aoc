import path from 'node:path';

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
