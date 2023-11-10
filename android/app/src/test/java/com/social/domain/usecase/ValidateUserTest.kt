package com.social.domain.usecase

import com.social.data.source.remote.dto.LoginDto
import com.social.domain.SocialAppRepository
import com.social.domain.model.LoginBody
import com.social.utils.Resource
import io.mockk.coEvery
import io.mockk.mockk
import kotlinx.coroutines.flow.toList
import kotlinx.coroutines.runBlocking
import org.junit.Before
import org.junit.Test
import java.net.UnknownHostException

class ValidateUserTest {

    private lateinit var socialAppRepository: SocialAppRepository
    private lateinit var validateUser: ValidateUser

    @Before
    fun setup() {
        socialAppRepository = mockk()
        validateUser = ValidateUser(socialAppRepository)
    }

    @Test
    fun emptyUserName() = runBlocking {
        // Given
        val loginBody = LoginBody(username = "", password = "Social@43")

        // When
        val result = validateUser(loginBody).toList()

        // Then
        assert(result[0] is Resource.Error && (result[0] as Resource.Error).message == "Ingrese usuario")
    }

    @Test
    fun `emptyPassword`() = runBlocking {
        // Given
        val loginBody = LoginBody(username = "test43test", password = "")

        // When
        val result = validateUser(loginBody).toList()

        // Then
        assert(result[0] is Resource.Error && (result[0] as Resource.Error).message == "Ingrese contraseña")
    }

    @Test
    fun `validReturnsOk`() = runBlocking {
        // Given
        val loginBody = LoginBody(username = "test43test", password = "Social@43")
        val mockedResponse = LoginDto("OK", "Success")

        coEvery { socialAppRepository.validateUser(loginBody) } returns mockedResponse

        // When
        val result = validateUser(loginBody).toList()

        // Then
        assert(result[0] is Resource.Loading)
        assert(result[1] is Resource.Success && (result[1] as Resource.Success).data == listOf(LoginDto(status = "OK",
            response = "Authentication successful")))

    }

    @Test
    fun `unknownHostExceptionReturnsException`() = runBlocking {
        // Given
        val loginBody = LoginBody(username = "user123", password = "123456")

        coEvery { socialAppRepository.validateUser(loginBody) } throws UnknownHostException()

        // When
        val result = validateUser(loginBody).toList()

        // Then
        assert(result[0] is Resource.Loading)
        assert(result[1] is Resource.Error && (result[1] as Resource.Error).message == "No se pudo conectar al servidor")
    }

    @Test
    fun `exceptionReturnsException`() = runBlocking {
        // Given
        val loginBody = LoginBody(username = "12345678", password = "inio7nv@cad")

        coEvery { socialAppRepository.validateUser(loginBody) } throws Exception("Some error message")

        // When
        val result = validateUser(loginBody).toList()

        // Then
        assert(result[0] is Resource.Loading)
        assert(result[1] is Resource.Error && (result[1] as Resource.Error).message == "Some error message")
    }
}
