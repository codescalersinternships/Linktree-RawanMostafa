<template>
    <div class="hero">
    <h1 class="title">Sign Up</h1>
    <form class="signup" method="post" @submit.prevent="submitForm">
        <div v-if='ChecksError.message' class="error-message">{{ ChecksError.message }}</div>
        <div class="form-field">
            <label for="first-name">First Name</label>
            <input type="text" id="first-name" class="first-name" v-model="data.first_name" required>
        </div>

        <div class="form-field">
            <label for="second-name">Second Name</label>
            <input type="text" id="second-name" class="second-name" v-model="data.second_name" required>
        </div>

        <div class="form-field">
            <label for="username">Username</label>
            <input type="text" id="username" class="username"  v-model="data.username" required>
        </div>

        <div class="form-field">
            <label for="password">Password</label>
            <input type="password" id="password" class="password" v-model="data.password" required>
        </div>

        <div class="form-field">
            <label for="conform-password">Confirm password</label>
            <input type="password" id="conform-password" class="conform-password"  v-model="data.confirmPassword" required>

        </div>
        <button class="submit-btn" type="submit">Sign Up</button>
    </form>
    <h4>Already have an account? 
        <a href="/login" class="login-link">Login</a>
    </h4>
</div>
</template>

<script setup>
import { reactive, defineEmits } from 'vue';

const data = reactive({
    username: '',
    password: '',
    first_name: '',
    second_name: '',
    confirmPassword: ''
});
const ChecksError = reactive({message:''});
const emit = defineEmits(['submitForm']);

function submitForm() {
    if (data.password !== data.confirmPassword) {
        ChecksError.message = "Passwords do not match";
    } else {
        const { confirmPassword, ...formData } = data;
        emit('submitForm', formData);
    }
}
</script>

<style>
html,
body {
    height: 100%;
}

body {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    background-color: rgba(209, 243, 240, 0.634);
}

.signup {
    width: fit-content;
    border: 3px solid rgb(51, 93, 107);
    border-radius: 20px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    background-color: rgba(147, 174, 172, 0.634);
}

.title {
    color: rgb(36, 73, 70);
    text-align: center;
    
}

.form-field {
    margin-bottom: 15px;
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
}

.form-field input {
    padding: 10px;
    margin: 20px;
    font-size: 16px;
    width: 60%;
    border-radius: 5px;
}

.form-field label {
    margin: 20px;
    width: 150px;
    color: rgb(36, 73, 70)
}

.submit-btn {
    padding: 10px 20px;
    font-size: 16px;
    margin: 20px;
    background-color: rgb(51, 93, 107);
    color: white;
    border: none;
    border-radius: 5px;
    cursor: pointer;
}

.submit-btn:hover {
    background-color: rgb(35, 64, 74);
}
a,h4 {
color: rgb(36, 73, 70);
}
.hero {
    padding-top: 100px; 
}
.error-message{
    color: red;
    margin-top: 20px;
}

</style>