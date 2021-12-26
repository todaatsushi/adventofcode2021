import getIncrementCountInNums from './getCount'

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
