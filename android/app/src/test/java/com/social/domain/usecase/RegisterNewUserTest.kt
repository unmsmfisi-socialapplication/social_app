import com.social.domain.SocialAppRepository
import com.social.domain.model.RegisterBody
import com.social.domain.model.RegisterData
import com.social.domain.usecase.RegisterNewUser
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

class RegisterNewUserTest {
    private lateinit var registerNewUser: RegisterNewUser
    private val socialAppRepository = mockk<SocialAppRepository>()

    @Before
    fun setup() {
        registerNewUser = RegisterNewUser(socialAppRepository)
    }

    @Test
    fun `registerNewUser with valid data should emit success`() =
        runBlocking {
            val registerBody =
                RegisterBody(
                    Phone = "",
                    Email = "josecalletest@gmail.com",
                    User_name = "josecalletest",
                    Password = "Social@43",
                )
            val registerData = RegisterData("josecalletest@gmail.com", "josecalletest")

            coEvery { socialAppRepository.registerNewUser(registerBody) } returns registerData

            val result = registerNewUser(registerBody)

            var successEmitted = false
            result.collect { resource ->
                when (resource) {
                    is Resource.Success -> {
                        successEmitted = true
                        assertEquals(listOf(registerData), resource.data)
                    }
                    is Resource.Loading -> {
                    }
                    else -> {
                        fail("Expected a success resource, but received $resource")
                    }
                }
            }

            coVerify { socialAppRepository.registerNewUser(registerBody) }

            assertTrue(successEmitted)
        }

    @Test
    fun `registerNewUser with invalid data should emit error`() =
        runBlocking {
            val registerBody = RegisterBody("", "email321@gmail.com", "", "Social@12")

            val result = registerNewUser(registerBody)

            result.collect { resource ->
                when (resource) {
                    is Resource.Error -> {
                        assert(resource.message == "Ingrese usuario")
                    }
                    else -> assert(false)
                }
            }
        }

    @Test
    fun `registerNewUser with invalid email should emit error`() =
        runBlocking {
            val registerBody = RegisterBody("", "", "josecalletest", "Social@12")

            val result = registerNewUser(registerBody)

            result.collect { resource ->
                when (resource) {
                    is Resource.Error -> {
                        assert(resource.message == "Ingrese correo")
                    }
                    else -> assert(false)
                }
            }
        }

    @Test
    fun `registerNewUser with invalid psw should emit error`() =
        runBlocking {
            val registerBody = RegisterBody("", "email321@gmail.com", "josecalletest", "")
            val result = registerNewUser(registerBody)
            val registerBody2 = RegisterBody("", "email321@gmail.com", "josecalletest", "dd")
            val result2 = registerNewUser(registerBody2)

            result.collect { resource ->
                when (resource) {
                    is Resource.Error -> {
                        assert(resource.message == "Ingrese contraseña")
                    }
                    else -> assert(false)
                }
            }

            result2.collect { resource2 ->
                when (resource2) {
                    is Resource.Error -> {
                        assert(resource2.message == "La contraseña debe ser segura")
                    }
                    else -> assert(false)
                }
            }
        }
}
