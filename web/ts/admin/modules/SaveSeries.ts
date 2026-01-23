import type { Series } from '@shared/types/models'
import type { EditSeriesFormElements } from '@shared/types'
import axios from 'axios'
import handleError from '@admin/modules/handleError'

export default class {
    public constructor(
        private elements: EditSeriesFormElements,
        private postId: number,
    ) { }

    public async save(description: string): Promise<Series | null> {
        const formData = new FormData()
        const postsElems = this.elements['posts[]'] as HTMLInputElement[] | HTMLInputElement
        let posts: HTMLInputElement[] = []

        if (postsElems) {
            posts = postsElems instanceof HTMLInputElement
                ? [postsElems]
                : Array.from(postsElems)
        }

        const imageFile = this.elements.image.files![0] as File

        formData.append('title', this.elements.title.value)
        formData.append('intro', this.elements.intro.value)
        formData.append('color_from', this.elements.color_from.value)
        formData.append('color_to', this.elements.color_to.value)
        formData.append('description', description)
        formData.append('image', imageFile)
        formData.append('_method', 'put')

        if (posts.length > 0) {
            formData.append('posts', JSON.stringify(posts.map(tag => tag.value)))
        }

        const params = {
            headers: { 'content-type': 'multipart/form-data' },
        }

        try {
            const response = await axios.post<Series>(`/admin/ajax/series/${this.postId}`, formData, params)
            return response.data
        } catch (err) {
            handleError(err)
            return null
        }
    }
}
