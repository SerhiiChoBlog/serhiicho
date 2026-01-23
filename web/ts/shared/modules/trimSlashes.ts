export default (path: string): string => {
    let result: string = path

    while (result.startsWith('/')) {
        result = result.slice(1)
    }

    while (result.endsWith('/')) {
        result = result.slice(0, -1)
    }

    return result
}
