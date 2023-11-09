// Top-level build file where you can add configuration options common to all sub-projects/modules.
import org.jlleitschuh.gradle.ktlint.reporter.ReporterType

plugins {
    id("com.android.application") version "8.1.1" apply false
    id("com.android.library") version "8.1.1" apply false
    id("org.jetbrains.kotlin.android") version "1.8.20" apply false
    id("org.jlleitschuh.gradle.ktlint") version "11.6.0"
}

buildscript {
    dependencies {
        classpath("com.android.tools.build:gradle:7.3.0")
        classpath("com.google.firebase:firebase-appdistribution-gradle:4.0.1")
        val navegacionVersion = "2.7.3"
        classpath("androidx.navigation:navigation-safe-args-gradle-plugin:$navegacionVersion")
        classpath("com.google.dagger:hilt-android-gradle-plugin:2.48")
        classpath("com.google.gms:google-services:4.4.0")
        // classpath("com.android.tools.build:gradle:7.4.1")
    }
}

allprojects {
    repositories {
        google()
        mavenCentral()
        maven {
            setUrl("https://jitpack.io")
        }
    }
}

subprojects {
    apply(plugin = "org.jlleitschuh.gradle.ktlint")

    configure<org.jlleitschuh.gradle.ktlint.KtlintExtension> {
        version.set("1.0.0")
        verbose.set(true)
        android.set(true)

        reporters {
            reporter(ReporterType.HTML)
        }

        filter {
            exclude("**/generated/**")
            include("**/kotlin/**")
        }
    }
}
