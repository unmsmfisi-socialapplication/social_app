// Top-level build file where you can add configuration options common to all sub-projects/modules.
plugins {
    id("com.android.application") version "8.1.1" apply false
    id("org.jetbrains.kotlin.android") version "1.8.10" apply false
}

buildscript {
    dependencies {
        val navegacionVersion="2.7.2"
        classpath("androidx.navigation:navigation-safe-args-gradle-plugin:$navegacionVersion")
        classpath("com.google.dagger:hilt-android-gradle-plugin:2.47")
        classpath("com.google.gms:google-services:4.4.0")
    }
}