package com.social.presentation.interaction.chats

import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.social.data.repository.ChatRepository
import com.social.data.repository.ChatRepositoryImp
import com.social.domain.model.ChatUserData
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext

class ListChatsViewModel : ViewModel() {
    private var _chats: MutableLiveData<List<ChatUserData>> = MutableLiveData()
    var chats: LiveData<List<ChatUserData>> = _chats

    private val repository: ChatRepository = ChatRepositoryImp()

    fun obtainChats() {
        viewModelScope.launch(Dispatchers.Main) {
            try {
                val response =
                    withContext(Dispatchers.IO) {
                        repository.getChats()
                    }
                response.let {
                    _chats.value = it
                }
            } catch (e: Exception) {
                e.message
            }
        }
    }
}
