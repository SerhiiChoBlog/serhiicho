import type { Post } from '@shared/types/models'
import type { EditPostFormElements } from '@shared/types'
import axios from 'axios'
import SweetAlert from 'sweetalert2'
import handleError from '@admin/modules/handleError'

export default class {
    public constructor(
        private elements: EditPostFormElements,
        private postId: number,
        private mainTagId: number,
    ) { }

    public async save(content: string): Promise<Post | null> {
        const formData = new FormData()
        const tagsElems = this.elements['tags[]'] as HTMLInputElement[] | HTMLInputElement

        if (!tagsElems) {
            SweetAlert.fire('Info', 'Choose tags before saving the post', 'info')
            return null
        }

        const tags = tagsElems instanceof HTMLInputElement
            ? [tagsElems]
            : Array.from(tagsElems)

        const imageFile = this.elements.image.files![0] as File

        formData.append('title', this.elements.title.value)
        formData.append('intro', this.elements.intro.value)
        formData.append('content', content)
        formData.append('tag_id', this.mainTagId.toString())
        formData.append('image', imageFile)
        formData.append('tags', JSON.stringify(tags.map(tag => tag.value)))
        formData.append('is-published', this.elements['is-published'].checked ? 'on' : 'off')
        formData.append('_method', 'put')

        const params = {
            headers: { 'content-type': 'multipart/form-data' },
        }

        try {
            const response = await axios.post<Post>(`/admin/ajax/posts/${this.postId}`, formData, params)
            return response.data
        } catch (err) {
            handleError(err)
            return null
        }
    }
}
