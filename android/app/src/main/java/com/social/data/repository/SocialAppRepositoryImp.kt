
package com.social.data.repository

import com.social.data.source.local.SocialDAO
import com.social.data.source.local.UserEntity
import com.social.data.source.remote.ApiInterface
import com.social.data.source.remote.dto.CreatePostDto
import com.social.data.source.remote.dto.LoginDto
import com.social.data.source.remote.dto.RegisterUserErrorDto
import com.social.domain.SocialAppRepository
import com.social.domain.model.CreatePostBody
import com.social.domain.model.LoginBody
import com.social.domain.model.RegisterBody
import com.social.domain.model.RegisterData
import java.util.Objects
import javax.inject.Inject

class SocialAppRepositoryImp
    @Inject
    constructor(
        private val api: ApiInterface,
        private val socialDAO: SocialDAO,
    ) : SocialAppRepository {
        override suspend fun validateUser(loginBody: LoginBody): LoginDto<Objects> {
            return api.validateUser(loginBody)
        }

        override suspend fun createPost(createPostBody: CreatePostBody): CreatePostDto {
            return api.createPost(createPostBody)
        }

        override suspend fun registerNewUser(registerBody: RegisterBody): RegisterData {
            return api.registerUser(registerBody)
        }

        override suspend fun registerNewUserError(registerBody: RegisterBody): RegisterUserErrorDto {
            return api.registerUserError(registerBody)
        }

        override suspend fun insertUser(userEntity: UserEntity) {
            socialDAO.insertUser(userEntity)
        }

        override suspend fun getUser(): UserEntity {
            return socialDAO.getUser()
        }

        override suspend fun updateUser(
            id: Int,
            username: String,
            photo: String,
        ) {
            socialDAO.updateUser(
                id,
                username,
                photo,
            )
        }

        override suspend fun deleteUser() {
            return socialDAO.deleteUser()
        }
    }
