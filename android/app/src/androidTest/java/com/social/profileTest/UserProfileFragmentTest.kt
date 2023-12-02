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
            "2k",
            result,
        )
    }

    @Test
    fun testConvertNumberM() {
        val userProfileFragment = UserProfileFragment()
        val result = userProfileFragment.convertNumberToK(1200000)
        assertEquals(
            "1M",
            result,
        )
    }
}
