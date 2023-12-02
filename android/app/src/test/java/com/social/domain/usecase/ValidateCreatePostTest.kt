import com.social.domain.SocialAppRepository
import com.social.domain.model.CreatePostBody
import com.social.domain.model.CreatePostResponse
import com.social.domain.usecase.ValidateCreatePost
import com.social.utils.Resource
import io.mockk.coEvery
import io.mockk.mockk
import kotlinx.coroutines.runBlocking
import org.junit.Assert.assertEquals
import org.junit.Test

class ValidateCreatePostTest {
    private val mockRepository: SocialAppRepository = mockk()
    private val validateCreatePost = ValidateCreatePost(mockRepository)

    @Test
    fun `invoke should return success resource when post creation is successful`() =
        runBlocking {
            val createPostBody =
                CreatePostBody(
                    userId = 2,
                    title = "hola",
                    description = "saludos",
                    hasMultimedia = false,
                    public = false,
                    multimedia = null,
                )
            val mockedResponse =
                CreatePostResponse(
                    Id = 1,
                    UserId = 2,
                    Title = "ddd",
                    Description = "ddd",
                    HasMultimedia = false,
                    Public = false,
                    Multimedia = null,
                    InsertionDate = "daa",
                    UpdateDate = "rrrr",
                )
            val expectedResource = Resource.Success(listOf(mockedResponse))
            coEvery { mockRepository.createPost(createPostBody) }
            val resultResource = mutableListOf<Resource<List<CreatePostResponse>>>()
            validateCreatePost(createPostBody).collect { resultResource.add(it) }
            assertEquals(expectedResource, resultResource.first())
        }

    @Test
    fun `invoke should return error resource when post creation fails`() =
        runBlocking {
            val createPostBody =
                CreatePostBody(
                    userId = null,
                    title = "aaa",
                    description = "ddd",
                    hasMultimedia = false,
                    public = false,
                    multimedia = null,
                )
            val expectedErrorMessage = "Error al crear el post"
            coEvery { mockRepository.createPost(createPostBody) } throws Exception("Error al crear el post")
            val resultResource = mutableListOf<Resource<List<CreatePostResponse>>>()
            validateCreatePost(createPostBody).collect { resultResource.add(it) }
            assertEquals(expectedErrorMessage, (resultResource.first() as Resource.Error).message)
        }
}
