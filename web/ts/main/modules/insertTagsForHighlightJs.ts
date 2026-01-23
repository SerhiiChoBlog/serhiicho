export default () => {
    const codeElems = document.querySelectorAll('code')

    codeElems.forEach(block => {
        const preElem = block.parentElement!
        const classes: string = preElem.getAttribute('class')!

        if (/^language-/.test(classes)) {
            const lang = classes.replace('language-', '').trim()
            const html = `<div class="absolute text-[#b9bec2] block right-[10px] top-[6px] text-xs" aria-hidden="true">${lang}</div>`

            preElem.insertAdjacentHTML('beforeend', html)
        }
    })
}
