package com.example.mobilesocialapp.di

import com.example.mobilesocialapp.data.source.remote.ApiInterface
import com.example.mobilesocialapp.util.Routes.API_SOCIAL_APP_PRD
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
    //function to get the api of service
    @Provides
    @Singleton
    fun getApi(@Named("social_app")okHttpClient: OkHttpClient): ApiInterface{
        return Retrofit.Builder()
            .baseUrl(API_SOCIAL_APP_PRD)
            .client(okHttpClient)
            .addConverterFactory(GsonConverterFactory.create())
            .build()
            .create(ApiInterface::class.java)
    }
}