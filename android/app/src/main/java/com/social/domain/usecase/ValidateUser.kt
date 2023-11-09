package com.social.domain.usecase

import com.social.data.source.remote.dto.aLogin
import com.social.domain.SocialAppRepository
import com.social.domain.model.InvalidUserException
import com.social.domain.model.LoginBody
import com.social.domain.model.LoginUserData
import com.social.utils.Resource
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.flow
import java.lang.Exception
import java.net.UnknownHostException
import javax.inject.Inject

class ValidateUser
    @Inject
    constructor(
        private val socialAppRepository: SocialAppRepository,
    ) {
        operator fun invoke(loginBody: LoginBody): Flow<Resource<List<LoginUserData>>> {
            return flow {
                try {
                    if (loginBody.username.isBlank()) {
                        throw InvalidUserException("Ingrese usuario")
                    }
                    if (loginBody.password.isBlank()) {
                        throw InvalidUserException("Ingrese contraseña")
                    }
                    emit(Resource.Loading())
                    val login = socialAppRepository.validateUser(loginBody).aLogin()
                    when (login.status) {
                        "OK" -> {
                            emit(Resource.Success(login.response))
                        }

                        else -> {
                            emit(Resource.Error(login.response))
                        }
                    }
                } catch (u: UnknownHostException) {
                    emit(Resource.Error("No se pudo conectar al servidor"))
                } catch (e: Exception) {
                    emit(Resource.Error(e.message.toString()))
                }
            }
        }
    }
