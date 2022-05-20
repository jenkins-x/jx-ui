export const HTMLInColumnCell = (link: string, text: string) => {
  if (isString(link) && link != '') {
    return `<a href=${link} target="_blank" rel="noopener"> ${text} </a>`
  }
  return text
}

const isString = (a: string) => {
  if (typeof a === 'string') {
    return true
  }
  return false
}
