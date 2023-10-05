package com.social.data.source.remote

import com.social.data.source.remote.dto.LoginDto
import com.social.domain.model.LoginBody
import retrofit2.http.Body
import retrofit2.http.POST

interface ApiInterface {

    /*****Endpoints*****/
    //Login
    @POST("/login")
    suspend fun validateUser(@Body loginBody : LoginBody): LoginDto

}