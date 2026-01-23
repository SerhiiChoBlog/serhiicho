import axios from 'axios'
import handleError from '@admin/modules/handleError'

export default () => {
    async function blockVisitor(visitorId: number, block: boolean): Promise<void> {
        const uriSuffix = block ? 'block' : 'unblock'

        try {
            await axios.post(`/admin/ajax/visitors/${visitorId}/${uriSuffix}`)
        } catch (err) {
            handleError(err)
        }
    }

    return {
        blockVisitor
    }
}
