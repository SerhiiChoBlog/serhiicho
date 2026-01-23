export default class {
    private activeClasses: string[] = [
        'border-b-main',
        '!text-main-dark',
        'hover:border-b-main-hover',
    ]

    public constructor(private tabs: NodeListOf<HTMLAnchorElement>) {
    }

    public handle(): void {
        this.handleEvent()
        window.addEventListener('hashchange', e => this.handleEvent())
    }

    private handleEvent(): void {
        const hash = window.location.hash.replace('#', '')

        this.hideAllTargets()

        if (hash === '') {
            this.setFirstElemActive()
            return
        }

        const target = document.getElementById(hash)

        if (!target)
            return

        this.hideAllTargets()
        this.changeActiveTabStyles(hash)

        target.classList.remove('hidden')
    }

    private changeActiveTabStyles(hash: string): void {
        this.tabs.forEach(tab => {
            if (tab.hash === `#${hash}`) {
                tab.classList.add(...this.activeClasses)
            } else {
                tab.classList.remove(...this.activeClasses)
            }
        })
    }

    private hideAllTargets(): void {
        this.tabs.forEach(tab => {
            const hash = tab.hash.replace('#', '')

            if (hash === '')
                return

            const target = document.getElementById(hash)

            if (!target)
                return

            target.classList.add('hidden')
        })
    }

    private setFirstElemActive(): void {
        const firstElems = document.querySelectorAll<HTMLAnchorElement>('._tab:first-child')

        if (firstElems.length === 0)
            return

        firstElems.forEach(firstTab => {
            const hash = firstTab.hash.replace('#', '')
            const target = document.getElementById(hash)

            if (!target)
                return

            firstTab.classList.add(...this.activeClasses)
            target.classList.remove('hidden')
        })
    }
}