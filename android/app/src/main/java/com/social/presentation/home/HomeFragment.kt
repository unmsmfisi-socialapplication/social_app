package com.social.presentation.home

import android.os.Bundle
import android.view.View
import androidx.core.content.ContextCompat
import androidx.fragment.app.Fragment
import com.social.R
import com.social.databinding.FragmentHomeBinding
import com.social.databinding.ItemCardPostBinding
import com.social.presentation.interaction.chats.ListChatsFragment
import com.social.utils.FragmentUtils

class HomeFragment : Fragment(R.layout.fragment_home) {
    private lateinit var binding: FragmentHomeBinding
    private lateinit var post: ItemCardPostBinding

    override fun onViewCreated(
        view: View,
        savedInstanceState: Bundle?,
    ) {
        super.onViewCreated(view, savedInstanceState)
        binding = FragmentHomeBinding.bind(view)
        post = ItemCardPostBinding.bind(view)

        action()
        setupLikes()
        setupComments()
        setupFavorites()
        setupLikes()
        setupShares()
    }

    private fun action() {
        binding.txtPub.setOnClickListener {
            FragmentUtils.replaceFragment(
                requireActivity().supportFragmentManager,
                NewPostFragment(),
            )
        }

        binding.messageIcon.setOnClickListener {
            FragmentUtils.replaceFragment(
                requireActivity().supportFragmentManager,
                ListChatsFragment(),
            )
        }
    }

    private fun setupLikes() {
        var isLiked = false
        var likeCount = 0

        post.iconLike.setOnClickListener {
            if (!isLiked) {
                likeCount++
                post.iconLike.setImageResource(R.drawable.likebutton)
                post.iconLike
            } else {
                likeCount--
                post.iconLike.setImageResource(R.drawable.likebutton)
                post.iconLike.clearColorFilter()
            }

            post.txtCountLikes.text = likeCount.toString()

            isLiked = !isLiked
        }
    }

    private fun setupShares() {
        var isShare = false
        var shareCount = 0

        post.iconShare.setOnClickListener {
            if (!isShare) {
                shareCount++
                post.iconShare.setImageResource(R.drawable.fowardbutton)
                post.iconShare.setColorFilter(
                    ContextCompat.getColor(
                        requireContext(),
                        R.color.color01,
                    ),
                )
            } else {
                shareCount--
                post.iconShare.setImageResource(R.drawable.fowardbutton)
                post.iconShare.clearColorFilter()
            }

            post.txtCountShare.text = shareCount.toString()

            isShare = !isShare
        }
    }

    private fun setupFavorites() {
        var isFavorite = false
        var favoriteCount = 0

        post.iconFavorite.setOnClickListener {
            if (!isFavorite) {
                favoriteCount++
                post.iconFavorite.setImageResource(R.drawable.favoritebutton)
                post.iconFavorite.setColorFilter(
                    ContextCompat.getColor(
                        requireContext(),
                        R.color.color01,
                    ),
                )
            } else {
                favoriteCount--
                post.iconFavorite.setImageResource(R.drawable.favoritebutton)
                post.iconFavorite.clearColorFilter()
            }

            post.txtCountFavorite.text = favoriteCount.toString()

            isFavorite = !isFavorite
        }
    }

    private fun setupComments() {
        var isComments = false
        var commentsCount = 0

        post.iconComment.setOnClickListener {
            if (!isComments) {
                commentsCount++
                post.iconComment.setImageResource(R.drawable.commentbutton)
                post.iconComment.setColorFilter(
                    ContextCompat.getColor(
                        requireContext(),
                        R.color.color01,
                    ),
                )
            } else {
                commentsCount--
                post.iconComment.setImageResource(R.drawable.commentbutton)
                post.iconComment.clearColorFilter()
            }

            post.txtCountComments.text = commentsCount.toString()

            isComments = !isComments
        }
    }
}
