
package com.social.domain.usecase.sqlite

import com.social.data.source.local.aUser
import com.social.domain.SocialAppRepository
import com.social.domain.model.UserM
import com.social.utils.Resource
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.flow
import javax.inject.Inject

class SQLiteGetUser
    @Inject
    constructor(
        private val socialAppRepository: SocialAppRepository,
    ) {
        operator fun invoke(): Flow<Resource<UserM>> {
            return flow {
                try {
                    val user = socialAppRepository.getUser().aUser()
                    emit(Resource.Success(user))
                } catch (e: Exception) {
                    emit(Resource.Error(e.message.toString()))
                }
            }
        }
    }
