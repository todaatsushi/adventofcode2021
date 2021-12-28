class TestBase {
  result: any
  expected: any

  getErrorString = () => `Failed test: expected ${this.expected}, got ${this.result}`
  getSuccessString = () => `Test passed: expected & got ${this.result}!`

  logError() {
    console.log(this.getErrorString())
  }

  logSuccess() {
    console.log(this.getSuccessString())
  }

  resultIsExpected(): boolean {
    return this.result === this.expected
  }

  test() {
    if (this.resultIsExpected() === false) {
      this.logError()
    }
    else {
      this.logSuccess()
    }
  }
}

class TestGetPositionTestCase extends TestBase {
  result: number
  expected: number

  constructor(inputString: string, expected: number, getFunc: Function) {
    super()

    this.expected = expected
    this.result = getFunc(inputString)
  }
}

export { TestBase, TestGetPositionTestCase }
