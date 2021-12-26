const getIncrementCountInNums = (nums: number[]): number => {
  if (nums.length === 0) {
    return 0
  }

  let current: number
  let prev: number = nums[0]
  let count: number = 0

  for (let index = 1; index < nums.length; index++){
    current = nums[index]
    if (current > prev) {
      count++
    }
    prev = current
  }
  return count
}

export default getIncrementCountInNums
