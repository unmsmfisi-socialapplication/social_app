package com.social.presentation.interaction.chats

import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.fragment.app.Fragment
import androidx.lifecycle.ViewModelProvider
import androidx.recyclerview.widget.LinearLayoutManager
import com.social.R
import com.social.databinding.ConteinerUserChatBinding
import com.social.databinding.FragmentListChatBinding
import com.social.domain.model.ChatUserData
import com.social.utils.BaseAdapter
import com.social.utils.Toast

class ListChatsFragment : Fragment(R.layout.fragment_list_chat) {
    private lateinit var binding: FragmentListChatBinding
    private lateinit var globalView: View

    private val adapter: BaseAdapter<ChatUserData> =
        object : BaseAdapter<ChatUserData>(emptyList()) {
            override fun getViewHolder(parent: ViewGroup): BaseViewHolder<ChatUserData> {
                val view =
                    LayoutInflater.from(parent.context)
                        .inflate(R.layout.conteiner_user_chat, parent, false)
                return object : BaseViewHolder<ChatUserData>(view) {
                    private val binding: ConteinerUserChatBinding = ConteinerUserChatBinding.bind(view)

                    override fun bind(chatUser: ChatUserData) =
                        with(binding) {
                            chatUserName.text = chatUser.name
                            chatMessage.text = chatUser.message
                            hourSendMessage.text = chatUser.hourSend
                            countNotificationsMessages.text = chatUser.countNotification
                        }
                }
            }
        }

    private val viewModel: ListChatsViewModel by lazy {
        ViewModelProvider(this)[ListChatsViewModel::class.java]
    }

    override fun onViewCreated(
        view: View,
        savedInstanceState: Bundle?,
    ) {
        super.onViewCreated(view, savedInstanceState)
        binding = FragmentListChatBinding.bind(view)
        globalView = view

        setupAdapter()
        observeViewModel()
        action()
    }

    private fun setupAdapter() {
        binding.rcyViewlistChats.layoutManager = LinearLayoutManager(requireContext())
        binding.rcyViewlistChats.adapter = adapter
    }

    private fun observeViewModel() {
        viewModel.chats.observe(viewLifecycleOwner) { chats ->
            adapter.updateList(chats)
        }
        viewModel.obtainChats()
    }

    private fun action() {
        binding.btnBar.setOnClickListener {
            barButton()
        }
    }

    private fun barButton() {
        Toast.showMessage(requireContext(), "En desarrollo")
    }
}
