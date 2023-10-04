package com.social.domain.use_case

import com.social.data.source.remote.dto.LoginDto
import com.social.domain.SocialAppRepository
import com.social.domain.model.LoginBody
import com.social.domain.model.LoginUserData
import com.social.utils.Resource
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.flow
import java.lang.Exception
import javax.inject.Inject

class ValidateUser @Inject constructor(
    private val socialAppRepository: SocialAppRepository
) {
    operator fun invoke(loginBody: LoginBody): Flow<Resource<List<LoginUserData>>> {
        return flow {
            try {

            }catch (e: Exception){

            }
        }
    }
}