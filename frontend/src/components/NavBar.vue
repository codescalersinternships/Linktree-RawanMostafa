<template>
    <div class="navbar">
        <div class="logo">
            <img src="../assets/logo.png" alt="">
            <a class="title" href="#">Linktree</a>
        </div>
        <div class="links">
            <a href="#" class="link">Home</a>
            <a href="#" class="link">Profile</a>
            <a href="#" class="link">Register</a>
        </div>
    </div>
</template>

<style>
body {
    box-sizing: border-box;
    width: 100%;
    height: 100%;
    font-family: Arial, sans-serif;
    background-color: rgba(209, 243, 240, 0.634);
}

.navbar {
    width: 100vw;
    background-color: rgba(147, 174, 172, 0.634);
    border-bottom: 3px solid rgb(51, 93, 107);
    display: flex;
    justify-content: space-between;
    align-items: center;
    position: sticky;
    top: 0;
    z-index: 1000;
}

.logo {
    font-weight: bold;
    display: flex;
    flex-direction: row;
    align-items: center;
}
.logo h2{
    color: rgb(36, 73, 70);
}
.logo img{
    height: 40px;
    margin-right: 10px;
}

.links {
    display: flex;
    gap: 30px;
    font-family:'Courier New', Courier, monospace;
}

.link {
    color: rgb(36, 73, 70);
    font-size: 18px;
    text-decoration: none;
    font-weight: 500;
    padding: 5px 5px;
    border-radius: 7px;
    transition: background-color 0.3s, color 0.3s;
}
.link:hover {
    background-color: rgba(52, 107, 102, 0.5);
    color: rgb(15, 40, 38);
}
.title{
    color: rgb(36, 73, 70);
    font-family:'Courier New', Courier, monospace;
    text-decoration: none;
    font-size: 25px;
}
</style>
