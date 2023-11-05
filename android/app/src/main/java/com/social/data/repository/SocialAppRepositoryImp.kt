package com.social.data.repository

import com.social.data.source.remote.ApiInterface
import com.social.data.source.remote.dto.CreatePostDto
import com.social.data.source.remote.dto.LoginDto
import com.social.domain.SocialAppRepository
import com.social.domain.model.CreatePostBody
import com.social.domain.model.LoginBody
import javax.inject.Inject

class SocialAppRepositoryImp
    @Inject
    constructor(
        private val api: ApiInterface,
    ) : SocialAppRepository {
        override suspend fun validateUser(loginBody: LoginBody): LoginDto {
            return api.validateUser(loginBody)
        }

        override suspend fun createPost(createPostBody: CreatePostBody): CreatePostDto {
            return api.createPost(createPostBody)
        }
    }
