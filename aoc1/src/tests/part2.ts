import getDataWithWindow from '../getData'
import { part1PuzzleInputs, part1TestInputs, part2TestInputs } from '../inputs/inputs'

class AggregatorTestCase {
  result: number[]
  expected: number[]

  constructor(nums: number[], expected: number[], windowSize: number) {
    this.result = getDataWithWindow(nums, windowSize)
    this.expected = expected
  }

  logError() {
    console.log(`Failed to correctly aggregate arr: expected ${this.expected}, got ${this.result}`)
  }

  logSuccess() {
    console.log(`Correctly aggregated arr: expected and got ${this.expected}`)
  }

  successful() {
    return (
      Array.isArray(this.expected) && Array.isArray(this.result)
      && this.expected.length === this.result.length
      && this.expected.every((val: number, index: number) => val === this.result[index])
    )
  }

  test() {
    if (!this.successful()) {
      this.logError()
    } else {
      this.logSuccess()
    }
  }
}

const testPart2 = () => {
  console.log("Testing part 2: \n")
  const numsToTotal: number[] = [1,2,3,1,2,3]
  const numsToTotalTestCase: AggregatorTestCase = new AggregatorTestCase(numsToTotal, [6,6,6,6,], 3)
  numsToTotalTestCase.test()

  const nullInputTestCase: AggregatorTestCase = new AggregatorTestCase([], [], 3)
  nullInputTestCase.test()

  const testInputExpectedData: number[] = [
    607,
    618,
    618,
    617,
    647,
    716,
    769,
    792,
  ]
  const testInputTestCase: AggregatorTestCase =  new AggregatorTestCase(part2TestInputs, testInputExpectedData, 3)
  testInputTestCase.test()
}

export default testPart2
