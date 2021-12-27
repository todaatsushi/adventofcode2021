const getInput = (input: string): number[] => input.split("\n").map((num: string) => parseInt(num)).filter((num: number) => num !== NaN)

export default getInput
