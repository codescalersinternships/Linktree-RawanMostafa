<script setup>
import NavBar from '../components/NavBar.vue'
import Login from '../components/Login.vue'
import axios from 'axios'
function handleSubmit(data) {
    axios.post('http://localhost:8083/user/login', data)
        .then(response => {
            console.log('login successful:', response.data);
            localStorage.setItem('authToken', response.data.token);
            window.location.href = '/'+data.username;
        })
        .catch(err => {
            console.error('login failed:', error);
        })
}
</script>
<template>
<NavBar/>
<Login @submitForm="handleSubmit"/>
</template>