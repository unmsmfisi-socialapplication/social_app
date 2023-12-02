package com.social.profileTest

import com.social.presentation.profile.UserProfileFragment
import org.junit.Assert.assertEquals
import org.junit.Test

class UserProfileFragmentTest {
    @Test
    fun testConvertNumber() {
        val userProfileFragment = UserProfileFragment()
        val result = userProfileFragment.convertNumberToK(500)
        assertEquals(
            "500",
            result,
        )
    }

    @Test
    fun testConvertNumberK() {
        val userProfileFragment = UserProfileFragment()
        val result = userProfileFragment.convertNumberToK(2500)
        assertEquals(
            "2.5k",
            result,
        )
    }

    @Test
    fun testConvertNumberM() {
        val userProfileFragment = UserProfileFragment()
        val result = userProfileFragment.convertNumberToK(1230000)
        assertEquals(
            "1.2M",
            result,
        )
    }
}
