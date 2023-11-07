package com.social.di

import com.social.data.repository.SocialAppRepositoryImp
import com.social.data.source.remote.ApiInterface
import com.social.domain.SocialAppRepository
import com.social.domain.usecase.ValidateUser
import com.social.presentation.authentication.AuthenticationUseCase
import dagger.Module
import dagger.Provides
import dagger.hilt.InstallIn
import dagger.hilt.components.SingletonComponent
import okhttp3.OkHttpClient
import retrofit2.Retrofit
import retrofit2.converter.gson.GsonConverterFactory
import java.util.concurrent.TimeUnit
import javax.inject.Named
import javax.inject.Singleton

@Module
@InstallIn(SingletonComponent::class)
object ApiSocialApp {
    @Provides
    @Singleton
    fun getMainApi(
        @Named("soccial_app") okHttpClient: OkHttpClient,
    ): ApiInterface {
        return Retrofit.Builder()
            .baseUrl("https://social-app-backend-service-7bg5siys2q-uc.a.run.app")
            .client(okHttpClient)
            .addConverterFactory(GsonConverterFactory.create())
            .build()
            .create(ApiInterface::class.java)
    }

    @Provides
    @Named("soccial_app")
    fun provideSocialAppOkHttpClient(): OkHttpClient {
        return OkHttpClient.Builder()
            .readTimeout(30, TimeUnit.SECONDS)
            .writeTimeout(30, TimeUnit.SECONDS)
            .build()
    }

    @Provides
    @Singleton
    fun supplierSocialAppRepository(apiInterface: ApiInterface): SocialAppRepository {
        return SocialAppRepositoryImp(apiInterface)
    }

    @Provides
    @Singleton
    fun authenticationUseCase(repository: SocialAppRepository): AuthenticationUseCase {
        return AuthenticationUseCase(
            validateUser = ValidateUser(repository),
        )
    }
}
