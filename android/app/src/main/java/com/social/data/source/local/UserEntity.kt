
package com.social.data.source.local

import androidx.room.Entity
import androidx.room.PrimaryKey
import com.social.domain.model.UserM

@Entity(tableName = "user")
data class UserEntity(
    @PrimaryKey(autoGenerate = true)
    val id: Int,
    val sLastName: String,
    val sName: String,
    val sEmail: String,
    val sUserName: String,
    val sPhoto: String,
    val sHeader: String,
    val sBiography: String,
)

fun UserEntity.aUser(): UserM {
    return UserM(
        id,
        sLastName,
        sName,
        sEmail,
        sUserName,
        sPhoto,
        sHeader,
        sBiography,
    )
}
