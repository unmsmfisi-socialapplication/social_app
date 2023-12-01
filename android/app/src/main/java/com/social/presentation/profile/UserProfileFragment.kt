package com.social.presentation.profile

import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import androidx.core.content.ContextCompat
import androidx.fragment.app.Fragment
import androidx.lifecycle.ViewModelProvider
import com.google.android.material.imageview.ShapeableImageView
import com.social.R
import com.social.databinding.FragmentUserProfileBinding
import com.social.databinding.ItemPostBinding
import com.social.domain.model.Post
import com.social.presentation.publications.ListPostViewModel
import com.social.utils.BaseAdapter
import com.social.utils.FragmentUtils.replaceFragment
import com.squareup.picasso.Picasso

class UserProfileFragment : Fragment(R.layout.fragment_user_profile) {
    private lateinit var binding: FragmentUserProfileBinding
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
        binding = FragmentUserProfileBinding.bind(view)
        globalView = view

        setupAdapter()
        observeViewModel()
        action()
        setupClickListeners()

        updateTextViewsWithKFormat(
            binding.textNumberPost,
            binding.textNumberPhotos,
            binding.textNumberFollowers,
            binding.textNumberFollowings,
        )
    }

    private fun action() {
        binding.buttonEditProfile.setOnClickListener {
            replaceFragment(
                requireActivity().supportFragmentManager,
                EditProfileFragment(),
            )
        }

        binding.buttonSettingProfile.setOnClickListener {
            replaceFragment(
                requireActivity().supportFragmentManager,
                SettingProfileFragment(),
            )
        }
    }

    private fun setupClickListeners() {
        val textPost = binding.textPost
        val textDestacados = binding.textDestacados
        val textActividad = binding.textActividad

        textPost.setOnClickListener {
            updateTextStyle(textPost)
            resetTextStyle(textDestacados, textActividad)
            binding.recyclerPostProfile.visibility = View.VISIBLE
        }

        textDestacados.setOnClickListener {
            updateTextStyle(textDestacados)
            resetTextStyle(textPost, textActividad)
            binding.recyclerPostProfile.visibility = View.GONE
        }

        textActividad.setOnClickListener {
            updateTextStyle(textActividad)
            resetTextStyle(textPost, textDestacados)
            binding.recyclerPostProfile.visibility = View.GONE
        }
    }

    private fun updateTextStyle(textView: TextView) {
        textView.setTextColor(ContextCompat.getColor(requireContext(), R.color.black))
    }

    private fun resetTextStyle(vararg textViews: TextView) {
        for (textView in textViews) {
            textView.setTextColor(ContextCompat.getColor(requireContext(), R.color.color03))
        }
    }

    private fun updateTextViewsWithKFormat(vararg textViews: TextView) {
        for (textView in textViews) {
            val textValue = textView.text.toString()
            val numberToConvert = textValue.toIntOrNull() ?: 0
            val convertedNumber = convertNumberToK(numberToConvert)
            textView.text = convertedNumber
        }
    }

    private fun convertNumberToK(number: Int): String {
        return when {
            number in 1000..999999 -> "${number / 1000}k"
            number >= 1000000 -> "${number / 1000000}M"
            else -> number.toString()
        }
    }

    private fun setupAdapter() {
        binding.recyclerPostProfile.adapter = adapter
    }

    private fun observeViewModel() {
        viewModel.data.observe(viewLifecycleOwner) { posts ->
            adapter.updateList(posts)
        }
        viewModel.obtainData()
    }

    private fun loadImage(
        imageURL: String,
        imageView: ShapeableImageView,
    ) {
        Picasso.get().load(imageURL).into(imageView)
    }
}
