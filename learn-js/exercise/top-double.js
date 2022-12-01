function topDouble(value, top) {
  let val = value;
  while (val < top/2) {
    val = val * 2;
  }
  return val;
}

console.log(topDouble(2, 100));