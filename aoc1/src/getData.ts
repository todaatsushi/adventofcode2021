const getDataWithWindow = (nums: number[], windowSize: number): number[] => {
  let startIndex: number = 0
  let endIndex: number = windowSize
  let aggregatedNums: number[] = []

  if (endIndex > nums.length) {
    return []
  }

  while (endIndex < nums.length) {
    aggregatedNums.push(nums.slice(startIndex, endIndex).reduce(
      (acc: number | typeof NaN, value: number): number => acc + value
    ))
    startIndex++
    endIndex++
  }
  return aggregatedNums
}

export default getDataWithWindow
