package com.social.domain.usecase

import com.social.data.source.remote.dto.InvalidUserException
import com.social.data.source.remote.dto.aCreatePost
import com.social.domain.SocialAppRepository
import com.social.domain.model.CreatePostBody
import com.social.domain.model.CreatePostResponse
import com.social.utils.Resource
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.flow
import java.net.UnknownServiceException
import javax.inject.Inject

class ValidateCreatePost
    @Inject
    constructor(
        private val socialAppRepository: SocialAppRepository,
    ) {
        operator fun invoke(createPostBody: CreatePostBody): Flow<Resource<List<CreatePostResponse>>> {
            return flow {
                try {
                    if (createPostBody.title.isBlank()) {
                        throw InvalidUserException("Ingrese título del post")
                    }
                    if (createPostBody.description.isBlank()) {
                        throw InvalidUserException("Ingrese descripción")
                    }
                    if (createPostBody.hasMultimedia) {
                        if (createPostBody.multimedia.isNullOrBlank()) {
                            throw InvalidUserException("Ingrese multimedia")
                        }
                    }
                    emit(Resource.Loading())
                    val post = socialAppRepository.createPost(createPostBody).aCreatePost()
                    when (post.status) {
                        "OK" -> {
                            emit(Resource.Success(listOf(post.response)))
                        }

                        else -> {
                            emit(Resource.Error("Error al crear el post"))
                        }
                    }
                } catch (e: Exception) {
                    emit(Resource.Error(e.message.toString()))
                } catch (u: UnknownServiceException) {
                    emit(Resource.Error(u.message.toString()))
                }
            }
        }
    }
