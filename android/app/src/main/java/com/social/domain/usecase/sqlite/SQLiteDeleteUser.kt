
package com.social.domain.usecase.sqlite

import com.social.domain.SocialAppRepository
import javax.inject.Inject

class SQLiteDeleteUser
    @Inject
    constructor(
        private val socialAppRepository: SocialAppRepository,
    ) {
        suspend operator fun invoke() {
            socialAppRepository.deleteUser()
        }
    }
