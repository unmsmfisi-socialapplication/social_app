package com.social.utilsTest

import com.social.utils.Validation
import org.junit.Assert.assertFalse
import org.junit.Assert.assertTrue
import org.junit.Test

class ValidationTest {
    @Test
    fun testEmailValid() {
        val validEmail = "test@test.com"
        assertTrue(Validation.isEmailValid(validEmail))
    }

    @Test
    fun testEmailWithoutSymbol() {
        val invalidEmail = "invalidEmail"
        assertFalse(Validation.isEmailValid(invalidEmail))
    }

    @Test
    fun testEmailWithoutCom() {
        val invalidEmail = "invalid@Email"
        assertFalse(Validation.isEmailValid(invalidEmail))
    }

    @Test
    fun testEmailWithSpace() {
        val invalidEmail = "invalid email"
        assertFalse(Validation.isEmailValid(invalidEmail))
    }

    @Test
    fun testPasswordValid() {
        val validPassword = "validpassword123"
        assertTrue(Validation.isPasswordValid(validPassword))
    }

    @Test
    fun testPasswordShort() {
        val invalidPassword = "short"
        assertFalse(Validation.isPasswordValid(invalidPassword))
    }

    @Test
    fun testPasswordEmpty() {
        val invalidPassword = ""
        assertFalse(Validation.isPasswordValid(invalidPassword))
    }

    @Test
    fun testPassword8spaces() {
        val invalidPassword = "          "
        assertFalse(Validation.isPasswordValid(invalidPassword))
    }

    @Test
    fun testPasswordSpaceBetween() {
        val invalidPassword = "o o"
        assertFalse(Validation.isPasswordValid(invalidPassword))
    }

    @Test
    fun testUserValid() {
        val validUser = "Social@12"
        assertTrue(Validation.isUserValid(validUser))
    }

    @Test
    fun testUserShort() {
        val invalidUser = "short"
        assertFalse(Validation.isUserValid(invalidUser))
    }

    @Test
    fun testUserEmpty() {
        val invalidUser = ""
        assertFalse(Validation.isUserValid(invalidUser))
    }

    @Test
    fun testUser6spaces() {
        val invalidUser = "        "
        assertFalse(Validation.isUserValid(invalidUser))
    }

    @Test
    fun testUserSpaceBetween() {
        val invalidUser = "o o"
        assertFalse(Validation.isUserValid(invalidUser))
    }
}
