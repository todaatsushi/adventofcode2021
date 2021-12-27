const getInput = (input: string): number[] =>
  input
    .split("\n")
    .map((num: string): string => num.trim())
    .filter((num: string): boolean => num !== '')
    .map((num: string) => parseInt(num))
    .filter((num: number | typeof NaN) => typeof num === 'number')


export default getInput
