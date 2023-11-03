package com.social.data.repository

import com.social.domain.model.Post

interface PostRepository {
    suspend fun obtainPost(): MutableList<Post>
}
