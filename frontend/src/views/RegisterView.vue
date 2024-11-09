<script setup>
import { reactive } from 'vue';
import NavBar from '../components/NavBar.vue'
import Register from '../components/Register.vue'
import axios from 'axios';

function handleSubmit(data) {
    axios.post('http://localhost:8083/user/register', data)
        .then(response => {
            console.log('Registration successful:', response.data);
        })
        .catch(err => {
            console.error('Registration failed:', error);
        })
}
</script>
<template>
    <NavBar />
    <Register @submitForm="handleSubmit" />
</template>