package com.social.domain.use_case

import android.util.Log
import com.social.data.source.remote.dto.aLogin
import com.social.domain.SocialAppRepository
import com.social.domain.model.InvalidUserException
import com.social.domain.model.Login
import com.social.domain.model.LoginBody
import com.social.utils.Resource
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.catch
import kotlinx.coroutines.flow.flow
import java.lang.Exception
import java.net.UnknownHostException
import javax.inject.Inject

class ValidateUser @Inject constructor(
    private val socialAppRepository: SocialAppRepository
) {
    operator fun invoke(loginBody: LoginBody): Flow<Resource<Login>> {
        return flow {
            var response : String =""
            try {
                if (loginBody.username.isBlank()) {
                    throw InvalidUserException("Ingrese usuario")
                }
                if (loginBody.password.isBlank()) {
                    throw InvalidUserException("Ingrese contraseña")
                }
                emit(Resource.Loading())
                val login = socialAppRepository.validateUser(loginBody).aLogin()
                Log.i("Gaaaa", login.status)
                Log.i("Gaaaa", login.response)
                response = login.status
                when (response) {
                    "OK" -> {
                        emit(Resource.Success(response))
                    }
                    else -> {
                        emit(Resource.Error(response))
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