package com.social.homeTest

import com.social.data.repository.PostRepository
import com.social.domain.model.Post
import kotlinx.coroutines.runBlocking
import org.junit.Assert.assertEquals
import org.junit.Test

class PostRepositoryImpTest {
    class MockPostRepository : PostRepository {
        override suspend fun obtainPost(): MutableList<Post> {
            return mutableListOf(
                Post(
                    "Usuario de prueba",
                    "Hora de prueba",
                    "Contenido de prueba",
                    "URL de imagen de prueba",
                ),
            )
        }
    }

    @Test
    fun obtainPost() {
        // Creamos una instancia del repositorio utilizando el stub
        val repository: PostRepository = MockPostRepository()

        // Ejecutamos el método que queremos probar
        val result = runBlocking { repository.obtainPost() }

        // Verificamos que el resultado obtenido coincida con lo esperado
        assertEquals(1, result.size) // Verifica que se haya obtenido un único Post en la lista
        assertEquals("Usuario de prueba", result[0].names)
        assertEquals("Hora de prueba", result[0].hour)
        assertEquals("Contenido de prueba", result[0].content)
        assertEquals("URL de imagen de prueba", result[0].image)
    }
}
