import getInput from '../inputs/getInput'
import rawTestInput from '../inputs/test'
import { TestBase, TestGetPositionTestCase } from './TestBase'
import Move from '../data/Move'
import getPosition from '../data/getPosition'

class InputTestCase extends TestBase {
  result: Move[]
  expected: Move[]

  constructor(expected: Move[], input: string) {
    super()
    this.expected = expected
    this.result = getInput(input)
  }

  resultIsExpected(): boolean {
    return (
      Array.isArray(this.expected) && Array.isArray(this.result)
      && this.expected.length === this.result.length
      && this.expected.every(
        (move: Move, index: number) => {
          return (
            move.direction === this.result[index].direction
            && move.movementValue === this.result[index].movementValue
          )
        }
      )
    )
  }
}

const testPart1 = () => {
  const expectedMoves = [
    new Move('forward', 5),
    new Move('down', 5),
    new Move('forward', 8),
    new Move('up', 3),
    new Move('down', 8),
    new Move('forward', 2),
  ]

  const testInputTestCase = new InputTestCase(expectedMoves, rawTestInput)
  testInputTestCase.test()

  const testGetPositionTestCase = new TestGetPositionTestCase(rawTestInput, 150, getPosition)
  testGetPositionTestCase.test()
}

export default testPart1
