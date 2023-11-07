package com.social.profileTest

import com.social.presentation.profile.UserProfileFragment
import org.junit.Assert.assertEquals
import org.junit.Test

class UserProfileFragmentTest {
    @Test
    fun testConvertNumberToK() {
        // Prueba con un número mayor o igual a 1000000
        assertEquals("1.2M", UserProfileFragment.convertNumberToK(1230000))

        // Prueba con un número mayor o igual a 1000
        assertEquals("1.2k", UserProfileFragment.convertNumberToK(1230))

        // Prueba con un número menor a 1000
        assertEquals("500", UserProfileFragment.convertNumberToK(500))
    }
}
