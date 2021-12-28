import getInput from '../inputs/getInput'
import Move from './Move'

const getPositionWithAim = (inputString: string): number => {
  const moves = getInput(inputString)
  let totalDepth: number = 0
  let totalHorizonal: number = 0
  let totalAim: number = 0

  let currentMove: Move
  for (let index: number = 0; index < moves.length; index++) {
    currentMove = moves[index]

    if (currentMove.isHorizontal()) {

      totalHorizonal += currentMove.movementValue
      totalDepth += currentMove.movementValue * totalAim

    } else {

      if (currentMove.direction === 'up') {
        totalAim -= currentMove.movementValue
      } else {
        totalAim += currentMove.movementValue
      }

    }

  }

  return totalHorizonal * totalDepth

}

export default getPositionWithAim
