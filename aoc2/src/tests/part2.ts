import { TestBase, TestGetPositionTestCase } from './TestBase'
import rawTestInput from '../inputs/test'
import getPositionWithAim from '../data/getPositionWithAim'

const testPart2 = () => {
  const testGetPositionTestCase = new TestGetPositionTestCase(rawTestInput, 900, getPositionWithAim)
  testGetPositionTestCase.test()
}

export default testPart2
