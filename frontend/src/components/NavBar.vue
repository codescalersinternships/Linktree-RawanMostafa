<template>
    
    <div class="navbar">
        <div class="logo">
            <img src="../assets/logo.png" alt="Logo">
            <a class="title" href="#">Linktree</a>
        </div>
        <div class="hamburger" @click="toggleDropdown">
            ☰
        </div>
        <div class="links" :class="{ 'show': dropdownVisible }">
            <a href="/" class="link">Home</a>
            <a href="#" class="link">Profile</a>
            <a href="/register" class="link">Register</a>
        </div>
    </div>
</template>


<script setup>
import { ref } from 'vue';

const dropdownVisible = ref(false);
function toggleDropdown() {
    dropdownVisible.value = !dropdownVisible.value;
}
</script>


<style>
html,
body {
    width: 100%;
    height: 100%;
    font-family: Arial, sans-serif;
    background-color: rgba(209, 243, 240, 0.634);
    margin: 0;
    padding: 0;
}

.navbar {
    width: 100%;
    background-color: rgba(147, 174, 172);
    border-bottom: 3px solid rgb(51, 93, 107);
    display: flex;
    justify-content: space-between;
    align-items: center;
    position: fixed;
    top: 0;
    left: 0;
    z-index: 1000;
    padding: 10px 20px;
}

.logo {
    display: flex;
    align-items: center;
}

.logo img {
    height: 40px;
    margin-right: 10px;
}

.links {
    display: flex;
    gap: 20px;
    margin-right: 20px;
    font-family:'Courier New', Courier, monospace;

}

.link {
    color: rgb(36, 73, 70);
    font-size: 18px;
    text-decoration: none;
    font-weight: 500;
    padding: 8px 12px;
    border: none;
    border-radius: 5px;
    transition: background-color 0.3s ease, color 0.3s ease;
}

.link:hover {
    background-color: rgba(52, 107, 102, 0.3);
    color: rgb(15, 40, 38);
}

.title{
    color: rgb(36, 73, 70);
    font-family:'Courier New', Courier, monospace;
    text-decoration: none;
    font-size: 25px;
    font-weight: 600;
}

.hamburger {
    display: none;
    font-size: 30px;
    cursor: pointer;
    margin-right: 50px;
    color: rgb(36, 73, 70);
}

@media (max-width: 600px) {
    .links {
        display: none;
        flex-direction: column;
        gap: 10px;
        position: absolute;
        top: 60px;
        right: 20px;
        background-color: rgba(147, 174, 172);
        border-radius: 5px;
        padding: 10px;
    }

    .links.show {
        display: flex;
    }

    .hamburger {
        display: block;
    }
}
</style>
