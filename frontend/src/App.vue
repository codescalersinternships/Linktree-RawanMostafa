<script setup>
import Register from './components/Register.vue'
import Login from './components/Login.vue'
import Link from './components/Link.vue'
import AddLink from './components/AddLink.vue'
import NavBar from './components/NavBar.vue'
</script>



<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: 'App',
};
</script>
