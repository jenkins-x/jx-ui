export const HTMLInColumnCell = (link: string, text: string): string => {
  if (isString(link) && link != '') {
    return `<a href=${link} target="_blank" rel="noopener"> ${text} </a>`
  }
  return text
}

export const BTNInColumnCell = (status: string, owner: string, branch: string, repository : string, build: string, name: string, clickHandler): HTMLElement => {
  if (status.toLowerCase() === "pending" || status.toLowerCase() === "running") {
    const p = document.createElement(`button`);
    p.setAttribute("class", "m-1 rounded-md focus:outline-none focus:shadow-outline-purple-500");
    p.setAttribute("aria-label", "Toggle color mode");
    p.innerHTML = `
        <svg class="mt-2 w-5 h-5" aria-hidden="true" fill="currentColor" viewBox="0 0 512 512">
          <path d="M0 256C0 114.6 114.6 0 256 0C397.4 0 512 114.6 512 256C512 397.4 397.4 512 256 512C114.6 512 0 397.4 0 256zM175 208.1L222.1 255.1L175 303C165.7 312.4 165.7 327.6 175 336.1C184.4 346.3 199.6 346.3 208.1 336.1L255.1 289.9L303 336.1C312.4 346.3 327.6 346.3 336.1 336.1C346.3 327.6 346.3 312.4 336.1 303L289.9 255.1L336.1 208.1C346.3 199.6 346.3 184.4 336.1 175C327.6 165.7 312.4 165.7 303 175L255.1 222.1L208.1 175C199.6 165.7 184.4 165.7 175 175C165.7 184.4 165.7 199.6 175 208.1V208.1z"/>
        </svg>
    `;

    const pipeline = {
      owner,
      branch,
      repository,
      build
    }

    p.addEventListener('click', () => clickHandler(name, pipeline));
    return p;
  }
  return null
}

const isString = (a: string): boolean => {
  if (typeof a === 'string') {
    return true
  }
  return false
}
