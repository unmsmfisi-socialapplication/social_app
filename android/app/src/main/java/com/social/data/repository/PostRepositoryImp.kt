package com.social.data.repository

import com.social.domain.model.Post

class PostRepositoryImp: PostRepository {
    override suspend fun obtainPost(): MutableList<Post> {
        val listPost = mutableListOf<Post>()
        //Probando con valores
        val response = listOf(
            Post("Mateo Pumacahua", "12:30 pm", "Contenido del post"),
            Post("Mateo Pachacutec", "12:30 pm", "Contenido del post"),
            Post("Mateo Pumacahua", "12:30 pm", "Contenido del post"),
            Post("Mateo Pachacutec", "12:30 pm", "Contenido del post"),
            Post("Mateo Pumacahua", "12:30 pm", "Contenido del post"),
            Post("Mateo Pachacutec", "12:30 pm", "Contenido del post"),
            Post("Mateo Pumacahua", "12:30 pm", "Contenido del post"),
            Post("Mateo Pachacutec", "12:30 pm", "Contenido del post"),
            Post("Mateo Valdelomar", "12:30 pm", "Contenido del post")
        )
        if(response.isNotEmpty()){
            response.forEach{ map ->
                val post = Post(
                    map.names,
                    map.hour,
                    map.content
                )
                listPost.add(post)
            }
        }
        return listPost
    }
}