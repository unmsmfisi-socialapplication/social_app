package com.social.di

import android.app.Application
import androidx.room.Room
import com.social.data.repository.SocialAppRepositoryImp
import com.social.data.source.local.SocialDB
import com.social.data.source.remote.ApiInterface
import com.social.domain.SocialAppRepository
import com.social.domain.usecase.RegisterNewUser
import com.social.domain.usecase.ValidateUser
import com.social.domain.usecase.sqlite.SQLiteDeleteUser
import com.social.domain.usecase.sqlite.SQLiteGetUser
import com.social.domain.usecase.sqlite.SQLiteInsertUser
import com.social.domain.usecase.sqlite.SQLiteUpdateUser
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
    fun supplierSocialAppRepository(
        apiInterface: ApiInterface,
        db: SocialDB,
    ): SocialAppRepository {
        return SocialAppRepositoryImp(apiInterface, db.socialDAO)
    }

    @Provides
    @Singleton
    fun authenticationUseCase(repository: SocialAppRepository): AuthenticationUseCase {
        return AuthenticationUseCase(
            validateUser = ValidateUser(repository),
            registerNewUser = RegisterNewUser(repository),
            sqliteInsertUser = SQLiteInsertUser(repository),
            sqliteGetUser = SQLiteGetUser(repository),
            sqliteDeleteUser = SQLiteDeleteUser(repository),
            sqliteUpdateUser = SQLiteUpdateUser(repository),
        )
    }

    @Provides
    @Singleton
    fun supplierSocialLocalDatabase(app: Application): SocialDB {
        return Room.databaseBuilder(app, SocialDB::class.java, SocialDB.DATABASE_NAME).fallbackToDestructiveMigration().build()
    }
}
