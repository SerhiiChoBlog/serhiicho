export type Pagination<T> = {
    current_page: number
    data: T
    first_page_url: string
    from: number
    next_page_url: string | null
    path: string
    per_page: number
    prev_page_url: string | null
    to: number
}

export type ServerError = {
    message: string
    errors: {
        [key: string]: string[]
    }
}

export type GetPostLikesResponse = {
    likes: number
    liked: boolean
}
