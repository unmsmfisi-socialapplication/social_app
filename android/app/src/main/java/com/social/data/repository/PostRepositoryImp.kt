package com.social.data.repository

import com.social.domain.model.Post

class PostRepositoryImp : PostRepository {
    override suspend fun obtainPost(): MutableList<Post> {
        val listPost = mutableListOf<Post>()

        // Hardcoded lists, then obtained from backend
        val response =
            listOf(
                Post(
                    "Juan Pérez",
                    "8:00 am",
                    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Erat vitae quis quam augue quam a.",
                    "https://firebasestorage.googleapis.com/v0/b/martes-ce.appspot.com/o/" +
                        "F3LNvxmUhjVsKfSi9DKVD5dnOO52%2Fimages%2F1690423586" +
                        ".jpg?alt=media&token=adafc9eb-6c22-4c53-accd-a0e079aa6d8b",
                ),
                Post(
                    "Maria López",
                    "9:30 am",
                    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Erat vitae quis quam augue quam a.",
                    "",
                ),
                Post(
                    "Pedro García",
                    "11:15 am",
                    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Erat vitae quis quam augue quam a.",
                    "https://firebasestorage.googleapis.com/v0/b/martes-ce.appspot.com/o/" +
                        "F3LNvxmUhjVsKfSi9DKVD5dnOO52%2Fimages%2F1690666127" +
                        ".jpg?alt=media&token=40f03f19-c134-4649-8ca2-ec6a23f0011f",
                ),
                Post(
                    "Luis Rodríguez",
                    "2:45 pm",
                    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Erat vitae quis quam augue quam a.",
                    "",
                ),
                Post(
                    "Ana Martínez",
                    "4:20 pm",
                    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Erat vitae quis quam augue quam a.",
                    "https://firebasestorage.googleapis.com/v0/b/martes-ce.appspot.com/o/" +
                        "31p0RbLYu3PEV7iCta4M3nM64j32%2Fimages%2F1691034116" +
                        ".jpg?alt=media&token=57e92e97-17bc-4392-9519-cdbdd0aa7838",
                ),
                Post(
                    "Sofia Fernández",
                    "6:10 pm",
                    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Erat vitae quis quam augue quam a.",
                    "https://firebasestorage.googleapis.com/v0/b/martes-ce.appspot.com/o/" +
                        "BHGVg2vNO2MacVzavMZ6AYbuxV23%2Fimages%2F1698368087" +
                        ".jpg?alt=media&token=c2664d6d-cfc3-420e-a6ad-827b26356603",
                ),
            )
        response.forEach { map ->
            val post =
                Post(
                    map.names,
                    map.hour,
                    map.content,
                    map.image,
                )
            listPost.add(post)
        }
        return listPost
    }
}
