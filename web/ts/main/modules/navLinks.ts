const activeClass = 'navbar-link--active bg-black/30 lg:bg-transparent'
const currPath = window.location.pathname

export default [
    {
        href: '/',
        class: currPath === '/' ? activeClass : '',
        title: 'Home',
    },
    {
        href: '/posts',
        class: currPath === '/posts' ? activeClass : '',
        title: 'Blog',
    },
    {
        href: '/series',
        class: currPath === '/series' ? activeClass : '',
        title: 'Series',
    },
    {
        href: '/projects',
        class: currPath === '/projects' ? activeClass : '',
        title: 'Open source',
    },
    {
        href: '/about-me',
        class: currPath === '/about-me' ? activeClass : '',
        title: 'About',
    },
]
