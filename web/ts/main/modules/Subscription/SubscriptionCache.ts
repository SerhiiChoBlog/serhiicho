import Cache from '@shared/modules/Cache'
import type { Tag } from '@shared/types/models'
import { storage } from '@shared/appConfig'

export default class {
    public static set(tags: Tag[]): void {
        Cache.set<Tag[]>(storage.selectedSubscriptionTags, tags)
    }

    public static get(): Tag[] | null {
        return Cache.get<Tag[]>(storage.selectedSubscriptionTags)
    }

    public static wantsAllTags(): boolean {
        return this.get() === null
    }

    public static isEmpty(): boolean {
        const items = this.get()
        return !items || items.length === 0
    }

    public static deleteAll(): void {
        Cache.delete(storage.selectedSubscriptionTags)
    }
}
