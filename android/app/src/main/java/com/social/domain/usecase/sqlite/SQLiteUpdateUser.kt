
package com.social.domain.usecase.sqlite

import com.social.domain.SocialAppRepository
import javax.inject.Inject

class SQLiteUpdateUser
    @Inject
    constructor(
        private val socialAppRepository: SocialAppRepository,
    ) {
        suspend operator fun invoke(
            id: Int,
            username: String,
            sPhoto: String,
        ) {
            try {
                socialAppRepository.updateUser(id, username, sPhoto)
            } catch (_: Exception) {
            }
        }
    }
