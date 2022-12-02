function endsWithX(string) {
  const strLen = string.length;
  if (string.charAt(strLen - 1).toLowerCase() === 'x') {
    return true;
  }
  return false;
}

function isAllX(string) {
  for (let i = 0; i < string.length; i++) {
    if (string.charAt(i).toLowerCase() !== 'x') {
      return false;
    }
  }
  return true;
}

function splitAtX(string) {
  if (string.indexOf('x') >= string.length / 2) {
    return string.slice(0, string.indexOf('x'))
  }

  return string.slice(string.indexOf('x') + 1, string.length)
}