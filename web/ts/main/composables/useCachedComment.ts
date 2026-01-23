import Cache from '@shared/modules/Cache'
import { storage } from '@shared/appConfig'

export default (postId: number, commentId?: number) => {
    const cacheKey = `${storage.cachedComment}-${postId}`

    function getCommentFromCache(): string {
        if (commentId) {
            return ''
        }

        return Cache.get<string>(cacheKey) || ''
    }

    function saveCommentToCache(comment: string): void {
        if (commentId) {
            return
        }

        if (comment === '') {
            Cache.delete(cacheKey)
            return
        }

        Cache.set<string>(cacheKey, comment)
    }

    function deleteCommentFromCache(): void {
        Cache.delete(cacheKey)
    }

    return {
        getCommentFromCache,
        saveCommentToCache,
        deleteCommentFromCache,
    }
}
