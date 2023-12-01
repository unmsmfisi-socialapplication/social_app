package com.social.presentation.home

import android.app.AlertDialog
import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.view.WindowManager
import android.widget.ImageView
import android.widget.TextView
import androidx.fragment.app.Fragment
import androidx.lifecycle.ViewModelProvider
import androidx.recyclerview.widget.LinearLayoutManager
import com.google.android.material.imageview.ShapeableImageView
import com.social.R
import com.social.databinding.FragmentHomeBinding
import com.social.databinding.ItemPostBinding
import com.social.domain.model.Post
import com.social.presentation.interaction.chats.ListChatsFragment
import com.social.presentation.publications.ListPostViewModel
import com.social.utils.BaseAdapter
import com.social.utils.FragmentUtils
import com.squareup.picasso.Picasso

class HomeFragment : Fragment(R.layout.fragment_home) {
    private lateinit var binding: FragmentHomeBinding
    private lateinit var globalView: View

    private val adapter: BaseAdapter<Post> =
        object : BaseAdapter<Post>(emptyList()) {
            override fun getViewHolder(parent: ViewGroup): BaseViewHolder<Post> {
                val view =
                    LayoutInflater.from(parent.context)
                        .inflate(R.layout.item_post, parent, false)
                return object : BaseViewHolder<Post>(view) {
                    private val binding: ItemPostBinding = ItemPostBinding.bind(view)

                    override fun bind(entity: Post) =
                        with(binding) {
                            textNames.text = entity.names
                            textHour.text = entity.hour
                            textContentPost.text = entity.content
                            if (entity.image.isNotEmpty()) {
                                loadImage(entity.image, binding.imagePost)
                            } else {
                                binding.imagePost.visibility = View.GONE
                            }
                            binding.iconLike.setOnClickListener {
                                handleIconLikeClick(entity, binding.iconLike, binding.txtCountLikes)
                            }
                            binding.txtCountLikes.text = entity.likeCount.toString()
                            binding.iconComment.setOnClickListener {
                                showCommentsFullScreen()
                            }
                        }
                }
            }
        }

    private val viewModel: ListPostViewModel by lazy {
        ViewModelProvider(this)[ListPostViewModel::class.java]
    }

    override fun onViewCreated(
        view: View,
        savedInstanceState: Bundle?,
    ) {
        super.onViewCreated(view, savedInstanceState)
        binding = FragmentHomeBinding.bind(view)
        globalView = view

        setupAdapter()
        observeViewModel()
        action()
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

    fun loadImage(
        imageURL: String,
        imageView: ShapeableImageView,
    ) {
        Picasso.get().load(imageURL).into(imageView)
    }

    private fun setupAdapter() {
        binding.rvListPostHome.layoutManager = LinearLayoutManager(requireContext())
        binding.rvListPostHome.adapter = adapter
    }

    private fun observeViewModel() {
        viewModel.data.observe(viewLifecycleOwner) { posts ->
            adapter.updateList(posts)
        }
        viewModel.obtainData()
    }

    fun handleIconLikeClick(
        post: Post,
        iconLike: ImageView,
        countLikes: TextView,
    ) {
        post.isLiked = !post.isLiked
        if (post.isLiked) {
            post.likeCount++
            iconLike.setImageResource(R.drawable.post_icon_heart_bold)
        } else {
            post.likeCount--
            iconLike.setImageResource(R.drawable.post_icon_like)
        }
        countLikes.text = post.likeCount.toString()
    }

    private fun showCommentsFullScreen() {
        val dialogView = LayoutInflater.from(requireContext()).inflate(R.layout.item_comment, null)
        val dialog =
            AlertDialog.Builder(requireContext()).apply {
                setView(dialogView)
                setCancelable(true)
            }.create()

        dialog.show()

        val layoutParams =
            WindowManager.LayoutParams().apply {
                copyFrom(dialog.window?.attributes)
                width = WindowManager.LayoutParams.MATCH_PARENT
                height = WindowManager.LayoutParams.MATCH_PARENT
            }

        dialog.window?.attributes = layoutParams

        val iconRetract = dialogView.findViewById<ImageView>(R.id.icon_retract)
        iconRetract.setOnClickListener {
            dialog.dismiss()
        }
    }
}
