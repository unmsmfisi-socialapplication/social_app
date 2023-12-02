package com.social.presentation.interaction.chats

import android.graphics.Bitmap
import android.graphics.BitmapFactory
import android.util.Base64
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.recyclerview.widget.RecyclerView
import com.social.databinding.ConteinerUserChatBinding
import com.social.domain.model.ChatUserData

class UserAdapter(val users: List<ChatUserData>) : RecyclerView.Adapter<UserAdapter.UserViewHolder>() {
    private lateinit var binding: ConteinerUserChatBinding
    private lateinit var globalView: View

    override fun onCreateViewHolder(
        parent: ViewGroup,
        viewType: Int,
    ): UserViewHolder {
        val inflater = LayoutInflater.from(parent.context)
        binding = ConteinerUserChatBinding.inflate(inflater, parent, false)
        globalView = binding.root
        return UserViewHolder(binding)
    }

    override fun onBindViewHolder(
        holder: UserAdapter.UserViewHolder,
        position: Int,
    ) {
        holder.bindUserData(users[position])
    }

    override fun getItemCount(): Int {
        return users.size
    }

    inner class UserViewHolder(private val binding: ConteinerUserChatBinding) :
        RecyclerView.ViewHolder(binding.root) {
        fun bindUserData(user: ChatUserData) {
            binding.chatUserName.text = user.name
            binding.chatMessage.text = user.message
            binding.hourSendMessage.text = user.hourSend
            binding.countNotificationsMessages.text = user.countNotification
            binding.contactChatProfile.setImageBitmap(getUserImage(user.image))
        }
    }

    private fun getUserImage(encodedImage: String): Bitmap? {
        val bytes = Base64.decode(encodedImage, Base64.DEFAULT)
        return BitmapFactory.decodeByteArray(bytes, 0, bytes.size)
    }
}
