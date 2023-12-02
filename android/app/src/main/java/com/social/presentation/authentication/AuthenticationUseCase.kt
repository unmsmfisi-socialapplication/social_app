package com.social.presentation.authentication

import com.social.domain.usecase.RegisterNewUser
import com.social.domain.usecase.ValidateUser
import com.social.domain.usecase.sqlite.SQLiteDeleteUser
import com.social.domain.usecase.sqlite.SQLiteGetUser
import com.social.domain.usecase.sqlite.SQLiteInsertUser
import com.social.domain.usecase.sqlite.SQLiteUpdateUser

data class AuthenticationUseCase(
    val validateUser: ValidateUser,
    val registerNewUser: RegisterNewUser,
    val sqliteInsertUser: SQLiteInsertUser,
    val sqliteGetUser: SQLiteGetUser,
    val sqliteDeleteUser: SQLiteDeleteUser,
    val sqliteUpdateUser: SQLiteUpdateUser,
)
