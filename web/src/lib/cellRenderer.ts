export const HTMLInColumnCell = (link: string, text: string): string => {
  if (isString(link) && link != '') {
    return `<a href=${link} target="_blank" rel="noopener"> ${text} </a>`
  }
  return text
}

const isString = (a: string): boolean => {
  if (typeof a === 'string') {
    return true
  }
  return false
}
