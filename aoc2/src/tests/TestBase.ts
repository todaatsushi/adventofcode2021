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

export default TestBase
