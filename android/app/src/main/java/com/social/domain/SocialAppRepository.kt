package com.social.domain

import com.social.data.source.remote.dto.LoginDto
import com.social.domain.model.LoginBody

interface SocialAppRepository {
    suspend fun validateUser(loginBody: LoginBody):LoginDto
}