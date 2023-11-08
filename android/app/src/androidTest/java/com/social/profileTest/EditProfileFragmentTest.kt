package com.social.profileTest

import com.social.presentation.profile.EditProfileFragment
import org.junit.Assert.assertEquals
import org.junit.Test

class EditProfileFragmentTest {
    @Test
    fun userUniqueTest() {
        val fragment = EditProfileFragment()
        val uniqueUsername = "uniqueUser"
        val isUnique = fragment.userUnique(uniqueUsername)
        assertEquals(true, isUnique)
    }
}
