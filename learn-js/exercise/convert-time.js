const time = '12:10AM';

function convertTo24HrsFormat(time) {
  let newTime
  if (time.length === 7) {
    newTime = time.substring(0, 5)
  } else {
    newTime = time.substring(0, 4)
  }

  let modifier = time.slice(-2)
  let [hours, minutes] = newTime.split(':')
  if (hours.length == 1) hours = '0' + hours
  if (minutes.length == 1) minutes = '0' + minutes

  if (hours === '12') {
    hours = '00';
  }

  if (modifier === 'PM') {
    hours = parseInt(hours, 10) + 12;
  }

  return `${hours}:${minutes}`
}

console.log(`Converted time: ${convertTo24HrsFormat(time)}`)
console.log(`Converted time: ${convertTo24HrsFormat("5:00AM")}`)
console.log(`Converted time: ${convertTo24HrsFormat("12:33PM")}`)
console.log(`Converted time: ${convertTo24HrsFormat("01:59PM")}`)
console.log(`Converted time: ${convertTo24HrsFormat("11:8PM")}`)
console.log(`Converted time: ${convertTo24HrsFormat("10:02PM")}`)