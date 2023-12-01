
package com.social.domain.usecase.sqlite

import com.social.data.source.local.UserEntity
import com.social.domain.SocialAppRepository
import javax.inject.Inject

class SQLiteInsertUser
    @Inject
    constructor(
        private val socialAppRepository: SocialAppRepository,
    ) {
        suspend operator fun invoke(userEntity: UserEntity) {
            socialAppRepository.insertUser(
                UserEntity(
                    userEntity.id,
                    userEntity.sLastName,
                    userEntity.sName,
                    userEntity.sEmail,
                    userEntity.sUserName,
                    userEntity.sPhoto,
                    userEntity.sHeader,
                    userEntity.sBiography,
                ),
            )
        }
    }
