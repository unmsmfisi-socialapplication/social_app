
package com.social.data.source.local

import androidx.room.Dao
import androidx.room.Insert
import androidx.room.OnConflictStrategy
import androidx.room.Query

@Dao
interface SocialDAO {
    @Insert(onConflict = OnConflictStrategy.REPLACE)
    suspend fun insertUser(userEntity: UserEntity)

    @Query("SELECT * FROM user")
    suspend fun getUser(): UserEntity

    @Query("UPDATE user SET sUserName = :username, sPhoto = :photo WHERE id = :id")
    suspend fun updateUser(
        id: Int,
        username: String,
        photo: String,
    )

    @Query("DELETE FROM user")
    suspend fun deleteUser()
}
