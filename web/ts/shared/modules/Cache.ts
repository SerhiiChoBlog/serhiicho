export default class {
    public static set<T>(key: string, item: T): void {
        localStorage.setItem(key, JSON.stringify(item))
    }

    public static get<T>(key: string): T | null {
        const result = localStorage.getItem(key)
        return result ? JSON.parse(result) : null
    }

    public static delete(key: string): void {
        localStorage.removeItem(key)
    }
}
