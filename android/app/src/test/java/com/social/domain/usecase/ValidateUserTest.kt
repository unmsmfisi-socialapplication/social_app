package com.social.domain.usecase

import com.social.data.source.remote.dto.LoginDto
import com.social.domain.SocialAppRepository
import com.social.domain.model.LoginBody
import com.social.domain.model.LoginUserData
import com.social.utils.Resource
import io.mockk.coEvery
import io.mockk.coVerify
import io.mockk.mockk
import kotlinx.coroutines.runBlocking
import org.junit.Assert.assertEquals
import org.junit.Assert.assertTrue
import org.junit.Assert.fail
import org.junit.Before
import org.junit.Test
import java.util.Objects

class ValidateUserTest {
    private lateinit var validateUser: ValidateUser
    private val repository: SocialAppRepository = mockk()

    @Before
    fun setup() {
        validateUser = ValidateUser(repository)
    }

    @Test
    fun serverResponseSuccess() =
        runBlocking {
            val loginBody =
                LoginBody(
                    username = "myuser12",
                    password = "Social@12",
                )

            val loginResponse =
                LoginDto<Objects>(
                    LoginUserData(
                        "yJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDAyOTA5NTcs" +
                            "ImlhdCI6MTcwMDIwNDU1NywianRpIjoiOGMwMjYzOTQtMDQ5NS00N2YyLWI5Z" +
                            "jctNGM1MzExYzI3NDA0Iiwicm5sIjoidXNlciIsInN1YiI6Im15dXNlcjEyIn0.Mk" +
                            "444bwx_IKrbA8-QiPyefen1Abb4sb3buOsF4H1VwA",
                    ),
                    "OK",
                )
            coEvery { repository.validateUser(loginBody) } returns loginResponse

            val result = validateUser(loginBody)
            var successEmitted = false
            result.collect { resource ->
                when (resource) {
                    is Resource.Success -> {
                        successEmitted = true
                        assertEquals(
                            listOf(resource.data?.get(0)?.let { LoginDto<Objects>(it, "OK") }),
                            resource.data?.map {
                                loginResponse
                            } ?: "",
                        )
                    }
                    is Resource.Loading -> {
                    }
                    else -> {
                        fail("Expected a success resource, but received $resource")
                    }
                }
            }
            coVerify { repository.validateUser(loginBody) }
            assertTrue(successEmitted)
        }

    @Test
    fun serverResponseUserNotFound() =
        runBlocking {
            val loginBody =
                LoginBody(
                    username = "myuser54321",
                    password = "Social@12",
                )

            val errorResponse = LoginDto<Objects>(LoginUserData(""), "NOTFOUND")
            coEvery { repository.validateUser(loginBody) } returns errorResponse

            val result = validateUser(loginBody)

            var errorEmitted = false
            result.collect { resource ->
                when (resource) {
                    is Resource.Error -> {
                        errorEmitted = true
                        assertEquals("NOTFOUND", resource.message)
                        assertEquals(null, resource.data)
                    }
                    is Resource.Loading -> {
                    }
                    else -> {
                        fail("Expected an error resource, but received $resource")
                    }
                }
            }

            coVerify { repository.validateUser(loginBody) }

            assertTrue(errorEmitted)
        }

    @Test
    fun serverResponseUnauthorized() =
        runBlocking {
            val loginBody =
                LoginBody(
                    username = "myuser12",
                    password = "social21",
                )

            val errorResponse = LoginDto<Objects>(LoginUserData(""), "BADCREDENTIALS")
            coEvery { repository.validateUser(loginBody) } returns errorResponse

            val result = validateUser(loginBody)

            var errorEmitted = false
            result.collect { resource ->
                when (resource) {
                    is Resource.Error -> {
                        errorEmitted = true
                        assertEquals("BADCREDENTIALS", resource.message)
                        assertEquals(null, resource.data)
                    }
                    is Resource.Loading -> {
                    }
                    else -> {
                        fail("Expected an error resource, but received $resource")
                    }
                }
            }

            coVerify { repository.validateUser(loginBody) }

            assertTrue(errorEmitted)
        }
}
