package com.social.presentation.publications

import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.social.data.repository.PostRepository
import com.social.data.repository.PostRepositoryImp
import com.social.domain.model.Post
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext

class ListPostViewModel : ViewModel() {
    private var _data: MutableLiveData<List<Post>> = MutableLiveData()
    val data: LiveData<List<Post>> = _data

    private val repository: PostRepository = PostRepositoryImp()

    fun obtainData() {
        viewModelScope.launch(Dispatchers.Main) {
            try {
                val response =
                    withContext(Dispatchers.IO) {
                        repository.obtainPost()
                    }
                response.let {
                    _data.value = it
                }
            } catch (e: Exception) {
                e.message
            }
        }
    }
}
