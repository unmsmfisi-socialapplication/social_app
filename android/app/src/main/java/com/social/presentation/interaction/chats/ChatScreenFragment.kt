package com.social.presentation.interaction.chats

import android.os.Bundle
import android.view.View
import androidx.fragment.app.Fragment
import com.social.R
import com.social.databinding.FragmentChatScreenBinding
import com.social.domain.model.ChatUserData

class ChatScreenFragment : Fragment(R.layout.fragment_chat_screen) {
    private lateinit var binding: FragmentChatScreenBinding
    private lateinit var globalView: View
    private lateinit var receiverUser: ChatUserData

    override fun onViewCreated(
        view: View,
        savedInstanceState: Bundle?,
    ) {
        super.onViewCreated(view, savedInstanceState)
        binding = FragmentChatScreenBinding.bind(view)
        globalView = view
    }

    // Backend
}
