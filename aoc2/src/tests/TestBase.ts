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

  test() {
    if (this.result !== this.expected) {
      this.logError()
    }
    else if (this.result === this.expected) {
      this.logSuccess()
    }
  }
}

export default TestBase
