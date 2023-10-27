package com.social.di

import com.social.data.source.remote.ApiInterface
import dagger.Module
import dagger.Provides
import dagger.hilt.InstallIn
import dagger.hilt.components.SingletonComponent
import okhttp3.OkHttpClient
import retrofit2.Retrofit
import retrofit2.converter.gson.GsonConverterFactory
import javax.inject.Named
import javax.inject.Singleton

@Module
@InstallIn(SingletonComponent::class)
object ApiSocialApp {
    @Provides
    @Singleton
    fun getMainApi(
        @Named("soccial_app")okHttpClient: OkHttpClient,
    ): ApiInterface {
        return Retrofit.Builder()
            .baseUrl("https://social-app-backend-service-7bg5siys2q-uc.a.run.app")
            .client(okHttpClient)
            .addConverterFactory(GsonConverterFactory.create())
            .build()
            .create(ApiInterface::class.java)
    }
}
