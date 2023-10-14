package com.social.utils

import kotlin.random.Random

object CodeGenerator {
    fun generateCode(): String {
        return (1000 + Random.nextInt(9000)).toString()
    }
}
