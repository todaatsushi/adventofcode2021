import getInput from '../inputs/getInput'
import Move from './Move'

const getPosition = (inputString: string): number => {
  const depthMoves = getInput(inputString)
  let totalDepth: number = 0
  let totalHorizonal: number = 0

  let currentMove: Move
  for (let index: number = 0; index < depthMoves.length; index++) {

    currentMove = depthMoves[index]
    if (currentMove.isHorizontal()) {
      totalHorizonal += currentMove.movementValue
    } else {

      if (currentMove.direction === 'up') {
        totalDepth -= currentMove.movementValue
      } else {
        totalDepth += currentMove.movementValue
      }

    }
  }
  return totalHorizonal * totalDepth
}

export default getPosition
