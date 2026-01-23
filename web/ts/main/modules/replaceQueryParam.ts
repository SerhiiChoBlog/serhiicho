export default (param: string, newVal: string): string => {
    const regex = new RegExp('([?;&])' + param + '[^&;]*[;&]?')
    const query = window.location.search.replace(regex, '$1').replace(/&$/, '')

    return (
        (query.length > 2 ? query + '&' : '?') + (newVal ? param + '=' + newVal : '')
    )
}
