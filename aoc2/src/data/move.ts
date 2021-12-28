const validDirections: string[] = ['forward', 'down', 'up']

class Move {
  direction: string
  movementValue: number

  constructor (direction: string, movementValue: number) {
    if (validDirections.includes(direction) === false) {
      throw new Error(
        `Value ${direction} must be a value direction: ${validDirections}`
      )
    }

    this.direction = direction
    this.movementValue = movementValue
  }

  isHorizontal(): boolean {
    return this.direction === 'forward'
  }

  toString() {
    return `[${this.direction},${this.movementValue}]`
  }
}

export default Move
