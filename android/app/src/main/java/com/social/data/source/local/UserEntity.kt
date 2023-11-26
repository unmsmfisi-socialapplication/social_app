package com.social.data.source.local

import androidx.room.Entity
import androidx.room.PrimaryKey
import com.social.domain.model.UserM

@Entity(tableName = "user")
data class UserEntity(
    @PrimaryKey(autoGenerate = true)
    val id:Int,
    val sApePat:String,
    val sNombre:String,
    val sCorreo:String,
    val sUsuario:String,
    val sFoto:String,
    val sHeader:String,
    val sBiografia:String
)

fun UserEntity.aUser(): UserM {
    return UserM(
        id,
        sApePat,
        sNombre,
        sCorreo,
        sUsuario,
        sFoto,
        sHeader,
        sBiografia
    )
}
