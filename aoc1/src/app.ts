import { part1TestInputs, part1PuzzleInputs } from './inputs/inputs'
import getIncrementCountInNums from './getCount'
import getDataWithWindow from './getData'

console.log(
  `Part 1 : answer for advent of code #1 is ${getIncrementCountInNums(part1PuzzleInputs)}`,
)

console.log(
  `Part 2 : answer for advent of code #1 is ${getIncrementCountInNums(getDataWithWindow(part1PuzzleInputs, 3))}`
)
