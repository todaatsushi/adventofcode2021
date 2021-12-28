import Move from '../data/Move'

interface directionArray extends Array< string | number> {
  0: string,
  1: number
}

const getInput = (inputString: string): Move[] => {
  return (
    inputString
      .split('\n')
      .map((input: string): string => input.trim())
      .map((input: string): string[] => input.split(' '))
      .map((input: string[]): directionArray => [input[0], parseInt(input[1])])
      .filter((inputs: directionArray): boolean => {
        return inputs.map((val): boolean => {
          return [null, NaN, '', ' '].includes(val) === false
        }).every((val: boolean): boolean => val === true)
      })
      .map((inputs: directionArray): Move => {
        return new Move(inputs[0], inputs[1])
      })
  )
}

export default getInput
