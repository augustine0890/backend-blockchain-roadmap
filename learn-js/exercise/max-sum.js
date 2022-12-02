function maxSum(num) {
  return num < 0 ? 0 : num * (num + 1) / 2;
}