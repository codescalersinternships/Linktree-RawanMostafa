<script setup>
import NavBar from '../components/NavBar.vue'
</script>

<template>
  <div class="homepage">
    <NavBar />

    <div class="main">
      <img src="../assets/logo.png" alt="">
      <h1 class="animated-phrase">Welcome to Your Personalized Link Tree!</h1>
      <h2 class="sub-phrase">Connect all your links in one place</h2>
    </div>
  </div>
</template>

<style scoped>
body, html {
  margin: 0;
  padding: 0;
  width: 100%;
  font-family: Arial, sans-serif;
  background-color: #f0f0f0;
}

.homepage {
  width: 100vw;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  background: linear-gradient(to right, #46bdad, #86c6b8);
  color: white;
  text-align: center;
  padding: 10px;
}

.main {
  max-width: 90%;
  padding: 20px;
    width: 100%;

}

.animated-phrase {
  font-size: 3rem;
  font-weight: bold;
  margin: 0;
  color: #fff;
  opacity: 0;
  animation: fade 1.5s ease-in-out forwards, changeColor 6s infinite alternate;
}

.sub-phrase {
  font-size: 1.5rem;
  margin-top: 20px;
  color: #e0f7fa;
  opacity: 0;
  animation: fade 1.5s ease-in-out 1.5s forwards;
}

.main img{
  height: 100px;
  margin-bottom: 50px;
  animation: fade 1.5s ease-in-out forwards;
}

@keyframes fade {
  from {
    transform: translateY(20px);
    opacity: 0;
  }
  to {
    transform: translateY(0);
    opacity: 1;
  }
}

@keyframes changeColor {
  0% { color: #ffffff; }
  50% { color: #bbdefb; }
  100% { color: #f0f0f0; }
}

@media (max-width: 768px) {
  .animated-phrase {
    font-size: 2.5rem;
  }
  .sub-phrase {
    font-size: 1.2rem;
  }
}

@media (max-width: 480px) {
  .animated-phrase {
    font-size: 2rem;
    padding: 0 10px;
  }
  .sub-phrase {
    font-size: 1rem;
    padding: 0 10px;
  }
}
</style>
