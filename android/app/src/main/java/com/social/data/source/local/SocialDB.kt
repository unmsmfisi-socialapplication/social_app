
package com.social.data.source.local

import androidx.room.Database
import androidx.room.RoomDatabase

@Database(
    entities = [
        UserEntity::class,
    ],
    version = 1,
)
abstract class SocialDB : RoomDatabase() {
    abstract val socialDAO: SocialDAO

    companion object {
        const val DATABASE_NAME = "social_db"
    }
}
