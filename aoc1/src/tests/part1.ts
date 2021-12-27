import getIncrementCountInNums from '../getCount'
import { part1TestInputs, part1PuzzleInputs } from '../inputs/inputs'

class TestCase {
  result: number
  expected: number
  testType: string

  constructor(expected: number, input: number[], testType: string) {
    this.result = getIncrementCountInNums(input)
    this.expected = expected
    this.testType = testType
  }

  logError() {
    console.log(`Failed to ${this.testType} correctly: got ${this.result}, should be ${this.expected}!`)
  }

  logSuccess() {
    console.log(`Test type ${this.testType} passed: expected and got ${this.result}!`)
  }

  test() {
    if (this.result !== this.expected) {
      this.logError()
    }
    else if (this.result === this.expected) {
      this.logSuccess()
    }
  }
}

const testPart1 = () => {
  const purelyIncrementingInput: number[] = [1, 2, 3]  // 2 increments
  const purelyIncrementTest: TestCase = new TestCase(2, purelyIncrementingInput, "increment")
  purelyIncrementTest.test()

  const nullInput: number[] = []  // 0 increments
  const nullTest: TestCase = new TestCase(0, nullInput, "handle null")
  nullTest.test()

  const flatInputs: number[] = [1,1,1,1,1,1]  // 0 increments
  const flatTest: TestCase = new TestCase(0, flatInputs, "handle flat")
  flatTest.test()

  const purelyDecreasingInput: number[] = [3,2,1]  // 0 increments
  const purelyDecreasingTest: TestCase = new TestCase(0, purelyDecreasingInput, "handle decreasing")
  purelyDecreasingTest.test()

  const functionalInput: number[] = [1, 0, 1]  // 1 increment
  const functionalTest = new TestCase(1, functionalInput, "full test")
  functionalTest.test()

  const testInputTest = new TestCase(7, part1TestInputs, "example test")
  testInputTest.test()

  console.log("\n")
}

export default testPart1

